package config_event

import (
	"encoding/json"
	"github.com/lowl11/lazyconfig/config_data"
	"github.com/lowl11/lazyconfig/services/env_helper"
	"github.com/lowl11/lazyfile/fileapi"
	"log"
	"os"
)

func (event *Event[T]) Read() (*T, error) {
	// read .env file
	var err error
	event.env, err = env_helper.Read(config_data.GetEnvFileName())
	if err != nil {
		return nil, err
	}

	// read environment name value
	environmentValue := os.Getenv(config_data.GetEnvironmentName())
	if environmentValue == "" {
		environmentValue = config_data.GetEnvironmentDefault()
	}

	// config file name
	fileName := event.fileName(environmentValue)

	// read config file
	fileContent, err := fileapi.Read(fileName)
	if err != nil {
		return nil, err
	}

	// replace variables like "{{variable_name}}" to .env values
	fileContent, err = env_helper.ReplaceVariables(fileContent, event.env)
	if err != nil {
		return nil, err
	}

	// parse file content to Configuration struct object
	if err = json.Unmarshal(fileContent, &event.object.Body); err != nil {
		return nil, err
	}

	log.Println("Loaded config", "'"+environmentValue+"'", "environment configuration")

	return &event.object.Body, nil
}

func (event *Event[T]) EnvironmentName(name string) *Event[T] {
	config_data.SetEnvironmentName(name)
	return event
}

func (event *Event[T]) EnvironmentDefault(value string) *Event[T] {
	config_data.SetEnvironmentDefault(value)
	return event
}

func (event *Event[T]) EnvFileName(fileName string) *Event[T] {
	config_data.SetEnvFileName(fileName)
	return event
}
