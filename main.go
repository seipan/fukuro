package fukuro

import (
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "fukuro"
	app.Usage = "fukuro is a simple container runtime"
	app.Commands = []cli.Command{
		createCommand,
	}
}
