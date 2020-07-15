package event

type ListenerList []Listener

func (ll ListenerList) IndexOf(element Listener) int {
	p := element.ptr()
	for i := range ll {
		if ll[i].ptr() == p {
			return i
		}
	}
	return -1
}
