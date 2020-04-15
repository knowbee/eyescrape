package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/knowbee/eyescrape/scrape"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for _, a := range args {
			a = strings.ToLower(a)
			switch {
			case a == "igihe":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := GetLatest("https://igihe.com/", ".homenews-title")
				if err != nil {
					fmt.Println("Network error")
					os.Exit(0)
				}
				fmt.Println("Blog titles: ")
				fmt.Println(blogTitles)
			case a == "inyarwanda":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := GetLatest("http://inyarwanda.com/", ".fonttitle")
				if err != nil {
					fmt.Println("Network error")
					os.Exit(0)
				}
				fmt.Println("Blog titles: ")
				fmt.Println(blogTitles)
			case a == "thechronicles":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := GetLatest("https://www.chronicles.rw/category/politics/", ".article-title")
				if err != nil {
					fmt.Println("Network error")
					os.Exit(0)
				}
				fmt.Println("Blog titles: ")
				fmt.Println(blogTitles)
			case a == "help" || a == "--help":
				fmt.Println("eyescrape")
				fmt.Println("Get headlines from multiple sources of news websites from Rwanda")
				fmt.Println("Example: eyescrape igihe")

			}

		}

	} else {
		fmt.Println("specify website name")
	}

}
