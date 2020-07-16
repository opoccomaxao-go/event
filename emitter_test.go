package event

import (
	. "github.com/opoccomaxao-go/helpers/test"
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
		CheckValue(t, "event argument", testObj, i)
	}

	ee.AddEventListener("Test", listener)
	ee.On("Test", listener)
	CheckValue(t, "ee.listeners[\"Test\"] length", 2, len(ee.listeners["Test"]))

	ee.Emit("Test", &testObj)

	ee.RemoveEventListener("Test", listener)
	CheckValue(t, "ee.listeners[\"Test\"] length", 1, len(ee.listeners["Test"]))

	ee.Off("Test", listener)
	CheckValue(t, "ee.listeners[\"Test\"] length", 0, len(ee.listeners["Test"]))

	ee.Emit("Test", &testObj)
}

func makeTestOnceListener(t *testing.T) Listener {
	counter := 0
	return func(...interface{}) {
		t.Helper()
		counter++
		if counter > 1 {
			t.Errorf("Second call of once listener")
		}
	}
}

func TestEmitter_Once(t *testing.T) {
	ee := NewEmitter()
	ee.Once("test", makeTestOnceListener(t))
	ee.Once("test", makeTestOnceListener(t))
	ee.Emit("test")
	ee.Once("test", makeTestOnceListener(t))
	ee.Emit("test")
	ee.Emit("test")
}

func TestEmitter_implementTarget(t *testing.T) {
	var target Target = NewEmitter()
	if target == nil {
		t.Errorf("Emitter doesn't implement Target interface")
	}
}
