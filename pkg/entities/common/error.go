package common

import (
	error_enum "challenge/pkg/enum/error"
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
