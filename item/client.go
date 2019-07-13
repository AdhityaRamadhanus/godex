package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Client interface {
	FindOneByName(name string) (godex.Item, error)
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

type pokeAPIItemResponse struct {
	Name  string `json:"name"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
		}
	} `json:"names"`
	Cost          uint `json:"cost"`
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

func (c client) FindOneByName(name string) (godex.Item, error) {
	url := fmt.Sprintf("%s/item/%s", c.baseURL, name)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return godex.Item{}, err
	}

	if resp.StatusCode == 404 {
		return godex.Item{}, ErrItemNotFound
	}

	if resp.StatusCode != 200 {
		return godex.Item{}, ErrUnknownError
	}

	decodedResponse := &pokeAPIItemResponse{}
	json.NewDecoder(resp.Body).Decode(&decodedResponse)

	itemEffects := []string{}
	for _, itemEffect := range decodedResponse.EffectEntries {
		effect := strings.Replace(itemEffect.Effect, "\n", "", -1)
		itemEffects = append(itemEffects, effect)
	}

	item := godex.Item{
		Name:    decodedResponse.Name,
		Cost:    decodedResponse.Cost,
		Effects: itemEffects,
	}
	return item, nil
}
