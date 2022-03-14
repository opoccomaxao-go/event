package event

import "io"

type Subscriber interface {
	io.Closer
	// Async modifier. These callbacks are called in goroutine.
	Async() Subscriber
	// Once modifier. These callbacks are only called once.
	Once() Subscriber
}

type subscriber struct {
	Listener func(interface{})
	Closed   bool
	async    bool
	once     bool
}

func (s *subscriber) Close() error {
	s.Closed = true

	return nil
}

func (s *subscriber) Exec(data interface{}) {
	if s.async {
		go s.Listener(data)
	} else {
		s.Listener(data)
	}

	if s.once {
		s.Closed = true
	}
}

func (s *subscriber) Async() Subscriber {
	s.async = true

	return s
}

func (s *subscriber) Once() Subscriber {
	s.once = true

	return s
}
