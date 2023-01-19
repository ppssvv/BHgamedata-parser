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
		result = append(result, e.Name())
	}
	return result
}

func uiChechError(err error) {
	if err != nil {
		log.Fatal("could not initialize UI: %w", err)
	}
}
