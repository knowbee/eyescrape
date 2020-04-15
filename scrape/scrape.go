package scrape

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetLatest is meant for scraping everything
func GetLatest(url, selector string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "Please try again", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", err
	}

	titles := ""

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		titles += "- " + s.Text() + "\n"
	})

	return titles, nil
}
