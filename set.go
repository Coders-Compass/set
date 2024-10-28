package set

type Set[T comparable] interface {
	Insert(elem T)
	Intersection(other Set[T]) Set[T]
	Equals(other Set[T]) bool
}

type hashSet[T comparable] struct {
	elements map[T]struct{}
}

func (h *hashSet[T]) Insert(elem T) {
	h.elements[elem] = struct{}{}
}

func (h *hashSet[T]) Intersection(other Set[T]) Set[T] {
	// Ensure that we are only working with other hash sets
	otherHashSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Intersection only supported between hashSets.")
	}

	result := NewHashSet[T]()

	var smaller, larger *hashSet[T]
	if len(h.elements) > len(otherHashSet.elements) {
		smaller, larger = otherHashSet, h
	} else {
		smaller, larger = h, otherHashSet
	}

	for elem := range smaller.elements {
		if _, exists := larger.elements[elem]; exists {
			result.Insert(elem)
		}
	}

	return result
}

func (h *hashSet[T]) Equals(other Set[T]) bool {
	otherHashSet, ok := other.(*hashSet[T])
	if !ok {
		panic("Intersection only supported between hashSets.")
	}

	if len(h.elements) != len(otherHashSet.elements) {
		return false
	}

	for elem := range h.elements {
		if _, exists := otherHashSet.elements[elem]; !exists {
			return false
		}
	}

	return true
}

func NewHashSet[T comparable]() Set[T] {
	return &hashSet[T]{
		elements: make(map[T]struct{}),
	}
}
