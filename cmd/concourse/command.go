package main

import "github.com/jessevdk/go-flags"

type ConcourseCommand struct {
	Version func() `short:"v" long:"version" description:"Print the version of Concourse and exit"`

	Web    WebCommand    `command:"web"    description:"Run the web UI and build scheduler."`
	Worker WorkerCommand `command:"worker" description:"Run and register a worker."`
}

func (cmd ConcourseCommand) lessenRequirements(parser *flags.Parser) {
	cmd.Worker.lessenRequirements(parser.Find("worker"))
}
