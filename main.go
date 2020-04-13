package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 {
		for _, a := range args {
			a = strings.ToLower(a)
			switch {
			case a == "igihe":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := getLatest("https://igihe.com/", ".homenews-title")
				if err != nil {
					fmt.Println("Please try again")
				}
				fmt.Println("Blog titles: ")
				fmt.Println(blogTitles)
			case a == "inyarwanda":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := getLatest("http://inyarwanda.com/", ".fonttitle")
				if err != nil {
					fmt.Println("Please try again")
				}
				fmt.Println("Blog titles: ")
				fmt.Println(blogTitles)
			case a == "thechronicles":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := getLatest("https://www.chronicles.rw/category/politics/", ".article-title")
				if err != nil {
					fmt.Println("Please try again")
				}
				fmt.Println("Blog titles: ")
				fmt.Println(blogTitles)
			case a == "help" || a == "--help":
				fmt.Println("eyescrape")
				fmt.Println("Get headlines from multiple sources of newswebsites from Rwanda")
				fmt.Println("Example: main.exe igihe")

			}

		}

	} else {
		fmt.Println("specify website name")
	}

}
func getLatest(url, selector string) (string, error) {
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
