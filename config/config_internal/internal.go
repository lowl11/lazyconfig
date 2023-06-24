package config_internal

import "sync"

var (
	// _configServicePool contains *config_service.Service
	_configServicePool sync.Pool
)

var (
	_baseFolder              = ""
	_environment             = ""
	_environmentFileName     = ""
	_environmentVariableName = ""
)
