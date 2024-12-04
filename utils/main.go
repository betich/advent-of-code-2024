package utils

import (
	"os"
)

// common

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) string {
	b1, err := os.ReadFile("./input.txt")
  check(err)

	data := string(b1)

	return data
}

// math

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}