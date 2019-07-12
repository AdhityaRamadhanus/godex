package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type PokeAPIItemResponse struct {
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

func WithClientTimeout(timeout time.Duration) func(client *Client) {
	return func(client *Client) {
		client.httpClient.Timeout = timeout
	}
}

func NewClient(baseURL string, options ...func(*Client)) *Client {
	client := &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
	for _, option := range options {
		option(client)
	}

	return client
}

func (c *Client) FindOneByName(name string) (godex.Item, error) {
	url := fmt.Sprintf("%s/item/%s", c.baseURL, name)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return godex.Item{}, err
	}

	if resp.StatusCode == 404 {
		return godex.Item{}, ErrItemNotFound
	}

	if resp.StatusCode != 200 {
		return godex.Item{}, err
	}

	decodedResponse := &PokeAPIItemResponse{}
	json.NewDecoder(resp.Body).Decode(&decodedResponse)

	itemEnglishName := decodedResponse.Name
	for _, name := range decodedResponse.Names {
		if name.Language.Name == "en" {
			itemEnglishName = name.Name
			break
		}
	}
	itemEffects := []string{}
	for _, itemEffect := range decodedResponse.EffectEntries {
		itemEffects = append(itemEffects, itemEffect.Effect)
	}
	item := godex.Item{
		Name:    itemEnglishName,
		Cost:    decodedResponse.Cost,
		Effects: itemEffects,
	}
	return item, nil
}
