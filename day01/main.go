package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func Abs (x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func DecodeDistance(rawInput string) int {
	// split input by spaces, then trim
	splitbyLine := strings.Split(rawInput, "\n")
  
  if (splitbyLine[len(splitbyLine) - 1] == "") {
    splitbyLine = splitbyLine[:len(splitbyLine) - 1]
  }

  listA := make([]int, len(splitbyLine))
  listB := make([]int, len(splitbyLine))

  for i, line := range splitbyLine {
    splitbySpace := strings.Split(line, "   ")

    A, err := strconv.Atoi(splitbySpace[0])
    check(err)
    
    listA[i] = A
    
    B, err := strconv.Atoi(splitbySpace[1])
    check(err)
    
    listB[i] = B
  }

  slices.Sort(listA)
  slices.Sort(listB)

  result := 0

  for i := 0; i < len(listA); i++ {
   if (len(listA) != len(listB)) {
      return -1
    }

    result += Abs(listA[i] - listB[i])
  }

	return result
}

func main() {
	data, err := os.ReadFile("./input.txt") // ans 1651298
  check(err)

	b1 := string(data)
	fmt.Println(DecodeDistance(b1))
}