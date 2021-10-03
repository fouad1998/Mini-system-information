package storage

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fouad1998/min-system-information/env"
)

func (s *Storage) extractEntry(line string) (*Entry, error) {
	arr := strings.Split(line, env.Setting.Seperator)
	if len(arr) != 10 {
		log.Println(line)
		return nil, fmt.Errorf("Couldn't get the required number of argument from given entry")
	}

	createdAt, _ := strconv.Atoi(arr[4])
	modifiedAt, _ := strconv.Atoi(arr[5])
	return &Entry{
		ID:          arr[0],
		Name:        arr[1],
		Description: arr[2],
		Owner:       arr[3],
		CreatedAt:   int64(createdAt),
		ModifiedAt:  int64(modifiedAt),
		Filename:    arr[6],
		Extension:   arr[7],
		Path:        arr[8],
		MimeType:    arr[9],
	}, nil
}
