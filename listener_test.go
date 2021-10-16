package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListener_Bind(t *testing.T) {
	var res []interface{}
	var original Listener = func(argv ...interface{}) {
		res = argv
	}
	bounded := original.Bind(1, 2, 3)

	original(1, 2, 3)
	assert.Equal(t, []interface{}{1, 2, 3}, res, "Original array")

	bounded(2, 3)
	assert.Equal(t, []interface{}{1, 2, 3, 2, 3}, res, "Bounded array")
}
