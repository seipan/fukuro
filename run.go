package fukuro

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "run a container",
	Flags: []cli.Flag{},
	Action: func(context *cli.Context) error {
		status, err := startContainer(context)
		if err == nil {
			os.Exit(status)
		}
		return fmt.Errorf("runc run failed: %w", err)
	},
}
