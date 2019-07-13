package ability

import (
	"testing"

	"github.com/AdhityaRamadhanus/godex"
	"github.com/AdhityaRamadhanus/godex/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAbilityByID(t *testing.T) {
	testAbility := godex.Ability{
		Name: "Telepathy",
		ID:   140,
		Effects: []string{
			"This Pokémon does not take damage from friendly Pokémon's moves, including single-target moves aimed at it.",
		},
	}
	client := mocks.AbilityClient{}
	client.On("FindOneByID", 140).Return(testAbility, nil)
	client.On("FindOneByID", 1000000).Return(godex.Ability{}, ErrAbilityNotFound)
	client.On("FindOneByID", -1).Return(godex.Ability{}, ErrUnknownError)

	testCases := []struct {
		AbilityID           int
		ExpectedToReturnErr bool
		ExpectedErr         error
	}{
		{
			AbilityID:           140,
			ExpectedToReturnErr: false,
			ExpectedErr:         nil,
		},
		{
			AbilityID:           1000000,
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrAbilityNotFound,
		},
		{
			AbilityID:           -1,
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrUnknownError,
		},
	}

	abilityService := service{
		Client: client,
	}

	for _, testCase := range testCases {
		ability, err := abilityService.GetAbilityByID(testCase.AbilityID)
		if testCase.ExpectedToReturnErr {
			assert.NotNil(t, err)
			assert.Equal(t, testCase.ExpectedErr, err)
		} else {
			assert.NotEmpty(t, ability.Name)
			assert.NotEmpty(t, ability.ID)
			assert.NotEmpty(t, ability.Effects)
		}
	}
}

func TestGetAbilitiesByIDs(t *testing.T) {
	testAbility1 := godex.Ability{
		Name: "Telepathy",
		ID:   140,
		Effects: []string{
			"This Pokémon does not take damage from friendly Pokémon's moves, including single-target moves aimed at it.",
		},
	}
	testAbility2 := godex.Ability{
		Name: "Electric Surge",
		ID:   226,
		Effects: []string{
			"When this Pokémon enters battle, it changes the terrain to electric terrain.",
		},
	}

	client := mocks.AbilityClient{}
	client.On("FindAllByIDs", []int{140, 226}).Return(godex.Abilities{testAbility1, testAbility2})
	client.On("FindAllByIDs", []int{140, 1000000}).Return(godex.Abilities{testAbility1})
	client.On("FindAllByIDs", []int{-1, -2}).Return(godex.Abilities{})

	testCases := []struct {
		AbilityIDs                    []int
		ExpectedToReturnNoOfAbilities int
	}{
		{
			AbilityIDs:                    []int{140, 226},
			ExpectedToReturnNoOfAbilities: 2,
		},
		{
			AbilityIDs:                    []int{140, 1000000},
			ExpectedToReturnNoOfAbilities: 1,
		},
		{
			AbilityIDs:                    []int{-1, -2},
			ExpectedToReturnNoOfAbilities: 0,
		},
	}

	abilityService := service{
		Client: client,
	}

	for _, testCase := range testCases {
		abilities := abilityService.GetAbilitiesByIDs(testCase.AbilityIDs)
		assert.Equal(t, testCase.ExpectedToReturnNoOfAbilities, len(abilities))
	}
}
