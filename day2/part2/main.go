package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	increasing direction = iota
	decreasing
)

func main() {
	reports := getReports("../reports.txt")

	numSafeReports := 0
	for _, report := range reports {
		var prevDirection direction
		if report[1] > report[0] {
			prevDirection = increasing
		} else if report[1] < report[0] {
			prevDirection = decreasing
		}

		// dec dec inc dec dec
		// inc dec dec dec
		// [0, 1, -4, -5, -7]

		// 1 3 2 4 5: Safe by removing the second level, 3.
		// 8 6 4 4 1: Safe by removing the third level, 4.

		numBadReports := 0
		for i := 1; i < len(report); i++ {
			var currDirection direction
			if report[i] > report[i-1] {
				currDirection = increasing
			} else if report[i] < report[i-1] {
				currDirection = decreasing
			}
			diff := abs(report[i] - report[i-1])
			if (diff < 1 || diff > 3) || currDirection != prevDirection {
				numBadReports += 1
				if numBadReports > 1 {
					break
				}
			}

			prevDirection = currDirection
		}

		if numBadReports <= 1 {
			numSafeReports += 1
		}
	}

	fmt.Printf("numSafeReports: %d\n", numSafeReports)
}

func getReports(filename string) [][]int {
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

	var reports [][]int
	for scanner.Scan() {
		var report []int
		if !empty(scanner.Text()) {
			for _, level := range strings.Fields(scanner.Text()) {
				if parsedLevel, err := strconv.Atoi(level); err == nil {
					report = append(report, parsedLevel)
				} else {
					panic(err)
				}
			}
			reports = append(reports, report)
		}
	}
	return reports
}

func empty(s string) bool {
	if strings.Trim(s, " ") == "" {
		return true
	} else {
		return false
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
