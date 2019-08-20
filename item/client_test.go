package item

import (
	"net/http"
	"testing"
	"time"

	"github.com/AdhityaRamadhanus/httpstub"
	"github.com/stretchr/testify/assert"
)

func TestFindOneByName(t *testing.T) {
	srv := httpstub.Server{}
	srv.StubRequest(http.MethodGet, "/item/ability-capsule", httpstub.WithResponseBodyJSONFile("../test/json/item.json"))
	srv.StubRequest(http.MethodGet, "/item/ability", httpstub.WithResponseCode(http.StatusNotFound))
	srv.StubRequest(http.MethodGet, "/item/s", httpstub.WithResponseCode(http.StatusInternalServerError))
	srv.Start()
	defer srv.Close()

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
		srv.URL(),
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
