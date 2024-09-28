package fukuro

import (
	"fmt"
	"os"

	"github.com/seipan/fukuro/internal"
	"github.com/seipan/fukuro/internal/utils"
	"github.com/urfave/cli"
)

var createCommand = cli.Command{
	Name:  "create",
	Usage: "create a container",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "bundle, b",
			Value: "",
			Usage: `path to the root of the bundle directory defaults to the current directory`,
		},
		cli.StringFlag{
			Name:  "pid-file",
			Value: "",
			Usage: "specify the file to write the process id to",
		},
		cli.BoolFlag{
			Name:  "no-pivot",
			Usage: "do not use pivot root to jail process inside rootfs.  This should be used whenever the rootfs is on top of a ramdisk",
		},
	},
	Action: func(context *cli.Context) error {
		if err := utils.CheckArgs(context, 1, internal.ExactArgs); err != nil {
			return err
		}

		status, err := startContainer(context)
		if err == nil {
			os.Exit(status)
		}
		return fmt.Errorf("runc create failed: %w", err)
	},
}
