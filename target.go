package event

type Target interface {
	Emit(name string, arguments ...interface{})
	AddEventListener(name string, listener Listener)
	On(name string, listener Listener)
	RemoveEventListener(name string, listener Listener)
	Off(name string, listener Listener)
}
