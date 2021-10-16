package event

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(m.t, m.expected, m.counter, "Call count")
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
		assert.Equal(t, &testObj, i[0], "event argument")
	}

	i1 := ee.AddEventListener("Test", listener)
	i2 := ee.On("Test", listener)
	assert.Equal(t, 2, len(ee.listeners["Test"]), `ee.listeners["Test"] length`)

	ee.Emit("Test", &testObj)

	ee.RemoveEventListener("Test", i1)
	assert.Equal(t, 1, len(ee.listeners["Test"]), `ee.listeners["Test"] length`)

	ee.Off("Test", i2)
	assert.Equal(t, 0, len(ee.listeners["Test"]), `ee.listeners["Test"] length`)

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
	assert.Equal(t, 2, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	ee.Emit("test")
	assert.Equal(t, 0, len(ee.listeners["test"]), `ee.listeners["test"] length`)

	ee.Once("test", o3.Listener)
	assert.Equal(t, 1, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	ee.Emit("test")
	assert.Equal(t, 0, len(ee.listeners["test"]), `ee.listeners["test"] length`)

	ee.Emit("test")
	assert.Equal(t, 0, len(ee.listeners["test"]), `ee.listeners["test"] length`)
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
	assert.Equal(t, 3, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	ee.Emit("test")
	assert.Equal(t, 1, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	time.Sleep(time.Millisecond * 100)

	ee.Once("test", o3.Listener)
	i2 := ee.On("test", l2.Listener)
	l1.Inc()
	l2.Inc()
	assert.Equal(t, 3, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	ee.Emit("test")
	assert.Equal(t, 2, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	time.Sleep(time.Millisecond * 100)

	ee.On("test", l3.Listener)
	assert.Equal(t, 3, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	ee.Off("test", i2)
	assert.Equal(t, 2, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	l1.Inc()
	l3.Inc()
	ee.Emit("test")
	assert.Equal(t, 2, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	time.Sleep(time.Millisecond * 100)

	l1.Inc()
	l3.Inc()
	ee.Emit("test")
	assert.Equal(t, 2, len(ee.listeners["test"]), `ee.listeners["test"] length`)
	time.Sleep(time.Millisecond * 100)
}

func TestEmitter_implementTarget(t *testing.T) {
	var target Target = NewEmitter()
	if target == nil {
		t.Errorf("Emitter doesn't implement Target interface")
	}
}
