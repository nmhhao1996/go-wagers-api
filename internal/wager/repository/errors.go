package repository

import (
	"errors"

	pkgErrors "github.com/go-errors/errors"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
)

func convErr(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		err = ErrNotFound
	}

	return pkgErrors.Wrap(err, 1)
}
