package crud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	foo string
	bar string
	baz int
}

func TestRemoveDuplicatePrimaryKeys(t *testing.T) {
	var primaryKeys = []string{"foo", "baz"}
	var models = []*TestModel{
		{foo: "foo1", bar: "bar1", baz: 1},
		{foo: "foo2", bar: "bar2", baz: 1},
		{foo: "foo3", bar: "bar2", baz: 1},
		{foo: "foo3", bar: "bar2", baz: 2},
		{foo: "foo3", bar: "bar2", baz: 2},
	}

	output := removeDuplicatePrimaryKeys[TestModel](models, primaryKeys)

	assert.Equal(t, len(output), 4)
}
