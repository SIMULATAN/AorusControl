package main

import (
	"os"

	"github.com/elliotchance/orderedmap/v2"
)

type Mode struct {
	name   string
	offset int64
	bit    int8
}

var MODE_QUIET = Mode{
	name:   "Quiet",
	offset: OFFSET_MODE_QUIET,
	bit:    BIT_MODE_QUIET,
}
var MODE_GAMING = Mode{
	name:   "Gaming",
	offset: OFFSET_MODE_GAMING,
	bit:    BIT_MODE_GAMING,
}
var MODE_DEEP_CONTROL = Mode{
	name:   "Deep Control",
	offset: OFFSET_MODE_DEEP_CONTROL,
	bit:    BIT_MODE_DEEP_CONTROL,
}
var MODE_AUTO = Mode{
	name:   "Auto Max",
	offset: OFFSET_MODE_AUTO,
	bit:    BIT_MODE_AUTO,
}
var MODE_FIXED = Mode{
	name:   "Fixed",
	offset: OFFSET_MODE_FIXED,
	bit:    BIT_MODE_FIXED,
}
var MODE_NORMAL = Mode{
	name: "Normal",
}

var modes = orderedmap.NewOrderedMap[string, Mode]()

func InitModes() {
	modes.Set("quiet", MODE_QUIET)
	modes.Set("gaming", MODE_GAMING)
	modes.Set("deep_control", MODE_DEEP_CONTROL)
	modes.Set("auto", MODE_AUTO)
	modes.Set("fixed", MODE_FIXED)
	modes.Set("normal", MODE_NORMAL)
}

func GetModeNames() []string {
	keys := make([]string, modes.Len())

	i := 0
	for el := modes.Front(); el != nil; el = el.Next() {
		keys[i] = el.Value.name
		i++
	}
	return keys
}

func SetModeByIndex(ec os.File, index int64) error {
	var currIndex int64
	for _, key := range modes.Keys() {
		mode, _ := modes.Get(key)
		err := SetBit(ec, mode.offset, mode.bit, index == currIndex)
		if err != nil {
			return err
		}
		currIndex++
	}
	return nil
}

func GetCurrentMode(ec os.File) (Mode, int) {
	var result Mode
	resultIndex := -1
	var normalIndex int
	index := 0
	for el := modes.Front(); el != nil; el = el.Next() {
		mode := el.Value
		if mode == MODE_NORMAL {
			normalIndex = index
		}

		// probably not normal mode - we can check for values
		if mode.offset != 0 {
			enabled, _ := mode.ReadModeValue(ec)
			if enabled {
				result = mode
				resultIndex = index
			}
		}
		index++
	}
	// fallback if no mode is set
	if resultIndex == -1 {
		return MODE_NORMAL, normalIndex
	}
	return result, resultIndex
}

func (m Mode) ReadModeValue(ec os.File) (bool, error) {
	return ReadBit(ec, m.offset, m.bit)
}
