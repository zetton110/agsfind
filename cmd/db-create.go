package cmd

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
	scrayping "github.com/zetton110/cmkish-cli/pkg/scrayping"
)

func CreateDB(c *cli.Context) error {

	zipUrlList := scrayping.GetZipUrlList("http://anison.info/data/download.html")
	programs, err := scrayping.ExtractText(zipUrlList[0]) // program.csv
	if err != nil {
		fmt.Println(err)
	}
	db, err := setUpDB("database.sqlite")
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range programs {
		err := insertProgram(db, p)
		if err != nil {
			return err
		}
	}

	return nil
}

func setUpDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	queries := []string{
		`CREATE TABLE IF NOT EXISTS programs(
			ID INT,
			category TEXT,
			game_type TEXT,
			name TEXT,
			name_ruby TEXT,
			name_sub TEXT,
			name_sub_ruby TEXT,
			episode_count TEXT,
			age_limit TEXT,
			start_date TEXT,
			PRIMARY KEY(ID)
		)`,
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func insertProgram(db *sql.DB, p scrayping.Program) error {
	_, err := db.Exec(`
		REPLACE INTO programs(
			ID, 
			category,
			game_type,
			name,
			name_ruby,
			name_sub,
			name_sub_ruby,
			episode_count,
			age_limit,
			start_date
		) values(?,?,?,?,?,?,?,?,?,?)
	`,
		p.ID,
		p.Category,
		p.GameType,
		p.Name,
		p.NameRuby,
		p.NameSub,
		p.NameSubRuby,
		p.EpisodeCount,
		p.AgeLimit,
		p.StartDate.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}
	return nil
}
