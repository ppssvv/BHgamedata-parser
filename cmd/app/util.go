package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	ProcessBatch(convertDirEntries(name, entries), f)
	return nil
}

func getFullName(folder string, filenames []string) ([]string, error) {
	results := []string{}

	entries, err := os.ReadDir("testdata")
	if err != nil {
		return nil, fmt.Errorf("can't read dir testdata: %w", err)
	}

	filenamesMap := convertToMap(filenames)

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		shortname, _, _ := strings.Cut(e.Name(), "_")

		_, ok := filenamesMap[shortname]
		if ok {
			results = append(results, e.Name())
		}
	}

	return results, nil
}

func convertToMap(obj []string) map[string]struct{} {
	result := map[string]struct{}{}

	for _, o := range obj {
		result[o] = struct{}{}
	}

	return result
}

func mapKeys(s map[string]string) []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
