package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/sotiri-geo/aod-2025/common"
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

func cloneGrid(input [][]byte) [][]byte {
	cloned := make([][]byte, len(input))
	for i := range input {
		cloned[i] = append([]byte(nil), input[i]...)
	}
	return cloned
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

// ProcessRollOfPaperWithUpdate will keep looping and attempting
// to remove in a greedy way and update the state of the grid.
// Will break out from loop after a full scan with no identifiable
// rolls that can be accessed
func (g *Grid) ProcessRollOfPaperWithUpdate() int {
	total := 0

	for {
		removed := false
		for i := range len(g.input) {
			for j := range len(g.input[0]) {
				if g.CanAccessRollOfPaper(i, j) {
					total++
					// update cell
					g.input[i][j] = '.'
					removed = true
				}
			}
		}
		if !removed {
			break
		}
	}
	return total
}

func QueueBasedUpdate(input [][]byte) (int, error) {
	if len(input) == 0 || len(input[0]) == 0 {
		return 0, nil
	}

	// Init datastructures
	inDegree := map[Point]int{}
	queue := common.NewQueue([]Point{})
	graph := map[Point][]Point{}
	seen := common.NewSet([]Point{})
	ans := 0
	n, m := len(input), len(input[0])
	directions := [8][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, 1}, {-1, -1}, {1, 1}, {1, -1},
	}

	// Populate the inDegree of each node that is a roll
	for row := range n {
		for col := range m {
			// Only process if current element is a roll
			if input[row][col] != '@' {
				continue
			}
			p := Point{row: row, col: col}
			count := 0
			// Check 8 directionally around roll to update indegree of any adjacent rolls
			for _, d := range directions {
				px := Point{row: row + d[0], col: col + d[1]}
				if inBounds(&px, input) && input[px.row][px.col] == '@' {
					count++
					// update graph: p directionally connect to px
					graph[p] = append(graph[p], px)
				}
			}

			inDegree[p] = count
			// If point is accessable, add to queue
			if count < 4 {
				queue.Enqueue(p)
				seen.Add(p)
			}
		}
	}

	// Process queue and update

	for queue.Len() > 0 {
		// pop a point off the queue
		p, err := queue.Dequeue()
		if err != nil {
			return 0, fmt.Errorf("dequeueing point %v: %w", p, err)
		}
		// stale queue entry
		if _, exists := inDegree[p]; !exists {
			continue
		}

		// Remove point from InDegree
		delete(inDegree, p)
		for _, nei := range graph[p] {
			if _, exists := inDegree[nei]; exists {
				// Decrement InDegree of all adjacent neighbours
				inDegree[nei]--
				// New node we can add to the queue
				if inDegree[nei] < 4 && !seen.Has(nei) {
					seen.Add(nei)
					queue.Enqueue(nei)
				}
			}
		}
		// Add to count
		ans++
	}

	return ans, nil
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
	grid := NewGrid(cloneGrid(input))

	fmt.Println("Part1:", grid.ProcessRollOfPaper())
	fmt.Println("Part2:", NewGrid(cloneGrid(input)).ProcessRollOfPaperWithUpdate())

	// Using Graph Algo with FIFO Queue
	ans, err := QueueBasedUpdate(cloneGrid(input))
	if err != nil {
		panic(err)
	}

	fmt.Println("Part2WithGraph:", ans)
}
