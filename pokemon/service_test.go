package pokemon

import (
	"testing"

	"github.com/AdhityaRamadhanus/godex"
	"github.com/AdhityaRamadhanus/godex/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemonByName(t *testing.T) {
	testPokemon := godex.Pokemon{
		Name:  "tapu-koko",
		Types: []string{"fairy", "electric"},
		ID:    785,
		Abilities: []godex.Ability{
			{
				Name: "telepathy",
				ID:   140,
			},
			{
				Name: "electric-surge",
				ID:   226,
			},
		},
	}
	client := mocks.PokemonClient{}
	client.On("FindOneByName", "tapu-koko").Return(testPokemon, nil)
	client.On("FindOneByName", "tapu").Return(godex.Pokemon{}, ErrPokemonNotFound)
	client.On("FindOneByName", "s").Return(godex.Pokemon{}, ErrUnknownError)

	testCases := []struct {
		PokemonName         string
		ExpectedToReturnErr bool
		ExpectedErr         error
	}{
		{
			PokemonName:         "tapu-koko",
			ExpectedToReturnErr: false,
			ExpectedErr:         nil,
		},
		{
			PokemonName:         "TAPU KoKo",
			ExpectedToReturnErr: false,
			ExpectedErr:         nil,
		},
		{
			PokemonName:         "Tapu",
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrPokemonNotFound,
		},
		{
			PokemonName:         "s",
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrUnknownError,
		},
	}

	pokemonService := service{
		Client: client,
	}

	for _, testCase := range testCases {
		pokemon, err := pokemonService.GetPokemonByName(testCase.PokemonName)
		if testCase.ExpectedToReturnErr {
			assert.NotNil(t, err)
			assert.Equal(t, testCase.ExpectedErr, err)
		} else {
			assert.NotEmpty(t, pokemon.Name)
			assert.NotEmpty(t, pokemon.ID)
			assert.NotEmpty(t, pokemon.Types)
			assert.NotEmpty(t, pokemon.Abilities)
		}
	}
}
