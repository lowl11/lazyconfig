package lazyconfig

import "os"

func createFolder(path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}

	return nil
}

func isFolderExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
