package action

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	model "github.com/zetton110/cmkish-cli/model"
)

func FindSongs(c *cli.Context) error {
	title := c.String("title")

	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()

	cmd := fmt.Sprintf("SELECT title, artist, program_name, op_ed, broadcast_order FROM anison where title LIKE '%%%s%%'", title)
	rows, _ := db.Query(cmd)
	defer rows.Close()

	var anisons []model.Anison
	for rows.Next() {
		var a model.Anison
		err := rows.Scan(&a.Title, &a.Artist, &a.ProgramName, &a.OpEd, &a.BroadcastOrder)
		if err != nil {
			fmt.Errorf("failed to parse anison. %w\n", err)
		}
		anisons = append(anisons, a)
	}

	if len(anisons) == 0 {
		fmt.Println("Nothig is found.")
		return nil
	}

	data := [][]string{}
	for _, a := range anisons {
		data = append(data, []string{
			a.Title,
			a.Artist,
			a.ProgramName,
			a.OpEd + " " + a.BroadcastOrder,
		})
	}
	header := []string{"曲名", "歌手", "作品名", "備考"}

	renderTable(data, header)

	return nil
}

func renderTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

}
