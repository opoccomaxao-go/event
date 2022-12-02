package event

import (
	"context"
	"sync"
)

// WithContext is container for all subscribers and bus for pub/sub.
//
// Usage:
//
//	event := NewEventWithContext()
//
// publish
//
//	event.Publish(ctx, data)
//
// subscribe
//
//	event.Subscribe(func(context.Context, interface{}){ ... })
type WithContext[T any] interface {
	// Subscribe to event. listener is a callback that is called when Publish is called.
	Subscribe(listener func(context.Context, T)) Subscriber
	// Publish call all callbacks for this event with specified argument.
	Publish(ctx context.Context, arg T)
}

type eventWithContext[T any] struct {
	Subscribers []*subscriber[T]
	mu          sync.Mutex
}

func (e *eventWithContext[T]) Subscribe(listener func(context.Context, T)) Subscriber {
	if listener == nil {
		return &subscriber[T]{
			Listener: func(T) {},
			Closed:   true,
		}
	}

	sub := &subscriber[T]{
		ListenerContext: listener,
	}

	e.mu.Lock()
	e.Subscribers = append(e.Subscribers, sub)
	e.mu.Unlock()

	return sub
}

func (e *eventWithContext[T]) Publish(ctx context.Context, arg T) {
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
		sub.ExecContext(ctx, arg)
	}
}

// NewEvent constructor for Event.
func NewEventWithContext[T any]() WithContext[T] {
	return &eventWithContext[T]{}
}
