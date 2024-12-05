package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var list1, list2 []int

	file, err := os.Open("../lists.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ids := strings.Fields(line)
		id1, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}
		id2, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, id1)
		list2 = append(list2, id2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num] += 1
	}

	similarityScore := 0
	for _, num := range list1 {
		count, ok := counts[num]
		if ok {
			similarityScore += num * count
		}
	}

	fmt.Printf("similarityScore: %d\n", similarityScore)
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
