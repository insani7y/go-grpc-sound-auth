package apiserver

import (
	"errors"
	"fmt"
)

var (
	NoUser = errors.New("incorrect email or bad sound")
)

func makeMissingOrIncorrectFileErr(ind int) error {
	return errors.New(fmt.Sprintf("missing or incorrect file at index %v", ind))
}