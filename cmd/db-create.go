package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	scrayping "github.com/zetton110/cmkish-cli/pkg/scrayping"
)

func CreateDB(c *cli.Context) error {

	zipUrlList := scrayping.GetZipUrlList("http://anison.info/data/download.html")
	fmt.Println(zipUrlList)
	programs, err := scrayping.ExtractText(zipUrlList[0])
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range programs {
		fmt.Println(p)
	}

	return nil
}
