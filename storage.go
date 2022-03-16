package event

import (
	"sync"
)

type internalID[T any] string

type WithStorage interface {
	// Storage - get internal storage for copy.
	Storage() *Storage
}

type Storage struct {
	data map[interface{}]interface{}
	mu   sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		data: map[interface{}]interface{}{},
	}
}
