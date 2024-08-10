package main

import (
	"log"
	"os"

	"github.com/zetton110/cmkish-cli/pkg/cli"
)

func main() {
	err := cli.NewCliApp().Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
