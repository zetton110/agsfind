package main

import (
	"os"

	"github.com/urfave/cli"
	action "github.com/zetton110/cmkish-cli/pkg/cmd_action"
)

func main() {

	app := cli.NewApp()
	app.Name = "cmkish-cli"
	app.Usage = ""
	app.Version = "0.0.0"

	app.Commands = []cli.Command{
		{
			Name:   "makedb",
			Usage:  "",
			Action: action.MakeDB,
		},
		{
			Name:   "updatedb",
			Usage:  "",
			Action: action.UpdateDB,
		},
	}

	app.Run(os.Args)
}
