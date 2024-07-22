package main

import (
	"os"

	"github.com/urfave/cli"
	action "github.com/zetton110/cmkish-cli/pkg/action"
)

func main() {

	app := cli.NewApp()
	app.Name = "agsfind"
	app.Usage = ""
	app.Version = "0.0.0"

	app.Commands = []cli.Command{
		{
			Name:   "updatedb",
			Usage:  "",
			Action: action.UpdateDB,
		},
	}

	app.Run(os.Args)
}
