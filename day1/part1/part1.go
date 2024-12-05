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

	sum := 0
	for i := 0; i < len(list1); i++ {
		sum += abs(list2[i] - list1[i])
	}

	fmt.Printf("sum: %d\n", sum)
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
