package ability

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Client interface {
	FindOneByID(id int) (godex.Ability, error)
	// allow partial result, suppress error
	FindAllByIDs(ids []int) godex.Abilities
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

type PokeAPIAbilityResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	EffectEntries []struct {
		Effect string `json:"effect"`
	} `json:"effect_entries"`
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

func (c client) FindOneByID(id int) (godex.Ability, error) {
	url := fmt.Sprintf("%s/ability/%d", c.baseURL, id)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return godex.Ability{}, err
	}

	if resp.StatusCode == 404 {
		return godex.Ability{}, ErrAbilityNotFound
	}

	if resp.StatusCode != 200 {
		return godex.Ability{}, ErrUnknownError
	}

	decodedResponse := &PokeAPIAbilityResponse{}
	json.NewDecoder(resp.Body).Decode(&decodedResponse)

	abilityEffects := []string{}
	for _, abilityEffect := range decodedResponse.EffectEntries {
		effect := strings.Replace(abilityEffect.Effect, "\n", "", -1)
		abilityEffects = append(abilityEffects, effect)
	}

	item := godex.Ability{
		Name:    decodedResponse.Name,
		ID:      decodedResponse.ID,
		Effects: abilityEffects,
	}
	return item, nil
}

func (c client) FindAllByIDs(ids []int) (abilities godex.Abilities) {
	type Result struct {
		Ability godex.Ability
		Error   error
	}
	resultChan := make(chan Result, 10)
	for _, abilityID := range ids {
		go func(id int) {
			ability, err := c.FindOneByID(id)
			resultChan <- Result{
				Ability: ability,
				Error:   err,
			}
		}(abilityID)
	}

	for range ids {
		result := <-resultChan
		if result.Error == nil {
			abilities = append(abilities, result.Ability)
		}
	}

	return abilities
}
