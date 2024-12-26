package set

import "fmt"

// Pair represents an ordered pair of elements for use in Cartesian products.
type Pair[T comparable] struct {
	First  T
	Second T
}

// String returns a string representation of the pair in the format "(First, Second)".
func (p Pair[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// CartesianProduct returns a new set containing all possible ordered pairs
// (x, y) where x is from the first set and y is from the second set.
// For sets **A** and **B**, it returns **A** × **B** = {(x, y) | x ∈ **A**, y ∈ **B**}.
//
// For example, if **A = {1, 2}** and **B = {3, 4}**, then
// CartesianProduct(**A**, **B**) = **{(1, 3), (1, 4), (2, 3), (2, 4)}**.
func CartesianProduct[T comparable](s1, s2 Set[T]) Set[Pair[T]] {
	result := NewHashSet[Pair[T]]()

	s1Elems := s1.ToSlice()
	s2Elems := s2.ToSlice()

	for _, elem1 := range s1Elems {
		for _, elem2 := range s2Elems {
			result.Insert(Pair[T]{
				First:  elem1,
				Second: elem2,
			})
		}
	}

	return result
}
