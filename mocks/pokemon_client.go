package mocks

import (
	"github.com/AdhityaRamadhanus/godex"
	"github.com/stretchr/testify/mock"
)

type PokemonClient struct {
	mock.Mock
}

func (m PokemonClient) FindOneByName(name string) (godex.Pokemon, error) {
	args := m.Called(name)
	if args.Get(1) == nil {
		return args.Get(0).(godex.Pokemon), nil
	}

	return args.Get(0).(godex.Pokemon), args.Get(1).(error)
}
