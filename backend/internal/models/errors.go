package models

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrInvalidReference = errors.New("invalid foreign key reference")
	ErrInvalidData      = errors.New("invalid data")
)

func ClassifyDBError(err error) error {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	switch pqErr.Code {
	case "23503":
		return ErrInvalidReference
	case "23502", "23514":
		return ErrInvalidData
	default:
		return err
	}
}
