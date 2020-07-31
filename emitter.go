package event

import (
	"sync"
)

// TODO add comments
type Emitter struct {
	listeners map[string]listenerList
	mu        sync.Mutex
	id        ListenerId
}

func NewEmitter() *Emitter {
	return &Emitter{
		listeners: map[string]listenerList{},
	}
}

func (e *Emitter) getId() ListenerId {
	e.id++
	return e.id
}

func (e *Emitter) emit(name string, arguments []interface{}) {
	e.mu.Lock()
	var listenersToProcess []Listener
	if listeners, ok := e.listeners[name]; ok {
		end := len(listeners) - 1
		for i := 0; i <= end; {
			l := listeners[i]
			listenersToProcess = append(listenersToProcess, l.Listener)
			if l.Once {
				listeners = listeners.RemoveByIndex(i)
				end--
			} else {
				i++
			}
		}
		e.listeners[name] = listeners
	}
	// another event can be processed now
	e.mu.Unlock()
	// single thread used
	for _, l := range listenersToProcess {
		l(arguments...)
	}
}

func (e *Emitter) on(name string, listener listenerWrapper) (res ListenerId) {
	e.mu.Lock()
	arr, ok := e.listeners[name]
	if !ok {
		t := make(listenerList, 0, 10)
		e.listeners[name] = t
		arr = t
	}
	res = e.getId()
	listener.Id = res
	e.listeners[name] = append(arr, listener)
	e.mu.Unlock()
	return
}

func (e *Emitter) Emit(name string, arguments ...interface{}) {
	e.emit(name, arguments)
}

func (e *Emitter) On(name string, listener Listener) ListenerId {
	return e.on(name, listenerWrapper{
		Listener: listener,
		Once:     false,
	})
}

func (e *Emitter) Once(name string, listener Listener) ListenerId {
	return e.on(name, listenerWrapper{
		Listener: listener,
		Once:     true,
	})
}

func (e *Emitter) Off(name string, id ListenerId) {
	e.mu.Lock()
	if arr, ok := e.listeners[name]; ok {
		e.listeners[name] = arr.Remove(id)
	}
	e.mu.Unlock()
}

func (e *Emitter) AddEventListener(name string, listener Listener) ListenerId {
	return e.On(name, listener)
}

func (e *Emitter) RemoveEventListener(name string, id ListenerId) {
	e.Off(name, id)
}
