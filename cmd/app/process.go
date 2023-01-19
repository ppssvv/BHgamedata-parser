package main

import (
	"dataparse"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

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
		}

		pterm.Success.Println(e)
		pbar.Increment()
	}

	pterm.Info.Printf("Done: %d ok, %d failed\n", total-fail, fail)
}

func ProcessFile(in string, f func(string) ([]byte, error)) error {
	if !filepath.IsAbs(in) {
		in = filepath.Join("testdata", in)
	}

	result, err := f(in)
	if err != nil {
		return err
	}

	out := filepath.Join("result", fmt.Sprintf("%s.json", dataparse.GetAsset(in).Name))

	return os.WriteFile(out, result, os.ModePerm)
}
