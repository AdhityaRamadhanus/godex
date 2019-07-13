package godex

import (
	"strings"
)

//Slugify return a human readable slug of a text
func Capitalize(text string, separator string) string {
	wordsCapitalized := []string{}
	for _, word := range strings.Split(text, separator) {
		wordsCapitalized = append(wordsCapitalized, strings.Title(word))
	}

	return strings.Join(wordsCapitalized, " ")
}
