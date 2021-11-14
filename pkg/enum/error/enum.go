package error

type ErrorEnum struct {
	Code ErrorCode
	Msg  string
}

type ErrorCode int

const (
	Warning  ErrorCode = 12
	Internal ErrorCode = 500

	NoError         ErrorCode = 0
	Undefined       ErrorCode = 1
	NotExist        ErrorCode = 2
	ZeroSizeFile    ErrorCode = 3
	EmptyFile       ErrorCode = 4
	WrongFormat     ErrorCode = 5
	WrongOrder      ErrorCode = 6
	WrongTimeChunk  ErrorCode = 7
	MissingRequired ErrorCode = 8

	NullPointer     ErrorCode = 9
	InvalidArgument ErrorCode = 10
	NotFound        ErrorCode = 11
)

var mapErrorMsg = map[ErrorCode]string{
	Warning:         "Waning",
	NoError:         "No Error",
	Undefined:       "Undefined ",
	NotExist:        "File not exist or unreadable",
	ZeroSizeFile:    "File zero size: 0kb",
	EmptyFile:       "Empty File",
	WrongFormat:     "Wrong format",
	WrongOrder:      "Wrong order",
	WrongTimeChunk:  "Wrong time between chunk files",
	MissingRequired: "Missing Required",
	NullPointer:     "Variable must not be null",
	InvalidArgument: "Invalid Argument",
}

func NewError(code ErrorCode) ErrorEnum {
	return ErrorEnum{
		Code: code,
		Msg:  mapErrorMsg[code],
	}
}

func NewErrorEnumWithMsg(code ErrorCode, msg string) ErrorEnum {
	if msg == "" {
		msg = mapErrorMsg[code]
	}
	return ErrorEnum{
		Code: code,
		Msg:  msg,
	}
}
