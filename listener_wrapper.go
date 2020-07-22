package event

type listenerWrapper struct {
	Listener
	Once bool
	Id   ListenerId
}

var nilWrapper = listenerWrapper{}
