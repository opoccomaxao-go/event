package event

// Pool is container for multiple untyped events.
//
// Usage:
//  pool := NewPool(cfg)
//  event := pool.Event("evt1")
type Pool[T any] interface {
	WithStorage

	// Event - get or create named event. Events with same names are equal.
	Event(name string) Event[T]
}

// PoolConfig is config for Pool with all optional parameters.
type PoolConfig struct {
	Storage *Storage // Storage for use with typed events.
}

type pool[T any] struct {
	storage *Storage
}

func (p *pool[T]) Event(name string) Event[T] {
	id := internalID[T](name)

	p.storage.mu.Lock()
	defer p.storage.mu.Unlock()

	event, ok := p.storage.data[id]
	if ok {
		res, ok := event.(Event[T])
		if ok {
			return res
		}
	}

	res2 := NewEvent[T]()
	p.storage.data[id] = res2

	return res2
}

func (p *pool[T]) Storage() *Storage {
	return p.storage
}

// NewTypedPool constructor for Pool.
func NewTypedPool[T any](cfg PoolConfig) Pool[T] {
	if cfg.Storage == nil {
		cfg.Storage = NewStorage()
	}

	return &pool[T]{
		storage: cfg.Storage,
	}
}

// NewPool constructor for Pool.
func NewPool(cfg PoolConfig) Pool[interface{}] {
	return NewTypedPool[interface{}](cfg)
}

// WithType bound pool to specified type with storage saving.
func WithType[T any](pool WithStorage) Pool[T] {
	return NewTypedPool[T](PoolConfig{
		Storage: pool.Storage(),
	})
}
