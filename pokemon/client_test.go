package pokemon

import (
	"net/http"
	"testing"
	"time"

	"github.com/AdhityaRamadhanus/httpstub"
	"github.com/stretchr/testify/assert"
)

func TestFindOneByName(t *testing.T) {
	srv := httpstub.Server{}
	srv.StubRequest(http.MethodGet, "/pokemon/tapu-koko", httpstub.WithResponseBodyJSONFile("../test/json/pokemon.json"))
	srv.StubRequest(http.MethodGet, "/pokemon/tapu", httpstub.WithResponseCode(http.StatusNotFound))
	srv.StubRequest(http.MethodGet, "/pokemon/s", httpstub.WithResponseCode(http.StatusInternalServerError))
	srv.Start()
	defer srv.Close()

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
			PokemonName:         "tapu",
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrPokemonNotFound,
		},
		{
			PokemonName:         "s",
			ExpectedToReturnErr: true,
			ExpectedErr:         ErrUnknownError,
		},
	}

	client := NewClient(
		srv.URL(),
		WithClientTimeout(5*time.Second),
	)

	for _, testCase := range testCases {
		pokemon, err := client.FindOneByName(testCase.PokemonName)
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
