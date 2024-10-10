package action

import (
	"database/sql"
	"time"

	"github.com/cheggaaa/pb/v3"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/zetton110/cmkish-cli/model"
	scraper "github.com/zetton110/cmkish-cli/pkg/file"
)

type UpdateDB struct {
	DatabasePath string
}

func (u *UpdateDB) Run() error {

	db, err := setUpDB(u.DatabasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	data, err := scraper.GetData("http://anison.info/data/download.html")
	if err != nil {
		return err
	}

	bar := pb.StartNew(data.Count())

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, p := range data.Programs {
		bar.Increment()
		err := insertProgram(tx, p)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}

	for _, s := range data.Anisons {
		bar.Increment()
		err := insertAnison(tx, s)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}

	for _, s := range data.SFs {
		bar.Increment()
		err := insertSideEffect(tx, s)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Nanosecond)
	}
	for _, s := range data.Games {
		bar.Increment()
		err := insertGame(tx, s)
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

func insertAnison(tx *sql.Tx, s model.Song) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	_, err := tx.Exec(
		`REPLACE INTO anison(
			ID, 
			program_id,
			program_name,
			category,
			op_ed,
			broadcast_order,
			title,
			artist
		) values(?,?,?,?,?,?,?,?)`,
		s.ID,
		s.ProgramID,
		s.ProgramName,
		s.Category,
		s.OpEd,
		s.BroadcastOrder,
		s.Title,
		s.Artist,
	)
	if err != nil {
		return err
	}
	return nil
}

func insertGame(tx *sql.Tx, s model.Song) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	_, err := tx.Exec(
		`REPLACE INTO game(
			ID, 
			program_id,
			program_name,
			category,
			op_ed,
			broadcast_order,
			title,
			artist
		) values(?,?,?,?,?,?,?,?)`,
		s.ID,
		s.ProgramID,
		s.ProgramName,
		s.Category,
		s.OpEd,
		s.BroadcastOrder,
		s.Title,
		s.Artist,
	)
	if err != nil {
		return err
	}
	return nil
}

func insertSideEffect(tx *sql.Tx, s model.Song) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	_, err := tx.Exec(
		`REPLACE INTO side_effect(
			ID, 
			program_id,
			program_name,
			category,
			op_ed,
			broadcast_order,
			title,
			artist
		) values(?,?,?,?,?,?,?,?)`,
		s.ID,
		s.ProgramID,
		s.ProgramName,
		s.Category,
		s.OpEd,
		s.BroadcastOrder,
		s.Title,
		s.Artist,
	)
	if err != nil {
		return err
	}
	return nil
}
