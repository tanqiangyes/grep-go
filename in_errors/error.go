package in_errors

import (
	"errors"
)

var (
	// ErrArgs just means that args num less than 1
	ErrArgs = errors.New("error: too few arguments")

	// ErrCreateRegexpFinder means that regexp is not valid
	ErrCreateRegexpFinder = errors.New("create regexp finder failed, please check your regexp")

	// ErrNoFilesFound is returned when the files pattern passed to the application
	// doesn't match any existing files.
	ErrNoFilesFound = errors.New("no files found")
)
