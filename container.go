package fukuro

import (
	"github.com/seipan/fukuro/internal"
	"github.com/urfave/cli"
)

func startContainer(context *cli.Context) (int, error) {
	err := internal.Create()
	if err != nil {
		return 1, err
	}
	return 0, nil
}
