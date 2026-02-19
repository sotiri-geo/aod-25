package main

import (
	"bytes"
	"fmt"
	"os"
)

// Grid stores a 2 Dimentional byte array
type Grid struct {
	input [][]byte
}

// Point represents a point on the grid
type Point struct {
	row int
	col int
}

// NewGrid creates a new Grid object and returns it
func NewGrid(input [][]byte) *Grid {
	return &Grid{input}
}

// NewPoint creates new point from row and col coordinates
func NewPoint(row, col int) *Point {
	return &Point{row, col}
}

func (g *Grid) CanAccessRollOfPaper(row, col int) bool {
	if len(g.input) == 0 {
		return false
	}
	if g.input[row][col] != '@' {
		return false
	}
	up := NewPoint(row-1, col)
	down := NewPoint(row+1, col)
	left := NewPoint(row, col-1)
	right := NewPoint(row, col+1)
	upRight := NewPoint(row-1, col+1)
	upLeft := NewPoint(row-1, col-1)
	downRight := NewPoint(row+1, col+1)
	downLeft := NewPoint(row+1, col-1)

	directions := []*Point{up, down, left, right, upRight, upLeft, downRight, downLeft}

	rollsOfPaper := 0

	for _, p := range directions {
		if ok := inBounds(p, g.input); ok && g.input[p.row][p.col] == '@' {
			rollsOfPaper++
		}
	}
	return rollsOfPaper < 4
}

// ProcessRollOfPaper counts the numbers of roll papers the forklift
// can access
func (g *Grid) ProcessRollOfPaper() int {
	total := 0

	for i := range len(g.input) {
		for j := range len(g.input[0]) {
			if g.CanAccessRollOfPaper(i, j) {
				total++
			}
		}
	}
	return total
}

// inBounds returns true if the coordinates are within bounds of
// the gride. False otherwise.
func inBounds(p *Point, input [][]byte) bool {
	rowSize, colSize := len(input), len(input[0])
	return 0 <= p.row && p.row < rowSize && 0 <= p.col && p.col < colSize
}

func main() {
	data, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		panic(err)
	}
	// Parse input
	input := bytes.Split(data, []byte{'\n'})
	input = input[:len(input)-1]
	grid := NewGrid(input)

	fmt.Println("Part1:", grid.ProcessRollOfPaper())
}
