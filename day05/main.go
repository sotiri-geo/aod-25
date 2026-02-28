package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Bounds [2]int

func (b Bounds) In(value int) bool {
	return b[0] <= value && value <= b[1]
}

func ParseBounds(bounds string) Bounds {
	b := strings.Split(bounds, "-")
	// Typically would not silence err's but I can guarentee input on this
	fst, _ := strconv.Atoi(b[0])
	snd, _ := strconv.Atoi(b[1])
	return Bounds{fst, snd}
}

func UnionMerge(bounds []Bounds) []Bounds {
	// First sort
	sort.Slice(bounds, func(i, j int) bool {
		if bounds[i][0] != bounds[j][0] {
			return bounds[i][0] < bounds[j][0]
		}
		return bounds[i][1] < bounds[j][1]
	})

	// Then we union
	// Algorithm behaves like a stack
	mergedIntervals := make([]Bounds, 0, len(bounds))

	for _, b := range bounds {
		if len(mergedIntervals) == 0 {
			mergedIntervals = append(mergedIntervals, b)
			continue
		}

		// Check for overlapping
		left := b[0]
		right := b[1]
		for len(mergedIntervals) > 0 && mergedIntervals[len(mergedIntervals)-1][1] >= left {
			top := mergedIntervals[len(mergedIntervals)-1]
			// Update unioned bounds
			right = max(right, top[1])
			left = min(left, top[0])
			// pop
			mergedIntervals = mergedIntervals[:len(mergedIntervals)-1]
		}

		// Add back to stack
		mergedIntervals = append(mergedIntervals, Bounds{left, right})
	}

	return mergedIntervals
}

func Part1(freshIDRanges []string, ids []int) int {
	total := 0

	bounds := make([]Bounds, 0, len(freshIDRanges))

	for _, r := range freshIDRanges {
		bounds = append(bounds, ParseBounds(r))
	}

	unionBounds := UnionMerge(bounds)

	for _, bound := range unionBounds {
		for _, id := range ids {
			if bound.In(id) {
				total++
			}
		}
	}
	return total
}

func main() {
	data, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		panic(err)
	}

	freshIDRanges := make([]string, 0, len(data))
	ids := make([]int, 0, len(data))

	withRanges := true
	// Parse input
	for r := range strings.SplitSeq(strings.TrimSpace(string(data)), "\n") {
		if r == "" {
			withRanges = false
			continue
		}
		if withRanges {
			freshIDRanges = append(freshIDRanges, r)
		} else {
			id, _ := strconv.Atoi(r)
			ids = append(ids, id)
		}
	}

	fmt.Println("Part1:", Part1(freshIDRanges, ids))
}
