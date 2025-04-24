package errors

import (
	"fmt"
)

type BodyError struct {
	message interface{}
}

func NewBodyError(msg interface{}) *BodyError {
	return &BodyError{
		message: msg,
	}
}

func (e *BodyError) Error() string {
	return fmt.Sprintf("invalid body request: %v", e.message)
}
