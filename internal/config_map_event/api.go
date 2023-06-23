package config_map_event

import (
	"github.com/lowl11/lazyconfig/config_data"
	"github.com/lowl11/lazyconfig/internal/helpers/env_helper"
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
	event.env, err = env_helper.Read(config_data.GetEnvFileName())
	if err != nil {
		return err
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
		return err
	}

	// replace variables like "{{variable_name}}" to .env values
	fileContent, err = env_helper.ReplaceVariables(fileContent, event.env)
	if err != nil {
		return err
	}

	// parse file content to map object
	if err = yaml.Unmarshal(fileContent, &event.object); err != nil {
		return err
	}

	log.Println("Loaded map", "'"+environmentValue+"'", "environment configuration")

	return nil
}

func (event *Event) EnvironmentName(name string) *Event {
	config_data.SetEnvironmentName(name)
	return event
}

func (event *Event) EnvironmentDefault(value string) *Event {
	config_data.SetEnvironmentDefault(value)
	return event
}

func (event *Event) EnvFileName(fileName string) *Event {
	config_data.SetEnvFileName(fileName)
	return event
}
