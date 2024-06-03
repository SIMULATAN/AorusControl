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
		SetRows(4, 2, -1).
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

	statusRow := tview.NewFlex()

	fan0Status:= tview.NewTextView().
			SetTextAlign(tview.AlignCenter)
	fan0 := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
			SetText("Fan0").
			SetTextAlign(tview.AlignCenter),
			0, 1, false,
		).
		AddItem(fan0Status, 0, 1, false)

	fan1Status:= tview.NewTextView().
			SetTextAlign(tview.AlignCenter)
	fan1 := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
			SetText("Fan1").
			SetTextAlign(tview.AlignCenter),
			0, 1, false,
		).
		AddItem(fan1Status, 0, 1, false)

	statusRow.AddItem(fan0, 0, 1, false)
	statusRow.AddItem(fan1, 0, 1, false)

	currentMode, currentModeIndex := GetCurrentMode(ec)
	log := LogView {
		tview.NewTextView().
			SetDynamicColors(true),
	}
	currentModeView := tview.NewTextView().
		SetText(currentMode.name).
		SetTextAlign(tview.AlignRight)

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

	modeRow := tview.NewFlex()
	modeRow.AddItem(dropdown, 0, 1, true)
	modeRow.AddItem(currentModeView, 0, 1, false)
	root.AddItem(statusRow, 0, 0, 1, 2, 0, 0, false)
	root.AddItem(modeRow, 1, 0, 1, 2, 0, 0, false)
	root.AddItem(log, 2, 0, 1, 2, 0, 0, false)

	go updateStatus(log, ec, app, fan0Status, fan1Status)

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
