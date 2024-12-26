package set

import "fmt"

type Pair[T comparable] struct {
	First  T
	Second T
}

func (p Pair[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

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
