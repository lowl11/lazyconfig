package config

import (
	"github.com/lowl11/lazyconfig/config/config_internal"
	"os"
	"strings"
)

func Get(key string) string {
	value := config_internal.Get(key)
	if value == "" {
		value = os.Getenv(key)
	}
	return value
}

func Env() string {
	return strings.ToLower(Get("env"))
}

func IsProduction() bool {
	return Env() == "production"
}

func IsTest() bool {
	return Env() == "test"
}

func IsDev() bool {
	return Env() == "dev"
}
