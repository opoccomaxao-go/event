package event

type Target interface {
	Emit(name string, arguments ...interface{})
	AddEventListener(name string, listener Listener) ListenerId
	On(name string, listener Listener) ListenerId
	Once(name string, listener Listener) ListenerId
	RemoveEventListener(name string, id ListenerId)
	Off(name string, id ListenerId)
}
