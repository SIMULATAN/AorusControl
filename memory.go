package main

import (
	"os"
)

func ReadByte(file os.File, offset int64) (byte, error) {
	file.Seek(offset, os.SEEK_SET)
	b := make([]byte, 1)
	_, error := file.Read(b)
	return b[0], error
}

func ReadBit(file os.File, offset int64, bit int8) (bool, error) {
	byte, error := ReadByte(file, offset)
	if error != nil {
		return false, error
	}
	return (byte & (1 << bit)) > 0, nil
}

func SetBit(file os.File, offset int64, bit int8, value bool) error {
	existing, err := ReadByte(file, offset)
	if err != nil {
		return err
	}

	if value {
		existing |= (1 << bit)
	} else {
		existing &= ^(1 << bit)
	}

	file.Seek(offset, os.SEEK_SET)
	_, err = file.Write([]byte{existing})

	return err
}
