package main

import (
	"log"
	"os"
)

func OpenEc() (*os.File, error) {
	return os.OpenFile("/sys/kernel/debug/ec/ec0/io", os.O_RDWR, 0644)
}

func main() {
	InitModes()

	ec, err := OpenEc()
	if err != nil {
		log.Fatal("Couldn't open EC: ", err)
	}
	TuiMain(*ec)
}
