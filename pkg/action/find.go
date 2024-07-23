package action

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func FindSongs(c *cli.Context) error {
	title := c.String("title")
	fmt.Printf("Finding song with title: %s\n", title)

	return nil
}
