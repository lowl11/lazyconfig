package config_map_event

import (
	"github.com/lowl11/lazyconfig/services/env_helper"
	"github.com/lowl11/lazyfile/fileapi"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func (event *Event) Get(key string) string {
	return event.object[key]
}

func (event *Event) Read() error {
	// read .env file
	var err error
	event.env, err = env_helper.Read(event.envFileName)

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
		return err
	}

	// replace variables like "{{variable_name}}" to .env values
	fileContent, err = env_helper.ReplaceVariables(fileContent, event.env)
	if err != nil {
		return err
	}

	// parse file content to Configuration struct object
	if err = yaml.Unmarshal(fileContent, &event.object); err != nil {
		return err
	}

	log.Println("Loaded", "'"+environmentValue+"'", "environment configuration")

	return nil
}

func (event *Event) EnvironmentName(name string) *Event {
	if name == "" {
		return event
	}

	event.environmentName = name
	return event
}

func (event *Event) EnvironmentDefault(value string) *Event {
	event.environmentDefault = value
	return event
}

func (event *Event) EnvFileName(fileName string) *Event {
	if fileName == "" {
		return event
	}

	event.envFileName = fileName
	return event
}
