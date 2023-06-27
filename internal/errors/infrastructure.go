package errors

import "errors"

var (
	ErrConnectingToPostgresDatabase = errors.New("could not open connection to the Postgres database")
	ErrInitializingServer           = errors.New("could not initialize the application server")
)
