package model

import "time"

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
