package ability

import (
	"time"

	"github.com/AdhityaRamadhanus/godex"
)

type Service interface {
	GetAbilityByID(id int) (godex.Ability, error)
	GetAbilitiesByIDs(ids []int) godex.Abilities
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

func (s service) GetAbilityByID(id int) (godex.Ability, error) {
	return s.Client.FindOneByID(id)
}

func (s service) GetAbilitiesByIDs(ids []int) godex.Abilities {
	return s.Client.FindAllByIDs(ids)
}
