package common

import (
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("NewSetDeduplicatesInitialValues", func(t *testing.T) {
		s := NewSet([]int{1, 1, 2, 3, 3})

		if got, want := s.Len(), 3; got != want {
			t.Errorf("got len %d, want %d", got, want)
		}

		for _, v := range []int{1, 2, 3} {
			if !s.Has(v) {
				t.Errorf("expected set to contain %d", v)
			}
		}
	})

	t.Run("AddNewElementIncreasesLen", func(t *testing.T) {
		s := NewSet([]int{})

		s.Add(10)

		if got, want := s.Len(), 1; got != want {
			t.Errorf("got len %d, want %d", got, want)
		}
		if !s.Has(10) {
			t.Errorf("expected set to contain %d", 10)
		}
	})

	t.Run("AddExistingElementDoesNotIncreaseLen", func(t *testing.T) {
		s := NewSet([]int{10})

		s.Add(10)

		if got, want := s.Len(), 1; got != want {
			t.Errorf("got len %d, want %d", got, want)
		}
	})

	t.Run("RemoveExistingElement", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})

		s.Remove(2)

		if s.Has(2) {
			t.Errorf("did not expect set to contain %d", 2)
		}
		if got, want := s.Len(), 2; got != want {
			t.Errorf("got len %d, want %d", got, want)
		}
	})

	t.Run("RemoveMissingElementDoesNotChangeLen", func(t *testing.T) {
		s := NewSet([]int{1, 2})

		s.Remove(999)

		if got, want := s.Len(), 2; got != want {
			t.Errorf("got len %d, want %d", got, want)
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		s := NewSet([]int{})

		if !s.IsEmpty() {
			t.Errorf("expected empty set")
		}

		s.Add(1)

		if s.IsEmpty() {
			t.Errorf("expected non-empty set")
		}
	})

	t.Run("InitConstructorWithZeroValueSlice", func(t *testing.T) {
		var set []int
		s := NewSet(set)

		s.Add(1)

		if got, want := s.Len(), 1; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
