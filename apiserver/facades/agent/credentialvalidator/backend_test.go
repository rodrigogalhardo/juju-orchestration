// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package credentialvalidator_test

import (
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/names.v2"

	"github.com/juju/juju/apiserver/facades/agent/credentialvalidator"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/state"
	statetesting "github.com/juju/juju/state/testing"
	coretesting "github.com/juju/juju/testing"
)

type BackendSuite struct {
	coretesting.BaseSuite

	state   *mockState
	backend credentialvalidator.Backend
}

var _ = gc.Suite(&BackendSuite{})

func (s *BackendSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)
	s.state = newMockState()

	s.backend = credentialvalidator.NewBackend(s.state)
}

func (s *BackendSuite) TestModelUsesCredential(c *gc.C) {
	uses, err := s.backend.ModelUsesCredential(s.state.aModel.credentialTag)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(uses, jc.IsTrue)
	s.state.CheckCallNames(c, "Model", "mockModel.CloudCredential")
}

func (s *BackendSuite) TestModelUsesCredentialUnset(c *gc.C) {
	s.state.aModel.credentialSet = false
	uses, err := s.backend.ModelUsesCredential(s.state.aModel.credentialTag)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(uses, jc.IsFalse)
	s.state.CheckCallNames(c, "Model", "mockModel.CloudCredential")
}

func (s *BackendSuite) TestModelUsesCredentialWrongCredential(c *gc.C) {
	uses, err := s.backend.ModelUsesCredential(names.NewCloudCredentialTag("foo/bob/two"))
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(uses, jc.IsFalse)
	s.state.CheckCallNames(c, "Model", "mockModel.CloudCredential")
}

func (s *BackendSuite) TestModelCredentialUnsetNotSupported(c *gc.C) {
	s.state.aModel.credentialSet = false
	mc, err := s.backend.ModelCredential()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(mc, gc.DeepEquals, &credentialvalidator.ModelCredential{
		Exists:     false,
		Credential: names.CloudCredentialTag{},
		Valid:      false,
	})
	s.state.CheckCallNames(c, "Model", "mockModel.CloudCredential", "ModelTag", "Cloud")
}

func (s *BackendSuite) TestModelCredentialUnsetSupported(c *gc.C) {
	s.state.aModel.credentialSet = false
	s.state.aCloud.AuthTypes = cloud.AuthTypes{cloud.EmptyAuthType}
	mc, err := s.backend.ModelCredential()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(mc, gc.DeepEquals, &credentialvalidator.ModelCredential{
		Exists:     false,
		Credential: names.CloudCredentialTag{},
		Valid:      true,
	})
	s.state.CheckCallNames(c, "Model", "mockModel.CloudCredential", "ModelTag", "Cloud")
}

func (s *BackendSuite) TestModelCredentialSetButCloudCredentialNotFound(c *gc.C) {
	s.state.SetErrors(
		nil,                      // Model
		errors.NotFoundf("lost"), // CloudCredential
	)
	mc, err := s.backend.ModelCredential()
	c.Assert(err, gc.ErrorMatches, "lost not found")
	c.Assert(mc, gc.IsNil)
	s.state.CheckCallNames(c, "Model", "mockModel.CloudCredential", "ModelTag", "mockState.CloudCredential")
}

func newMockState() *mockState {
	b := &mockState{
		Stub:        &testing.Stub{},
		aCredential: statetesting.NewEmptyCredential(),
		aCloud: cloud.Cloud{
			Name:      "stratus",
			Type:      "low",
			AuthTypes: cloud.AuthTypes{cloud.AccessKeyAuthType, cloud.UserPassAuthType},
		},
	}
	b.aModel = &mockModel{
		Stub:          b.Stub,
		credentialTag: names.NewCloudCredentialTag("foo/bob/one"),
		credentialSet: true,
	}
	return b
}

type mockState struct {
	*testing.Stub

	aCloud      cloud.Cloud
	aModel      *mockModel
	aCredential state.Credential
}

func (b *mockState) Model() (credentialvalidator.ModelAccessor, error) {
	b.AddCall("Model")
	if err := b.NextErr(); err != nil {
		return nil, err
	}
	return b.aModel, nil
}

func (b *mockState) CloudCredential(tag names.CloudCredentialTag) (state.Credential, error) {
	b.AddCall("mockState.CloudCredential", tag)
	if err := b.NextErr(); err != nil {
		return state.Credential{}, err
	}
	return b.aCredential, nil
}

func (b *mockState) WatchCredential(tag names.CloudCredentialTag) state.NotifyWatcher {
	b.AddCall("WatchCredential", tag)
	return apiservertesting.NewFakeNotifyWatcher()
}

func (b *mockState) InvalidateModelCredential(reason string) error {
	b.AddCall("InvalidateModelCredential", reason)
	return b.NextErr()
}

func (b *mockState) Cloud(name string) (cloud.Cloud, error) {
	b.AddCall("Cloud", name)
	if err := b.NextErr(); err != nil {
		return cloud.Cloud{}, err
	}
	return b.aCloud, nil
}

type mockModel struct {
	*testing.Stub

	modelTag names.ModelTag

	credentialTag names.CloudCredentialTag
	credentialSet bool

	cloud string
}

func (m *mockModel) CloudCredential() (names.CloudCredentialTag, bool) {
	m.MethodCall(m, "mockModel.CloudCredential")
	return m.credentialTag, m.credentialSet
}

func (m *mockModel) ModelTag() names.ModelTag {
	m.MethodCall(m, "ModelTag")
	return m.modelTag
}

func (m *mockModel) Cloud() string {
	return m.cloud
}
