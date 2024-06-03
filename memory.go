package main

import "os"

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
