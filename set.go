// Package set provides a generic set data interface and implementations of this interface.
//
// The interface is based on the mathematical definition of a set and provides operations like
// union, intersection, difference, and subset relationships. The package also provides a map-based
// in-memory implementation of the Set interface called hashSet.
//
// The CartesianProduct and PowerSet functions are also provided separately to avoid type
// dependency cycles while still maintaining the complete set of operations from set theory.
package set

// Set is a generic interface that defines the operations that can be performed on a set.
// A set is defined as an unordered collection of unique, arbitrary elements.
// The zero value of a set is an empty set.
type Set[T comparable] interface {
	// Insert adds the element to the set.
	// If the element already exists, the set remains unchanged.
	Insert(elem T)

	// Remove deletes the element from the set.
	// If the element does not exist, the set remains unchanged.
	Remove(elem T)

	// Contains reports whether the element exists in the set.
	Contains(elem T) bool

	// Cardinality returns the number of elements (a natural, counting number) in this finite set.
	Cardinality() int

	// IsEmpty reports whether the set has no elements.
	IsEmpty() bool

	// Equals reports whether this set contains exactly the same elements as the other set.
	Equals(other Set[T]) bool

	// IsSubsetOf reports whether this set is a subset of the other set.
	// A set **X** is a subset of set **Y** if every element of **X** is also an element of **Y**.
	IsSubsetOf(other Set[T]) bool

	// IsSupersetOf reports whether this set is a superset of the other set.
	// A set **X** is a superset of set **Y** if every element of **Y** is also an element of **X**.
	IsSupersetOf(other Set[T]) bool

	// IsProperSubsetOf reports whether this set is a proper subset of the other set.
	// A set **X** is a proper subset of set **Y** if **X** is subset of **Y** and **X** ≠ **Y**.
	IsProperSubsetOf(other Set[T]) bool

	// IsProperSupersetOf reports whether this set is a proper superset of the other set.
	// A set **X** is a proper superset of set **Y** if **X** is superset of **Y** and **X** ≠ **Y**.
	IsProperSupersetOf(other Set[T]) bool

	// Union returns a new set containing all elements that are in either this set
	// or the other set (**X** ∪ **Y**).
	Union(other Set[T]) Set[T]

	// Intersection returns a new set containing all elements that are in both this set
	// and the other set (**X** ∩ **Y**).
	Intersection(other Set[T]) Set[T]

	// Difference returns a new set containing all elements that are in this set
	// but not in the other set (**X** \ **Y**).
	Difference(other Set[T]) Set[T]

	// SymmetricDifference returns a new set containing all elements that are in either this set
	// or the other set, but not in both (**X** Δ **Y**).
	SymmetricDifference(other Set[T]) Set[T]

	// ToSlice returns a slice containing all elements in the set.
	// **Note**: The order of elements is not guaranteed to be stable between calls.
	ToSlice() []T

	// String returns a string representation of the set.
	// For sets of strings, the elements are sorted lexicographically.
	String() string
}
