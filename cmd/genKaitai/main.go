// package genKantai is meant to simplify work
// with kaitai struct generation.
//
// Simply call genKaitai.exe to generate code for ALL
// templates in `./templates/` folder.
//
// Alternatively, you can provide an argument with path to file,
// e.g. `genKaitai.exe TextMap.ksy`
// Those can be more than one
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		entries, err := os.ReadDir("templates")
		if err != nil {
			log.Fatalf("can't read dir templates: %s\n", err)
		}
		args = convertDirEntries("", entries)
	}

	for _, e := range args {
		if !filepath.IsAbs(e) {
			e = filepath.Join("templates", e)
		}
		Generate(e)
	}
}

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

func Generate(f string) {
	fmt.Printf("Generating code for %s..\n", f)

	cmd := exec.Command("kaitai-struct-compiler",
		"-t", "go",
		"--go-package", "ksygen",
		"--outdir", "internal",
		f,
	)

	if err := cmd.Run(); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Println("\tDone!")
}
