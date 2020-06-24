// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caas

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/juju/cmd"
	"github.com/juju/errors"

	"github.com/juju/juju/caas"
	"github.com/juju/juju/caas/kubernetes/clientconfig"
	"github.com/juju/juju/cmd/juju/interact"
)

const eksDomain = "eksctl.io"

type eks struct {
	tool string
	CommandRunner
}

func newEKSCluster() k8sCluster {
	return &eks{
		tool:          "eksctl",
		CommandRunner: &defaultRunner{},
	}
}

func (e *eks) cloud() string {
	return caas.K8sCloudEC2
}

func (e *eks) ensureExecutable() error {
	cmd := []string{"which", e.tool}
	err := collapseRunError(runCommand(e, cmd, ""))
	errAnnotationMessage := fmt.Sprintf(`%q not found. Please install %q (see: https://eksctl.io/introduction/#installation), login, and try again`, e.tool, e.tool)
	if err != nil {
		return errors.Errorf(errAnnotationMessage)
	}

	// check that we are logged in, there is no way to provide login details to a separate command.
	cmd = []string{e.tool, "get", "cluster"}
	err = collapseRunError(runCommand(e, cmd, ""))
	if err != nil {
		return errors.Errorf("please ensure the account has been setup and re-run this command")
	}
	return nil
}

func (e *eks) getKubeConfig(p *clusterParams) (io.ReadCloser, string, error) {
	kubeconfig := clientconfig.GetKubeConfigPath()
	cmd := []string{
		e.tool, "utils", "write-kubeconfig",
		"--cluster", p.name,
		"--kubeconfig", kubeconfig,
	}
	if len(p.region) > 0 {
		cmd = append(cmd, "--region", p.region)
	}

	err := collapseRunError(runCommand(e, cmd, kubeconfig))
	if err != nil {
		return nil, "", errors.Trace(err)
	}
	reader, err := os.Open(kubeconfig)
	return reader, fmt.Sprintf("%s.%s.%s", p.name, p.region, eksDomain), err
}

func (e *eks) interactiveParams(ctx *cmd.Context, p *clusterParams) (*clusterParams, error) {
	errout := interact.NewErrWriter(ctx.Stdout)
	pollster := interact.New(ctx.Stdin, ctx.Stdout, errout)

	if len(p.region) == 0 {
		var err error
		if p.region, err = pollster.Enter("region"); err != nil {
			return nil, errors.Trace(err)
		}
	}
	if len(p.name) == 0 {
		clusters, err := e.listClusters(p.region)
		if err != nil {
			return nil, errors.Trace(err)
		}

		if len(clusters) == 0 {
			return nil, errors.NewNotFound(nil, fmt.Sprintf("no cluster found in region %q", p.region))
		}
		if len(clusters) == 1 {
			p.name = clusters[0]
			return p, nil
		}
		if p.name, err = pollster.Select(interact.List{
			Singular: "cluster",
			Plural:   fmt.Sprintf("Available clusters in %s", p.region),
			Options:  clusters,
			Default:  clusters[0],
		}); err != nil {
			return nil, errors.Trace(err)
		}
	}
	return p, nil
}

type eksDetail struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

func (e *eks) listClusters(region string) (clusters []string, err error) {
	cmd := []string{
		e.tool, "get", "cluster", "--region", region, "-o", "json",
	}
	result, err := runCommand(e, cmd, "")
	if err != nil {
		return nil, errors.Trace(err)
	}
	if result.Code != 0 {
		return nil, errors.New(string(result.Stderr))
	}

	var eksClusters []eksDetail
	if err := json.Unmarshal(result.Stdout, &eksClusters); err != nil {
		return nil, errors.Trace(err)
	}

	for _, item := range eksClusters {
		clusters = append(clusters, item.Name)
	}
	return clusters, nil
}