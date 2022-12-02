package event

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testListenerContext[T any] struct {
	counter  int64
	expected int64
	t        *testing.T
}

func (m *testListenerContext[T]) Listener(context.Context, T) {
	m.t.Helper()
	atomic.AddInt64(&m.counter, 1)
	assert.Equal(m.t, m.expected, m.counter, "Invalid listener call count")
}

func (m *testListenerContext[T]) LongListener(ctx context.Context, data T) {
	m.t.Helper()
	time.Sleep(time.Millisecond * 100)
	m.Listener(ctx, data)
}

func (m *testListenerContext[T]) Inc() {
	atomic.AddInt64(&m.expected, 1)
}

func TestEventWithContext(t *testing.T) {
	t.Parallel()

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	eventInstance := &eventWithContext[int]{}

	listener := testListenerContext[int]{t: t}

	sub1 := eventInstance.Subscribe(listener.Listener)
	assert.Len(t, eventInstance.Subscribers, 1)

	listener.Inc()
	eventInstance.Publish(ctx, 0)

	listener.Inc()
	eventInstance.Publish(ctx, 0)

	sub1.Close()
	eventInstance.Publish(ctx, 0)
	assert.Len(t, eventInstance.Subscribers, 0)

	eventInstance.Publish(ctx, 0)

	time.Sleep(time.Millisecond * 500)
}
