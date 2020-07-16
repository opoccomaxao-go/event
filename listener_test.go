package event

import (
	. "github.com/opoccomaxao-go/helpers/test"
	"testing"
)

func TestListener_Bind(t *testing.T) {
	var res []interface{}
	var original Listener = func(argv ...interface{}) {
		res = argv
	}
	bounded := original.Bind(1, 2, 3)

	original(1, 2, 3)
	CheckValue(t, "Original array", []interface{}{1, 2, 3}, res)

	bounded(2, 3)
	CheckValue(t, "Bounded array", []interface{}{1, 2, 3, 2, 3}, res)
}
