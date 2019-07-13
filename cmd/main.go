package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AdhityaRamadhanus/godex/item"
	"github.com/AdhityaRamadhanus/godex/pokemon"
	"github.com/urfave/cli"
)

var (
	apiBaseURL     = "https://pokeapi.co/api/v2"
	itemService    item.Service
	pokemonService pokemon.Service
)

func init() {
	itemService = item.NewService(item.ServiceConfig{
		APIBaseURL: apiBaseURL,
	})
	pokemonService = pokemon.NewService(pokemon.ServiceConfig{
		APIBaseURL: apiBaseURL,
	})
}

func main() {
	// Init Config
	app := cli.NewApp()
	app.Name = "godex"
	app.Author = "Adhitya Ramadhanus"
	app.Email = "adhitya.ramadhanus@gmail.com"
	app.Usage = "godex [pokemon/item name]"
	app.Version = "0.0.0"

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Println("Please provide item/pokemon name")
			os.Exit(0)
		}

		args := strings.Join(c.Args(), " ")

		foundItem, err := itemService.GetItemByName(args)
		if err != nil && err != item.ErrItemNotFound {
			fmt.Println("Sorry, encountering problem")
			return err
		}

		if err == nil {
			fmt.Println(foundItem)
			return nil
		}

		// can't find item, move to search pokemon
		foundPokemon, err := pokemonService.GetPokemonByName(args)

		switch err {
		case pokemon.ErrPokemonNotFound:
			err = nil
			fmt.Println("Sorry, we could not find any item or pokemon with that name")
		case nil:
			fmt.Println(foundPokemon)
		default:
			fmt.Println("Sorry, encountering problem")
		}

		return err
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
