package fukuro

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var createCommand = cli.Command{
	Name:  "create",
	Usage: "create a container",
	Flags: []cli.Flag{},
	Action: func(context *cli.Context) error {
		status, err := startContainer(context)
		if err == nil {
			os.Exit(status)
		}
		return fmt.Errorf("runc create failed: %w", err)
	},
}
