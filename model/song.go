package model

import "strconv"

type Song struct {
	ID             int
	ProgramID      int
	ProgramName    string
	Category       string
	OpEd           string
	BroadcastOrder string
	Title          string
	Artist         string
}

func Records2Songs(records [][]string) ([]Song, error) {
	songs := []Song{}

	for _, record := range records {

		id, err := strconv.Atoi(record[5])
		if err != nil {
			return nil, err
		}

		programId, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		s := Song{
			ID:             id,
			ProgramID:      programId,
			Category:       record[1],
			ProgramName:    record[2],
			OpEd:           record[3],
			BroadcastOrder: record[4],
			Title:          record[6],
			Artist:         record[7],
		}

		songs = append(songs, s)
	}
	return songs, nil
}
