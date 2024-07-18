package scraype

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/csv"
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

func GetZipUrlList(siteURL string) []string {
	doc, _ := goquery.NewDocument(siteURL)
	fmt.Println(doc)

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

func ExtractText(zipUrl string) ([]Program, error) {
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

			var fixedLines []string

			scanner := bufio.NewScanner(f)

			isHeader := true
			for scanner.Scan() {
				if isHeader {
					isHeader = false
					continue
				}
				line := scanner.Text()
				line = strings.ReplaceAll(line, "\\\"", "\"\"")
				fixedLines = append(fixedLines, line)
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
				return nil, err
			}

			fixedReader := csv.NewReader(strings.NewReader(strings.Join(fixedLines, "\n")))
			records, err := fixedReader.ReadAll()
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}

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
					StartDate:    str2time(record[9]),
				})
			}
			return programs, nil
		}
	}
	return nil, errors.New("contents not found")
}

func str2time(t string) time.Time {
	tz, _ := time.LoadLocation("Asia/Tokyo")
	timeJST, _ := time.ParseInLocation("2006-01-02", t, tz)
	return timeJST
}
