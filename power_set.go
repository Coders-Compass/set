package set

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
