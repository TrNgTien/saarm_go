package utilities

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"saarm/pkg/common"
)

func GetFilePath(pathDir, fileName string) string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("[GetFilePath] Failed to read file: %v", err)
	}

	return filepath.Join(dir, pathDir+fileName)
}

func removeContents(path, excludePattern string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Check if file matches exclude pattern
			if matched, err := filepath.Match(excludePattern, info.Name()); err != nil {
				return err
			} else if matched {
				fmt.Println("Skipping:", path)
				return nil
			}

			return os.RemoveAll(path)
		}

		return nil
	})

	return err
}

func RemoveAllAssets() error {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("[RemoveAllAssets] Failed to remove files: %v", err)
	}

	p := filepath.Join(dir, common.WATER_METER_PATH)

	excludePattern := ".keep"

	if err := removeContents(p, excludePattern); err != nil {
		return err
	}

	return nil
}
