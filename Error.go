package PointerList

//index out of range [2] with length 2

import (
	"errors"
	"fmt"
)

type Error int

const (
	IndexOutOfRange Error = iota
)

var statusText = map[Error]string{
	IndexOutOfRange: "index out of range [%d] with length %d",
}

func GetError(errType Error) error {
	return errors.New(statusText[errType])
}

func GetErrorf(errType Error, params ...any) error {
	return fmt.Errorf(statusText[errType], params...)
}

func GetErrorType(err error) Error {
	for index, status := range statusText {
		if status == err.Error() {
			return index
		}
	}

	return -1
}
