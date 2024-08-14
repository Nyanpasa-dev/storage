package main

import (
	"fmt"
	"log"

	"github.com/Nyanpasa-dev/storage/internal/storage"
)

func main() {
	storageService := storage.NewStorage()

	file, err := storageService.Upload("test.txt", []byte("Hello, World!"))

	if err != nil {
		log.Fatal(err)
	}
	returnedFile, err := storageService.GetFileById(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Uploaded file:", file)
	fmt.Println("Returned file:", returnedFile)
}
