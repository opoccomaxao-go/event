package event

import "context"

// Signal like event.WithContext, but without value.
type Signal Event[context.Context]

func NewSignal() Signal {
	return NewEvent[context.Context]()
}
