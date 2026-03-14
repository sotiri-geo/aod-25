package main

import (
	"fmt"
	"math/big"
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

func RollingSumWithStrings(nums ...string) int {
	input := make([]int, 0, len(nums))

	for _, val := range nums {
		num, _ := strconv.Atoi(val)
		input = append(input, num)
	}
	return RollingSum(input...)
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

func RollingMulWithStrings(nums ...string) int {
	input := make([]int, 0, len(nums))

	for _, val := range nums {
		num, _ := strconv.Atoi(val)
		input = append(input, num)
	}

	return RollingMul(input...)
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

func toString(input []int) []string {
	output := make([]string, 0, len(input))

	for _, num := range input {
		val := strconv.Itoa(num)
		output = append(output, val)
	}
	return output
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

func Part2(input [][]string) int {
	operators, operands := ExtractColumns(input)
	total := 0

	for i, operand := range operands {
		if operand == "+" {
			total += RightToLeftAdd(toString(operators[i]))
		} else {
			total += RightToLeftMultiply(toString(operators[i]))
		}
	}

	return total
}

func charAt(line string, idx int) byte {
	if idx < 0 || idx >= len(line) {
		return ' '
	}
	return line[idx]
}

func isSeparatorColumn(lines []string, col int) bool {
	for _, line := range lines {
		if charAt(line, col) != ' ' {
			return false
		}
	}
	return true
}

func Part2FromLines(lines []string) *big.Int {
	if len(lines) == 0 {
		return big.NewInt(0)
	}

	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	total := big.NewInt(0)
	lastRow := len(lines) - 1

	for col := 0; col < width; {
		if isSeparatorColumn(lines, col) {
			col++
			continue
		}

		start := col
		for col < width && !isSeparatorColumn(lines, col) {
			col++
		}
		end := col - 1

		var operator byte
		for c := start; c <= end; c++ {
			ch := charAt(lines[lastRow], c)
			if ch == '+' || ch == '*' {
				operator = ch
				break
			}
		}

		problemTotal := big.NewInt(0)
		if operator == '*' {
			problemTotal.SetInt64(1)
		}

		for c := end; c >= start; c-- {
			var b strings.Builder
			b.Grow(lastRow)
			for r := 0; r < lastRow; r++ {
				ch := charAt(lines[r], c)
				if ch >= '0' && ch <= '9' {
					b.WriteByte(ch)
				}
			}

			if b.Len() == 0 {
				continue
			}

			val := new(big.Int)
			val.SetString(b.String(), 10)

			if operator == '+' {
				problemTotal.Add(problemTotal, val)
			} else {
				problemTotal.Mul(problemTotal, val)
			}
		}

		total.Add(total, problemTotal)
	}

	return total
}

func maxLength(arr []string) int {
	output := 0

	for _, val := range arr {
		if len(val) > output {
			output = len(val)
		}
	}
	return output
}

func RightToLeftAdd(input []string) int {
	// We need to some how pad to the right length
	// but it needs to operate from right to left
	ml := maxLength(input)

	columns := make([]string, 0, ml)

	for i := range ml {
		// buffer string value
		var b strings.Builder
		b.Grow(len(input))
		for _, val := range input {
			if i >= len(val) {
				continue
			}
			b.WriteByte(val[i])
		}
		columns = append(columns, b.String())
	}

	return RollingSumWithStrings(columns...)
}

func RightToLeftMultiply(input []string) int {
	// For multiplication we need to pad to the right

	ml := maxLength(input)

	columns := make([]string, ml)

	for i := 0; i < ml; i++ {
		var b strings.Builder
		b.Grow(len(input))
		for _, val := range input {
			idx := len(val) - i - 1
			if idx < 0 {
				// guard against overflow indx
				continue
			}
			b.WriteByte(val[idx])
		}
		// multiplication is associative, so either way round is applicable
		columns[i] = b.String()
	}

	return RollingMulWithStrings(columns...)
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
	fmt.Println("Part2:", Part2FromLines(lines).String())
}
