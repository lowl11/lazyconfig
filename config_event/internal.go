package config_event

import (
	"github.com/lowl11/lazyfile/fileapi"
	"os"
	"regexp"
	"strings"
)

const (
	environmentName    = "env"
	environmentDefault = "test"

	configFolderBase   = "profiles/"
	envFileNameDefault = ".env"
)

type Object[T any] struct {
	Body T
}

func (event *Event[T]) fileName(value string) string {
	return configFolderBase + value + ".json"
}

func (event *Event[T]) readEnv() error {
	if !fileapi.Exist(event.envFileName) {
		return nil
	}

	fileContent, err := fileapi.Read(event.envFileName)
	if err != nil {
		return err
	}

	combinations := strings.Split(string(fileContent), "\n")

	for _, combo := range combinations {
		dividedCombo := strings.Split(combo, "=")
		if len(dividedCombo) == 1 {
			continue
		}

		event.env[dividedCombo[0]] = dividedCombo[1]
	}

	return nil
}

func (event *Event[T]) replaceEnvVariables(fileContent []byte) ([]byte, error) {
	// define variables
	variableRegex, err := regexp.Compile("{{(.*?)}}")
	if err != nil {
		return nil, err
	}

	// convert file content to string once
	fileContentString := string(fileContent)

	// replace variables
	variables := variableRegex.FindAllString(fileContentString, -1)
	for _, envVariable := range variables {
		envKeyName := strings.ReplaceAll(envVariable, "{{", "")
		envKeyName = strings.ReplaceAll(envKeyName, "}}", "")

		if value, ok := event.env[envKeyName]; ok {
			fileContentString = strings.ReplaceAll(fileContentString, envVariable, value)
		} else {
			if osValue := os.Getenv(envKeyName); osValue != "" {
				fileContentString = strings.ReplaceAll(fileContentString, envVariable, osValue)
			} else {
				fileContentString = strings.ReplaceAll(fileContentString, envVariable, "")
			}
		}
	}

	// convert file content to bytes once
	return []byte(fileContentString), nil
}
