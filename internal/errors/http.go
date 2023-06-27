package errors

import "errors"

var (
	ErrBuildingResponse              = errors.New("could not build the request operation's response")
	ErrExecutingRequestedOperation   = errors.New("could not execute the requested operation")
	ErrBindingRequestBody            = errors.New("could not bind the given request body")
	ErrBindingRequestQueryParameters = errors.New("could not bind the given request query parameters")
	ErrHandlingRequestBody           = errors.New("could not handle the given request body")
)
