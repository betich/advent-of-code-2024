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

func ReadFile(path string) string {
	b1, err := os.ReadFile(path)
	check(err)

	data := string(b1)

	return data
}

func CreateRule(rulesInput []string) [][]int {
	rule := [][]int{}

	for _, line := range rulesInput {
		splitBySeperator := strings.Split(line, "|")
		before := strings.TrimSpace(splitBySeperator[0])
		after := strings.TrimSpace(splitBySeperator[1])

		b, err := strconv.Atoi(before)
		check(err)

		a, err := strconv.Atoi(after)
		check(err)

		ruleSlice := []int{b, a}
		rule = append(rule, ruleSlice)
	}

	return rule
}

func CreateUpdateInstructions(updatesInput []string) [][]int {
	updates := [][]int{}

	for i, line := range updatesInput {
		splitByComma := strings.Split(line, ",")

		updates = append(updates, []int{})

		for _, num := range splitByComma {
			n, err := strconv.Atoi(num)
			check(err)

			updates[i] = append(updates[i], n)
		}
	}

	return updates
}

func ProcessInput(data string) ([]string, []string) {
	rulesInput := []string{}
	updatesInput := []string{}

	splitByLines := strings.Split(data, "\n")

	// find the first empty line
	for i, line := range splitByLines {
		if line == "" {
			rulesInput = splitByLines[:i]
			updatesInput = splitByLines[i+1:]
			break
		}
	}

	return rulesInput, updatesInput
}

func FindInSlice(value int, slice []int, fromIndex int) (bool, int) {
	for i, v := range slice[fromIndex:] {
		if v == value {
			return true, i
		}
	}

	return false, -1
}

func FindMultipleInSlice(before int, after int, slice []int) (bool, []int) {
	found := []int{}
	beforeIndex := 0
	afterIndex := 0

	for i, v := range slice {
		if v == before {
			found = append(found, i)
			beforeIndex = i
		}

		if v == after {
			found = append(found, i)
			afterIndex = i
		}
	}

	return len(found) == 2, []int{beforeIndex, afterIndex}
}

func All(a []bool) bool {
	for _, v := range a {
		if !v {
			return false
		}
	}

	return true
}

func FilterValidUpdates(rules [][]int, updates [][]int) [][]int {
	validUpdates := [][]int{}

	for _, update := range updates {
		valid := make([]bool, len(rules))

		for i, rule := range rules {
			before := rule[0]
			after := rule[1]

			found, indexes := FindMultipleInSlice(before, after, update)

			valid[i] = !found || (indexes[0] < indexes[1])
		}

		if All(valid) {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}

func SumMiddleValues(values [][]int) int {
	sum := 0

	for _, a := range values {
		middleIndex := len(a) / 2
		if len(a)%2 == 0 {
			middleIndex--
		}
		sum += a[middleIndex]
	}

	return sum
}

func PrinterUpdateValue(data string) int {
	rulesInput, updatesInput := ProcessInput(data)

	rules := CreateRule(rulesInput)
	updates := CreateUpdateInstructions(updatesInput)

	validUpdates := FilterValidUpdates(rules, updates)

	result := SumMiddleValues(validUpdates)

	return result
}

func main() {
	data := ReadFile("./input.txt")

	fmt.Println(PrinterUpdateValue(data)) // 4790
}
