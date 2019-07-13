package godex

import (
	"fmt"
)

//Item domain entity representing https://pokeapi.co/docs/v2.html/#items-section
type Item struct {
	Effects []string
	Name    string
	Cost    uint
}

func (i Item) String() string {
	format := "Item : %s\nCost : %d\nEntries : \n%s"

	var effects string
	for idx, effect := range i.Effects {
		effects += fmt.Sprintf("%d. %s\n", (idx + 1), effect)
	}
	return fmt.Sprintf(format, Capitalize(i.Name, "-"), i.Cost, effects)
}
