package event

import (
	"fmt"
	"testing"
	"time"
)

type testStruct struct {
	A int
	B string
	C time.Time
}

func TestEventEmitter(t *testing.T) {
	ee := NewEmitter()
	testObj := testStruct{
		A: 10,
		B: "str",
		C: time.Now(),
	}
	var listener Listener = func(i ...interface{}) {
		fmt.Println("Test", i)
		if &testObj != i[0].(*testStruct) {
			t.Errorf("Event has %v argument; want %v", i, testObj)
		}
	}
	ee.On("Test", listener)
	if l := len(ee.listeners["Test"]); l != 1 {
		t.Errorf("EE has %d Test-listeners; want 1", l)
	}
	go ee.Emit("Test", &testObj)
	ee.Off("Test", listener)
	if l := len(ee.listeners["Test"]); l != 0 {
		t.Errorf("EE has %d Test-listeners; want 0", l)
	}
	ee.Off("Test", listener)
	if l := len(ee.listeners["Test"]); l != 0 {
		t.Errorf("EE has %d Test-listeners; want 0", l)
	}
}
