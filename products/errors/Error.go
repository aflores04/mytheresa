package errors

import (
	"errors"
	"fmt"
)

func NewError(message string, err error) error {
	if err != nil {
		return errors.New(fmt.Sprintf("%s error: %s", message, err.Error()))
	}

	return errors.New(message)
}
