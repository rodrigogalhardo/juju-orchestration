package main

import (
	"fmt"
	"launchpad.net/gnuflag"
	"launchpad.net/juju-core/cmd"
	"launchpad.net/juju-core/environs"
	"os"
)

// BootstrapCommand is responsible for launching the first machine in a juju
// environment, and setting up everything necessary to continue working.
type BootstrapCommand struct {
	EnvName     string
	UploadTools bool
}

func (c *BootstrapCommand) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "bootstrap",
		Purpose: "start up an environment from scratch",
	}
}

func (c *BootstrapCommand) SetFlags(f *gnuflag.FlagSet) {
	addEnvironFlags(&c.EnvName, f)
	f.BoolVar(&c.UploadTools, "upload-tools", false, "upload local version of tools before bootstrapping")
}

func (c *BootstrapCommand) Init(args []string) error {
	return cmd.CheckEmpty(args)
}

// Run connects to the environment specified on the command line and bootstraps
// a juju in that environment if none already exists. If there is as yet no environments.yaml file,
// the user is informed how to create one.
func (c *BootstrapCommand) Run(context *cmd.Context) error {
	environ, err := environs.NewFromName(c.EnvName)
	if err == nil {
		return environs.Bootstrap(environ, c.UploadTools, nil)
	}
	if !os.IsNotExist(err) {
		return err
	}
	out := context.Stderr
	fmt.Fprintln(out, "No juju environment configuration file exists.")
	fmt.Fprintln(out, "Please create a configuration by running:")
	fmt.Fprintln(out, "    juju init -w")
	fmt.Fprintln(out, "then edit the file to configure your juju environment.")
	fmt.Fprintln(out, "You can then re-run bootstrap.")
	return err
}
