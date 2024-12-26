// Package set provides a generic set data interface and implementations of this interface.
//
// The interface is based on the mathematical definition of a set and provides a wide range of operations on sets.
// The package also provides a map-based in-memory implementation of the Set interface called hashSet.
package set

// Set is a generic interface that defines the operations that can be performed on a set.
// A set is defined an unordered collection of unique, arbitrary elements.
// The zero value of a set is an empty set.
//
// The CartesianProduct and PowerSet functions are also separately provided in this package.
type Set[T comparable] interface {
	Insert(elem T)
	Remove(elem T)
	Contains(elem T) bool

	Cardinality() int
	IsEmpty() bool

	Equals(other Set[T]) bool
	IsSubsetOf(other Set[T]) bool
	IsSupersetOf(other Set[T]) bool
	IsProperSubsetOf(other Set[T]) bool
	IsProperSupersetOf(other Set[T]) bool

	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]          // A \ B
	SymmetricDifference(other Set[T]) Set[T] // (A \ B) ∪ (B \ A)

	ToSlice() []T
	String() string
}
