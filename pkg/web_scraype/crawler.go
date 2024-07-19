package scraype

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
	model "github.com/zetton110/cmkish-cli/model"
	csv_access "github.com/zetton110/cmkish-cli/pkg/csv_access"
)

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

func ExtractPrograms(zipUrl string) ([]model.Program, error) {
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

			records, err := csv_access.GetRecords(f)
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}

			programs, err := records2Programs(records)
			if err != nil {
				fmt.Println("Error serialize Program:", err)
				return nil, err
			}

			return programs, nil
		}
	}
	return nil, errors.New("contents not found")
}

func records2Programs(records [][]string) ([]model.Program, error) {
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
