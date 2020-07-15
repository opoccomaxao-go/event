package event

import (
	"testing"
)

func InterfaceArrEqual(arr1, arr2 []interface{}) bool {
	l := len(arr1)
	if l != len(arr2) {
		return false
	}
	for i := 0; i < l; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestListener_Bind(t *testing.T) {
	var res []interface{}
	var original Listener = func(argv ...interface{}) {
		res = argv
	}
	bounded := original.Bind(1, 2, 3)

	expected := []interface{}{1, 2, 3}
	original(1, 2, 3)
	if !InterfaceArrEqual(expected, res) {
		t.Errorf("Original call is not successful\nexpected: %v, received: %v", expected, res)
	}

	expected = []interface{}{1, 2, 3, 2, 3}
	bounded(2, 3)
	if !InterfaceArrEqual(expected, res) {
		t.Errorf("Bounded call is not successful\nexpected: %v, received: %v", expected, res)
	}
}
