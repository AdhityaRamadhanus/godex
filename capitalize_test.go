package godex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		Text                     string
		Separator                string
		ExpectedCapitalizedWords string
	}{
		{
			Text:                     "adhitya Ramadhanus",
			Separator:                "-",
			ExpectedCapitalizedWords: "Adhitya Ramadhanus",
		},
		{
			Text:                     "adhitya Ramadhanus",
			Separator:                " ",
			ExpectedCapitalizedWords: "Adhitya Ramadhanus",
		},
		{
			Text:                     "tapu-koko",
			Separator:                "-",
			ExpectedCapitalizedWords: "Tapu Koko",
		},
		{
			Text:                     "fairy",
			Separator:                "-",
			ExpectedCapitalizedWords: "Fairy",
		},
		{
			Text:                     "Tapu Koko",
			Separator:                "-",
			ExpectedCapitalizedWords: "Tapu Koko",
		},
	}

	for _, testCase := range testCases {
		capitalizedWord := Capitalize(testCase.Text, testCase.Separator)
		assert.Equal(t, capitalizedWord, testCase.ExpectedCapitalizedWords)
	}
}
