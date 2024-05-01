package utilities

import (
	"log"
	"os"
	"path/filepath"
)

func GetFilePath(pathDir, fileName string) string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("[GetFilePath] Failed to read file: %v", err)
	}

	return filepath.Join(dir, pathDir+fileName)
}
