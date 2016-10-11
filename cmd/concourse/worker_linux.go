package main

import (
	"github.com/jessevdk/go-flags"
	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/guardian/guardiancmd"
	"github.com/concourse/atc"
	"github.com/tedsuo/ifrit"
)

type GardenBackend guardiancmd.GuardianCommand

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

func (cmd *WorkerCommand) gardenRunner(logger lager.Logger, args []string) (atc.Worker, ifrit.Runner, error) {
	err := cmd.checkRoot()
	if err != nil {
		return atc.Worker{}, nil, err
	}

	return cmd.houdiniRunner(logger, "linux")
}

func (cmd *WorkerCommand) baggageclaimRunner(logger lager.Logger) (ifrit.Runner, error) {
	return cmd.naiveBaggageclaimRunner(logger)
}
