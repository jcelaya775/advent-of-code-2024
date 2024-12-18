package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	row, col int
}

type Puzzle struct {
	Map     [][]string
	Visited map[Coord]bool
}

func main() {
	puzzle := NewPuzzle("../map.txt")

	var initialPos *Coord
	var guard string
	for row := range puzzle.Map {
		for col := range puzzle.Map[row] {
			if isGuard(puzzle.Map[row][col]) {
				initialPos = &Coord{row, col}
				guard = puzzle.Map[row][col]
				puzzle.Map[row][col] = "."
			}
		}
	}

	if initialPos == nil {
		fmt.Printf("# positions visited: %d\n", 0)
		return
	}

	slowPos, fastPos := *initialPos, *initialPos
	slowGuard, fastGuard := guard, guard
	outOfBounds, cycleDetected := false, false
	for !outOfBounds && !cycleDetected {
		puzzle.Map[slowPos.row][slowPos.col] = "X"
		puzzle.Visited[slowPos] = true
		if !puzzle.isOutOfBounds(fastPos) {
			puzzle.Map[fastPos.row][fastPos.col] = "X"
			puzzle.Visited[fastPos] = true
		}

		slowPos, slowGuard = puzzle.moveGuard(slowPos, slowGuard, 1)
		if !puzzle.isOutOfBounds(fastPos) {
			fastPos, fastGuard = puzzle.moveGuard(fastPos, fastGuard, 2)
		}

		if puzzle.isOutOfBounds(slowPos) {
			outOfBounds = true
		}
		if slowPos == fastPos {
			cycleDetected = true
		}
	}

	fmt.Printf("# positions visited: %d\n", len(puzzle.Visited))
}

func (puzzle *Puzzle) mapToStr() string {
	s := strings.Builder{}
	for _, row := range puzzle.Map {
		for _, char := range row {
			s.WriteString(char)
		}
		s.WriteString("\n")
	}
	return s.String()
}

func (puzzle *Puzzle) moveGuard(pos Coord, guard string, numSteps int) (Coord, string) {
	var newRow, newCol int
	for range numSteps {
		direction := guardToDirection(guard)
		newRow, newCol = pos.row+direction[0], pos.col+direction[1]
		if !puzzle.isOutOfBounds(Coord{newRow, newCol}) && puzzle.Map[newRow][newCol] == "#" {
			guard = rotateGuard(guard)
			newDirection := guardToDirection(guard)
			newRow, newCol = pos.row+newDirection[0], pos.col+newDirection[1]
		}
		pos = Coord{newRow, newCol}
	}

	return Coord{newRow, newCol}, guard
}

func guardToDirection(guard string) []int {
	if guard == "<" {
		return []int{0, -1}
	} else if guard == "^" {
		return []int{-1, 0}
	} else if guard == ">" {
		return []int{0, 1}
	} else {
		return []int{1, 0}
	}
}

func rotateGuard(guard string) string {
	if guard == "<" {
		return "^"
	} else if guard == "^" {
		return ">"
	} else if guard == ">" {
		return "v"
	} else {
		return "<"
	}
}

func (puzzle *Puzzle) isOutOfBounds(pos Coord) bool {
	if (pos.row < 0 || pos.row >= len(puzzle.Map)) || (pos.col < 0 || pos.col >= len(puzzle.Map[pos.row])) {
		return true
	} else {
		return false
	}
}

func isGuard(char string) bool {
	if char == "<" || char == "^" || char == ">" || char == "v" {
		return true
	} else {
		return false
	}
}

func NewPuzzle(filename string) Puzzle {
	return Puzzle{
		Map:     readMap(filename),
		Visited: make(map[Coord]bool),
	}
}

func readMap(filename string) [][]string {
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

	var result [][]string
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}
	return result
}
