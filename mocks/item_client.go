package mocks

import (
	"github.com/AdhityaRamadhanus/godex"
	"github.com/stretchr/testify/mock"
)

type ItemClient struct {
	mock.Mock
}

func (m ItemClient) FindOneByName(name string) (godex.Item, error) {
	args := m.Called(name)
	if args.Get(1) == nil {
		return args.Get(0).(godex.Item), nil
	}

	return args.Get(0).(godex.Item), args.Get(1).(error)
}
