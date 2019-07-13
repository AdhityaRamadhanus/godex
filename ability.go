package godex

import "fmt"

type Ability struct {
	ID      int
	Name    string
	Effects []string
}

func (a Ability) String() string {
	format := "%s\n%s"

	var effects string
	for _, effect := range a.Effects {
		effects += fmt.Sprintf("- %s\n", effect)
	}
	return fmt.Sprintf(format, a.Name, effects)
}

type Abilities []Ability
