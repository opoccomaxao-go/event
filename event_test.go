package event

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testListener struct {
	counter  int64
	expected int64
	t        *testing.T
}

func (m *testListener) Listener(interface{}) {
	m.t.Helper()
	atomic.AddInt64(&m.counter, 1)
	assert.Equal(m.t, m.expected, m.counter, "Invalid listener call count")
}

func (m *testListener) LongListener(data interface{}) {
	m.t.Helper()
	time.Sleep(time.Millisecond * 100)
	m.Listener(data)
}

func (m *testListener) Inc() {
	atomic.AddInt64(&m.expected, 1)
}

func TestEvent(t *testing.T) {
	t.Parallel()

	eventInstance := &event{}

	listener := testListener{t: t}

	sub1 := eventInstance.Subscribe(listener.Listener)
	assert.Len(t, eventInstance.Subscribers, 1)

	listener.Inc()
	eventInstance.Publish(nil)

	listener.Inc()
	eventInstance.Publish(nil)

	sub1.Close()
	eventInstance.Publish(nil)
	assert.Len(t, eventInstance.Subscribers, 0)

	eventInstance.Publish(nil)

	time.Sleep(time.Millisecond * 500)
}

func TestEvent_Once(t *testing.T) {
	t.Parallel()

	eventInstance := &event{}

	listener1 := testListener{t: t}
	listener2 := testListener{t: t}
	listener3 := testListener{t: t}

	eventInstance.Subscribe(listener1.Listener).Once()
	eventInstance.Subscribe(listener2.Listener).Once()
	assert.Len(t, eventInstance.Subscribers, 2)

	listener1.Inc()
	listener2.Inc()
	eventInstance.Publish(nil)
	assert.Len(t, eventInstance.Subscribers, 2)

	eventInstance.Publish(nil)
	assert.Len(t, eventInstance.Subscribers, 0)

	eventInstance.Subscribe(listener3.Listener).Once()
	assert.Len(t, eventInstance.Subscribers, 1)

	listener3.Inc()
	eventInstance.Publish(nil)

	eventInstance.Publish(nil)
	assert.Len(t, eventInstance.Subscribers, 0)

	time.Sleep(time.Millisecond * 500)
}

func TestEvent_Async(t *testing.T) {
	t.Parallel()

	eventInstance := &event{}

	listener := testListener{t: t}

	sub1 := eventInstance.Subscribe(listener.LongListener).Async()
	assert.Len(t, eventInstance.Subscribers, 1)

	listener.Inc()
	eventInstance.Publish(nil)

	time.Sleep(time.Millisecond * 200)

	listener.Inc()
	eventInstance.Publish(nil)

	time.Sleep(time.Millisecond * 200)

	sub1.Close()
	eventInstance.Publish(nil)
	assert.Len(t, eventInstance.Subscribers, 0)

	eventInstance.Publish(nil)

	time.Sleep(time.Millisecond * 500)
}

func TestEvent_pubType(t *testing.T) {
	t.Parallel()

	type test struct {
		A string
		B int64
		C float64
	}

	eventInstance := &event{}

	listener := func(data interface{}) {
		assert.Equal(t, &test{
			A: "1",
			B: 2,
			C: 3,
		}, data)
	}

	eventInstance.Subscribe(listener)
	assert.Len(t, eventInstance.Subscribers, 1)

	eventInstance.Publish(&test{
		A: "1",
		B: 2,
		C: 3,
	})

	time.Sleep(time.Millisecond * 500)
}
