package godex

import (
	"strings"
)

//Capitalize return string capitalized with specific separator
func Capitalize(text string, separator string) string {
	wordsCapitalized := []string{}
	for _, word := range strings.Split(text, separator) {
		wordsCapitalized = append(wordsCapitalized, strings.Title(word))
	}

	return strings.Join(wordsCapitalized, " ")
}
