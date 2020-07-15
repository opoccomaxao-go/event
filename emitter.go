package event

import (
	"sync"
)

type Emitter struct {
	listeners map[string]listenerList
	mu        sync.Mutex
}

func NewEmitter() *Emitter {
	return &Emitter{
		listeners: map[string]listenerList{},
	}
}

func (e *Emitter) emit(name string, arguments []interface{}) {
	e.mu.Lock()
	if listeners, ok := e.listeners[name]; ok {
		end := len(listeners) - 1
		for i := 0; i <= end; {
			l := listeners[i]
			go l.Listener(arguments...)
			if l.Once {
				listeners = listeners.RemoveByIndex(i)
				end--
			} else {
				i++
			}
		}
		e.listeners[name] = listeners
	}
	e.mu.Unlock()
}

func (e *Emitter) on(name string, listener listenerWrapper) {
	e.mu.Lock()
	arr, ok := e.listeners[name]
	if !ok {
		t := make(listenerList, 0, 10)
		e.listeners[name] = t
		arr = t
	}
	e.listeners[name] = append(arr, listener)
	e.mu.Unlock()
}

func (e *Emitter) Emit(name string, arguments ...interface{}) {
	e.emit(name, arguments)
}

func (e *Emitter) On(name string, listener Listener) {
	e.on(name, listenerWrapper{
		Listener: listener,
		Once:     false,
	})
}

func (e *Emitter) Once(name string, listener Listener) {
	e.on(name, listenerWrapper{
		Listener: listener,
		Once:     true,
	})
}

func (e *Emitter) Off(name string, listener Listener) {
	e.mu.Lock()
	if arr, ok := e.listeners[name]; ok {
		e.listeners[name] = arr.Remove(listener)
	}
	e.mu.Unlock()
}

func (e *Emitter) AddEventListener(name string, listener Listener) {
	e.On(name, listener)
}

func (e *Emitter) RemoveEventListener(name string, listener Listener) {
	e.Off(name, listener)
}
