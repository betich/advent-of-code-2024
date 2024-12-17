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

func IsAllValid(rules [][]int, update []int) (bool, [][]int) {
	valid := make([]bool, len(rules))
	wrongRules := [][]int{}

	for i, rule := range rules {
		before := rule[0]
		after := rule[1]

		found, indexes := FindMultipleInSlice(before, after, update)

		valid[i] = !found || (indexes[0] < indexes[1])
		if !valid[i] {
			wrongRules = append(wrongRules, rule)
		}
	}
	return All(valid), wrongRules
}

func FilterValidUpdates(rules [][]int, updates [][]int) [][]int {
	validUpdates := [][]int{}

	for _, update := range updates {
		allValid, _ := IsAllValid(rules, update)
		if allValid {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}

func Find(a []int, value int) int {
	for i, v := range a {
		if v == value {
			return i
		}
	}

	return -1
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

func Rearrange(sample []int, rules [][]int) []int {
	// Create adjacency list and in-degree map, scoped to elements in the sample
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Initialize in-degree map for elements in the sample
	for _, num := range sample {
		inDegree[num] = 0
	}

	// Build graph and in-degree map, only for elements in the sample
	for _, rule := range rules {
		before, after := rule[0], rule[1]
		if _, existsBefore := inDegree[before]; existsBefore {
			if _, existsAfter := inDegree[after]; existsAfter {
				graph[before] = append(graph[before], after)
				inDegree[after]++
			}
		}
	}

	// Topological sort using Kahn's algorithm
	queue := []int{}
	for num, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if the sorted result contains all elements from the sample
	if len(result) != len(sample) {
		return nil
	}

	return result
}

func ProcessIncorrectUpdates(rules [][]int, updates [][]int) [][]int {
	processedUpdates := [][]int{}

	for _, update := range updates {
		allValid, wrongRules := IsAllValid(rules, update)
		if !allValid {
			// Rearrange the update to make it valid
			rearranged := Rearrange(update, wrongRules)
			if rearranged != nil {
				processedUpdates = append(processedUpdates, rearranged)
			}
			fmt.Println("Rearranged:", rearranged)
		}
	}

	return processedUpdates
}

func PrinterUpdateValue(data string) int {
	rulesInput, updatesInput := ProcessInput(data)

	rules := CreateRule(rulesInput)
	updates := CreateUpdateInstructions(updatesInput)

	// validUpdates := FilterValidUpdates(rules, updates) // 4790

	// result := SumMiddleValues(validUpdates)

	processedUpdates := ProcessIncorrectUpdates(rules, updates) // 10717

	result := SumMiddleValues(processedUpdates)

	return result
}

func main() {
	data := ReadFile("./input.txt")

	fmt.Println(PrinterUpdateValue(data))

	// sample := [][]int{{97, 13, 75, 29, 47}}
	// rules := [][]int{{47, 53}, {97, 13}, {97, 61}, {97, 47}, {75, 29}, {61, 13}, {75, 53}, {29, 13}, {97, 29}, {53, 29}, {61, 53}, {97, 53}, {61, 29}, {47, 13}, {75, 47}, {97, 75}, {47, 61}, {75, 61}, {47, 29}, {75, 13}, {53, 13}}

	// fmt.Println(ProcessIncorrectUpdates(rules, sample))
	// fmt.Println(Rearrange(sample[0], rules))
}
