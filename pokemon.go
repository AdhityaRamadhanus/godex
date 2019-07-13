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
		abilities += fmt.Sprintf("%d. %s\n", (idx + 1), ability.Name)
		abilities += fmt.Sprintf("%d\n\n", ability.ID)
	}
	return fmt.Sprintf(format, p.ID, p.Name, strings.Join(p.Types, " - "), abilities)
}
