package confapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
)

const (
	configFolderPath     = "config"
	configFileTest       = "config/test.json"
	configFileProduction = "config/production.json"
)

func convertObjectToJSON(obj interface{}) ([]byte, error) {
	objJSON, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return objJSON, nil
}

func convertJSONtoMap(jsonContent []byte) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{}, 0)
	if err := json.Unmarshal(jsonContent, &jsonMap); err != nil {
		return nil, err
	}

	return jsonMap, nil
}

func convertMapToJson(configMap map[string]interface{}) ([]byte, error) {
	return json.Marshal(configMap)
}

func initConfigFiles() error {
	if folderapi.NotExist(configFolderPath) {
		if err := folderapi.Create("./", configFolderPath); err != nil {
			return err
		}
	}

	if err := createConfigFile(configFileTest, configFileProduction); err != nil {
		return err
	}

	return nil
}

func checkAllTransformations(paths ...string) (bool, error) {
	configMapList := make([]map[string]interface{}, 0)
	for _, path := range paths {
		configMap, err := getConfigurationMap(path)
		if err != nil {
			return false, errors.New(path + " file read error: " + err.Error())
		}

		configMapList = append(configMapList, configMap)
	}

	configMapKeyList := make([][]string, 0)
	for _, configMap := range configMapList {
		keysList := make([]string, 0)
		for key := range configMap {
			keysList = append(keysList, key)
		}
		configMapKeyList = append(configMapKeyList, keysList)
	}

	if len(configMapKeyList) > 0 {
		firstMapKeyList := configMapKeyList[0]
		for _, keyList := range configMapKeyList {
			for _, key := range firstMapKeyList {
				if !contains(keyList, key) {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

func contains(array []string, searchValue string) bool {
	for _, value := range array {
		if value == searchValue {
			return true
		}
	}

	return false
}

func getConfigurationMap(path string) (map[string]any, error) {
	fileContent, err := fileapi.Read(path)
	if err != nil {
		return nil, err
	}

	configMap := make(map[string]any)
	if err = json.Unmarshal(fileContent, &configMap); err != nil {
		fmt.Println("still here")
		return nil, err
	}

	return configMap, nil
}

func createConfigFile(paths ...string) error {
	for _, path := range paths {
		if !fileapi.Exist(path) {
			if err := fileapi.Create(path, nil); err != nil {
				return err
			}
		}
	}

	return nil
}
