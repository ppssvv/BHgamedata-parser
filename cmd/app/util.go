package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func getFullName(folder string, filenames []string) ([]string, error) {
	results := []string{}

	entries, err := os.ReadDir("testdata")
	if err != nil {
		return nil, fmt.Errorf("can't read dir testdata: %w", err)
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		shortname, _, _ := strings.Cut(e.Name(), "_")

		ok := slices.Contains(filenames, shortname)
		if ok {
			results = append(results, e.Name())
		}
	}

	return results, nil
}

func converDirEntries(src []os.DirEntry) []string {
	result := []string{}

	for _, e := range src {
		if e.IsDir() {
			continue
		}
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}
		result = append(result, e.Name())
	}
	return result
}

func uiChechError(err error) {
	if err != nil {
		log.Fatal("could not initialize UI: %w", err)
	}
}

func getShortname(s string) string {
	res, _, _ := strings.Cut(s, "_")
	return res
}

func getExisting() map[string]interface{} {
	src, err := os.ReadDir("result")
	if err != nil {
		log.Fatal("can't read result folder")
	}

	result := map[string]interface{}{}

	for _, e := range src {
		if e.IsDir() {
			continue
		}
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}

		result[strings.TrimSuffix(e.Name(), ".json")] = nil
	}

	return result
}
