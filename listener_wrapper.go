package event

type listenerWrapper struct {
	Listener
	Once bool
}

var nilWrapper = listenerWrapper{}
