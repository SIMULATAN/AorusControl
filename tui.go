package main

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TuiMain(ec os.File) {
	app := tview.NewApplication()
	root := tview.NewGrid().
		SetBorders(true).
		SetRows(1, 3, -1).
		SetColumns(-1)

	root.SetTitle("AorusControl")

	root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
			os.Exit(0)
		case 'j', 'l':
			return nil
		}
		return event
	})

	currentMode, currentModeIndex := GetCurrentMode(ec)
	log := LogView {
		tview.NewTextView().
			SetDynamicColors(true),
	}
	currentModeView := tview.NewTextView().
		SetText(currentMode.name)

	dropdown := tview.NewDropDown().
		SetLabel("Select a mode: ").
		SetOptions(GetModeNames(), nil).
		SetCurrentOption(currentModeIndex)
	// set *after* initializing the initial option to avoid writing
	dropdown.SetSelectedFunc(func(text string, index int) {
			err := SetModeByIndex(ec, int64(index))
			if err != nil {
				log.Error(err)
			} else {
				log.Success("Set mode to", text)
			}
			// re-fetch the mode in order to make sure we correctly wrote
			currentMode, _ := GetCurrentMode(ec)
			currentModeView.SetText(currentMode.name)
		})

	root.AddItem(currentModeView, 0, 0, 1, 2, 0, 0, false)
	root.AddItem(dropdown, 1, 0, 1, 2, 0, 0, false)
	root.AddItem(log, 2, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(root, true).SetFocus(dropdown).Run(); err != nil {
		panic(err)
	}
}

type LogView struct {
	*tview.TextView
}

func (view *LogView) Error(err error) {
	view.Write([]byte("[red]" + err.Error() + "\n"))
}

func (view *LogView) Success(msg ...string) {
	view.Write([]byte("[green]" + strings.Join(msg, " ") + "\n"))
}
