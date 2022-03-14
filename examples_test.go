package event

func Example() {
	// create common pool. empty config - use defaults
	pool := NewPool(PoolConfig{})

	// subscribe for event
	sub := pool.
		Event("event1").                // get named event
		Subscribe(func(interface{}) {}) // callback on event was published

	// publish event
	pool.
		Event("event1").
		Publish(nil) // emit event, call all callbacks

	// caution: avoid Pool.Event() frequent usage
	event := pool.Event("event1") // get event and store
	event.Publish(nil)

	// event equality
	eventCopy := pool.Event("event1")
	if event == eventCopy {
		// events with same names are single object and always equal
	}

	// unsubscribe
	sub.Close()

	// modifiers
	pool.
		Event("event1").
		Subscribe(func(interface{}) {}).
		Async(). // async mod - callback will be called in goroutine
		Once()   // once mod - callback will be called only one time

	// modifiers could be applied later, quantity and order doesn't matter
	sub2 := pool.
		Event("event1").
		Subscribe(func(interface{}) {})

	sub2.Async()
	sub2.Once()
	sub2.Async().Once()
	sub2.Once().Async()
}

func ExampleEvent() {
	event := NewEvent()

	// subscribe for event
	sub := event.
		Subscribe(func(interface{}) {}).
		Async().
		Once()

	// publish event
	event.
		Publish(nil)

	// unsubscribe
	sub.Close()
}
