package model

import (
	"errors"

	"github.com/wuhc1010/go-zero/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
