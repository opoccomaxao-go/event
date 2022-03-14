package event

import "sync"

type Pool interface {
	// Event - get named event. Events with same names are equal.
	Event(name string) Event
}

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

func NewPool(config PoolConfig) Pool {
	if config.EventConstructor == nil {
		config.EventConstructor = NewEvent
	}

	return &pool{
		PoolConfig: config,
		Data:       map[string]Event{},
	}
}
