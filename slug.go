package godex

import (
	"regexp"
	"strings"
)

//Slugify return a human readable slug of a text
func Slugify(text string) string {
	wordsLowerCased := []string{}
	for _, word := range strings.Split(text, " ") {
		regex, _ := regexp.Compile("[a-zA-Z0-9]+")
		matchedStrings := regex.FindAllString(word, -1)

		if len(matchedStrings) > 0 {
			wordsLowerCased = append(wordsLowerCased, strings.ToLower(strings.Join(matchedStrings, "")))
		}
	}

	return strings.Join(wordsLowerCased, "-")
}
