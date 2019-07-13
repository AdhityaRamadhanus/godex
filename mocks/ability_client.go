package mocks

import (
	"github.com/AdhityaRamadhanus/godex"
	"github.com/stretchr/testify/mock"
)

type AbilityClient struct {
	mock.Mock
}

func (m AbilityClient) FindOneByID(id int) (godex.Ability, error) {
	args := m.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).(godex.Ability), nil
	}

	return args.Get(0).(godex.Ability), args.Get(1).(error)
}

func (m AbilityClient) FindAllByIDs(ids []int) godex.Abilities {
	args := m.Called(ids)

	return args.Get(0).(godex.Abilities)
}
