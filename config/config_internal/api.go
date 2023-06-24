package config_internal

import (
	"github.com/lowl11/lazyconfig/internal/services/config_service"
	"log"
)

var (
	_configService *config_service.Service
)

func Init() {
	_configService = config_service.
		New()

	if err := _configService.Read(); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) string {
	if _configService == nil {
		panic("configService is NULL")
	}

	return _configService.Get(key)
}
