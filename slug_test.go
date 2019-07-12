package godex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	testCases := []struct {
		Text         string
		ExpectedSlug string
	}{
		{
			"!@#!@!@!@!@ Adhitya Ramadhanus",
			"adhitya-ramadhanus",
		},
		{
			"How are you?",
			"how-are-you",
		},
		{
			"       ",
			"",
		},
		{
			" How are!@!*@&!*@1923123!*@!@() ",
			"how-are1923123",
		},
		{
			"Ability Capsule",
			"ability-capsule",
		},
	}

	for _, testCase := range testCases {
		slug := Slugify(testCase.Text)
		assert.Equal(t, slug, testCase.ExpectedSlug, "Incorrect Slug")
	}
}
