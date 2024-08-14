package storage

import (
	"errors"

	"github.com/Nyanpasa-dev/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(name string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(name, blob)

	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetFileById(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileID]

	if !ok {
		return nil, errors.New("file not found")
	}

	return foundFile, nil
}
