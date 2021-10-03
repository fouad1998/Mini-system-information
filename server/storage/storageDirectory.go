package storage

import (
	"os"

	"github.com/fouad1998/min-system-information/env"
)

func (s *Storage) storageDirectory() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := workingDirectory + PATH_SEPERATOR + env.Setting.StorageDirectory
	return path, nil
}
