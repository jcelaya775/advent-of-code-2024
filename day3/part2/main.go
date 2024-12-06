package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type state int

const (
	do state = iota
	dont
)

func main() {
	memory := getMemory("../memory.txt")

	result := 0

	matchMul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	mulIndexes := matchMul.FindAllStringIndex(memory, -1)
	matchDo := regexp.MustCompile(`do\(\)`)
	doIndexes := matchDo.FindAllStringIndex(memory, -1)
	matchDont := regexp.MustCompile(`don't\(\)`)
	dontIndexes := matchDont.FindAllStringIndex(memory, -1)

	currState := do
	mulIndex := 0
	mulComparer := mulIndexes[mulIndex][0]
	doIndex := 0
	doComparer := doIndexes[doIndex][0]
	dontIndex := 0
	dontComparer := dontIndexes[dontIndex][0]
	for mulIndex < len(mulIndexes) || doIndex < len(doIndexes) || dontIndex < len(dontIndexes) {
		switch currIndex := min(mulComparer, doComparer, dontComparer); currIndex {
		case mulComparer:
			if currState == do {
				mulIndexPair := mulIndexes[mulIndex]
				mulStr := memory[mulIndexPair[0]:mulIndexPair[1]]
				numMatch := regexp.MustCompile(`[0-9]{1,3}`)
				nums := strArrToIntArr(numMatch.FindAllString(mulStr, -1))
				result += nums[0] * nums[1]
			}

			mulIndex++
			if mulIndex < len(mulIndexes) {
				mulComparer = mulIndexes[mulIndex][0]
			} else {
				mulComparer = math.MaxInt
			}
		case doComparer:
			currState = do
			doIndex++
			if doIndex < len(doIndexes) {
				doComparer = doIndexes[doIndex][0]
			} else {
				doComparer = math.MaxInt
			}
		case dontComparer:
			currState = dont
			dontIndex++
			if dontIndex < len(dontIndexes) {
				dontComparer = dontIndexes[dontIndex][0]
			} else {
				dontComparer = math.MaxInt
			}
		}
	}

	fmt.Printf("result: %d\n", result)
}

func getMemory(filename string) string {
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

	memory := ""
	for scanner.Scan() {
		if !empty(scanner.Text()) {
			memory += strings.Trim(scanner.Text(), " ")
		}
	}
	return memory
}

func empty(s string) bool {
	if strings.Trim(s, " ") == "" {
		return true
	} else {
		return false
	}
}

func strArrToIntArr(sArr []string) []int {
	var result []int
	for _, s := range sArr {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return result
}
