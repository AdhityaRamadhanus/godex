package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AdhityaRamadhanus/godex"
	"github.com/AdhityaRamadhanus/godex/ability"
	"github.com/AdhityaRamadhanus/godex/item"
	"github.com/AdhityaRamadhanus/godex/pokemon"
	"github.com/urfave/cli"
)

var (
	apiBaseURL     = "https://pokeapi.co/api/v2"
	itemService    item.Service
	pokemonService pokemon.Service
	abilityService ability.Service
)

func init() {
	itemService = item.NewService(item.ServiceConfig{
		APIBaseURL: apiBaseURL,
	})
	pokemonService = pokemon.NewService(pokemon.ServiceConfig{
		APIBaseURL: apiBaseURL,
	})
	abilityService = ability.NewService(ability.ServiceConfig{
		APIBaseURL: apiBaseURL,
	})
}

func main() {
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
			return nil
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
			pokemonAbilityIDs := []int{}
			for _, ability := range foundPokemon.Abilities {
				pokemonAbilityIDs = append(pokemonAbilityIDs, ability.ID)
			}
			completeAbilities := abilityService.GetAbilitiesByIDs(pokemonAbilityIDs)

			// merge abilities from pokemon service and ability service
			mergedAbilities := godex.Abilities{}
			for _, foundPokemonAbility := range foundPokemon.Abilities {
				completeAbility := foundPokemonAbility
				for _, ability := range completeAbilities {
					if ability.ID == foundPokemonAbility.ID {
						completeAbility = ability
						break
					}
				}
				mergedAbilities = append(mergedAbilities, completeAbility)
			}

			foundPokemon.Abilities = mergedAbilities
			fmt.Println(foundPokemon)
		default:
			fmt.Println("Sorry, encountering problem")
		}

		return err
	}

	app.Run(os.Args)
}
