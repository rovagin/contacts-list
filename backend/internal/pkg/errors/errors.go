package errors

import (
	"errors"
)

func New(reason string) error {
	return errors.New(reason)
}

func ProcessError(err error) (int, []byte) {
	return 0, []byte(``)
}
