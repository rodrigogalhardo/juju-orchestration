// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package space_test

import (
	"bytes"

	"github.com/juju/cmd"
	"github.com/juju/cmd/cmdtesting"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/cmd/juju/space"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/jujuclient"
)

type RemoveSuite struct {
	BaseSpaceSuite

	store *jujuclient.MemStore
}

var _ = gc.Suite(&RemoveSuite{})

func (s *RemoveSuite) SetUpTest(c *gc.C) {
	s.BaseSpaceSuite.SetUpTest(c)
	s.newCommand = space.NewRemoveCommand

	s.store = jujuclient.NewMemStore()
	s.store.Controllers["foo"] = jujuclient.ControllerDetails{}
	s.store.CurrentControllerName = "foo"
	s.store.Accounts["foo"] = jujuclient.AccountDetails{
		User: "bar", Password: "hunter2",
	}
	err := s.store.UpdateModel("foo", "bar/currentfoo",
		jujuclient.ModelDetails{ModelUUID: "uuidfoo1", ModelType: model.IAAS})
	c.Assert(err, jc.ErrorIsNil)

	err = s.store.SetCurrentModel("foo", "bar/currentfoo")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *RemoveSuite) runCommand(c *gc.C, api space.SpaceAPI, args ...string) (*cmd.Context, *space.RemoveCommand, error) {
	spaceCmd := &space.RemoveCommand{
		SpaceCommandBase: space.NewSpaceCommandBase(api),
	}
	cmd := modelcmd.Wrap(spaceCmd)
	cmd.SetClientStore(s.store)
	ctx, err := cmdtesting.RunCommand(c, cmd, args...)
	return ctx, spaceCmd, err
}

func (s *RemoveSuite) TestInit(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	for i, test := range []struct {
		about      string
		args       []string
		expectName string
		expectErr  string
	}{{
		about:     "no arguments",
		expectErr: "space name is required",
	}, {
		about:     "invalid space name",
		args:      s.Strings("%inv$alid", "new-name"),
		expectErr: `"%inv\$alid" is not a valid space name`,
	}, {
		about:      "multiple space names aren't allowed",
		args:       s.Strings("a-space", "another-space"),
		expectErr:  `unrecognized args: \["another-space"\]`,
		expectName: "a-space",
	}, {
		about:      "delete a valid space name",
		args:       s.Strings("myspace"),
		expectName: "myspace",
	}} {
		c.Logf("test #%d: %s", i, test.about)
		if test.expectErr == "" {
			api.EXPECT().RemoveSpace(test.expectName, false, false).Return(network.RemoveSpace{}, nil)
		}
		_, cmd, err := s.runCommand(c, api, test.args...)
		if test.expectErr != "" {
			prefixedErr := "invalid arguments specified: " + test.expectErr
			c.Check(err, gc.ErrorMatches, prefixedErr)
		} else {
			c.Check(err, jc.ErrorIsNil)
			c.Check(cmd.Name(), gc.Equals, test.expectName)
		}
	}
}

func (s *RemoveSuite) TestRunWithValidSpaceSucceeds(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	spaceName := "default"
	api.EXPECT().RemoveSpace(spaceName, false, false).Return(network.RemoveSpace{}, nil)
	ctx, _, err := s.runCommand(c, api, spaceName)

	c.Assert(err, gc.IsNil)
	c.Assert(ctx.Stderr.(*bytes.Buffer).String(), gc.Equals, "removed space \"default\"\n")
}

func (s *RemoveSuite) TestRunWithForceNoConfirmation(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	spaceName := "default"

	api.EXPECT().RemoveSpace(spaceName, true, false).Return(network.RemoveSpace{}, nil)

	_, _, err := s.runCommand(c, api, spaceName, "--force", "-y")

	c.Assert(err, jc.ErrorIsNil)
}

func (s *RemoveSuite) TestRunWithForceWithConfirmation(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	spaceName := "myspace"

	spaceRemove := network.RemoveSpace{
		HasModelConstraint: true,
		Space:              spaceName,
		Constraints:        []string{"mysql"},
		Bindings:           []string{"mysql", "mediawiki"},
		ControllerConfig:   []string{"jujuhaspace", "juuuu-space"},
	}
	api.EXPECT().RemoveSpace(spaceName, false, true).Return(spaceRemove, nil)
	expectedErrMsg := `
WARNING! This command will remove the space with the following existing boundaries:


- "myspace" is used as a constraint on: mysql
- "myspace" is used as a model constraint: bar/currentfoo
- "myspace" is used as a binding on: mysql, mediawiki
- "myspace" is used for controller config(s): jujuhaspace, juuuu-space


Continue [y/N]?`[1:]

	ctx, _, err := s.runCommand(c, api, spaceName, "--force")

	c.Assert(cmdtesting.Stdout(ctx), gc.Equals, expectedErrMsg)
	c.Assert(err, gc.ErrorMatches, "cannot remove space \"myspace\": space removal: aborted")
}

func (s *RemoveSuite) TestRunWithForceWithNoError(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	spaceName := "default"
	api.EXPECT().RemoveSpace(spaceName, false, true).Return(network.RemoveSpace{}, nil)
	expectedErrMsg := `
WARNING! This command will remove the space. 
Safe removal possible. No constraints, bindings or controller config found with dependency on the given space.

Continue [y/N]? `[1:]

	ctx, _, err := s.runCommand(c, api, spaceName, "--force")

	c.Assert(cmdtesting.Stdout(ctx), gc.Equals, expectedErrMsg)
	c.Assert(err, gc.ErrorMatches, "cannot remove space \"default\": space removal: aborted")
}

func (s *RemoveSuite) TestRunWhenSpacesAPIFails(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	spaceName := "default"
	apiErr := &params.Error{Code: params.CodeOperationBlocked, Message: "nope"}
	api.EXPECT().RemoveSpace(spaceName, false, false).Return(network.RemoveSpace{}, apiErr)
	ctx, _, err := s.runCommand(c, api, spaceName)

	c.Assert(err, gc.ErrorMatches, "cannot remove space \"default\": nope")
	c.Assert(ctx.Stderr.(*bytes.Buffer).String(), gc.Equals, "")
	c.Assert(ctx.Stdout.(*bytes.Buffer).String(), gc.Equals, "")

}
