package errors

import (
	"github.com/pkg/errors"
)

// NewBadRequest creates a BadRequest error
func NewBadRequest(msg string) error {
	return customError{errorType: BadRequest, originalError: errors.New(msg)}
}

// New creates a no type error
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}
