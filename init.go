package main

import (
	"github.com/seipan/fukuro/internal"
	"github.com/urfave/cli"
)

var initCommand = cli.Command{
	Name:  "init",
	Usage: "initialize the container process",
	Flags: []cli.Flag{},
	Action: func(context *cli.Context) error {
		return internal.Init()
	},
}
