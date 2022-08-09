package transformers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtilsGetFunctionName(t *testing.T) {
	function := getFunctionName(getFunctionName)

	assert.Equal(t, function, "getFunctionName")
}
