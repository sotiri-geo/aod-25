package main

import (
	"errors"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var ErrInvalidInput = errors.New("invalid input")

// IsRepeatedTwice returns true if the string is repeated twice such as "1212". For odd length this is impossible
// and returns false.
func IsRepeatedTwice(input string) bool {
	size := len(input)
	if size%2 != 0 {
		return false
	}

	mid := size / 2
	return input[:mid] == input[mid:]
}

func IsRepeatedBy(input string, factor int) bool {
	size := len(input)
	toMatch := input[:factor]

	for i := range size / factor {
		if toMatch != input[factor*i:factor*(i+1)] {
			return false
		}
	}
	return true
}

// factors extracts all unique factors of a value into an array and returns it.
func factors(n int) []int {
	factorSet := map[int]struct{}{}
	sqrtN := int(math.Sqrt(float64(n)))

	// One is always a factor
	factorSet[1] = struct{}{}

	// Start on 2
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			if _, exists := factorSet[i]; !exists {
				factorSet[i] = struct{}{}
			}
			// Also check that we have x for i * x == N
			x := n / i
			if _, exists := factorSet[x]; !exists {
				factorSet[x] = struct{}{}
			}
		}
	}

	return slices.Collect(maps.Keys(factorSet))
}

// IsRepeated returns true if the input string is made of a sequence that is repeated
// at least twice.
// E.g. "1212", "123123123"
func IsRepeated(input string) bool {
	size := len(input)
	if size <= 1 {
		return false
	}
	uniqueFactors := factors(size)
	for _, f := range uniqueFactors {
		if IsRepeatedBy(input, f) {
			return true
		}
	}
	return false
}

// GetRange accepts a string with signature "<int1>-<int2>" and parses it to
// create an array of numerical strings between int1 and int2.
func GetRange(input string) ([]string, error) {
	left, right, err := parseRange(input)
	if err != nil {
		return nil, fmt.Errorf("parsing range: %w", err)
	}

	// prealloc arrange
	output := make([]string, right-left+1)
	for i := 0; i < right-left+1; i++ {
		output[i] = strconv.Itoa(i + left)
	}
	return output, nil
}

// parseRange expects an input of the format "<int1>-<int2>" and returns
// both the left and right integer values. Returns an error if it fails to parse.
func parseRange(input string) (int, int, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return 0, 0, ErrInvalidInput
	}

	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, ErrInvalidInput
	}
	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, ErrInvalidInput
	}

	return left, right, nil
}

func Part1(input []string) (int, error) {
	ans := 0

	for _, r := range input {
		interval, err := GetRange(r)
		if err != nil {
			return 0, fmt.Errorf("getting range %q: %w", r, err)
		}
		for _, num := range interval {
			if IsRepeatedTwice(num) {
				n, err := strconv.Atoi(num)
				if err != nil {
					return 0, fmt.Errorf("converting numeric string %q to int: %w", num, err)
				}
				ans += n
			}
		}
	}
	return ans, nil
}

func Part2(input []string) (int, error) {
	ans := 0

	for _, r := range input {
		interval, err := GetRange(r)
		if err != nil {
			return 0, fmt.Errorf("getting range %q: %w", r, err)
		}
		for _, num := range interval {
			if IsRepeated(num) {
				n, err := strconv.Atoi(num)
				if err != nil {
					return 0, fmt.Errorf("converting numeric string %q to int: %w", num, err)
				}
				ans += n
			}
		}
	}
	return ans, nil
}

func main() {
	data, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		panic(err)
	}
	// Parse input
	input := strings.Split(strings.TrimSpace(string(data)), ",")

	p1, err := Part1(input)
	if err != nil {
		panic(err)
	}

	p2, err := Part2(input)
	if err != nil {
		panic(err)
	}

	// Results:
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
