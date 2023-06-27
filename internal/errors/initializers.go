package errors

import "errors"

var (
	ErrLoadingDotEnvFile = errors.New("could not load the dot env file")
)
