package lazyconfig

import (
	"io/ioutil"
	"os"
)

func isFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func createFile(path string) error {
	if _, err := os.Create(path); err != nil {
		return err
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
