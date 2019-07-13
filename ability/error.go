package ability

import (
	"github.com/pkg/errors"
)

var (
	ErrUnknownError    = errors.New("Something is wrong")
	ErrAbilityNotFound = errors.New("Pokedex ability cannot be found")
)
