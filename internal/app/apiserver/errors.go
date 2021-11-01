package apiserver

import (
	"errors"
	"fmt"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

func makeMissingOrIncorrectFileErr(ind int) error {
	return errors.New(fmt.Sprintf("missing or incorrect file at index %v", ind))
}