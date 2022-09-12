package utils

import (
	"github.com/stretchr/testify/assert"
	"math"
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

func TestStringHexToFloat64(t *testing.T) {
	inputString := "-0x33b1f1c23424ff41be4000"
	output := StringHexToFloat64(inputString, 18)

	assert.Equal(t, -62495535, int(math.Round(output)), "negative case")
}
