package main

import (
	"dataparse/internal/animegame"
	"dataparse/internal/decode"
	"log"
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
	"Parse TextMap",
	"Parse DialogueData",
	"Parse DormData",
	"Enter file manually and try to parse it",
	"Decode file",
	"Exit",
}

func showUI() {
	options := map[string]func(){
		mainMenuOptions[0]: uiTextMap,
		mainMenuOptions[1]: uiDialogueData,
		mainMenuOptions[2]: uiDorm,
		mainMenuOptions[3]: uiManualMode,
		mainMenuOptions[4]: uiDecode,
		mainMenuOptions[5]: func() { os.Exit(0) },
	}

	choice, err := pterm.DefaultInteractiveSelect.WithOptions(mainMenuOptions).Show("Make a choice")
	uiChechError(err)

	options[choice]()

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

func uiDorm() {
	options := []string{
		"1435715286 - DormitoryEventSequence",
		"606639383 - DormitoryFurnitureData",
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
	}

	ProcessBatch(files)

	pterm.Success.Println("Everything done!")
}

func uiManualMode() {
	file, err := pterm.DefaultInteractiveTextInput.Show(`enter path to the file`)
	uiChechError(err)

	err = ProcessFile(file, animegame.GetAsset(file).Parser)

	if err != nil {
		pterm.Error.Printfln("error parsing %s: %s", file, err)
	} else {
		pterm.Success.Println("Everything done!")
	}
}

func uiChechError(err error) {
	if err != nil {
		log.Fatal("could not initialize UI: %w", err)
	}
}

func uiDecode() {
	file, err := pterm.DefaultInteractiveTextInput.Show("enter path to the file - " +
		"relative to 'testdata' folder or absolute path")
	uiChechError(err)

	if file == "" {
		pterm.Error.Printfln("can't find file: %s", file)
	}

	if !filepath.IsAbs(file) {
		file = filepath.Join("testdata", file)
	}

	result, err := decode.Parse(file)
	if err != nil {
		pterm.Error.Printfln("error while decoding file [%s]: %s", file, err)
		return
	}

	outFolder := filepath.Join("result", "decoded")
	if err := os.MkdirAll(outFolder, os.ModePerm); err != nil {
		pterm.Error.Printfln("can't create output folder: %s", err)
		return
	}
	outFile := filepath.Join(outFolder, filepath.Base(file))

	f, err := os.Create(outFile)
	if err != nil {
		pterm.Error.Printfln("can't create output file: %s", err)
		return
	}

	if _, err := f.Write(result); err != nil {
		pterm.Error.Printfln("can't write result to file: %s", err)
		return
	}

	pterm.Success.Printfln("%s - ok!", outFile)
}
