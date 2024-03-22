package object_values

import "errors"

const (
	_METHOD_NOT_ALLOWED = "method not allowed"
)

func MethodNotAllowed() error {
	return errors.New(_METHOD_NOT_ALLOWED)
}
