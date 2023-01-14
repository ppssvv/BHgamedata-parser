package main

import (
	"dataparse/internal/animegame"
	"dataparse/internal/decode"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

func main() {
	showUI()
}

var mainMenuOptions = map[string]func(){
	"Parse TextMap":      uiTextMap,
	"Parse DialogueData": uiDialogueData,
	"Parse DormData":     uiDorm,
	"enter file manually and try to guess type": uiManualMode,
	"Decode file": uiDecode,
}

func showUI() {
	keys := make([]string, 0, len(mainMenuOptions))
	for k := range mainMenuOptions {
		keys = append(keys, k)
	}

	choice, err := pterm.DefaultInteractiveSelect.WithOptions(keys).Show("Make a choice")
	uiChechError(err)

	mainMenuOptions[choice]()

	pterm.DefaultInteractiveConfirm.Show("press any key to close the window..")
}

func uiTextMap() {
	options := map[string]string{
		"2578607515 - cn": "2578607515",
		"2578607537 - de": "2578607537",
		"2578607577 - en": "2578607577",
		"2578607612 - fr": "2578607612",
	}

	uiProcessMulti(options, animegame.ProcessTextMap)
}

func uiDialogueData() {
	options := map[string]string{
		"816421621 - cn": "816421621",
		"816421643 - de": "816421643",
		"816421683 - en": "816421683",
		"816421718 - fr": "816421718",
	}

	uiProcessMulti(options, animegame.ProcessDialogueData)
}

func uiDorm() {
	options := map[string]string{
		"1435715286 - DormEventSequence": "1435715286",
	}

	uiProcessMulti(options, animegame.ProcessDormEvent)
}

func uiProcessMulti(options map[string]string, f func(string) error) {
	files, err := pterm.DefaultInteractiveMultiselect.WithOptions(mapKeys(options)).
		Show("choose what files to parse")
	uiChechError(err)

	for i, f := range files {
		files[i] = options[f]
	}

	files, err = getFullName("testdata", files)
	if err != nil {
		pterm.Error.Printfln("can't find files in testdata folder: %s", err)
	}

	ProcessBatch(files, f)

	pterm.Success.Println("Everything done!")
}

func uiManualMode() {
	file, err := pterm.DefaultInteractiveTextInput.Show(`enter path to the file`)
	uiChechError(err)

	switch animegame.GetAsset(file).Parser {
	case "dialogueData":
		err = animegame.ProcessDialogueData(file)
	case "textMap":
		err = animegame.ProcessTextMap(file)
	case "dormEvent":
		err = animegame.ProcessDormEvent(file)
	default:
		err = fmt.Errorf("this type is not implemented yet")
	}

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
