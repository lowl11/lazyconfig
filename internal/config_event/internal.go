package config_event

import "github.com/lowl11/lazyconfig/config_data"

type Object[T any] struct {
	Body T
}

func (event *Event[T]) fileName(value string) string {
	return config_data.ConfigFolder + value + ".json"
}
