package models

import (
	"strconv"
)

//Convert string to int
func StringtoInt(str string) int {
	value, err := strconv.ParseInt(str, 10, 32)
	CheckErr(err)
	return int(value)
}

//Convert string to int8
func StringtoInt8(str string) int8 {
	value, err := strconv.ParseInt(str, 10, 8)
	CheckErr(err)
	return int8(value)
}

//Convert string to int16
func StringtoInt16(str string) int16 {
	value, err := strconv.ParseInt(str, 10, 16)
	CheckErr(err)
	return int16(value)
}

//Convert string to int64
func StringtoInt64(str string) int64 {
	value, err := strconv.ParseInt(str, 10, 64)
	CheckErr(err)
	return int64(value)
}

