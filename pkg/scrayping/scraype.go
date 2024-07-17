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

func ExtractText(zipUrl string) (string, error) {
	resp, err := http.Get(zipUrl)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return "", err
	}

	for _, file := range r.File {
		if path.Ext(file.Name) == ".csv" {
			f, err := file.Open()
			if err != nil {
				return "", err
			}
			b, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				return "", err
			}
			return string(b), nil
		}
	}

	return "", errors.New("contents not found")
}
