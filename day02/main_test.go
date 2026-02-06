package main_test

import (
	"errors"
	"slices"
	"testing"

	main "github.com/sotiri-geo/aod-2025/day02"
)

func TestIsRepeatedTwice(t *testing.T) {
	testCases := map[string]struct {
		Input string
		Want  bool
	}{
		"single character": {
			Input: "1",
			Want:  false,
		},
		"odd length is non repeated": {
			Input: "122",
			Want:  false,
		},
		"non repeated even characters": {
			Input: "12",
			Want:  false,
		},
		"repeated even characters": {
			Input: "1212",
			Want:  true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := main.IsRepeatedTwice(tc.Input)

			if got != tc.Want {
				t.Errorf("got %v, want %v", got, tc.Want)
			}
		})
	}
}

func TestGetRange(t *testing.T) {
	t.Run("range of length 1", func(t *testing.T) {
		got, err := main.GetRange("1-1")
		want := []string{"1"}
		if err != nil {
			t.Fatalf("should not error: %v", err)
		}

		if !slices.Equal(got, want) {
			t.Errorf("got range %+v, want %+v", got, want)
		}
	})

	t.Run("range of length 3", func(t *testing.T) {
		got, err := main.GetRange("1-3")
		want := []string{"1", "2", "3"}

		if err != nil {
			t.Fatalf("should not error: %v", err)
		}

		if !slices.Equal(got, want) {
			t.Errorf("got range %+v, want %+v", got, want)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := main.GetRange("~~~")

		if !errors.Is(err, main.ErrInvalidInput) {
			t.Errorf("should error with %v, got %v", main.ErrInvalidInput, err)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		_, err := main.GetRange("")

		if !errors.Is(err, main.ErrInvalidInput) {
			t.Errorf("should error with %v, got %v", main.ErrInvalidInput, err)
		}
	})
	t.Run("non numeric boundaries", func(t *testing.T) {
		_, err := main.GetRange("a-b")

		if !errors.Is(err, main.ErrInvalidInput) {
			t.Errorf("should error with %v, got %v", main.ErrInvalidInput, err)
		}
	})
}

func TestIsRepeatedBy(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Factor int
		Want   bool
	}{
		"repeated by a factor of 2": {
			Input:  "121212",
			Factor: 2,
			Want:   true,
		},
		"repeated by factor of 3": {
			Input:  "123123",
			Factor: 3,
			Want:   true,
		},
		"not repeated": {
			Input:  "124123",
			Factor: 3,
			Want:   false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := main.IsRepeatedBy(tc.Input, tc.Factor)

			if got != tc.Want {
				t.Errorf("got %v, want %v", got, tc.Want)
			}
		})
	}
}

func TestIsRepeated(t *testing.T) {
	testCases := map[string]struct {
		Input string
		Want  bool
	}{
		"hasRepeatedSequence": {
			Input: "121212",
			Want:  true,
		},
		"notRepeatingSequence": {
			Input: "121412",
			Want:  false,
		},
		"singleCharacter": {
			Input: "1",
			Want:  false,
		},
		"empty": {
			Input: "",
			Want:  false,
		},
		"singleFactorRepetition": {
			Input: "111",
			Want:  true,
		},
		"longRepeatingSequence": {
			Input: "1188511885",
			Want:  true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := main.IsRepeated(tc.Input)

			if got != tc.Want {
				t.Errorf("for sequence %q: got %v, want %v", tc.Input, got, tc.Want)
			}
		})
	}
}
