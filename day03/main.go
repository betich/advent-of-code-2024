package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func ProcessMulMemory(data string) int { // Part 1
	r := regexp.MustCompile(`(mul\(\d+,\d+\))`)
	matches := r.FindAllString(data, -1)

	result := 0

	for _, match := range matches {
		splitResult := strings.Split(match, ",")

		leftStr := splitResult[0]
		rightStr := splitResult[1]

		leftStr = strings.ReplaceAll(leftStr, "mul(", "")
		rightStr = strings.ReplaceAll(rightStr, ")", "")

		left, err := strconv.Atoi(leftStr)
		check(err)

		right, err := strconv.Atoi(rightStr)
		check(err)

		result += left * right
	}

	return result
}

func main() {
	data := ReadFile("./input.txt")

	fmt.Println(ProcessMulMemory(data)) // 164730528
}