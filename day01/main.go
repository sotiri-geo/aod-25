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

type Incrementor interface {
	mustIncrement(current, amount int) (int, int)
}

// ZeroIncrementor will increment when dial is at zero
type ZeroIncrementor struct {
	capacity int
}

// mustIncrement returns incrementCount and remainder
func (i *ZeroIncrementor) mustIncrement(current, amount int) (int, int) {
	r := (current + amount) % i.capacity
	if r == 0 {
		return 1, 0
	}
	return 0, r
}

type PassThroughZeroIncrementor struct {
	capacity int
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

// mustIncrement returns incrementCount and remainder
func (i *PassThroughZeroIncrementor) mustIncrement(current, amount int) (int, int) {
	// We need to factor in if current > 0 and then the amount brings below 0
	// this means we havent accounted for the dial going through zero
	additional := 0
	if current > 0 {
		additional = 1
	}
	total := current + amount
	// early exit
	if total == 0 {
		return 1, 0
	}
	q, r := total/i.capacity, mod(total, i.capacity)
	if total < 0 {
		// truncates towards zero so need to turn positive + 1
		q := (-q) + additional
		return q, r
	}
	return q, r
}

func NewZeroIncrementor() *ZeroIncrementor {
	return &ZeroIncrementor{capacity: DialCapacity}
}

func NewPassThroughZeroIncrementor() *PassThroughZeroIncrementor {
	return &PassThroughZeroIncrementor{capacity: DialCapacity}
}

func NewDial(number int) *Dial {
	return &Dial{Number: number, capacity: DialCapacity}
}

func (d *Dial) Rotate(amount int, incrementor Incrementor) {
	// diff needs to be either below 0 or above capacity
	inc, num := incrementor.mustIncrement(d.Number, amount)
	d.Number = num
	d.Counter += inc
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
	zeroIncrementor := NewZeroIncrementor()
	dial := NewDial(50)

	for _, rotation := range input {
		amount, err := ParseRotation(rotation)

		if err != nil {
			return 0, err
		}
		dial.Rotate(amount, zeroIncrementor)
	}
	return dial.Counter, nil
}

func Part2(input []string) (int, error) {
	passThroughZeroIncrementor := NewPassThroughZeroIncrementor()
	dial := NewDial(50)

	for _, rotation := range input {
		amount, err := ParseRotation(rotation)

		if err != nil {
			return 0, err
		}
		dial.Rotate(amount, passThroughZeroIncrementor)
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

	d2, err := Part2(input)

	if err != nil {
		panic(err)
	}

	// Pipe out results to stdout
	fmt.Println("Part 1:", d1)
	fmt.Println("Part 2:", d2)

}
