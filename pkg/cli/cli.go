package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/zetton110/cmkish-cli/action"
	"github.com/zetton110/cmkish-cli/pkg/config"
)

func NewCliApp() *cli.App {
	app := &cli.App{
		Name:    "agsf",
		Usage:   "",
		Version: "0.0.0",
		Flags: config.ConcatFlags(
			config.LogFlags,
			config.DatabaseFlags,
		),
	}

	app.Before = func(c *cli.Context) error {
		logDirPath := c.String("agsf-log-dir-path")
		if len(logDirPath) == 0 {
			return nil
		}
		logFilePath := filepath.Join(logDirPath, "agsf.log")
		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("cannnot open log file: %w", err)
		}
		log.SetOutput(logFile)
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:   "updatedb",
			Usage:  "",
			Action: updateDB,
		},
		{
			Name:    "song",
			Aliases: []string{"s"},
			Usage:   "Find song by title",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Usage:   "Find information about songs by part of its name.",
				},
				&cli.StringFlag{
					Name:    "xlookup-by-program",
					Aliases: []string{"x"},
					Usage:   "Find information about theme song by part of the program name.",
				},
				&cli.StringFlag{
					Name:    "singer",
					Aliases: []string{"s"},
					Usage:   "Find songs by artist name.",
				},
				&cli.BoolFlag{
					Name:    "verbose",
					Aliases: []string{"v"},
					Value:   false,
					Usage:   "Find information about the programs with details.",
				},
			},
			Action: findSongs,
		},
	}
	return app
}

func findSongs(c *cli.Context) error {
	title := c.String("name")
	programTitle := c.String("xlookup-by-program")
	artist := c.String("singer")
	databasePath := filepath.Join(c.String("agsf-db-base-path"), "database.sqlite")
	verbose := c.Bool("verbose")

	a := &action.FindSong{
		Title:        title,
		ProgramTitle: programTitle,
		Artist:       artist,
		DatabasePath: databasePath,
		Verbose:      verbose,
	}
	return a.Run()
}

func updateDB(c *cli.Context) error {
	databasePath := filepath.Join(c.String("agsf-db-base-path"), "database.sqlite")
	a := &action.UpdateDB{
		DatabasePath: databasePath,
	}
	return a.Run()
}
