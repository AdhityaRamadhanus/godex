package ability

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AdhityaRamadhanus/httpstub"
	"github.com/stretchr/testify/assert"
)

func TestFindOneById(t *testing.T) {
	srv := httpstub.Server{}
	srv.StubRequest(http.MethodGet, "/ability/140", httpstub.WithResponseBodyJSONFile("../test/json/item.json"))
	srv.StubRequest(http.MethodGet, "/ability/1000000", httpstub.WithResponseCode(http.StatusNotFound))
	srv.StubRequest(http.MethodGet, "/ability/-1", httpstub.WithResponseCode(http.StatusInternalServerError))
	srv.Start()
	defer srv.Close()

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

	client := NewClient(
		srv.URL(),
		WithClientTimeout(5*time.Second),
	)

	for _, testCase := range testCases {
		ability, err := client.FindOneByID(testCase.AbilityID)
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

func TestFindAllByIds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch url := req.URL.String(); url {
		case "/ability/140":
			jsonResponseBytes, _ := ioutil.ReadFile("../test/json/item.json")
			res.Header().Set("Content-Type", "application/json; charset=utf-8")
			res.WriteHeader(200)
			res.Write(jsonResponseBytes)
		case "/ability/226":
			jsonResponseBytes, _ := ioutil.ReadFile("../test/json/item.json")
			res.Header().Set("Content-Type", "application/json; charset=utf-8")
			res.WriteHeader(200)
			res.Write(jsonResponseBytes)
		case "/ability/1000000":
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
		AbilityIDs                    []int
		ExpectedToReturnNoOfAbilities int
	}{
		{
			AbilityIDs:                    []int{140, 226},
			ExpectedToReturnNoOfAbilities: 2,
		},
		{
			AbilityIDs:                    []int{140, -1},
			ExpectedToReturnNoOfAbilities: 1,
		},
		{
			AbilityIDs:                    []int{-1, 0},
			ExpectedToReturnNoOfAbilities: 0,
		},
	}

	client := NewClient(
		server.URL,
		WithClientTimeout(5*time.Second),
	)

	for _, testCase := range testCases {
		abilities := client.FindAllByIDs(testCase.AbilityIDs)
		assert.Equal(t, testCase.ExpectedToReturnNoOfAbilities, len(abilities))
	}
}
