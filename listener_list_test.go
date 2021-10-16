package event

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenerList(t *testing.T) {
	e1 := listenerWrapper{
		Listener: Listener(func(i ...interface{}) { fmt.Println(1) }),
		Once:     false,
		Id:       1,
	}
	e2 := listenerWrapper{
		Listener: Listener(func(i ...interface{}) { fmt.Println(1) }),
		Once:     false,
		Id:       2,
	}
	e3 := listenerWrapper{
		Listener: Listener(func(i ...interface{}) { fmt.Println(1) }),
		Once:     false,
		Id:       3,
	}
	list := listenerList{e1, e2}

	assert.Equal(t, 1, list.IndexOf(2), "Index")
	assert.Equal(t, -1, list.IndexOf(3), "Index")

	assert.Equal(
		t,
		fmt.Sprintf("%v", listenerList{e3, e2}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(0)),
		"List",
	)
	assert.Equal(
		t,
		fmt.Sprintf("%v", listenerList{e1, e3}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(1)),
		"List",
	)
	assert.Equal(
		t,
		fmt.Sprintf("%v", listenerList{e1, e2, e3}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(-1)),
		"List",
	)
	assert.Equal(
		t,
		fmt.Sprintf("%v", listenerList{e1, e2, e3}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(3)),
		"List",
	)
}
