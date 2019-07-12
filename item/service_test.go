package item

import (
	"testing"

	"github.com/AdhityaRamadhanus/godex"
	"github.com/AdhityaRamadhanus/godex/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetItemByName(t *testing.T) {
	testItem := godex.Item{
		Name: "Ability Capsule",
		Cost: 1000,
		Effects: []string{
			"Switches a Pokémon between its two possible (non-Hidden) Abilities.",
		},
	}
	client := mocks.ItemClient{}
	client.On("FindOneByName", "ability-capsule").Return(testItem, nil)
	client.On("FindOneByName", "ability").Return(godex.Item{}, ErrItemNotFound)
	client.On("FindOneByName", "s").Return(godex.Item{}, ErrUnknownError)

	testCases := []struct {
		ItemName            string
		ExpectedToReturnErr bool
		ExpectedErr         error
	}{
		{
			ItemName:            "ability-capsule",
			ExpectedToReturnErr: false,
			ExpectedErr:         nil,
		},
		{
			ItemName:            "ABility CApsule",
			ExpectedToReturnErr: false,
			ExpectedErr:         nil,
		},
		{
			ItemName:            "ability",
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrItemNotFound,
		},
		{
			ItemName:            "s",
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrUnknownError,
		},
	}

	itemService := service{
		Client: client,
	}

	for _, testCase := range testCases {
		item, err := itemService.GetItemByName(testCase.ItemName)
		if testCase.ExpectedToReturnErr {
			assert.NotNil(t, err)
			assert.Equal(t, testCase.ExpectedErr, err)
		} else {
			assert.NotEmpty(t, item.Name)
			assert.NotEmpty(t, item.Cost)
			assert.NotEmpty(t, item.Effects)
		}
	}
}
