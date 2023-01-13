package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// pathPrefix optional subfolder
func convertDirEntries(pathPrefix string, source []os.DirEntry) []string {
	result := []string{}

	for _, e := range source {
		if e.IsDir() {
			continue
		}

		result = append(result, filepath.Join(pathPrefix, e.Name()))
	}

	return result
}

func ProcessAllInFolder(name string, f func(string) error) error {
	entries, err := os.ReadDir(filepath.Join("testdata", name))
	if err != nil {
		return fmt.Errorf("can't read dir testdata/%s: %w", name, err)
	}

	return ProcessBatch(convertDirEntries(name, entries), f)
}
