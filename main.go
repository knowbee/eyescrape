package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	. "github.com/knowbee/eyescrape/mailer"
	. "github.com/knowbee/eyescrape/scrape"
)

func main() {
	args := os.Args[1:]
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
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

				Send(blogTitles)
			case a == "inyarwanda":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := GetLatest("http://inyarwanda.com/", ".fonttitle")
				if err != nil {
					fmt.Println("Network error")
					os.Exit(0)
				}

				Send(blogTitles)
			case a == "thechronicles":
				fmt.Println("Fetching headlines from ", a)
				blogTitles, err := GetLatest("https://www.chronicles.rw/category/politics/", ".article-title")
				if err != nil {
					fmt.Println("Network error")
					os.Exit(0)
				}

				Send(blogTitles)
			case a == "help" || a == "--help":
				fmt.Println("eyescrape")
				fmt.Println("Get headlines from multiple sources of news websites from Rwanda")
				fmt.Println("Example: eyescrape igihe")

			}

		}

	} else {

		fmt.Println("failed")

	}

}
