package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	action "github.com/zetton110/cmkish-cli/action"
	"github.com/zetton110/cmkish-cli/config"
)

func main() {

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
					Name:    "title",
					Aliases: []string{"t"},
					Usage:   "Title of the song",
				},
				&cli.StringFlag{
					Name:    "program-title",
					Aliases: []string{"p"},
					Usage:   "Title of the program",
				},
				&cli.StringFlag{
					Name:    "artist",
					Aliases: []string{"a"},
					Usage:   "artist of the song",
				},
			},
			Action: findSongs,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func findSongs(c *cli.Context) error {
	title := c.String("title")
	programTitle := c.String("program-title")
	artist := c.String("artist")
	databasePath := filepath.Join(c.String("agsf-db-base-path"), "database.sqlite")

	a := &action.FindSong{
		Title:        title,
		ProgramTitle: programTitle,
		Artist:       artist,
		DatabasePath: databasePath,
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
