package lazyconfig

import "encoding/json"

const (
	configFolderPath  = "config"
	configFileDebug   = "config/config_debug.json"
	configFileRelease = "config/config_release.json"
)

func initConfigFiles() error {
	if !isFolderExist(configFolderPath) {
		if err := createFolder(configFolderPath); err != nil {
			return err
		}
	}

	if err := createConfigFile(configFileDebug, configFileRelease); err != nil {
		return err
	}

	return nil
}

func checkAllTransformations(paths ...string) (bool, error) {
	configMapList := make([]map[string]interface{}, 0)
	for _, path := range paths {
		configMap, err := getConfigurationMap(path)
		if err != nil {
			return false, err
		}

		configMapList = append(configMapList, configMap)
	}

	configMapKeyList := make([]string, 0)
	for _, configMap := range configMapList {
		var keys string
		for key, _ := range configMap {
			keys += key
		}
		configMapKeyList = append(configMapKeyList, keys)
	}

	if len(configMapKeyList) > 0 {
		previousKeyList := configMapKeyList[0]
		for _, keyList := range configMapKeyList {
			if previousKeyList != keyList {
				return false, nil
			}
			previousKeyList = keyList
		}
	}

	return true, nil
}

func getConfigurationMap(path string) (map[string]interface{}, error) {
	fileContent, err := readFile(path)
	if err != nil {
		return nil, err
	}

	configMap := make(map[string]interface{}, 0)
	if err := json.Unmarshal(fileContent, &configMap); err != nil {
		return nil, err
	}

	return configMap, nil
}

func createConfigFile(paths ...string) error {
	for _, path := range paths {
		if !isFileExist(path) {
			if err := createFile(path); err != nil {
				return err
			}
		}
	}
	return nil
}
