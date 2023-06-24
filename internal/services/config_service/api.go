package config_service

import (
	"errors"
	"github.com/lowl11/lazyconfig/internal/helpers/env_helper"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"gopkg.in/yaml.v3"
)

func (service *Service) Read() error {
	if !folderapi.Exist(service.baseFolder) {
		return errors.New("base folder does not exist: " + service.baseFolder)
	}

	// read .env file
	envFileContent, err := env_helper.Read(service.environmentFileName)
	if err != nil {
		return err
	}

	baseVariables := make(map[string]string)

	if fileapi.Exist(service.environmentBase) {
		// read base config.yml file
		envBaseContent, err := fileapi.Read(service.environmentBase)
		if err != nil {
			return err
		}

		if err = yaml.Unmarshal(envBaseContent, &baseVariables); err != nil {
			return err
		}
	}

	// read <environment>.yml file
	envContent, err := fileapi.Read(service.environment + ".yml")
	if err != nil {
		return err
	}

	// replace variables from file
	config, err := env_helper.ReplaceVariables(envContent, envFileContent)
	if err != nil {
		return err
	}

	// set data to values map
	if err = yaml.Unmarshal(config, &service.variables); err != nil {
		return err
	}

	// check if there is no such variable
	// even if variable with such key exist, need to check if in current it is empty
	for key, baseValue := range baseVariables {
		if currentValue, ok := service.variables[key]; (!ok || currentValue == "") && baseValue != "" {
			service.variables[key] = baseValue
		}
	}

	return nil
}

func (service *Service) Get(key string) string {
	return service.variables[key]
}

func (service *Service) BaseFolder(baseFolder string) *Service {
	if baseFolder == "" {
		return service
	}

	if baseFolder[len(baseFolder)-1] != '/' {
		baseFolder += "/"
	}

	// update all paths
	service.baseFolder = baseFolder
	service.environment = service.baseFolder + service.environment
	service.environment = service.baseFolder + service.environmentBase
	return service
}

func (service *Service) Environment(environment string) *Service {
	service.environment = environment
	return service
}

func (service *Service) EnvironmentVariableName(variableName string) *Service {
	service.environmentVariableName = variableName
	return service
}

func (service *Service) EnvironmentFileName(fileName string) *Service {
	service.environmentFileName = fileName
	return service
}
