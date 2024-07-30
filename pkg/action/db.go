package action

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb/v3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	model "github.com/zetton110/cmkish-cli/model"
	scraype "github.com/zetton110/cmkish-cli/pkg/web"
)

func UpdateDB(c *cli.Context) error {
	zipUrlList := scraype.GetZipUrlList("http://anison.info/data/download.html")
	databasePath := c.String("agsf-db-base-path")

	db, err := setUpDB(filepath.Join(databasePath, "database.sqlite"))
	if err != nil {
		return err
	}

	programs, err := scraype.ExtractPrograms(zipUrlList[0]) // program.csv
	if err != nil {
		return err
	}

	animeSongs, err := scraype.ExtractSongs(zipUrlList[1]) // anison.csv
	if err != nil {
		return err
	}

	sfSongs, err := scraype.ExtractSongs(zipUrlList[2]) // sf.csv
	if err != nil {
		return err
	}

	gameSongs, err := scraype.ExtractSongs(zipUrlList[3]) // game.csv
	if err != nil {
		return err
	}

	count := len(programs) + len(animeSongs) + len(sfSongs) + len(gameSongs)
	bar := pb.StartNew(count)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, p := range programs {
		bar.Increment()
		err := insertProgram(tx, p)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}

	for _, s := range animeSongs {
		bar.Increment()
		err := insertSongTo(tx, s, "anison")
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}

	for _, s := range sfSongs {
		bar.Increment()
		err := insertSongTo(tx, s, "side_effect")
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}
	for _, s := range gameSongs {
		bar.Increment()
		err := insertSongTo(tx, s, "game")
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}

	tx.Commit()
	bar.Finish()

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
		`CREATE TABLE IF NOT EXISTS game(
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
		`CREATE TABLE IF NOT EXISTS side_effect(
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

func insertProgram(tx *sql.Tx, p model.Program) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	_, err := tx.Exec(`
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

func insertSongTo(tx *sql.Tx, a model.Song, tableName string) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	_, err := tx.Exec(
		fmt.Sprintf(`
			REPLACE INTO %s(
				ID, 
				program_id,
				program_name,
				category,
				op_ed,
				broadcast_order,
				title,
				artist
			) values(?,?,?,?,?,?,?,?)
	`, tableName),
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
