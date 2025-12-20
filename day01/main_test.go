package main

import (
	"testing"
)

func TestDialRotation(t *testing.T) {
	testCases := map[string]struct {
		Rotation int
		Want     int
	}{
		"turns right without cycle": {
			Rotation: 5,
			Want:     55,
		},
		"turns left without cycle": {
			Rotation: -5,
			Want:     45,
		},
		"turns right with cycle": {
			Rotation: 55,
			Want:     5,
		},
		"turns left with cycle": {
			Rotation: 60,
			Want:     10,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			dial := NewDial(50)
			dial.Rotate(tc.Rotation)

			if dial.Number != tc.Want {
				t.Errorf("got %d, want %d", dial.Number, tc.Want)
			}
		})
	}
}

func TestDialCount(t *testing.T) {
	t.Run("adds 1 to count when dial rotates back to 0", func(t *testing.T) {
		// GIVEN dial d initially pointing at 50 with count 0
		dial := NewDial(50)
		if dial.Counter != 0 {
			t.Fatalf("should initialise counter to 0")
		}

		// WHEN dial d rotates right by 50
		dial.Rotate(50)

		// THEN dial d
		// - increments count by 1 after doing a full cycle back to 0
		if dial.Number != 0 {
			t.Fatalf("should reset back to zero after full cycle")
		}

		if dial.Counter != 1 {
			t.Errorf("should have increased counter by 1")
		}
	})
}

func TestParseRotation(t *testing.T) {
	t.Run("parses left rotation", func(t *testing.T) {
		got, err := ParseRotation("L20")
		want := -20

		if err != nil {
			t.Fatalf("should not error: %s", err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("parses right rotation", func(t *testing.T) {
		got, err := ParseRotation("R20")
		want := 20

		if err != nil {
			t.Fatalf("should not error: %s", err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("must be a valid direction", func(t *testing.T) {
		_, err := ParseRotation("P20")

		if err == nil {
			t.Fatalf("should return error")
		}
	})
}

func TestPart1(t *testing.T) {
	input := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	want := 3

	got, err := Part1(input)

	if err != nil {
		t.Fatalf("should not error: %s", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
