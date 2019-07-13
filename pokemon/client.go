package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Client interface {
	FindOneByName(name string) (godex.Pokemon, error)
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

type pokeAPIPokemonResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		}
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	} `json:"abilities"`
}

func WithClientTimeout(timeout time.Duration) func(client *client) {
	return func(client *client) {
		client.httpClient.Timeout = timeout
	}
}

func NewClient(baseURL string, options ...func(*client)) Client {
	client := &client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
	for _, option := range options {
		option(client)
	}

	return client
}

func (c client) FindOneByName(name string) (godex.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", c.baseURL, name)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return godex.Pokemon{}, err
	}

	if resp.StatusCode == 404 {
		return godex.Pokemon{}, ErrPokemonNotFound
	}

	if resp.StatusCode != 200 {
		return godex.Pokemon{}, ErrUnknownError
	}

	decodedResponse := &pokeAPIPokemonResponse{}
	json.NewDecoder(resp.Body).Decode(&decodedResponse)

	pokemonTypes := []string{}
	for _, pokemonType := range decodedResponse.Types {
		pokemonTypes = append(pokemonTypes, pokemonType.Type.Name)
	}

	pokemonAbilities := []godex.Ability{}
	for _, pokemonAbility := range decodedResponse.Abilities {
		var pokemonAbilityID int
		regex := regexp.MustCompile(`(?m)(?:\/ability\/)([\d]+)`)
		matchedString := regex.FindStringSubmatch(pokemonAbility.Ability.URL)

		if len(matchedString) > 0 {
			id, _ := strconv.Atoi(matchedString[1])
			pokemonAbilityID = id
		}
		pokemonAbilities = append(pokemonAbilities, godex.Ability{
			ID:   pokemonAbilityID,
			Name: pokemonAbility.Ability.Name,
		})
	}

	pokemon := godex.Pokemon{
		ID:        decodedResponse.ID,
		Name:      decodedResponse.Name,
		Types:     pokemonTypes,
		Abilities: pokemonAbilities,
	}
	return pokemon, nil
}
