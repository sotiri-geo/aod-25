package common_test

import (
	"errors"
	"testing"

	"github.com/sotiri-geo/aod-2025/common"
)

func TestQueue(t *testing.T) {
	t.Run("LenWithSingleElement", func(t *testing.T) {
		q := common.NewQueue([]int{1})

		got, want := q.Len(), 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("LenAfterEnqueueAndDequeue", func(t *testing.T) {
		q := common.NewQueue([]int{1, 2})
		q.Enqueue(3)

		first, err := q.Dequeue()
		if err != nil {
			t.Fatalf("should not fail first dequeue: %+v", err)
		}

		if first != 1 {
			t.Errorf("first dequeue: got %d, want %d", first, 1)
		}

		if q.Len() != 2 {
			t.Errorf("first dequeue: got len %d, want %d", q.Len(), 2)
		}

		second, err := q.Dequeue()
		if err != nil {
			t.Fatalf("should not fail second dequeue: %+v", err)
		}

		if second != 2 {
			t.Errorf("second dequeue: got %d, want %d", second, 2)
		}

		if q.Len() != 1 {
			t.Errorf("second dequeue: got len %d, want %d", q.Len(), 1)
		}
	})

	t.Run("EnqueueSingleElementFromRight", func(t *testing.T) {
		q := common.NewQueue([]int{})
		q.Enqueue(1)

		if q.Len() != 1 {
			t.Errorf("got %d, want %d", q.Len(), 1)
		}
	})

	t.Run("DequeueSingleElementFromLeft", func(t *testing.T) {
		q := common.NewQueue([]int{5})
		got, err := q.Dequeue()
		if err != nil {
			t.Fatal("should not error")
		}

		if got != 5 {
			t.Errorf("got %d, want %d", got, 5)
		}
	})

	t.Run("DequeueEmptyQueue", func(t *testing.T) {
		q := common.NewQueue([]int{})
		_, err := q.Dequeue()

		if !errors.Is(err, common.ErrDequeueEmptyQueue) {
			t.Errorf("got error %q, wanted %q", err, common.ErrDequeueEmptyQueue)
		}
	})

	t.Run("DequeueMultipleTimes", func(t *testing.T) {
		q := common.NewQueue([]int{1, 2})
		gotFirst, err := q.Dequeue()
		if err != nil {
			t.Fatalf("should not fail after first dequeue: %+v", err)
		}

		if gotFirst != 1 {
			t.Errorf("got %d, want %d", gotFirst, 1)
		}
		gotSecond, err := q.Dequeue()
		if err != nil {
			t.Fatalf("should not fail after second dequeue: %+v", err)
		}

		if gotSecond != 2 {
			t.Errorf("got %d, want %d", gotSecond, 2)
		}
	})

	t.Run("IsEmptyOnEmptyQueue", func(t *testing.T) {
		q := common.NewQueue([]int{})

		if !q.IsEmpty() {
			t.Errorf("got %v, want %v", q.IsEmpty(), true)
		}
	})

	t.Run("IsEmptyOnNonEmtpyQueue", func(t *testing.T) {
		q := common.NewQueue([]int{1})

		if q.IsEmpty() {
			t.Errorf("got %v, want %v", q.IsEmpty(), false)
		}
	})

	t.Run("PeekWithEmptyQueue", func(t *testing.T) {
		q := common.NewQueue([]int{})

		_, err := q.Peek()

		if !errors.Is(err, common.ErrPeakEmptyQueue) {
			t.Errorf("got err %+v, want %+v", err, common.ErrPeakEmptyQueue)
		}
	})

	t.Run("PeekWithNonEmptyQueue", func(t *testing.T) {
		q := common.NewQueue([]int{1})

		got, err := q.Peek()
		if err != nil {
			t.Fatalf("should not error when peaking: %+v", err)
		}

		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
	})
}
