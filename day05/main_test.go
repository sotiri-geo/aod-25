package main_test

import (
	"slices"
	"testing"

	main "github.com/sotiri-geo/aod-2025/day05"
)

func TestParseBounds(t *testing.T) {
	got := main.ParseBounds("3-5")
	want := [2]int{3, 5}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestUnionMergeFor(t *testing.T) {
	testCases := []struct {
		name   string
		bounds []main.Bounds
		want   []main.Bounds
	}{
		{
			name: "StrictlySortedOverlappingIntervals",
			bounds: []main.Bounds{
				{1, 3}, {2, 4},
			},
			want: []main.Bounds{
				{1, 4},
			},
		},
		{
			name: "TouchingSortedOverlappingIntervals",
			bounds: []main.Bounds{
				{1, 3}, {3, 4},
			},
			want: []main.Bounds{
				{1, 4},
			},
		},
		{
			name: "MultipleSortedOverlappingIntervals",
			bounds: []main.Bounds{
				{1, 2}, {3, 4}, {3, 6}, {5, 8},
			},
			want: []main.Bounds{
				{1, 2}, {3, 8},
			},
		},
		{
			name: "MultipleUnsortedOverlappingIntervals",
			bounds: []main.Bounds{
				{3, 4}, {1, 2}, {5, 8}, {3, 6},
			},
			want: []main.Bounds{
				{1, 2}, {3, 8},
			},
		},
		{
			name: "NonOverlappingSortedIntervals",
			bounds: []main.Bounds{
				{1, 2}, {4, 5}, {7, 10},
			},
			want: []main.Bounds{
				{1, 2}, {4, 5}, {7, 10},
			},
		},
		{
			name: "NonOverlappingUnsortedIntervals",
			bounds: []main.Bounds{
				{4, 5}, {1, 2}, {7, 10},
			},
			want: []main.Bounds{
				{1, 2}, {4, 5}, {7, 10},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := main.UnionMerge(tc.bounds)
			if !slices.EqualFunc(got, tc.want, func(x, y main.Bounds) bool { return x == y }) {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	freshIDRanges := []string{
		"3-5", "10-14", "16-20", "12-18",
	}
	ids := []int{1, 5, 8, 11, 17, 32}

	got := main.Part1(freshIDRanges, ids)
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
