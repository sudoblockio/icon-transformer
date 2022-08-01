package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringHexToInt64(t *testing.T) {
	inputString := "0x1234567890"
	output := StringHexToInt64(inputString)

	assert.Equal(t, int64(78187493520), output, "Test success")

	inputString = "0xXX"
	output = StringHexToInt64(inputString)

	assert.Equal(t, int64(0), output, "Test failure")
}
