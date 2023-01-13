package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	Generate(filepath.Join("templates", "DialogueData.ksy"))
	Generate(filepath.Join("templates", "TextMap.ksy"))
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
