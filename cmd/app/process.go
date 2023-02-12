package main

import (
	"dataparse"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"go.uber.org/multierr"
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

	exist := getExisting("result")

	for _, e := range entries {
		pbar.UpdateTitle(e)

		short := getShortname(e)

		errr := multierr.Errors(ProcessFile(e, exist))
		if errr == nil {
			pterm.Success.Println(e)
			pbar.Increment()
			continue
		}

		for _, procerr := range errr {
			switch er := procerr.(type) {
			case nil:
				pterm.Success.Println(short)
			case ErrAlreadyParsed:
				pterm.Success.Println(short, er.Error())
			case dataparse.ErrNotSupported:
				pterm.Error.Println(short, er.Error())
				missing++
			default:
				pterm.Error.Println(er.Error())
				fail++
			}
		}

		pbar.Increment()
	}

	pterm.Info.Printf("Done: %d ok, %d failed, %d missing\n", total-fail-missing, fail, missing)
}

type ErrAlreadyParsed string

func (e ErrAlreadyParsed) Error() string {
	return fmt.Sprintf("already parsed as %s", string(e))
}

func ProcessFile(entry string, exist map[string]interface{}) error {
	assets, ok := dataparse.GetAsset(entry)
	if !ok {
		return dataparse.ErrNotSupported("")
	}

	subfolder := ""

	if len(assets) > 1 {
		// make additional folder for multiparser
		subfolder = getShortname(entry)
	} else {
		// check early if no need to parse
		if _, ok := exist[assets[0].Name]; ok {
			return ErrAlreadyParsed(assets[0].Name)
		}
	}

	output := filepath.Join("result", subfolder)
	if err := os.MkdirAll(output, os.ModePerm); err != nil {
		pterm.Fatal.Printfln("can't create result folder: %s", err)
	}

	var errr error

	dec := dataparse.GetTestData(entry)

	for _, ass := range assets {
		dec.SetPosition(0)
		name := ass.Name
		if subfolder != "" {
			name = subfolder + "\\" + ass.Name
		}
		if _, ok := exist[name]; ok {
			multierr.AppendInto(&errr, ErrAlreadyParsed(ass.Name))
			continue
		}

		if ass.Parser == nil {
			multierr.AppendInto(&errr, dataparse.ErrNotSupported(""))
			continue
		}

		if multierr.AppendInto(&errr, dec.Decode(ass.Parser)) {
			continue
		}

		data, jsonErr := json.MarshalIndent(ass.Parser, "", "  ")
		if multierr.AppendInto(&errr, jsonErr) {
			continue
		}

		outPath := filepath.Join(output, fmt.Sprintf("%s.json", ass.Name))
		multierr.AppendInto(&errr, os.WriteFile(outPath, data, os.ModePerm))

	}

	return errr
}
