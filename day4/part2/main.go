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
	marked [][]int
}

func main() {
	wordSearch := NewWordSearch("../puzzle.txt", "MAS")

	for row := 0; row < len(wordSearch.puzzle); row++ {
		for col := 0; col < len(wordSearch.puzzle[0]); col++ {
			for _, neighborDirection := range [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
				drow, dcol := neighborDirection[0], neighborDirection[1]
				if wordSearch.puzzle[row][col] == string(wordSearch.target[0]) {
					if wordSearch.isValidSearchDirection(row, col, drow, dcol) {
						midRow, midCol := row+drow, col+dcol
						wordSearch.marked[midRow][midCol]++
					}
				}
			}
		}
	}

	numOccurrences := 0
	for row := 0; row < len(wordSearch.marked); row++ {
		for col := 0; col < len(wordSearch.marked[0]); col++ {
			if wordSearch.marked[row][col] >= 2 {
				numOccurrences++
			}
		}
	}

	fmt.Printf("numOccurrences: %d\n", numOccurrences)
}

func (ws *WordSearch) printMarked() {
	for row := 0; row < len(ws.puzzle); row++ {
		for col := 0; col < len(ws.puzzle[0]); col++ {
			fmt.Print(ws.marked[row][col])
		}
		fmt.Println()
	}
}

func (ws *WordSearch) isValidSearchDirection(row, col, drow, dcol int) bool {
	for i := 0; i < len(ws.target); i++ {
		r, c := row+(drow*i), col+(dcol*i)
		if ((r < 0 || r >= len(ws.puzzle)) || (c < 0 || c >= len(ws.puzzle[0]))) || ws.puzzle[r][c] != string(ws.target[i]) {
			return false
		}
	}
	return true
}

func NewWordSearch(inputFilename string, target string) *WordSearch {
	wordSearch := &WordSearch{puzzle: getPuzzle(inputFilename), target: target}
	wordSearch.marked = Make2D[int](len(wordSearch.puzzle), len(wordSearch.puzzle[0]))
	return wordSearch
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for row := range n {
		matrix[row] = make([]T, m)
	}
	return matrix
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
