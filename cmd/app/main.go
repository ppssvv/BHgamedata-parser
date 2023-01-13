package main

import (
	"dataparse/internal/animegame"
	"fmt"
	"log"

	"github.com/pterm/pterm"
)

func main() {
	showUI()
}

var mainMenuOptions = []string{
	"Parse TextMap",
	"Parse DialogueData",
	"enter file manually and try to guess type",
}

func showUI() {
	choice, err := pterm.DefaultInteractiveSelect.WithOptions(mainMenuOptions).Show("Make a choice")
	uiChechError(err)

	switch choice {
	case mainMenuOptions[0]:
		uiTextMap()
	case mainMenuOptions[1]:
		uiDialogueData()
	default:
		uiManualMode()
	}

	pterm.DefaultInteractiveConfirm.Show("press any key to close the window..")
}

func uiTextMap() {
	file, err := pterm.DefaultInteractiveTextInput.Show("enter path to the file or leave empty" +
		"to parse everything inside 'testdata/TextMap' folder")
	uiChechError(err)

	if file == "" {
		ProcessTextMapAll()
		return
	}

	if err := animegame.ProcessTextMap(file); err != nil {
		pterm.Error.Printfln("error parsing %s: %s", file, err)
	}

	pterm.Success.Println("Everything done!")
}

func uiDialogueData() {
	file, err := pterm.DefaultInteractiveTextInput.Show("enter path to the file or leave empty" +
		"to parse everything inside 'testdata/DialogueData' folder")
	uiChechError(err)

	if file == "" {
		ProcessDialogueDataAll()
		return
	}

	if err := animegame.ProcessDialogueData(file); err != nil {
		pterm.Error.Printfln("error parsing %s: %s", file, err)
	}
}

func uiManualMode() {
	file, err := pterm.DefaultInteractiveTextInput.Show(`enter path to the file`)
	uiChechError(err)

	switch animegame.GetAsset(file).Parser {
	case "dialogueData":
		err = animegame.ProcessDialogueData(file)
	case "textMap":
		err = animegame.ProcessTextMap(file)
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
