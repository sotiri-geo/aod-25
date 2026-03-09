package main_test

import (
	"slices"
	"testing"

	main "github.com/sotiri-geo/aod-2025/day06"
)

func TestRollingSum(t *testing.T) {
	t.Run("MultiplePositiveNumbers", func(t *testing.T) {
		got := main.RollingSum([]int{1, 2, 3}...)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("MultipleNegativeNumbers", func(t *testing.T) {
		got := main.RollingSum([]int{-1, -2, 3}...)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestRollingMul(t *testing.T) {
	t.Run("MultiplePositiveNumbers", func(t *testing.T) {
		got := main.RollingMul([]int{1, 2, 3}...)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("MultipleNegativeNumbers", func(t *testing.T) {
		got := main.RollingMul([]int{1, -2, 3}...)
		want := -6
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		got := main.RollingMul([]int{}...)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestExtractColumns(t *testing.T) {
	grid := [][]string{
		{"123", "328", "51", "64"},
		{"45", "64", "387", "23"},
		{"6", "98", "215", "314"},
		{"*", "+", "*", "+"},
	}
	gotOperands, gotOperators := main.ExtractColumns(grid)
	wantOperands := [][]int{
		{123, 45, 6},
		{328, 64, 98},
		{51, 387, 215},
		{64, 23, 314},
	}
	wantOperators := []string{"*", "+", "*", "+"}

	if !slices.Equal(gotOperators, wantOperators) {
		t.Errorf("got operators: %v, want %v", gotOperators, wantOperators)
	}

	if !slices.EqualFunc(gotOperands, wantOperands, func(x, y []int) bool { return slices.Equal(x, y) }) {
		t.Errorf("got operands: %v, want %v", gotOperands, wantOperands)
	}
}

func TestPart1(t *testing.T) {
	input := [][]string{
		{"123", "328", "51", "64"},
		{"45", "64", "387", "23"},
		{"6", "98", "215", "314"},
		{"*", "+", "*", "+"},
	}

	got := main.Part1(input)
	want := 4277556

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
