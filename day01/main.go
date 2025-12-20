package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DialCapacity = 100

type Dial struct {
	Number   int
	Counter  int
	capacity int
}

func NewDial(number int) *Dial {
	return &Dial{Number: number, capacity: DialCapacity}
}

func (d *Dial) Rotate(amount int) {
	d.Number = (d.Number + amount) % d.capacity
	if d.Number == 0 {
		d.Counter++
	}
}

func ParseRotation(rotation string) (int, error) {
	if len(rotation) == 0 {
		return 0, fmt.Errorf("cannot parse rotation %s", rotation)
	}

	isRight := strings.HasPrefix(rotation, "R")
	isLeft := strings.HasPrefix(rotation, "L")

	if !isLeft && !isRight {
		return 0, fmt.Errorf("unknown direction %s", rotation)
	}
	amount, err := strconv.Atoi(rotation[1:])
	if err != nil {
		return 0, fmt.Errorf("converting to int base 10: %w", err)
	}
	if isLeft {
		return -1 * amount, nil
	}
	if isRight {
		return amount, nil
	}

	return 0, errors.New("failed to parse rotation")
}

func Part1(input []string) (int, error) {
	dial := NewDial(50)

	for _, rotation := range input {
		amount, err := ParseRotation(rotation)

		if err != nil {
			return 0, err
		}
		dial.Rotate(amount)
	}
	return dial.Counter, nil
}

func main() {
	data, err := os.ReadFile("./day01/input.txt")

	if err != nil {
		panic(err)
	}
	// Parse input
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	// Process results
	d1, err := Part1(input)

	if err != nil {
		panic(err)
	}

	// Pipe out results to stdout
	fmt.Println("Part 1:", d1)

}
