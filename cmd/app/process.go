package main

import (
	"dataparse"
	"dataparse/dump"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

func ProcessAllInFolder(folder string) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		pterm.Error.Printfln("can't read %s folder - %s", folder, err)
		return
	}

	ProcessBatch(converDirEntries(entries))
}

func ProcessBatch(entries []string) {
	total := len(entries)
	pterm.Info.Printf("found %d entries\n", total)

	var fail int
	pbar, err := pterm.DefaultProgressbar.WithTotal(total).Start()
	uiChechError(err)

	for _, e := range entries {
		pbar.UpdateTitle(e)

		if err = ProcessFile(e, dataparse.GetAsset(e).Parser); err != nil {
			pterm.Error.Printfln("\t%s - %s", e, err)
			fail++
		} else {
			pterm.Success.Println(e)
		}

		pbar.Increment()
	}

	pterm.Info.Printf("Done: %d ok, %d failed\n", total-fail, fail)
}

func ProcessFile(in string, obj dump.ReaderWrapper) error {
	result, err := dataparse.ProcessStructNew(in, obj)
	if err != nil {
		return err
	}

	out := filepath.Join("result", fmt.Sprintf("%s.json", dataparse.GetAsset(in).Name))
	os.Mkdir("result", os.ModePerm)

	return os.WriteFile(out, result, os.ModePerm)
}
