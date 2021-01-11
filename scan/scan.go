package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

// Scan scans for files
func Scan() ([]string, error) {
	files := []string{}

	fmt.Println("Scanning for markdown files...\n")

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, nil
	}

	fmt.Printf("Found %v files\n", len(files))

	return files, nil
}
