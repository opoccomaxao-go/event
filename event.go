package event

import (
	"sync"
)

type Event interface {
	// Subscribe to event. listener is a callback that is called when Publish is called.
	Subscribe(listener func(interface{})) Subscriber
	// Publish call all callbacks for this event with specified argument.
	Publish(arg interface{})
}

type event struct {
	Subscribers []*subscriber
	mu          sync.Mutex
}

func (e *event) Subscribe(listener func(interface{})) Subscriber {
	if listener == nil {
		return &subscriber{
			Listener: func(interface{}) {},
			Closed:   true,
		}
	}

	sub := &subscriber{
		Listener: listener,
	}

	e.mu.Lock()
	e.Subscribers = append(e.Subscribers, sub)
	e.mu.Unlock()

	return sub
}

func (e *event) Publish(arg interface{}) {
	e.mu.Lock()
	subs := e.Subscribers
	e.Subscribers = e.Subscribers[0:0]

	for _, sub := range subs {
		if !sub.Closed {
			e.Subscribers = append(e.Subscribers, sub)
		}
	}

	subs = e.Subscribers
	e.mu.Unlock()

	for _, sub := range subs {
		sub.Exec(arg)
	}
}

func NewEvent() Event {
	return &event{}
}
