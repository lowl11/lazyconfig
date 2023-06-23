package confapi

import (
	"github.com/lowl11/lazyconfig/internal/config_event"
	"github.com/lowl11/lazyconfig/internal/config_map_event"
)

var (
	_configMap *config_map_event.Event
)

func NewConfig[T any]() *config_event.Event[T] {
	return config_event.New[T]()
}

func NewMap() *config_map_event.Event {
	_configMap = config_map_event.New()
	return _configMap
}

func Get(key string) string {
	if _configMap == nil {
		panic("config map is NULL")
	}
	return _configMap.Get(key)
}
