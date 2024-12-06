package main

import (
	"fmt"
	"os"
	"strings"
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

func WriteFile(path string, data string) {
	f, err := os.Create(path)
	check(err)

	defer f.Close()

	_, err = f.WriteString(data)
	check(err)
}

func WriteOutputMatchMap(matchMap [][]string, path string) string {
	output := ""
	for _, line := range matchMap {
		for _, char := range line {
			output += char
		}
		output += "\n"
	}

	WriteFile(path, output)

	return output
}

func CreateMatchMap(charMap *[][]string) [][]string {
	matchMap := make([][]string, len(*charMap))

	// initialize matchMap
	for i, line := range *charMap {
		matchMap[i] = make([]string, len(line))
		for j := range line {
			matchMap[i][j] = "."
		}
	}

	return matchMap
}

func XMASHorizontalSearch(charMap [][]string, matchMap [][]string) int {
	matches := 0

	// find all X occurrences
	for i, line := range charMap {
		for j, char := range line {
			if char == "X" {
				// j, j+1, j+2, j+3
				if j+3 < len(line) {
					if line[j+1] == "M" && line[j+2] == "A" && line[j+3] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i][j+1] = "M"
						matchMap[i][j+2] = "A"
						matchMap[i][j+3] = "S"
					}
				}
				// reverse condition
				if j-3 >= 0 {
					if line[j-1] == "M" && line[j-2] == "A" && line[j-3] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i][j-1] = "M"
						matchMap[i][j-2] = "A"
						matchMap[i][j-3] = "S"
					}
				}
			}
		}
	}

	WriteOutputMatchMap(matchMap, "./horizontal_output.txt")
	return matches
}

func XMASVerticalSearch(charMap [][]string, matchMap [][]string) int {
	matches := 0
	// find all X occurrences
	for i, line := range charMap {
		for j, char := range line {
			if char == "X" {
				// i, i+1, i+2, i+3
				if i+3 < len(charMap) {
					if charMap[i+1][j] == "M" && charMap[i+2][j] == "A" && charMap[i+3][j] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i+1][j] = "M"
						matchMap[i+2][j] = "A"
						matchMap[i+3][j] = "S"
					}
				}
				// SAMX condition
				if i-3 >= 0 {
					if charMap[i-1][j] == "M" && charMap[i-2][j] == "A" && charMap[i-3][j] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i-1][j] = "M"
						matchMap[i-2][j] = "A"
						matchMap[i-3][j] = "S"
					}
				}
			}
		}
	}

	WriteOutputMatchMap(matchMap, "./vertical_output.txt")
	return matches
}

func XMASDiagonalSearch(charMap [][]string, matchMap [][]string) int {
	matches := 0

	// find all X occurrences
	for i, line := range charMap {
		for j, char := range line {
			if char == "X" {
				// bottom right diagonal
				if i+3 < len(charMap) && j+3 < len(line) {
					if charMap[i+1][j+1] == "M" && charMap[i+2][j+2] == "A" && charMap[i+3][j+3] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i+1][j+1] = "M"
						matchMap[i+2][j+2] = "A"
						matchMap[i+3][j+3] = "S"
					}
				}
				// top left diagonal
				if i-3 >= 0 && j-3 >= 0 {
					if charMap[i-1][j-1] == "M" && charMap[i-2][j-2] == "A" && charMap[i-3][j-3] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i-1][j-1] = "M"
						matchMap[i-2][j-2] = "A"
						matchMap[i-3][j-3] = "S"
					}
				}
				// bottom left diagonal
				if i+3 < len(charMap) && j-3 >= 0 {
					if charMap[i+1][j-1] == "M" && charMap[i+2][j-2] == "A" && charMap[i+3][j-3] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i+1][j-1] = "M"
						matchMap[i+2][j-2] = "A"
						matchMap[i+3][j-3] = "S"
					}
				}
				// top right diagonal
				if i-3 >= 0 && j+3 < len(line) {
					if charMap[i-1][j+1] == "M" && charMap[i-2][j+2] == "A" && charMap[i-3][j+3] == "S" {
						matches += 1
						matchMap[i][j] = "X"
						matchMap[i-1][j+1] = "M"
						matchMap[i-2][j+2] = "A"
						matchMap[i-3][j+3] = "S"
					}
				}
			}
		}
	}

	WriteOutputMatchMap(matchMap, "./diagonal_output.txt")
	return matches
}


func XMASSearch(data string) int {
	wordLines := strings.Split(data, "\n")
	charMap := make([][]string, len(wordLines))

	for i, line := range wordLines {
		charMap[i] = strings.Split(line, "")
	}

	totalMatch := 0

	matchMap := CreateMatchMap(&charMap)

	horizontalMatch := XMASHorizontalSearch(charMap, matchMap)
	verticalMatch := XMASVerticalSearch(charMap, matchMap)
	diagonalMatch := XMASDiagonalSearch(charMap, matchMap)

	WriteOutputMatchMap(matchMap, "./output.txt")

	totalMatch = horizontalMatch + verticalMatch + diagonalMatch
	
	// WriteFile("./output.txt", output)
	return totalMatch
}

func main() {
	data := ReadFile("./input.txt")

	fmt.Println(XMASSearch(data)) // 1811
}