package main

import (
	"dataparse"
	"fmt"
	"os"
	"strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

// app.exe input output [decode-only]
func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		interactiveMode()
		os.Exit(0)
	}

	if len(args) >= 3 {
		dec := args[2]
		if dec == "decode-only" {
			if err := decodeInput(args[0], args[1]); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done!")
			os.Exit(0)
		}
	}

	err := dataparse.Parse(args[0], args[1])
	if err != nil {
		exitWithWait(err.Error())
	}

	// in cli mode not need to wait for input
	fmt.Println("Done!")
}

func interactiveMode() {
	results, err := cfdutil.ShowOpenMultipleFilesDialog(cfd.DialogConfig{
		Title: "Choose files to parse",
		Role:  "GameDataParserFiles",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Unity files",
				Pattern:     "*.unity3d",
			},
			{
				DisplayName: "All files",
				Pattern:     "*.*",
			},
		},
	})

	// if err == cfd.ErrorCancelled {
	if err != nil { // on any error
		results = simpleMode()
	}

	if results == nil || len(results) < 1 {
		exitWithWait("bad input")
	}

	for _, f := range results {
		dataparse.ParseFile(f, "out")
	}

	exitWithWait("Done!")
}

func simpleMode() []string {
	fmt.Print("Enter input manually (file or folder): ")
	var in string
	if _, err := fmt.Scanln(&in); err != nil {
		fmt.Println("try again")
		simpleMode()
	}

	in = strings.Trim(in, `"`)

	info, err := os.Stat(in)
	if err != nil {
		return nil
	}

	if info.IsDir() {
		tmp, err := dataparse.GetAllFiles(in)
		if err != nil {
			return nil
		}

		return tmp
	}

	return []string{in}
}
