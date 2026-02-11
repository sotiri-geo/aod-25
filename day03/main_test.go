package main_test

import (
	"testing"

	main "github.com/sotiri-geo/aod-2025/day03"
)

func TestMaxTwoDigit(t *testing.T) {
	testCases := map[string]struct {
		Input []int
		Want  int
	}{
		"increasingBank": {
			Input: []int{1, 2, 3},
			Want:  23,
		},
		"decreasingBank": {
			Input: []int{3, 2, 1},
			Want:  32,
		},
		"unorderedBank": {
			Input: []int{4, 5, 7, 3},
			Want:  73,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := main.MaxTwoDigit(tc.Input)

			if got != tc.Want {
				t.Errorf("got %d, want %d", got, tc.Want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := [][]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
		{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
		{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
		{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
	}

	want := 357

	got := main.Part1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMaxTwelveDigit(t *testing.T) {
	testCases := map[string]struct {
		Input []int
		Want  int
	}{
		"decreasingBank": {
			Input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			Want:  987654321111,
		},
		"unorderedBank": {
			Input: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			Want:  434234234278,
		},
		"repeatedDigitBank": {
			Input: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			Want:  888911112111,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := main.MaxTwelveDigit(tc.Input)

			if got != tc.Want {
				t.Errorf("got %d, want %d", got, tc.Want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	input := [][]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
		{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
		{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
		{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
	}

	want := 3121910778619

	got := main.Part2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
