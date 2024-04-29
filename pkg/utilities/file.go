package utilities

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func GetFilePath(path string) string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("[GetFilePath] Failed to read file: %v", err)
	}

	return filepath.Join(dir, path)
}

func CreateDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)

	if err == nil {
		return nil
	}

	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)

		if err != nil {
			return err
		}

		if !info.IsDir() {
			return errors.New("[CreateDir] | Path exists but is not a directory")
		}

		return nil
	}

	return err
}
