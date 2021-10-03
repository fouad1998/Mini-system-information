package storage

import "os"

func (s *Storage) Init() error {
	storagefilepath, err := s.storagefilepath()
	if err != nil {
		return err
	}

	storagedirectory, err := s.storageDirectory()
	if err != nil {
		return err
	}

	err = os.MkdirAll(storagedirectory, 0777)
	if err != nil {
		return err
	}

	_, err = os.Stat(storagefilepath)
	if os.IsNotExist(err) {
		file, err := os.Create(storagefilepath)
		if err != nil {
			return err
		}

		file.Close()
	}

	return nil
}
