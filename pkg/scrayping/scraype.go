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
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Program struct {
	ID          int
	category    string
	gameType    string
	Name        string
	NameYomi    string
	NameSub     string
	NameSubYomi string
	startDate   string
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

func ExtractText(zipUrl string) ([][]string, error) {
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
			for scanner.Scan() {
				line := scanner.Text()
				line = strings.ReplaceAll(line, "\\\"", "\"\"")
				fixedLines = append(fixedLines, line)
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
				return nil, err
			}

			fixedReader := csv.NewReader(strings.NewReader(strings.Join(fixedLines, "\n")))

			// CSVデータを読み込む
			records, err := fixedReader.ReadAll()
			if err != nil {
				fmt.Println("Error parsing CSV:", err)
				return nil, err
			}
			return records, nil
		}
	}
	return nil, errors.New("contents not found")
}
