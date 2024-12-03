package utils

import (
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ReadFile(s string) string {
	data, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}
