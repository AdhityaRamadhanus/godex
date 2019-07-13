package pokemon

import (
	"github.com/pkg/errors"
)

var (
	ErrUnknownError    = errors.New("Something is wrong")
	ErrPokemonNotFound = errors.New("Pokedex pokemon cannot be found")
)
