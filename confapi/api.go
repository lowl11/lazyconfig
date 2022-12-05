package confapi

import (
	"encoding/json"
	"errors"
	"github.com/lowl11/lazyfile/fileapi"
)

func Read(configObj interface{}, production bool) error {
	if err := initConfigFiles(); err != nil {
		return err
	}

	if !production {
		configurationsChecked, err := checkAllTransformations(configFileDebug, configFileRelease)
		if err != nil {
			return err
		}

		if !configurationsChecked {
			return errors.New("[lazyconfig] Configurations has no the same keys")
		}
	}

	var configFilePath string
	if !production {
		configFilePath = configFileDebug
	} else {
		configFilePath = configFileRelease
	}

	// reading config file
	configFileContent, err := fileapi.Read(configFilePath)
	if err != nil {
		return err
	}

	// converting it to key-value map
	configMap, err := convertJSONtoMap(configFileContent)
	if err != nil {
		return err
	}

	// application config JSON
	appConfigJSON, err := convertObjectToJSON(configObj)
	if err != nil {
		return err
	}

	// application config map
	appConfigMap, err := convertJSONtoMap(appConfigJSON)
	if err != nil {
		return err
	}

	// joining two maps
	for key, value := range configMap {
		appConfigMap[key] = value
	}

	// read config parse to json
	readConfigJSON, err := convertMapToJson(appConfigMap)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(readConfigJSON, &configObj); err != nil {
		return err
	}

	return nil
}

func ReadCertain(configObj interface{}, filePath string) error {
	// reading config file
	configFileContent, err := fileapi.Read(filePath)
	if err != nil {
		return err
	}

	// converting it to key-value map
	configMap, err := convertJSONtoMap(configFileContent)
	if err != nil {
		return err
	}

	// application config JSON
	appConfigJSON, err := convertObjectToJSON(configObj)
	if err != nil {
		return err
	}

	// application config map
	appConfigMap, err := convertJSONtoMap(appConfigJSON)
	if err != nil {
		return err
	}

	// joining two maps
	for key, value := range configMap {
		appConfigMap[key] = value
	}

	// read config parse to json
	readConfigJSON, err := convertMapToJson(appConfigMap)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(readConfigJSON, &configObj); err != nil {
		return err
	}

	return nil
}
