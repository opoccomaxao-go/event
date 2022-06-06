package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	t.Parallel()

	pool := NewPool(PoolConfig{})

	event1 := pool.Event("test")
	event2 := pool.Event("test")

	assert.True(t, event1 == event2)
}

//nolint:varnamelen // test var names
func TestWithType(t *testing.T) {
	t.Parallel()

	p1 := NewPool(PoolConfig{})
	p2 := WithType[int](p1)
	p3 := WithType[any](p2)
	p4 := WithType[int](p1)

	e1 := p1.Event("test")
	e3 := p3.Event("test")
	e2 := p2.Event("test")
	e4 := p4.Event("test")

	assert.True(t, e1 == e3)
	assert.True(t, e2 == e4)
}
