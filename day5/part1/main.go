package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := 0
	rules, updates := readRulesAndUpdates("../manual.txt")

	for _, update := range updates {
		numsToIndexes := make(map[int]int)
		for i, num := range update {
			numsToIndexes[num] = i
		}

		isValidUpdate := true
		for _, rule := range rules {
			if leftIndex, ok := numsToIndexes[rule[0]]; ok {
				if rightIndex, ok := numsToIndexes[rule[1]]; ok && leftIndex > rightIndex {
					isValidUpdate = false
					break
				}
			}
		}

		if isValidUpdate {
			result += update[len(update)/2]
		}
	}

	fmt.Printf("result: %d\n", result)
}

func readRulesAndUpdates(filename string) (rules [][2]int, updates [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if isEmpty(scanner.Text()) {
			break
		}
		tokens := strings.Split(scanner.Text(), "|")
		leftNum, _ := strconv.Atoi(tokens[0])
		rightNum, _ := strconv.Atoi(tokens[1])
		rules = append(rules, [2]int{leftNum, rightNum})
	}

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		var update []int
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return
}

func isEmpty(s string) bool {
	if strings.Trim(s, " ") == "" {
		return true
	} else {
		return false
	}
}
