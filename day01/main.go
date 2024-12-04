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

func DecodeNumbers(rawInput string) ([]int, []int) {
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

  return listA, listB
}

func Part1(rawInput string) int {
	listA, listB := DecodeNumbers(rawInput)

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

func Part2(rawInput string) int {
  listA, listB := DecodeNumbers(rawInput)

  // unique values of listA

  unique := map[int]int{}

  for _, value := range listA {
    if _, ok := unique[value]; !ok {
      // find all occurences of value in listB
      count := 0
      for _, valueB := range listB {
        if valueB == value {
          count++
        }
      }
      unique[value] = count
    }
  }

  result := 0

  for key, value := range unique {
    result += key * value
  }

  return result
}

func main() {
	data, err := os.ReadFile("./input.txt")
  check(err)

	b1 := string(data)
	fmt.Println(Part1(b1)) // ans 1651298

  fmt.Println(Part2(b1)) // ans 21306195
}