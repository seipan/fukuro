package fukuro

import (
	"github.com/seipan/fukuro/internal"
	"github.com/seipan/fukuro/internal/utils"
	"github.com/urfave/cli"
)

func startContainer(context *cli.Context) (int, error) {
	if err := utils.RevisePidFile(context); err != nil {
		return -1, err
	}

	err := internal.Create()
	if err != nil {
		return 1, err
	}
	return 0, nil
}

func runContaier(context *cli.Context) (int, error) {
	internal.Run()
	return 0, nil
}
