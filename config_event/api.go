package config_event

import (
	"encoding/json"
	"github.com/lowl11/lazyconfig/services/env_helper"
	"github.com/lowl11/lazyfile/fileapi"
	"log"
	"os"
)

func (event *Event[T]) Read() (*T, error) {
	// read .env file
	var err error
	event.env, err = env_helper.Read(event.envFileName)
	if err != nil {
		return nil, err
	}

	// read environment name value
	environmentValue := os.Getenv(event.environmentName)
	if environmentValue == "" {
		environmentValue = event.environmentDefault
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

	log.Println("Loaded", "'"+environmentValue+"'", "environment configuration")

	return &event.object.Body, nil
}

func (event *Event[T]) EnvironmentName(name string) *Event[T] {
	if name == "" {
		return event
	}

	event.environmentName = name
	return event
}

func (event *Event[T]) EnvironmentDefault(value string) *Event[T] {
	event.environmentDefault = value
	return event
}

func (event *Event[T]) EnvFileName(fileName string) *Event[T] {
	if fileName == "" {
		return event
	}

	event.envFileName = fileName
	return event
}
