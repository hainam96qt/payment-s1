package errors

type ErrorEnum struct {
	Error   ErrorCode
	Message string
}

type ErrorCode int

const (
	Warning  ErrorCode = 12
	Internal ErrorCode = 500

	Undefined       ErrorCode = 1
	MissingRequired ErrorCode = 2

	NullPointer      ErrorCode = 3
	InvalidArgument  ErrorCode = 4
	NotFound         ErrorCode = 404
	ErrorDescription ErrorCode = 400
)

var mapErrorMsg = map[ErrorCode]string{
	Warning:          "Waning",
	Undefined:        "Undefined ",
	MissingRequired:  "Missing Required",
	NullPointer:      "Variable must not be null",
	InvalidArgument:  "Invalid Argument",
	ErrorDescription: "ERROR_DESCRIPTION",
}

func NewError(code ErrorCode) ErrorEnum {
	return ErrorEnum{
		Error:   code,
		Message: mapErrorMsg[code],
	}
}

func NewErrorEnumWithMsg(code ErrorCode, msg string) ErrorEnum {
	if msg == "" {
		msg = mapErrorMsg[code]
	}
	return ErrorEnum{
		Error:   code,
		Message: msg,
	}
}

type Error struct {
	ErrorEnum
	Note string
}

func (e *Error) Error() string {
	return e.Note
}

func NewErr(code ErrorCode, note string) *Error {
	err := &Error{
		ErrorEnum: NewError(code),
		Note:      note,
	}
	return err
}
