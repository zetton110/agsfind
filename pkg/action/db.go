package action

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	model "github.com/zetton110/cmkish-cli/model"
	scraype "github.com/zetton110/cmkish-cli/pkg/web"
)

func UpdateDB(c *cli.Context) error {
	zipUrlList := scraype.GetZipUrlList("http://anison.info/data/download.html")

	db, err := setUpDB("database.sqlite")
	if err != nil {
		return err
	}

	programs, err := scraype.ExtractPrograms(zipUrlList[0]) // program.csv
	if err != nil {
		return err
	}

	for _, p := range programs {
		err := insertProgram(db, p)
		if err != nil {
			return err
		}
	}

	anisons, err := scraype.ExtractAnisons(zipUrlList[1]) // anison.csv
	if err != nil {
		return err
	}

	for _, a := range anisons {
		err := insertAnison(db, a)
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
		`CREATE TABLE IF NOT EXISTS program(
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
		`CREATE TABLE IF NOT EXISTS anison(
			ID INT,
			program_id INT,
			program_name TEXT,
			category TEXT,
			op_ed TEXT,
			broadcast_order TEXT,
			title TEXT,
			artist TEXT,
			FOREIGN KEY (program_id) REFERENCES program(ID),
			PRIMARY KEY(ID, program_id)
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

func insertProgram(db *sql.DB, p model.Program) error {
	_, err := db.Exec(`
		REPLACE INTO program(
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

func insertAnison(db *sql.DB, a model.Anison) error {
	_, err := db.Exec(`
		REPLACE INTO anison(
			ID, 
			program_id,
			program_name,
			category,
			op_ed,
			broadcast_order,
			title,
			artist
		) values(?,?,?,?,?,?,?,?)
	`,
		a.ID,
		a.ProgramID,
		a.ProgramName,
		a.Category,
		a.OpEd,
		a.BroadcastOrder,
		a.Title,
		a.Artist,
	)
	if err != nil {
		return err
	}
	return nil
}
