package main

import (
	"dataparse/internal/decode"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	for {
		showUI()
	}
}

var mainMenuOptions = []string{
	"Parse file",
	"Decode file",
	"Exit",
}

func showUI() {
	options := map[string]func(){
		mainMenuOptions[0]: uiParse,
		mainMenuOptions[1]: uiDecode,
		mainMenuOptions[2]: func() { os.Exit(0) },
	}

	choice, err := pterm.DefaultInteractiveSelect.WithOptions(mainMenuOptions).Show("Make a choice")
	uiChechError(err)

	options[choice]()

	pterm.Println()
}

func uiParse() {
	options := []string{
		"Everything in 'testdata' folder",
		"Enter file manually",
		"TextMaps",
		"DialogueData",
	}

	choice, err := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Choose what to parse")
	uiChechError(err)

	routes := map[string]func(){
		options[0]: func() { ProcessAllInFolder("testdata") },
		options[1]: uiManualMode,
		options[2]: uiTextMap,
		options[3]: uiDialogueData,
	}

	routes[choice]()

	pterm.Println()
}

func uiTextMap() {
	options := []string{
		"2578607515 - cn",
		"2578607537 - de",
		"2578607577 - en",
		"2578607612 - fr",
	}

	uiProcessMulti(options)
}

func uiDialogueData() {
	options := []string{
		"816421621 - cn",
		"816421643 - de",
		"816421683 - en",
		"816421718 - fr",
	}

	uiProcessMulti(options)
}

func uiProcessMulti(options []string) {
	files, err := pterm.DefaultInteractiveMultiselect.WithOptions(options).
		Show("choose what files to parse")
	uiChechError(err)

	for i := 0; i < len(files); i++ {
		files[i], _, _ = strings.Cut(files[i], " - ")
	}

	files, err = getFullName("testdata", files)
	if err != nil {
		pterm.Error.Printfln("can't find files in testdata folder: %s", err)
		return
	}

	ProcessBatch(files)

	pterm.Success.Println("Everything done!")
}

func uiManualMode() {
	file, err := pterm.DefaultInteractiveTextInput.Show(`enter path to the file`)
	uiChechError(err)

	err = ProcessFile(file, getExisting("result"))

	if err != nil {
		pterm.Error.Printfln("error parsing %s: %s", file, err)
	} else {
		pterm.Success.Println("Everything done!")
	}
}

func uiDecode() {
	options := []string{
		"Everything in 'testdata' folder",
		"Enter file manually",
	}

	choice, err := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Choose what to decode")
	uiChechError(err)

	files := []string{}
	if choice == options[1] {
		file, err := pterm.DefaultInteractiveTextInput.Show("enter path to the file - " +
			"relative to 'testdata' folder or absolute path")
		uiChechError(err)

		if file == "" {
			pterm.Error.Printfln("can't find file: %s", file)
			return
		}

		if !filepath.IsAbs(file) {
			file = filepath.Join("testdata", file)
		}

		files = append(files, file)
	} else {
		entries, err := os.ReadDir("testdata")
		if err != nil {
			pterm.Fatal.Println("can't read testdata folder: ", err)
		}
		files = converDirEntries(entries)
	}

	decodeMulti(files)
}

func decodeMulti(files []string) {
	total := len(files)
	fail := 0

	outFolder := filepath.Join("result", "decoded")
	if err := os.MkdirAll(outFolder, os.ModePerm); err != nil {
		pterm.Error.Println("can't create output folder: ", err)
		return
	}

	pbar, err := pterm.DefaultProgressbar.WithTotal(total).WithRemoveWhenDone(true).Start()
	uiChechError(err)

	for _, file := range files {
		pbar.UpdateTitle(file)

		if !filepath.IsAbs(file) {
			file = filepath.Join("testdata", file)
		}

		result, err := decode.Parse(file)
		if err != nil {
			pterm.Error.Printfln("[%s]: %s", file, err)
			fail++
			pbar.Increment()
			continue
		}

		outFile := filepath.Join(outFolder, filepath.Base(file))
		if os.WriteFile(outFile, result, os.ModePerm); err != nil {
			pterm.Error.Printfln("can't write result to file: %s", err)
			fail++
			pbar.Increment()
			continue
		}

		pbar.Increment()
		pterm.Success.Printfln("%s - ok!", outFile)
	}

	pterm.Info.Printf("Done: %d ok, %d failed\n", total-fail, fail)
}
