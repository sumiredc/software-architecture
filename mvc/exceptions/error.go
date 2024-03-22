package exceptions

import "errors"

const (
	method_not_allowed = "method not allowed"
)

func MethodNotAllowed() error {
	return errors.New(method_not_allowed)
}
