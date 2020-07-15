package event

type listenerList []listenerWrapper

func (ll listenerList) IndexOf(element Listener) int {
	p := element.ptr()
	for i := range ll {
		if ll[i].ptr() == p {
			return i
		}
	}
	return -1
}

func (ll listenerList) RemoveByIndex(index int) listenerList {
	if index < 0 {
		return ll
	}
	last := len(ll) - 1
	if index > last {
		return ll
	}
	ll[index] = ll[last]
	ll[last] = nilWrapper
	return ll[:last]
}

func (ll listenerList) Remove(listener Listener) listenerList {
	return ll.RemoveByIndex(ll.IndexOf(listener))
}
