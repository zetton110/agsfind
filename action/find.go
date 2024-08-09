package action

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	model "github.com/zetton110/cmkish-cli/model"
)

type FindSong struct {
	Title        string
	ProgramTitle string
	Artist       string
	DatabasePath string
}

func (f *FindSong) Run() error {
	title := f.Title
	programTitle := f.ProgramTitle
	artist := f.Artist
	databasePath := f.DatabasePath

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	conditions := map[string]bool{
		"findByTitle":        len(title) > 0,
		"findByProgramTitle": len(programTitle) > 0,
		"findByArtist":       len(artist) > 0,
	}

	queries := []string{
		buildQuery("anison", title, programTitle, artist, conditions),
		buildQuery("game", title, programTitle, artist, conditions),
		buildQuery("side_effect", title, programTitle, artist, conditions),
	}

	var songs []model.Song
	for _, q := range queries {

		rows, err := db.Query(q)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var s model.Song
			err := rows.Scan(&s.Title, &s.Artist, &s.ProgramName, &s.OpEd, &s.BroadcastOrder)
			if err != nil {
				fmt.Errorf("failed to parse anison. %w\n", err)
			}
			songs = append(songs, s)
		}
	}

	if len(songs) == 0 {
		fmt.Println("Nothig is found.")
		return nil
	}

	data := [][]string{}
	for _, a := range songs {
		data = append(data, []string{
			a.Title,
			a.Artist,
			a.ProgramName,
			a.OpEd + " " + a.BroadcastOrder,
		})
	}
	header := []string{"曲名", "歌手", "作品名", "備考"}

	renderTable(data, header)

	fmt.Printf("%d hits.\n", len(songs))

	return nil
}

func buildQuery(table string, title string, programTitle string, artist string, conditons map[string]bool) string {
	condition := ""
	join := fmt.Sprintf("INNER JOIN program ON %s.program_id = program.ID", table)
	order := "ORDER BY program.start_date ASC"
	for k, v := range conditons {
		if v {
			if condition != "" {
				condition += " AND "
			}
			switch k {
			case "findByTitle":
				condition += fmt.Sprintf("title LIKE '%%%s%%'", title)
			case "findByProgramTitle":
				condition += fmt.Sprintf("program_name LIKE '%%%s%%'", programTitle)
			case "findByArtist":
				condition += fmt.Sprintf("artist LIKE '%%%s%%'", artist)
			}
		}
	}
	return getQuery(table, join, condition, order)
}

func getQuery(table, join, condition, order string) string {
	TARGET_COLUMNS := "title, artist, program_name, op_ed, broadcast_order"
	return fmt.Sprintf("SELECT %s FROM %s %s where %s %s", TARGET_COLUMNS, table, join, condition, order)
}

func renderTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

}
