package main_test

import (
	"testing"

	main "github.com/sotiri-geo/aod-2025/day04"
)

// Create a struct that holds point to x, y positions
// holds the boolean if it has roll
// has methods to check if roll is to the: left, right, up, down, upLeft, upRight, downLeft, downRight

func TestGrid(t *testing.T) {
	t.Run("PointHasRollCanAccessRollOfPaper", func(t *testing.T) {
		grid := main.NewGrid([][]byte{
			{'@', '.'},
			{'.', '@'},
		})

		got := grid.CanAccessRollOfPaper(0, 0)

		if got != true {
			t.Errorf("got %v, want %v", got, true)
		}
	})
	t.Run("PointHasRollCanNotAccessRollOfPaper", func(t *testing.T) {
		grid := main.NewGrid([][]byte{
			{'.', '@', '.'},
			{'@', '@', '@'},
			{'.', '@', '.'},
		})

		got := grid.CanAccessRollOfPaper(1, 1)

		if got != false {
			t.Errorf("got %v, want %v", got, false)
		}
	})
	t.Run("PointHasNoRollCanNotAccessRollOfPaper", func(t *testing.T) {
		grid := main.NewGrid([][]byte{
			{'.', '@', '.'},
			{'@', '@', '@'},
			{'.', '@', '.'},
		})

		got := grid.CanAccessRollOfPaper(0, 0)

		if got != false {
			t.Errorf("got %v, want %v", got, false)
		}
	})
	t.Run("CanNotAccessRollOfPaperWithEmptyGrid", func(t *testing.T) {
		grid := main.NewGrid([][]byte{})

		got := grid.CanAccessRollOfPaper(0, 0)

		if got != false {
			t.Errorf("got %v, want %v", got, false)
		}
	})
	t.Run("ProcessRollsOfPaper", func(t *testing.T) {
		grid := main.NewGrid([][]byte{
			[]byte("..@@.@@@@."),
			[]byte("@@@.@.@.@@"),
			[]byte("@@@@@.@.@@"),
			[]byte("@.@@@@..@."),
			[]byte("@@.@@@@.@@"),
			[]byte(".@@@@@@@.@"),
			[]byte(".@.@.@.@@@"),
			[]byte("@.@@@.@@@@"),
			[]byte(".@@@@@@@@."),
			[]byte("@.@.@@@.@."),
		})

		got := grid.ProcessRollOfPaper()

		if got != 13 {
			t.Errorf("got %d, want %d", got, 13)
		}
	})
}
