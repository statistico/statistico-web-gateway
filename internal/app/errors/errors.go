package errors

import "errors"

var ErrorBadGateway = errors.New("error response returned from external service")
var ErrInternalServerError = errors.New("internal server error")
var ErrNotFound = errors.New("the resource requested does not exist")
