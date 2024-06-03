package main

import (
	"os"
	"strconv"
	"time"

	"github.com/rivo/tview"
)

func updateStatus(
	log LogView,
	ec os.File,
	app *tview.Application,
	fan0 *tview.TextView,
	fan1 *tview.TextView,
) {
	for {
		fan0Rpm, error := ReadInt16(ec, OFFSET_FAN0_RPM)
		if error != nil {
			log.Error(error)
		}
		fan1Rpm, error := ReadInt16(ec, OFFSET_FAN1_RPM)
		if error != nil {
			log.Error(error)
		}
		app.QueueUpdateDraw(func() {
			fan0.SetText(strconv.FormatUint(uint64(fan0Rpm), 10) + "RPM")
			fan1.SetText(strconv.FormatUint(uint64(fan1Rpm), 10) + "RPM")
		})
		time.Sleep(2500)
	}
}
