package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/zetton110/cmkish-cli/pkg"
)

func CreateDB(c *cli.Context) error {

	zipUrlList := pkg.GetZipUrlList("http://anison.info/data/download.html")
	fmt.Println(zipUrlList)

	return nil
}
