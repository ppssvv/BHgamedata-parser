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

	var fail, missing int
	pbar, err := pterm.DefaultProgressbar.WithTotal(total).Start()
	uiChechError(err)

	exist := getExisting()

	for _, e := range entries {
		pbar.UpdateTitle(e)

		ass := dataparse.GetAsset(e)

		if _, ok := exist[ass.Name]; ok {
			pterm.Success.Printfln("%s already parsed as %s, skipping", getShortname(e), ass.Name)
			pbar.Increment()
			continue
		}

		if ass.Parser.GetData() == nil {
			pterm.Error.Printfln("\t%s not supported yet", e)
			missing++
			pbar.Increment()
			continue
		}

		if err = ProcessFile(e, ass.Parser); err != nil {
			pterm.Error.Printfln("\t%s - %s", e, err)
			fail++
		} else {
			pterm.Success.Println(e)
		}

		pbar.Increment()
	}

	pterm.Info.Printf("Done: %d ok, %d failed, %d missing\n", total-fail-missing, fail, missing)
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
