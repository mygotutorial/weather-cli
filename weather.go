package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/gotutorial/weather-cli/services"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp() // := is the Short variable declarations
	app.Name = "Weather CLI Tool"
	app.Usage = "Prompt the current weather condition for a city, zipcode or georaphical coordination"
	app.Version = "1.0.0"
	app.HideHelp = false

	var city string
	var country string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "city, ct",
			Usage:       "The city name you are looking for its current weather condition (defaut is 'Atlanta')",
			Destination: &city,
		},
		cli.StringFlag{
			Name:        "country, co",
			Usage:       "The country name you are looking for its current weather condition (default is empty)",
			Destination: &country,
		},
		cli.BoolFlag{
			Name:  "temp, tm",
			Usage: "The location current temperature",
		},
		cli.BoolFlag{
			Name:  "desc, ds",
			Usage: "The location current description",
		},
		cli.BoolFlag{
			Name:  "hum, hm",
			Usage: "The location current humidity",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(city) > 0 {
			var reponse, err = services.CurrentWeatherByCity(city, country, true, true, true)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(reponse)
			}
		} else {
			fmt.Println("City is required. Pass the city value using --city")
			cli.ShowAppHelp(c)
		}

		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
