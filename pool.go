package event

import "sync"

// Pool is container for multiple events.
//
// Usage:
//  pool := NewPool(cfg)
//  event := pool.Event("evt1")
type Pool interface {
	// Event - get or create named event. Events with same names are equal.
	Event(name string) Event
}

// PoolConfig is config for Pool with all optional parameters.
type PoolConfig struct {
	// EventConstructor is optional. Default value: NewEvent
	EventConstructor func() Event
}

type pool struct {
	PoolConfig
	Data map[string]Event
	mu   sync.Mutex
}

func (p *pool) Event(name string) Event {
	p.mu.Lock()

	event, ok := p.Data[name]
	if !ok {
		event = p.EventConstructor()
		p.Data[name] = event
	}

	p.mu.Unlock()

	return event
}

// NewPool constructor for Pool.
func NewPool(config PoolConfig) Pool {
	if config.EventConstructor == nil {
		config.EventConstructor = NewEvent
	}

	return &pool{
		PoolConfig: config,
		Data:       map[string]Event{},
	}
}
