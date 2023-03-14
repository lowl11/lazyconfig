package config_event

import (
	"encoding/json"
	"github.com/lowl11/lazyfile/fileapi"
	"os"
)

func (event *Event[T]) Read() (*T, error) {
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

	// parse file content to Configuration struct object
	if err = json.Unmarshal(fileContent, &event.object.Body); err != nil {
		return nil, err
	}

	return &event.object.Body, nil
}

func (event *Event[T]) Configuration(name, value string) *Event[T] {
	if name == "" || value == "" {
		return event
	}

	event.configurations[name] = value
	return event
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
