package config_internal

import (
	"github.com/lowl11/lazyconfig/internal/services/config_service"
	"log"
	"sync"
)

func Init() {
	configService := config_service.
		New()

	if _environmentVariableName != "" {
		configService.EnvironmentVariableName(_environmentVariableName)
	}

	if _environmentFileName != "" {
		configService.EnvironmentFileName(_environmentFileName)
	}

	if _environment != "" {
		configService.Environment(_environment)
	}

	if _baseFolder != "" {
		configService.BaseFolder(_baseFolder)
	}

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

func SetEnvironment(value string) {
	_environment = value
}

func SetBaseFolder(folderPath string) {
	_baseFolder = folderPath
}

func SetEnvironmentFileName(fileName string) {
	_environmentFileName = fileName
}

func SetEnvironmentVariableName(variableName string) {
	_environmentVariableName = variableName
}
