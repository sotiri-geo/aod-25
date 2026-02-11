package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func MaxTwoDigit(input []int) int {
	maxSeen := 0
	ans := 0

	for _, val := range input {
		curr := (maxSeen * 10) + val
		ans = max(ans, curr)
		maxSeen = max(maxSeen, val)
	}
	return ans
}

func MaxTwelveDigit(input []int) int {
	stack := []int{}
	buffer := 12
	fromEnd := len(input) - 1

	for _, val := range input {
		// Keep mono decreasing stack constraint
		for len(stack) > 0 && val > stack[len(stack)-1] && fromEnd >= buffer {
			// Recreate pop in go
			stack = stack[:len(stack)-1]
			buffer++
		}
		fromEnd--
		if buffer == 0 {
			continue
		}
		stack = append(stack, val)
		buffer--
	}
	// Recreate the final value
	total := 0

	for i, num := range stack {
		total += num * int(math.Pow(10, float64(len(stack)-i-1)))
	}

	return total
}

func Part1(input [][]int) int {
	total := 0

	for _, bank := range input {
		total += MaxTwoDigit(bank)
	}

	return total
}

func Part2(input [][]int) int {
	total := 0

	for _, bank := range input {
		total += MaxTwelveDigit(bank)
	}

	return total
}

func parseDigits(input string) []int {
	output := make([]int, 0, len(input))

	for _, val := range input {
		output = append(output, int(val-'0'))
	}

	return output
}

func main() {
	data, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		panic(err)
	}
	// Parse input
	banks := strings.Split(strings.TrimSpace(string(data)), "\n")
	input := make([][]int, 0, len(banks))

	for _, bank := range banks {
		input = append(input, parseDigits(bank))
	}

	p1 := Part1(input)
	p2 := Part2(input)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
}
