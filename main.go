package main

import (
	"log"
	"os"
)

func OpenEc() (*os.File, error) {
	err := ExecCommand("modprobe", "ec_sys", "write_support=1")
	if err != nil {
		return nil, err
	}
	return os.OpenFile("/sys/kernel/debug/ec/ec0/io", os.O_RDWR, 0644)
}

func main() {
	InitModes()

	ec, err := OpenEc()
	if err != nil {
		log.Fatal("Couldn't open EC (did you forget to use sudo?): ", err)
	}
	TuiMain(*ec)
}
