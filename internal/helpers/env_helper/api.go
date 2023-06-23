package env_helper

import (
	"errors"
	"github.com/lowl11/lazyfile/fileapi"
	"os"
	"regexp"
	"strings"
)

func Read(fileName string) (map[string]string, error) {
	envVariables := make(map[string]string)

	if !fileapi.Exist(fileName) {
		return envVariables, nil
	}

	fileContent, err := fileapi.Read(fileName)
	if err != nil {
		return envVariables, err
	}

	combinations := strings.Split(string(fileContent), "\n")

	for _, combo := range combinations {
		dividedCombo := strings.Split(combo, "=")
		if len(dividedCombo) == 1 {
			continue
		}

		envVariables[dividedCombo[0]] = dividedCombo[1]
	}

	return envVariables, nil
}

func ReplaceVariables(fileContent []byte, envVariables map[string]string) ([]byte, error) {
	if envVariables == nil {
		return nil, errors.New("env variables map is NULL")
	}

	if fileContent == nil {
		return nil, errors.New("file content is NULL")
	}

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

		if value, ok := envVariables[envKeyName]; ok {
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
