package main

import (
	"github.com/seipan/fukuro/internal"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "run a container",
	Flags: []cli.Flag{},
	Action: func(context *cli.Context) error {
		internal.Run()
		return nil
	},
}
