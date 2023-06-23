package config_event

type Event[T any] struct {
	object Object[T]
	env    map[string]string
}

func New[T any]() *Event[T] {
	return &Event[T]{
		env: make(map[string]string),
	}
}
