package event

import (
	"fmt"
	. "github.com/opoccomaxao-go/helpers/test"
	"testing"
)

func TestListenerList(t *testing.T) {
	e1 := listenerWrapper{
		Listener: Listener(func(i ...interface{}) { fmt.Println(1) }),
		Once:     false,
	}
	e2 := listenerWrapper{
		Listener: Listener(func(i ...interface{}) { fmt.Println(1) }),
		Once:     false,
	}
	e3 := listenerWrapper{
		Listener: Listener(func(i ...interface{}) { fmt.Println(1) }),
		Once:     false,
	}
	list := listenerList{e1, e2}

	CheckValue(t, "Index", 1, list.IndexOf(e2.Listener))
	CheckValue(t, "Index", -1, list.IndexOf(e3.Listener))

	CheckValue(t, "List",
		fmt.Sprintf("%v", listenerList{e3, e2}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(0)),
	)
	CheckValue(t, "List",
		fmt.Sprintf("%v", listenerList{e1, e3}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(1)),
	)
	CheckValue(t, "List",
		fmt.Sprintf("%v", listenerList{e1, e2, e3}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(-1)),
	)
	CheckValue(t, "List",
		fmt.Sprintf("%v", listenerList{e1, e2, e3}),
		fmt.Sprintf("%v", listenerList{e1, e2, e3}.RemoveByIndex(3)),
	)
}
