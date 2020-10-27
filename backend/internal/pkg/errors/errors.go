package errors

import (
	"errors"

	errors2 "github.com/pkg/errors"
)

func New(reason string) error {
	return errors.New(reason)
}

func ProcessError(err error) (int, []byte) {
	return 0, []byte(``)
}

func Errorf(fmt string, val ...interface{}) error {
	return errors2.Errorf(fmt, val)
}
