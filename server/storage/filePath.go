package storage

func (s *Storage) filePath(filename string) (string, error) {
	directory, err := s.storageDirectory()
	if err != nil {
		return "", err
	}

	return directory + PATH_SEPERATOR + filename, nil
}
