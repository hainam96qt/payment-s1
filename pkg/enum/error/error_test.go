package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateError(t *testing.T) {
	testcase := []struct {
		value    ErrorCode
		expected string
	}{
		{
			value:    WrongOrder,
			expected: mapErrorMsg[WrongOrder],
		},
		{
			value:    WrongFormat,
			expected: mapErrorMsg[WrongFormat],
		},
		{
			value:    Undefined,
			expected: mapErrorMsg[Undefined],
		},
	}
	for _, v := range testcase {
		err := NewError(v.value)
		assert.Equal(t, mapErrorMsg[v.value], v.expected)

		err = NewErrorEnumWithMsg(v.value, "This is test new msg")
		assert.Equal(t, "This is test new msg", err.Msg)
		assert.NotNil(t, testcase)
	}
}
