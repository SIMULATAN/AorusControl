package main

import (
	"fmt"
	"log"
	"os"
)

func OpenEc() (*os.File, error) {
	return os.Open("/sys/kernel/debug/ec/ec0/io")
}

func main() {
	ecPt, err := OpenEc()
	if err != nil {
		log.Fatal("Couldn't open EC: ", err)
	}
	ec := *ecPt

	quiet, err := ReadBit(ec, OFFSET_MODE_QUIET, BIT_MODE_QUIET)
	gaming, err := ReadBit(ec, OFFSET_MODE_GAMING, BIT_MODE_GAMING)
	deep, err := ReadBit(ec, OFFSET_MODE_DEEP_CONTROL, BIT_MODE_DEEP_CONTROL)
	auto, err := ReadBit(ec, OFFSET_MODE_AUTO, BIT_MODE_AUTO)
	fixed, err := ReadBit(ec, OFFSET_MODE_FIXED, BIT_MODE_FIXED)
	fmt.Println("Current values:\n  Fan mode")
	fmt.Println("    Quiet:", quiet)
	fmt.Println("    Gaming:", gaming)
	fmt.Println("    Deep contol:", deep)
	fmt.Println("    Auto:", auto)
	fmt.Println("    Fixed:", fixed)
}
