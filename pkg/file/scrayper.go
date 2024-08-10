package file

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/zetton110/cmkish-cli/model"
)

type ExtractedData struct {
	Programs []model.Program
	Anisons  []model.Song
	SFs      []model.Song
	Games    []model.Song
}

func (e *ExtractedData) Count() int {
	return len(e.Programs) + len(e.Anisons) + len(e.SFs) + len(e.Games)
}

func GetData(siteURL string) (*ExtractedData, error) {
	zipUrls := getZipUrls(siteURL)
	if len(zipUrls) != 4 {
		return nil, errors.New("zip urls not found")
	}

	programUrl := zipUrls[0]
	anisonUrl := zipUrls[1]
	sfUrl := zipUrls[2]
	gameUrl := zipUrls[3]

	programs, err := extractPrograms(programUrl)
	anisons, err := extractSongs(anisonUrl)
	sfs, err := extractSongs(sfUrl)
	games, err := extractSongs(gameUrl)

	if err != nil {
		return nil, err
	}

	return &ExtractedData{
		Programs: programs,
		Anisons:  anisons,
		SFs:      sfs,
		Games:    games,
	}, nil

}

func getZipUrls(siteURL string) []string {
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

			records, err := CSV2Records(f)
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}
			return model.Records2Programs(records)
		}
	}
	return nil, errors.New("contents not found")
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

			records, err := CSV2Records(f)
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}
			return model.Records2Songs(records)
		}
	}
	return nil, errors.New("contents not found")
}
