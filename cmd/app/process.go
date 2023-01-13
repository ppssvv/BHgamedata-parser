package main

import (
	"dataparse/internal/animegame"

	"github.com/pterm/pterm"
)

func ProcessTextMapAll() (xerr error) {
	return ProcessAllInFolder("TextMap", animegame.ProcessTextMap)
}

func ProcessDialogueDataAll() (xerr error) {
	return ProcessAllInFolder("DialogueData", animegame.ProcessDialogueData)
}

func ProcessBatch(entries []string, f func(string) error) (xerr error) {
	total := len(entries)
	pterm.Info.Printf("found %d entries\n", total)

	var fail int
	pbar, err := pterm.DefaultProgressbar.WithTotal(total).Start()
	uiChechError(err)

	for _, e := range entries {
		pbar.UpdateTitle(e)
		if err := f(e); err != nil {
			pterm.Error.Printfln("\t%s - %s", e, err)
			fail++
		}

		pterm.Success.Println(e)
		pbar.Increment()
	}

	pterm.Info.Printf("Done: %d ok, %d failed\n", total-fail, fail)

	return xerr
}
