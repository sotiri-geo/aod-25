package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RollingSum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func RollingMul(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	total := 1

	for _, num := range nums {
		total *= num
	}
	return total
}

// ExtractColumns takes in input grid and returns
// an array of integer arrays where the index corresponds to the column
// of the operands and the index of string values corresponding to the index
// of the operator. i.e. index = 0 is first column.
func ExtractColumns(grid [][]string) ([][]int, []string) {
	rowSize := len(grid)
	if rowSize == 0 {
		return [][]int{}, []string{}
	}
	colSize := len(grid[0])
	// Init response data structures
	operands := make([][]int, 0, colSize)
	operators := make([]string, 0, colSize)

	// Exclude the last row as this will be the operator
	for j := range colSize {
		currentOperand := make([]int, 0, rowSize-1)
		for i := range rowSize - 1 {
			num, _ := strconv.Atoi(grid[i][j])
			currentOperand = append(currentOperand, num)
		}
		operands = append(operands, currentOperand)
		operators = append(operators, grid[rowSize-1][j])
	}
	return operands, operators
}

func Part1(input [][]string) int {
	operators, operands := ExtractColumns(input)
	total := 0
	for i, operand := range operands {
		if operand == "+" {
			total += RollingSum(operators[i]...)
		} else {
			total += RollingMul(operators[i]...)
		}
	}
	return total
}

func main() {
	data, err := os.ReadFile("./day06/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	grid := make([][]string, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		grid = append(grid, fields)
	}

	// Results
	fmt.Println("Part1:", Part1(grid))
}
