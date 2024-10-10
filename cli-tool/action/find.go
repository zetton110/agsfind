package action

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	model "github.com/zetton110/cmkish-cli/model"
	"github.com/zetton110/cmkish-cli/pkg/util"
)

type FindSong struct {
	Title        string
	ProgramTitle string
	Artist       string
	DatabasePath string
	Verbose      bool
	WordRegexp   bool
}

func (f *FindSong) Run() error {
	title := f.Title
	programTitle := f.ProgramTitle
	artist := f.Artist
	databasePath := f.DatabasePath
	verbose := f.Verbose
	wordRegexp := f.WordRegexp

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
		buildQuery("anison", title, programTitle, artist, conditions, wordRegexp),
		buildQuery("game", title, programTitle, artist, conditions, wordRegexp),
		buildQuery("side_effect", title, programTitle, artist, conditions, wordRegexp),
	}

	var songs []model.SongFindResult
	for _, q := range queries {

		rows, err := db.Query(q)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var s model.SongFindResult
			var startDateStr string
			err := rows.Scan(
				&s.Title,
				&s.Artist,
				&s.ProgramName,
				&s.OpEd,
				&s.BroadcastOrder,
				&startDateStr)
			if err != nil {
				fmt.Errorf("failed to parse song. %w\n", err)
			}
			s.Program.StartDate = util.Str2timeWithTime(startDateStr)
			songs = append(songs, s)
		}
	}

	if len(songs) == 0 {
		fmt.Println("Nothig is found.")
		return nil
	}

	data := [][]string{}
	for _, a := range songs {
		rec := []string{
			a.Title,
			a.Artist,
			a.ProgramName,
			a.OpEd + " " + a.BroadcastOrder,
		}
		if verbose {
			rec = append(rec, a.Program.StartDate.Format("2006-01-02"))
		}
		data = append(data, rec)

	}
	header := []string{"曲名", "歌手", "作品名", "備考"}
	if verbose {
		header = append(header, "放送日")
	}

	renderTable(data, header)

	fmt.Printf("%d hits.\n", len(songs))

	return nil
}

func buildQuery(table string, title string, programTitle string, artist string, conditons map[string]bool, wordRegexp bool) string {
	condition := ""
	join := fmt.Sprintf("INNER JOIN program ON %s.program_id = program.ID", table)
	order := "ORDER BY program.start_date ASC"
	columns := "title, artist, program_name, op_ed, broadcast_order, program.start_date"
	for k, v := range conditons {
		if v {
			if condition != "" {
				condition += " AND "
			}
			switch k {
			case "findByTitle":
				condition += buildWhereClause("title", title, wordRegexp)
			case "findByProgramTitle":
				condition += buildWhereClause("program_name", programTitle, wordRegexp)
			case "findByArtist":
				condition += buildWhereClause("artist", artist, wordRegexp)
			}
		}
	}
	return fmt.Sprintf(
		"SELECT %s FROM %s %s where %s %s",
		columns,
		table,
		join,
		condition,
		order,
	)
}

func buildWhereClause(tableName string, value string, wordRegexp bool) string {
	if wordRegexp {
		return fmt.Sprintf("%s = '%s'", tableName, value)
	}
	return fmt.Sprintf("%s LIKE '%%%s%%'", tableName, value)
}

func renderTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

}
