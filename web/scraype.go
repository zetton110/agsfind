package web

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	csv "github.com/zetton110/cmkish-cli/file"
	model "github.com/zetton110/cmkish-cli/model"
)

type Results struct {
	Programs []model.Program
	Anisons  []model.Song
	SFs      []model.Song
	Games    []model.Song
}

func GetZipUrlList(siteURL string) []string {
	doc, _ := goquery.NewDocument(siteURL)

	zipUrlList := []string{}

	doc.Find("table td.list a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURl := getZipUrl(siteURL, href)
			zipUrlList = append(zipUrlList, zipURl)
		}
	})

	return zipUrlList
}

func getZipUrl(siteURL, zipFilePath string) string {
	u, err := url.Parse(siteURL)
	if err != nil {
		return ""
	}
	u.Path = path.Join(path.Dir(u.Path), zipFilePath)
	return u.String()
}

func Extract(targetUrlMap map[string]string) (Results, error) {
	r := Results{}
	for k, v := range targetUrlMap {
		switch k {
		case "program":
			p, err := extractPrograms(v)
			if err != nil {
				return r, err
			}
			r.Programs = append(r.Programs, p...)
		case "anison":
			s, err := extractSongs(v)
			if err != nil {
				return r, err
			}
			r.Anisons = append(r.Anisons, s...)
		case "sf":
			s, err := extractSongs(v)
			if err != nil {
				return r, err
			}
			r.SFs = append(r.SFs, s...)
		case "game":
			s, err := extractSongs(v)
			if err != nil {
				return r, err
			}
			r.Games = append(r.Games, s...)
		}
	}
	return r, nil
}

func extractPrograms(zipUrl string) ([]model.Program, error) {
	resp, err := http.Get(zipUrl)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	for _, file := range r.File {
		if path.Ext(file.Name) == ".csv" {

			f, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer f.Close()

			records, err := csv.GetRecords(f)
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}

			programs, err := toPrograms(records)
			if err != nil {
				fmt.Println("Error serialize Program:", err)
				return nil, err
			}

			return programs, nil
		}
	}
	return nil, errors.New("contents not found")
}

func toPrograms(records [][]string) ([]model.Program, error) {
	programs := []model.Program{}

	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		programs = append(programs, model.Program{
			ID:           id,
			Category:     record[1],
			GameType:     record[2],
			Name:         record[3],
			NameRuby:     record[4],
			NameSub:      record[5],
			NameSubRuby:  record[6],
			EpisodeCount: record[7],
			AgeLimit:     record[8],
			StartDate:    str2time(record[9]),
		})
	}
	return programs, nil
}

func str2time(t string) time.Time {
	tz, _ := time.LoadLocation("Asia/Tokyo")
	timeJST, _ := time.ParseInLocation("2006-01-02", t, tz)
	return timeJST
}

func extractSongs(zipUrl string) ([]model.Song, error) {
	resp, err := http.Get(zipUrl)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	for _, file := range r.File {
		if path.Ext(file.Name) == ".csv" {

			f, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer f.Close()

			records, err := csv.GetRecords(f)
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}

			anisons, err := toSongs(records)
			if err != nil {
				fmt.Println("Error serialize Program:", err)
				return nil, err
			}

			return anisons, nil
		}
	}
	return nil, errors.New("contents not found")
}

func toSongs(records [][]string) ([]model.Song, error) {
	songs := []model.Song{}

	for _, record := range records {

		id, err := strconv.Atoi(record[5])
		if err != nil {
			return nil, err
		}

		programId, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		songs = append(songs, model.Song{
			ID:             id,
			ProgramID:      programId,
			Category:       record[1],
			ProgramName:    record[2],
			OpEd:           record[3],
			BroadcastOrder: record[4],
			Title:          record[6],
			Artist:         record[7],
		})
	}
	return songs, nil
}
