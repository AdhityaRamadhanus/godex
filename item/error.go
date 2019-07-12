package item

import (
	"github.com/pkg/errors"
)

var (
	ErrUnknownError = errors.New("Something is wrong")
	ErrItemNotFound = errors.New("Pokedex item cannot be found")
)
