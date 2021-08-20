package lazyconfig

import (
	"encoding/json"
	"errors"
)

func ReadConfig(configObj interface{}, debug bool) error {
	if err := initConfigFiles(); err != nil {
		return err
	}

	configurationsChecked, err := checkAllTransformations(configFileDebug, configFileRelease)
	if err != nil {
		return err
	}

	if !configurationsChecked {
		return errors.New("[lazyconfig] Configurations has no the same keys")
	}

	var configFilePath string
	if debug {
		configFilePath = configFileDebug
	} else {
		configFilePath = configFileRelease
	}

	// reading config file
	configFileContent, err := readFile(configFilePath)
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

	if err := json.Unmarshal(readConfigJSON, &configObj); err != nil {
		return err
	}

	return nil
}

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
