package errors

import "errors"

var (
	ErrTodoNotFound = errors.New("could not find the requested todo")
)
