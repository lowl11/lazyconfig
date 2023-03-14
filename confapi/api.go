package confapi

import "github.com/lowl11/lazyconfig/config_event"

func New[T any]() *config_event.Event[T] {
	return config_event.New[T]()
}
