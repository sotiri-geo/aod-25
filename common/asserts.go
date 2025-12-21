// This package is to provide utilities for assertions
package common

import "testing"

func assertEqual[T comparable](t *testing.T, got, want T, msg string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v: %s", got, want, msg)
	}
}
