package model

import (
	"strconv"
	"time"

	util "github.com/zetton110/cmkish-cli/util"
)

type Program struct {
	ID           int
	Category     string
	GameType     string
	Name         string
	NameRuby     string
	NameSub      string
	NameSubRuby  string
	EpisodeCount string
	AgeLimit     string
	StartDate    time.Time
}

func Records2Programs(records [][]string) ([]Program, error) {
	programs := []Program{}

	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		programs = append(programs, Program{
			ID:           id,
			Category:     record[1],
			GameType:     record[2],
			Name:         record[3],
			NameRuby:     record[4],
			NameSub:      record[5],
			NameSubRuby:  record[6],
			EpisodeCount: record[7],
			AgeLimit:     record[8],
			StartDate:    util.Str2time(record[9]),
		})
	}
	return programs, nil
}
