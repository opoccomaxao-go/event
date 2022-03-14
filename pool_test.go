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

	assert.Equal(t, event1, event2)
}
