package errors

import "errors"

var (
	ErrGeneratingMigrate = errors.New("could not generate the migrations")
)
