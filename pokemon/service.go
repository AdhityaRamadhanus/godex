package pokemon

import (
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Service interface {
	GetPokemonByName(name string) (godex.Pokemon, error)
}

type ServiceConfig struct {
	APIBaseURL string
}

type service struct {
	Client Client
}

func NewService(config ServiceConfig) Service {
	return &service{
		Client: NewClient(
			config.APIBaseURL,
			WithClientTimeout(5*time.Second),
		),
	}
}

func (s service) GetPokemonByName(name string) (godex.Pokemon, error) {
	sluggedName := godex.Slugify(name)
	return s.Client.FindOneByName(sluggedName)
}
