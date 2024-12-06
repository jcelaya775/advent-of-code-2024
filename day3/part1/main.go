package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	memory := getMemory("../memory.txt")

	// Method 1: regex /mul\([0-9]{1,3},[0-9]{1,3}\)/g
	// Method 2: manually loop through characters (pretty much manual regex) & mul(*,*) patterns w/ two pointers

	// Method 1: regex
	result := 0

	mulMatch := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	mulStrs := mulMatch.FindAllString(memory, -1)
	for _, mulStr := range mulStrs {
		numMatch := regexp.MustCompile(`[0-9]{1,3}`)
		nums := strArrToIntArr(numMatch.FindAllString(mulStr, -1))
		result += nums[0] * nums[1]
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
