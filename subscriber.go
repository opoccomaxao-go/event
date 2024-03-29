package event

import (
	"context"
	"io"
)

// Subscriber is controller for subscription. Commonly used for subscription cancelling.
//
// Usage:
//
//	sub := event.Subscribe(fn)
//
// or with mods
//
//	sub := event.Subscribe(fn).Async().Once()
//
// cancel:
//
//	sub.Close()
type Subscriber interface {
	io.Closer
	// Async modifier. These callbacks are called in goroutine.
	Async() Subscriber
	// Once modifier. These callbacks are only called once.
	Once() Subscriber
}

type subscriber[T any] struct {
	Listener        func(T)
	ListenerContext func(context.Context, T)
	Closed          bool
	async           bool
	once            bool
}

func (s *subscriber[T]) Close() error {
	s.Closed = true

	return nil
}

func (s *subscriber[T]) Exec(data T) {
	if s.async {
		go s.Listener(data)
	} else {
		s.Listener(data)
	}

	if s.once {
		s.Closed = true
	}
}

func (s *subscriber[T]) ExecContext(ctx context.Context, data T) {
	if s.async {
		go s.ListenerContext(ctx, data)
	} else {
		s.ListenerContext(ctx, data)
	}

	if s.once {
		s.Closed = true
	}
}

func (s *subscriber[T]) Async() Subscriber {
	s.async = true

	return s
}

func (s *subscriber[T]) Once() Subscriber {
	s.once = true

	return s
}

// NilSubscriber constructs empty valid Subscriber to pass as argument or default not initialized value.
func NilSubscriber() Subscriber {
	return &subscriber[struct{}]{}
}
