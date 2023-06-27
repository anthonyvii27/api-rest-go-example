package errors

import "errors"

var (
	ErrProcessingEnvironmentIntoConfiguration = errors.New("could not process configuration from environment")
)
