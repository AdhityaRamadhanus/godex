package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AdhityaRamadhanus/godex/item"
	"github.com/urfave/cli"
)

var (
	apiBaseURL  = "https://pokeapi.co/api/v2"
	itemService item.Service
)

func init() {
	itemService = item.NewService(item.ServiceConfig{
		APIBaseURL: apiBaseURL,
	})
}

func main() {
	// Init Config
	app := cli.NewApp()
	app.Name = "godex"
	app.Author = "Adhitya Ramadhanus"
	app.Email = "adhitya.ramadhanus@gmail.com"

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Println("Please provide item/pokemon name")
			os.Exit(0)
		}

		args := strings.Join(c.Args(), " ")

		item, err := itemService.GetItemByName(args)
		if err != nil && err != item.ErrItemNotFound {
			fmt.Println("Sorry, encountering problem")
			os.Exit(1)
		}

		fmt.Println("Item", item.Name)
		fmt.Println("Cost", item.Cost)
		fmt.Println("Entries", item.Effects)
		return nil
	}
	app.Usage = "godex cli"
	app.Version = "0.0.0"

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
