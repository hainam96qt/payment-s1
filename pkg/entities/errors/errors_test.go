package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateError(t *testing.T) {
	testcase := []struct {
		value    ErrorCode
		expected string
	}{
		{
			value:    Undefined,
			expected: mapErrorMsg[Undefined],
		},
	}
	for _, v := range testcase {
		err := NewError(v.value)
		assert.Equal(t, mapErrorMsg[v.value], v.expected)

		err = NewErrorEnumWithMsg(v.value, "This is test new msg")
		assert.Equal(t, "This is test new msg", err.Message)
		assert.NotNil(t, testcase)
	}
}
