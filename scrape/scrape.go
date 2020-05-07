package scrape

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetLatest is meant for scraping headlines from news website
func GetLatest(url, selector, limit string) (string, error) {
	fmt.Println("please wait...")
	resp, err := http.Get(url)

	if err != nil {
		return "Please try again", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", err
	}

	titles := ""
	l, _ := strconv.Atoi(limit)
	total := doc.Find(selector).Length()
	headlines := doc.Find(selector)

	if l > total {
		headlines = headlines.Slice(0, total)
	} else {

		headlines = headlines.Slice(0, l)
	}

	headlines.Each(func(i int, s *goquery.Selection) {
		titles += "- " + strings.TrimSpace(s.Text()) + "\n"
	})

	return titles, nil
}
