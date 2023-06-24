package config_internal

import (
	"github.com/lowl11/lazyconfig/internal/services/config_service"
	"log"
	"sync"
)

var (
	// _configServicePool contains *config_service.Service
	_configServicePool sync.Pool
)

func Init() {
	configService := config_service.
		New()

	_configServicePool = sync.Pool{
		New: func() any {
			return configService
		},
	}

	if err := configService.Read(); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) string {
	configPtr := _configServicePool.Get().(*config_service.Service)
	if configPtr == nil {
		panic("configService is NULL")
	}

	return configPtr.Get(key)
}
