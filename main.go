package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/zetton110/cmkish-cli/cmd"
)

func main() {

	app := cli.NewApp()
	app.Name = "cmkish-cli"
	app.Usage = ""
	app.Version = "0.0.0"

	app.Commands = []cli.Command{
		{
			Name:   "create-db",
			Usage:  "",
			Action: cmd.CreateDB,
		},
	}

	app.Run(os.Args)
}
