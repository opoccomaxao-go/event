package event

type listenerList []listenerWrapper

func (ll listenerList) IndexOf(id ListenerId) int {
	for i := range ll {
		if ll[i].Id == id {
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

func (ll listenerList) Remove(id ListenerId) listenerList {
	return ll.RemoveByIndex(ll.IndexOf(id))
}
