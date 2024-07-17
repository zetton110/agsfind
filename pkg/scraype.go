package pkg

import (
	"fmt"
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
