package fileroller

import "errors"

var (
	// ErrInvalidFile is the error returned when the file specified
	// is not valid.
	ErrInvalidFile = errors.New("invalid file")
)
