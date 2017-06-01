package main

import (
	"errors"
	"os/user"
	"code.cloudfoundry.org/guardian/guardiancmd"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/jessevdk/go-flags"
	"github.com/tedsuo/ifrit"
)

type GardenBackend guardiancmd.ServerCommand

func (cmd WorkerCommand) lessenRequirements(command *flags.Command) {
	command.FindOptionByLongName("garden-depot").Required = false
	command.FindOptionByLongName("garden-graph").Required = false
	command.FindOptionByLongName("garden-runc-bin").Required = false
	command.FindOptionByLongName("garden-dadoo-bin").Required = false
	command.FindOptionByLongName("garden-init-bin").Required = false
	command.FindOptionByLongName("garden-nstar-bin").Required = false
	command.FindOptionByLongName("garden-tar-bin").Required = false

	command.FindOptionByLongName("baggageclaim-volumes").Required = false
	command.FindOptionByLongName("baggageclaim-driver").Default = []string{"btrfs"}
}

func (cmd *WorkerCommand) gardenRunner(logger lager.Logger, hasAssets bool) (atc.Worker, ifrit.Runner, error) {
	err := cmd.checkRoot()
	if err != nil {
		return atc.Worker{}, nil, err
	}

	return cmd.houdiniRunner(logger, "linux")
}

var ErrNotRoot = errors.New("worker must be run as root")

func (cmd *WorkerCommand) checkRoot() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	if currentUser.Uid != "0" {
		return ErrNotRoot
	}

	return nil
}