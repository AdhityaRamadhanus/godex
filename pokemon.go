package godex

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	Abilities Abilities
	Name      string
	ID        int
	Types     []string
}

func (p Pokemon) String() string {
	format := "#%d - %s\nType : %s\nAbilities :\n%s"

	var abilities string
	for idx, ability := range p.Abilities {
		abilities += fmt.Sprintf("%d. %v\n", (idx + 1), ability)
	}

	types := []string{}
	for _, pokemonType := range p.Types {
		types = append(types, Capitalize(pokemonType, "-"))
	}
	return fmt.Sprintf(format, p.ID, Capitalize(p.Name, "-"), strings.Join(types, " - "), abilities)
}
