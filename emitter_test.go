package event

import (
	. "github.com/opoccomaxao-go/helpers/test"
	"sync"
	"testing"
	"time"
)

type testStruct struct {
	A int
	B string
	C time.Time
}

type testListener struct {
	mu       sync.Mutex
	counter  int
	expected int
	t        *testing.T
}

func NewTestListener(t *testing.T) *testListener {
	return &testListener{
		t: t,
	}
}

func (m *testListener) Listener(...interface{}) {
	m.t.Helper()
	m.mu.Lock()
	m.counter++
	if m.counter != m.expected {
		CheckValue(m.t, "Call count", m.expected, m.counter)
	}
	m.mu.Unlock()
}

func (m *testListener) Inc() {
	m.mu.Lock()
	m.expected++
	m.mu.Unlock()
}

func TestEventEmitter(t *testing.T) {
	ee := NewEmitter()
	testObj := testStruct{
		A: 10,
		B: "str",
		C: time.Now(),
	}
	var listener Listener = func(i ...interface{}) {
		CheckValue(t, "event argument", &testObj, i[0])
	}

	i1 := ee.AddEventListener("Test", listener)
	i2 := ee.On("Test", listener)
	CheckValue(t, "ee.listeners[\"Test\"] length", 2, len(ee.listeners["Test"]))

	ee.Emit("Test", &testObj)

	ee.RemoveEventListener("Test", i1)
	CheckValue(t, "ee.listeners[\"Test\"] length", 1, len(ee.listeners["Test"]))

	ee.Off("Test", i2)
	CheckValue(t, "ee.listeners[\"Test\"] length", 0, len(ee.listeners["Test"]))

	ee.Emit("Test", &testObj)
	time.Sleep(time.Millisecond * 100)
}

func TestEmitter_Once(t *testing.T) {
	ee := NewEmitter()

	o1 := NewTestListener(t)
	o2 := NewTestListener(t)
	o3 := NewTestListener(t)
	o1.Inc()
	o2.Inc()
	o3.Inc()

	ee.Once("test", o1.Listener)
	ee.Once("test", o2.Listener)
	CheckValue(t, "ee.listeners[\"test\"] length", 2, len(ee.listeners["test"]))
	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 0, len(ee.listeners["test"]))

	ee.Once("test", o3.Listener)
	CheckValue(t, "ee.listeners[\"test\"] length", 1, len(ee.listeners["test"]))
	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 0, len(ee.listeners["test"]))

	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 0, len(ee.listeners["test"]))
	time.Sleep(time.Millisecond * 100)
}

func TestEmitter_Mixed(t *testing.T) {
	ee := NewEmitter()

	l1 := NewTestListener(t)
	l2 := NewTestListener(t)
	l3 := NewTestListener(t)

	o1 := NewTestListener(t)
	o2 := NewTestListener(t)
	o3 := NewTestListener(t)
	o1.Inc()
	o2.Inc()
	o3.Inc()

	ee.Once("test", o1.Listener)
	ee.Once("test", o2.Listener)
	ee.On("test", l1.Listener)
	l1.Inc()
	CheckValue(t, "ee.listeners[\"test\"] length", 3, len(ee.listeners["test"]))
	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 1, len(ee.listeners["test"]))
	time.Sleep(time.Millisecond * 100)

	ee.Once("test", o3.Listener)
	i2 := ee.On("test", l2.Listener)
	l1.Inc()
	l2.Inc()
	CheckValue(t, "ee.listeners[\"test\"] length", 3, len(ee.listeners["test"]))
	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 2, len(ee.listeners["test"]))
	time.Sleep(time.Millisecond * 100)

	ee.On("test", l3.Listener)
	CheckValue(t, "ee.listeners[\"test\"] length", 3, len(ee.listeners["test"]))
	ee.Off("test", i2)
	CheckValue(t, "ee.listeners[\"test\"] length", 2, len(ee.listeners["test"]))
	l1.Inc()
	l3.Inc()
	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 2, len(ee.listeners["test"]))
	time.Sleep(time.Millisecond * 100)

	l1.Inc()
	l3.Inc()
	ee.Emit("test")
	CheckValue(t, "ee.listeners[\"test\"] length", 2, len(ee.listeners["test"]))
	time.Sleep(time.Millisecond * 100)
}

func TestEmitter_implementTarget(t *testing.T) {
	var target Target = NewEmitter()
	if target == nil {
		t.Errorf("Emitter doesn't implement Target interface")
	}
}
