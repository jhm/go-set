// Package set provides an implementation of a set; that is, a collection of a
// comparable type that contains no duplicate elements. This implementation is
// not thread safe and is backed by a map with empty structs as values
// (map[T]struct{}).
//
// This package will be deprecated once the standard library includes a generic
// Set type.
package set

// A Set is a collection of a comparable type that contains no duplicate values.
// The zero value is usable as an empty set.
type Set[T comparable] struct {
	elements map[T]struct{}
}

// New returns a new empty set.
func New[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// Of returns a new set with the given elements.
func Of[T comparable](elements ...T) *Set[T] {
	s := &Set[T]{elements: make(map[T]struct{}, len(elements))}
	for _, e := range elements {
		s.Add(e)
	}
	return s
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.elements)
}

// Add adds the given element to the set.
func (s *Set[T]) Add(element T) {
	if s.elements == nil {
		s.elements = make(map[T]struct{})
	}
	s.elements[element] = struct{}{}
}

// AddAll adds all of the elements of the given set to the receiver set.
func (s *Set[T]) AddAll(other *Set[T]) {
	for key := range other.elements {
		s.Add(key)
	}
}

// AsSlice returns a slice that includes all of the elements in the set. Order
// is not guaranteed and should not be relied upon.
func (s *Set[T]) AsSlice() []T {
	xs := make([]T, 0, s.Len())
	for key := range s.elements {
		xs = append(xs, key)
	}
	return xs
}

// Contains returns true iff the given element is in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// ContainsAll returns true iff all of the elements of the given set are
// contained within the set.
func (s *Set[T]) ContainsAll(xs *Set[T]) bool {
	for key := range xs.elements {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

// Equal returns true iff the given set and this set are both nil or they have
// the same length and include exactly the same elements.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s == other {
		return true
	}
	if s == nil || other == nil || s.Len() != other.Len() {
		return false
	}
	return s.ContainsAll(other)
}

// IsEmpty returns true iff the set has no elements.
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// IsSubsetOf returns true iff the receiver set is a subset of the given set
// `other`.
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}
	return other.ContainsAll(s)
}

// IsProperSubsetOf returns true iff the receiver set is a proper subset of the
// given set `other`.
func (s *Set[T]) IsProperSubsetOf(other *Set[T]) bool {
	if s.Len() >= other.Len() {
		return false
	}
	return other.ContainsAll(s)
}

// IsSupersetOf returns true iff the receiver set is a superset of the given set
// `other`.
func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {
	if s.Len() < other.Len() {
		return false
	}
	return s.ContainsAll(other)
}

// IsProperSupersetOf returns true iff the receiver set is a proper superset of
// the given set `other`.
func (s *Set[T]) IsProperSupersetOf(other *Set[T]) bool {
	if s.Len() <= other.Len() {
		return false
	}
	return s.ContainsAll(other)
}

// Remove removes the given element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// RemoveAll removes all of the elements contained within the given set from the
// receiver set.
func (s *Set[T]) RemoveAll(other *Set[T]) {
	for key := range other.elements {
		s.Remove(key)
	}
}

// Difference returns a new set that contains all of the elements of the first
// set that are not in the second set.
func Difference[T comparable](a, b *Set[T]) *Set[T] {
	s := New[T]()
	s.AddAll(a)
	for key := range b.elements {
		s.Remove(key)
	}
	return s
}

// SymmetricDifference returns a new set that contains the elements that are in
// either set but not both.
func SymmetricDifference[T comparable](a, b *Set[T]) *Set[T] {
	s := New[T]()
	s.AddAll(a)
	for key := range b.elements {
		if s.Contains(key) {
			s.Remove(key)
		} else {
			s.Add(key)
		}
	}
	return s
}

// Intersection returns a new set that contains all of the elements that are in
// both sets.
func Intersection[T comparable](a, b *Set[T]) *Set[T] {
	s := New[T]()
	for key := range a.elements {
		if b.Contains(key) {
			s.Add(key)
		}
	}
	return s
}

// Union returns a new set that contains all of the elements in both sets.
func Union[T comparable](a, b *Set[T]) *Set[T] {
	s := New[T]()
	s.AddAll(a)
	s.AddAll(b)
	return s
}
