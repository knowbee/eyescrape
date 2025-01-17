package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/knowbee/eyescrape/scrape"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "eyescrape"
	app.Author = "knowbee"
	app.Email = "knowbeeinc@gmail.com"
	app.Usage = "Get headlines from multiple sources of news websites from Rwanda"
	app.Commands = []cli.Command{
		{
			Name:  "igihe",
			Usage: "fetch igihe headlines and specify your limit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "5",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("https://www.igihe.com/", ".homenews-title a", c.String("limit"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
		{
			Name:  "inyarwanda",
			Usage: "fetch inyarwanda headlines and specify your limit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "5",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("http://inyarwanda.com/", ".fonttitle a", c.String("limit"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
		{
			Name:  "chronicles",
			Usage: "fetch chronicles headlines and specify your limit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "5",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("https://www.chronicles.rw/category/politics/", ".article-title a", c.String("limit"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
		{
			Name:  "umuryango",
			Usage: "fetch umuryango headlines and specify your limit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "5",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("http://umuryango.rw/amakuru/", ".artticle a", c.String("limit"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
		{
			Name:  "bwiza",
			Usage: "fetch bwiza headlines and specify your limit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "5",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("http://bwiza.com", ".media-heading a", c.String("limit"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
		{
			Name:  "kt",
			Usage: "fetch kigali today headlines and specify your limit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "5",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("https://www.kigalitoday.com/", ".headline a", c.String("limit"))
				if err != nil {
					return nil
				}
				fmt.Println(headlines)
				return nil
			},
		},
		{
			Name:  "covid",
			Usage: "fetch covid statistics for Rwanda",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "limit",
					Value: "1",
				},
			},
			Action: func(c *cli.Context) error {
				headlines, err := GetLatest("https://www.worldometers.info/coronavirus/country/rwanda/", ".news_li", c.String("limit"))
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
}
