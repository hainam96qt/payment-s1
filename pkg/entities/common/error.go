package common

import (
	error_enum "payment-s1/pkg/entities/errors"
)

type Error struct {
	error_enum.ErrorEnum
	Note string
}

func (e *Error) Error() string {
	return e.Note
}

func NewErr(code error_enum.ErrorCode, note string) *Error {
	err := &Error{
		ErrorEnum: error_enum.NewError(code),
		Note:      note,
	}
	return err
}
