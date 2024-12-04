package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
  if x >= 0 {
    return x
  } else {
    return -x
  }
}

func IsUnsualSequence(sequence []int) bool {
  if len(sequence) < 2 { return true }
  
  isAscending := sequence[1] - sequence[0] > 0
  pastNumber := sequence[0]

  for _, currentNumber := range sequence[1:] {
    grad := currentNumber - pastNumber

    if grad == 0 {
      return false // neither an increase or a decrease
    } else if (grad > 0 && !isAscending) || (grad < 0 && isAscending) { // ascending or desecnding case
      return false // wrong direction
    }
    if !(Abs(grad) >= 1 && Abs(grad) <= 3) {
      return false // an increase / derease of more than 3
    }

    pastNumber = currentNumber
  }

  return true
}

func Part1(rawInput string) int {
  splitByLine := strings.Split(rawInput, "\n")

  safeReports := 0

  for _, line := range splitByLine {
    splitBySpace := strings.Split(line, " ")

    numLevels := make([]int, len(splitBySpace))

    for i:=0; i<len(splitBySpace); i++ {
      intValue, err := strconv.Atoi(splitBySpace[i])
      check(err)
      numLevels[i] = intValue
    }

    if IsUnsualSequence(numLevels) {
      safeReports += 1
    }
  }

  return safeReports
}

func main() {
	b1, err := os.ReadFile("./input.txt")
  check(err)

	data := string(b1)

	fmt.Println(Part1(data))

  // a := []int{7, 6, 4, 2, 1}
  // fmt.Printf("a %t\n", IsUnsualSequence(a))

  // b := []int{1, 2, 7, 8, 9}
  // fmt.Printf("b %t\n", IsUnsualSequence(b))

  // c := []int{9, 7, 6, 2, 1}
  // fmt.Printf("c %t\n", IsUnsualSequence(c))

  // d := []int{1, 3, 2, 4, 5}
  // fmt.Printf("d %t\n", IsUnsualSequence(d))

  // e := []int{8, 6, 4, 4, 1}
  // fmt.Printf("e %t\n", IsUnsualSequence(e))

  // f := []int{1, 3, 6, 7, 9}
  // fmt.Printf("f %t\n", IsUnsualSequence(f))
}