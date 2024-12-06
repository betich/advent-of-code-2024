package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) string {
	b1, err := os.ReadFile(path)
  check(err)

	data := string(b1)

	return data
}

func main() {
	data := ReadFile("./input.txt")

	fmt.Println(data)
}