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

func main() {
	b, err := os.ReadFile("./input.txt")
  check(err)

	data := string(b)

	fmt.Println("halo", data)
}