package item

import (
	"fmt"
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Service interface {
	GetItemByName(name string) (godex.Item, error)
}

type ServiceConfig struct {
	APIBaseURL string
}

type service struct {
	client *Client
}

func NewService(config ServiceConfig) Service {
	return &service{
		client: NewClient(
			config.APIBaseURL,
			WithClientTimeout(5*time.Second),
		),
	}
}

func (s *service) GetItemByName(name string) (godex.Item, error) {
	sluggedName := godex.Slugify(name)
	fmt.Println("Item Name ", sluggedName)
	return s.client.FindOneByName(sluggedName)
}
