package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	. "github.com/knowbee/eyescrape/scrape"
	"github.com/urfave/cli"
)

func main() {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatalf("Error loading .env file")
	}

	app := cli.NewApp()
	app.Name = "eyescrape"
	app.Author = "knowbee"
	app.Email = "knowbeeinc@gmail.com"
	app.Usage = "Get headlines from multiple sources of news websites from Rwanda"

	app.Commands = []cli.Command{
		{
			Name:  "fetch",
			Usage: "fetching latest headlines",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "https://igihe.com",
				},
				cli.StringFlag{
					Name:  "selector",
					Value: ".homenews-title",
				},
			},
			Action: func(c *cli.Context) error {
				// switch c.FlagNames()[0]
				headlines, err := GetLatest(c.String("host"), c.String("selector"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// for _, a := range args {
	// 	a = strings.ToLower(a)
	// 	switch {
	// 	case a == "igihe":
	// 		fmt.Println("Fetching headlines from ", a)
	// 		blogTitles, err := GetLatest("https://igihe.com/", ".homenews-title")
	// 		if err != nil {
	// 			fmt.Println("Network error")
	// 			os.Exit(0)
	// 		}
	// 		fmt.Println("Blog titles: ")
	// 		print(blogTitles)
	// 		// Send(blogTitles, a)
	// 	case a == "inyarwanda":
	// 		fmt.Println("Fetching headlines from ", a)
	// 		blogTitles, err := GetLatest("http://inyarwanda.com/", ".fonttitle")
	// 		if err != nil {
	// 			fmt.Println("Network error")
	// 			os.Exit(0)
	// 		}
	// 		fmt.Println("Blog titles: ")
	// 		print(blogTitles)
	// 		// Send(blogTitles, a)
	// 	case a == "thechronicles":
	// 		fmt.Println("Fetching headlines from ", a)
	// 		blogTitles, err := GetLatest("https://www.chronicles.rw/category/politics/", ".article-title")
	// 		if err != nil {
	// 			fmt.Println("Network error")
	// 			os.Exit(0)
	// 		}
	// 		fmt.Println("Blog titles: ")
	// 		print(blogTitles)
	// 		// Send(blogTitles, a)
	// 	case a == "help" || a == "--help":
	// 		fmt.Println("eyescrape")
	// 		fmt.Println("Get headlines from multiple sources of news websites from Rwanda")
	// 		fmt.Println("Example: eyescrape igihe")

	// 	}

	// }

}
