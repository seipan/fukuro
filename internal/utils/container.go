package utils

import (
	"path/filepath"

	"github.com/urfave/cli"
)

func RevisePidFile(context *cli.Context) error {
	pidFile := context.String("pid-file")
	if pidFile == "" {
		return nil
	}

	pidFile, err := filepath.Abs(pidFile)
	if err != nil {
		return err
	}
	return context.Set("pid-file", pidFile)
}
