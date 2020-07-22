package event

type Listener func(...interface{})

func (l Listener) Bind(boundedArgv ...interface{}) Listener {
	return func(native ...interface{}) {
		bLen := len(boundedArgv)
		args := make([]interface{}, bLen+len(native))
		copy(args[:bLen], boundedArgv)
		copy(args[bLen:], native)
		l(args...)
	}
}
