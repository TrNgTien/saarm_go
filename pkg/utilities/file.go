package utilities

import (
	"log"
	"os"
	"path/filepath"
)

func GetFilePath(path string) string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return filepath.Join(dir, path)
}
