package file

import (
	"os"
)

func Create(path string, fileName string, data []byte) error {
	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}

	filePath := path + "/" + fileName

	if Exists(filePath) {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	return nil
}
