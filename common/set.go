package common

type Set[T comparable] struct {
	Set map[T]bool
}

// NewSet is a constructor for type Set. Limited to constructor-only usage
func NewSet[T comparable](s []T) *Set[T] {
	// Capacity hint to minimise the no. rehashing from prealloc
	newSet := &Set[T]{make(map[T]bool, len(s))}

	for _, v := range s {
		newSet.Set[v] = true
	}

	return newSet
}

// Len returns the no. of elements in Set
func (s *Set[T]) Len() int {
	return len(s.Set)
}

// Remove will remove element v from Set and ignore otherwise
func (s *Set[T]) Remove(v T) {
	if s.Set[v] {
		delete(s.Set, v)
	}
}

// Has returns true if v in Set
func (s *Set[T]) Has(v T) bool {
	return s.Set[v]
}

// IsEmpty returns true if Set is empty
func (s *Set[T]) IsEmpty() bool {
	return len(s.Set) == 0
}

// Add will add element v to Set
func (s *Set[T]) Add(v T) {
	s.Set[v] = true
}
