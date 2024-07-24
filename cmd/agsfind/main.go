package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/zetton110/cmkish-cli/config"
	action "github.com/zetton110/cmkish-cli/pkg/action"
)

func main() {

	app := &cli.App{
		Name:    "agsfind",
		Usage:   "",
		Version: "0.0.0",
		Flags:   config.ConcatFlags(config.LogFlags),
		Commands: []*cli.Command{
			{
				Name:   "updatedb",
				Usage:  "",
				Action: action.UpdateDB,
			},
			{
				Name:    "song",
				Aliases: []string{"s"},
				Usage:   "Find song by title",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "title",
						Aliases:  []string{"t"},
						Usage:    "Title of the song",
						Required: true,
					},
				},
				Action: action.FindSongs,
			},
		},
		Before: func(c *cli.Context) error {
			logDirPath := c.String("agsfind-log-dir-path")
			if len(logDirPath) == 0 {
				return nil
			}
			logFilePath := filepath.Join(logDirPath, "agsfind.log")
			logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("cannnot open log file: %w", err)
			}
			log.SetOutput(logFile)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
