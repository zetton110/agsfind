package action

import (
	"database/sql"
	"fmt"

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

	for _, a := range anisons {
		fmt.Println(
			"'"+a.Title+"'",
			"by "+a.Artist,
			"【"+a.ProgramName+"//"+a.OpEd+a.BroadcastOrder+"】",
		)
	}

	return nil
}
