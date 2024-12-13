package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordSearch struct {
	puzzle [][]string
	target string
}

func main() {
	ws := WordSearch{puzzle: getPuzzle("../puzzle.txt"), target: "XMAS"}

	numOccurrences := 0
	for row := 0; row < len(ws.puzzle); row++ {
		for col := 0; col < len(ws.puzzle[0]); col++ {
			for _, neighborDirection := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
				drow, dcol := neighborDirection[0], neighborDirection[1]
				if ws.puzzle[row][col] == string(ws.target[0]) && ws.isValidSearchDirection(row, col, drow, dcol) {
					numOccurrences++
				}
			}
		}
	}

	fmt.Printf("numOccurrences: %d\n", numOccurrences)
}

func (ws WordSearch) isValidSearchDirection(row, col, drow, dcol int) bool {
	for i := 0; i < len(ws.target); i++ {
		r, c := row+(drow*i), col+(dcol*i)
		if ((r < 0 || r >= len(ws.puzzle)) || (c < 0 || c >= len(ws.puzzle[0]))) || ws.puzzle[r][c] != string(ws.target[i]) {
			return false
		}
	}
	return true
}

func getPuzzle(filename string) [][]string {
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

	var puzzle [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		puzzle = append(puzzle, line)
	}
	return puzzle
}
