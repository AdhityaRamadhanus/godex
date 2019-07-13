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
		abilities += fmt.Sprintf("%d. %v\n\n", (idx + 1), ability)
	}
	return fmt.Sprintf(format, p.ID, p.Name, strings.Join(p.Types, " - "), abilities)
}
