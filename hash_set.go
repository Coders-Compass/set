package set

import (
	"fmt"
	"sort"
	"strings"
)

// hashSet implements the Set interface using a map for O(1) operations.
// The empty struct value for map entries uses no additional memory.
type hashSet[T comparable] struct {
	elements map[T]struct{}
}

// NewHashSet creates and returns a new empty set.
func NewHashSet[T comparable]() Set[T] {
	return &hashSet[T]{
		elements: make(map[T]struct{}),
	}
}

func (h *hashSet[T]) Insert(elem T) {
	h.elements[elem] = struct{}{}
}

func (h *hashSet[T]) Remove(elem T) {
	delete(h.elements, elem)
}

func (h *hashSet[T]) Contains(elem T) bool {
	_, exists := h.elements[elem]
	return exists
}

func (h *hashSet[T]) Cardinality() int {
	return len(h.elements)
}

func (h *hashSet[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *hashSet[T]) ToSlice() []T {
	result := make([]T, 0, len(h.elements))
	for elem := range h.elements {
		result = append(result, elem)
	}
	return result
}

func (h *hashSet[T]) String() string {
	if h.IsEmpty() {
		return "{}"
	}

	elems := h.ToSlice()

	// We'll sort the elements if they are strings.
	if len(elems) > 0 {
		if _, ok := any(elems[0]).(string); ok {
			strElems := make([]string, len(elems))
			for i, v := range elems {
				strElems[i] = any(v).(string)
			}
			sort.Strings(strElems)
			for i, v := range strElems {
				elems[i] = any(v).(T)
			}
		}
	}

	var sb strings.Builder
	sb.WriteString("{")
	for i, elem := range elems {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", elem))
	}
	sb.WriteString("}")
	return sb.String()
}

func (h *hashSet[T]) Equals(other Set[T]) bool {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}

	if h.Cardinality() != otherSet.Cardinality() {
		return false
	}

	for elem := range h.elements {
		if !otherSet.Contains(elem) {
			return false
		}
	}
	return true
}

func (h *hashSet[T]) IsSubsetOf(other Set[T]) bool {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}

	if h.Cardinality() > otherSet.Cardinality() {
		return false
	}

	for elem := range h.elements {
		if !otherSet.Contains(elem) {
			return false
		}
	}
	return true
}

func (h *hashSet[T]) IsSupersetOf(other Set[T]) bool {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}
	return otherSet.IsSubsetOf(h)
}

func (h *hashSet[T]) IsProperSubsetOf(other Set[T]) bool {
	return h.Cardinality() < other.(*hashSet[T]).Cardinality() && h.IsSubsetOf(other)
}

func (h *hashSet[T]) IsProperSupersetOf(other Set[T]) bool {
	return h.Cardinality() > other.(*hashSet[T]).Cardinality() && h.IsSupersetOf(other)
}

func (h *hashSet[T]) Union(other Set[T]) Set[T] {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}

	unionSet := NewHashSet[T]()
	for elem := range h.elements {
		unionSet.Insert(elem)
	}
	for elem := range otherSet.elements {
		unionSet.Insert(elem)
	}
	return unionSet
}

func (h *hashSet[T]) Intersection(other Set[T]) Set[T] {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}

	cardinality1 := h.Cardinality()
	cardinality2 := otherSet.Cardinality()

	intersectionSet := NewHashSet[T]()

	smaller, larger := h, otherSet
	if cardinality1 > cardinality2 {
		smaller, larger = otherSet, h
	}

	for elem := range smaller.elements {
		if larger.Contains(elem) {
			intersectionSet.Insert(elem)
		}
	}

	return intersectionSet
}

func (h *hashSet[T]) Difference(other Set[T]) Set[T] {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}

	differenceSet := NewHashSet[T]()
	for elem := range h.elements {
		if !otherSet.Contains(elem) {
			differenceSet.Insert(elem)
		}
	}
	return differenceSet
}

func (h *hashSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	otherSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Operation only supported between hashSets")
	}

	// Simple implementation:
	// difference1 := h.Difference(otherSet)
	// difference2 := otherSet.Difference(h)
	// return difference1.Union(difference2)

	// One-pass implementation:
	symmetricDifferenceSet := NewHashSet[T]()
	for elem := range h.elements {
		if !otherSet.Contains(elem) {
			symmetricDifferenceSet.Insert(elem)
		}
	}
	for elem := range otherSet.elements {
		if !h.Contains(elem) {
			symmetricDifferenceSet.Insert(elem)
		}
	}
	return symmetricDifferenceSet
}
