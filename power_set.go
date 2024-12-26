package set

// PowerSet returns a slice containing all possible subsets of the input set.
// For a set S, it returns P(S) = {T | T ⊆ S}.
//
// For a set with n elements, the power set contains 2^n elements.
// For example, if S = {1, 2}, then PowerSet(S) = {{}, {1}, {2}, {1, 2}}.
func PowerSet[T comparable](s Set[T]) Set[Set[T]] {
	result := NewHashSet[Set[T]]()

	result.Insert(NewHashSet[T]())

	elements := s.ToSlice()
	for _, elem := range elements {
		currentSets := result.ToSlice() // Get current subsets
		for _, subset := range currentSets {
			newSubset := NewHashSet[T]()
			subsetElems := subset.ToSlice()
			for _, e := range subsetElems {
				newSubset.Insert(e)
			}
			newSubset.Insert(elem)
			result.Insert(newSubset)
		}
	}

	return result
}
