package pokemon

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindOneByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch url := req.URL.String(); url {
		case "/pokemon/tapu-koko":
			jsonResponseBytes, _ := ioutil.ReadFile("../test/json/pokemon.json")
			res.Header().Set("Content-Type", "application/json; charset=utf-8")
			res.WriteHeader(200)
			res.Write(jsonResponseBytes)
		case "/pokemon/tapu":
			res.Header().Set("Content-Type", "text/plain; charset=utf-8")
			res.WriteHeader(404)
			res.Write([]byte("Not Found"))
		default:
			res.Header().Set("Content-Type", "text/plain; charset=utf-8")
			res.WriteHeader(500)
			res.Write([]byte("Internal Server Error"))
		}
	}))
	defer server.Close()

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
		server.URL,
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
