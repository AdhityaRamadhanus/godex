package item

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
		case "/item/ability-capsule":
			jsonResponseBytes, _ := ioutil.ReadFile("../test/json/item.json")
			res.Header().Set("Content-Type", "application/json; charset=utf-8")
			res.WriteHeader(200)
			res.Write(jsonResponseBytes)
		case "/item/ability":
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

	client := NewClient(
		server.URL,
		WithClientTimeout(5*time.Second),
	)

	for _, testCase := range testCases {
		item, err := client.FindOneByName(testCase.ItemName)
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
