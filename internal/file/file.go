package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	newFile := &File{
		ID:   fileID,
		Name: name,
		Data: blob,
	}

	return newFile, nil
}

func (f *File) String() string {
	return fmt.Sprintf(("ID: %d, Name: %s, Data: %s"), f.ID, f.Name, f.Data)
}
