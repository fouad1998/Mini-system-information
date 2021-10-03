package storage

import (
	"os"

	"github.com/fouad1998/min-system-information/env"
)

func (s *Storage) storagefilepath() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := workingDirectory + PATH_SEPERATOR + env.Setting.StorageDirectory + PATH_SEPERATOR + env.Setting.StorageFile
	return path, nil
}
