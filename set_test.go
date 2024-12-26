package set

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestStringSets(t *testing.T) {
	s1 := NewHashSet[string]()
	s1.Insert("Frodo")
	s1.Insert("Sam")
	s1.Insert("Took")
	s1.Insert("Merry")

	s2 := NewHashSet[string]()
	s2.Insert("Frodo")
	s2.Insert("Sam")
	s2.Insert("Gandalf")
	s2.Insert("Legolas")

	s3 := NewHashSet[string]()
	s3.Insert("Frodo")
	s3.Insert("Sam")

	intersection := s1.Intersection(s2)

	if !intersection.Equals(s3) {
		t.Fatalf("Intersection did not work for string sets.")
	}
}

func TestEmptyIntersections(t *testing.T) {
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "empty intersection empty",
			set1:     []int{},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "empty intersection non-empty",
			set1:     []int{},
			set2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "non-empty intersection empty",
			set1:     []int{1, 2, 3},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "disjoint sets",
			set1:     []int{1, 2, 3},
			set2:     []int{4, 5, 6},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewHashSet[int]()
			s2 := NewHashSet[int]()
			expected := NewHashSet[int]()

			for _, v := range tt.set1 {
				s1.Insert(v)
			}
			for _, v := range tt.set2 {
				s2.Insert(v)
			}
			for _, v := range tt.expected {
				expected.Insert(v)
			}

			result := s1.Intersection(s2)
			if !result.Equals(expected) {
				t.Errorf("%s: expected empty set, got non-empty set", tt.name)
			}
		})
	}
}

func TestSmallSets(t *testing.T) {
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "single element sets with intersection",
			set1:     []int{1},
			set2:     []int{1},
			expected: []int{1},
		},
		{
			name:     "single element sets without intersection",
			set1:     []int{1},
			set2:     []int{2},
			expected: []int{},
		},
		{
			name:     "small sets with partial overlap",
			set1:     []int{1, 2, 3},
			set2:     []int{2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "small sets with complete overlap",
			set1:     []int{1, 2, 3},
			set2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewHashSet[int]()
			s2 := NewHashSet[int]()
			expected := NewHashSet[int]()

			for _, v := range tt.set1 {
				s1.Insert(v)
			}
			for _, v := range tt.set2 {
				s2.Insert(v)
			}
			for _, v := range tt.expected {
				expected.Insert(v)
			}

			result := s1.Intersection(s2)
			if !result.Equals(expected) {
				t.Errorf("%s: intersection result did not match expected", tt.name)
			}
		})
	}
}

func TestLargeSets(t *testing.T) {
	tests := []struct {
		name    string
		size1   int
		size2   int
		overlap float64 // percentage of overlap between sets
	}{
		{
			name:    "large sets with no overlap",
			size1:   1000,
			size2:   1000,
			overlap: 0,
		},
		{
			name:    "large sets with 50% overlap",
			size1:   1000,
			size2:   1000,
			overlap: 0.5,
		},
		{
			name:    "large sets with complete overlap",
			size1:   1000,
			size2:   1000,
			overlap: 1.0,
		},
		{
			name:    "asymmetric large sets with partial overlap",
			size1:   100,
			size2:   1000,
			overlap: 0.3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewHashSet[int]()
			s2 := NewHashSet[int]()
			expected := NewHashSet[int]()

			// Calculate overlap size
			overlapSize := int(float64(tt.size1) * tt.overlap)

			// Fill first set and overlapping elements
			for i := 0; i < tt.size1; i++ {
				s1.Insert(i)
				if i < overlapSize {
					expected.Insert(i)
				}
			}

			// Fill second set
			for i := 0; i < overlapSize; i++ {
				s2.Insert(i)
			}
			for i := overlapSize; i < tt.size2; i++ {
				s2.Insert(i + tt.size1) // Ensure no unintended overlap
			}

			result := s1.Intersection(s2)
			if !result.Equals(expected) {
				t.Errorf("%s: intersection size expected %d, got %d",
					tt.name,
					getSize(expected),
					getSize(result))
			}
		})
	}
}

// Helper function to get size of a set
func getSize[T comparable](s Set[T]) int {
	// Type assert to access the underlying map
	if hs, ok := s.(*hashSet[T]); ok {
		return len(hs.elements)
	}
	return 0
}

func TestSpecialValues(t *testing.T) {
	tests := []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected []interface{}
	}{
		{
			name:     "integer edge values",
			set1:     []interface{}{math.MaxInt64, math.MinInt64, 0},
			set2:     []interface{}{math.MaxInt64, 0},
			expected: []interface{}{math.MaxInt64, 0},
		},
		{
			name:     "zero values",
			set1:     []interface{}{"", 0, false},
			set2:     []interface{}{"", 0, false, "non-zero"},
			expected: []interface{}{"", 0, false},
		},
		{
			name:     "time values",
			set1:     []interface{}{time.Time{}, time.Now()},
			set2:     []interface{}{time.Time{}},
			expected: []interface{}{time.Time{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewHashSet[interface{}]()
			s2 := NewHashSet[interface{}]()
			expected := NewHashSet[interface{}]()

			for _, v := range tt.set1 {
				s1.Insert(v)
			}
			for _, v := range tt.set2 {
				s2.Insert(v)
			}
			for _, v := range tt.expected {
				expected.Insert(v)
			}

			result := s1.Intersection(s2)
			if !result.Equals(expected) {
				t.Errorf("%s: intersection result did not match expected", tt.name)
			}
		})
	}
}

func TestImmutability(t *testing.T) {
	// Create and populate original sets
	orig1 := NewHashSet[int]()
	orig2 := NewHashSet[int]()
	for i := 0; i < 5; i++ {
		orig1.Insert(i)
		orig2.Insert(i + 3)
	}

	// Create copies for comparison
	copy1 := NewHashSet[int]()
	copy2 := NewHashSet[int]()
	for i := 0; i < 5; i++ {
		copy1.Insert(i)
		copy2.Insert(i + 3)
	}

	// Perform intersection
	intersection := orig1.Intersection(orig2)

	// Verify original sets weren't modified
	if !orig1.Equals(copy1) {
		t.Error("First set was modified during intersection")
	}
	if !orig2.Equals(copy2) {
		t.Error("Second set was modified during intersection")
	}

	// Modify intersection result and verify it doesn't affect originals
	intersection.(*hashSet[int]).elements[100] = struct{}{}
	if !orig1.Equals(copy1) || !orig2.Equals(copy2) {
		t.Error("Modifying intersection result affected original sets")
	}
}

func TestIdempotency(t *testing.T) {
	s1 := NewHashSet[int]()
	s2 := NewHashSet[int]()

	// Add some elements
	for i := 0; i < 10; i++ {
		s1.Insert(i)
		if i%2 == 0 {
			s2.Insert(i)
		}
	}

	// Perform intersection multiple times
	result1 := s1.Intersection(s2)
	result2 := s1.Intersection(s2)
	result3 := result1.Intersection(s2)

	// All results should be equal
	if !result1.Equals(result2) || !result2.Equals(result3) {
		t.Error("Intersection operation is not idempotent")
	}
}

func TestCommutativity(t *testing.T) {
	tests := []struct {
		name string
		set1 []int
		set2 []int
	}{
		{
			name: "empty sets",
			set1: []int{},
			set2: []int{},
		},
		{
			name: "different sized sets",
			set1: []int{1, 2, 3},
			set2: []int{2, 3, 4, 5, 6},
		},
		{
			name: "large difference in size",
			set1: []int{1},
			set2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewHashSet[int]()
			s2 := NewHashSet[int]()

			for _, v := range tt.set1 {
				s1.Insert(v)
			}
			for _, v := range tt.set2 {
				s2.Insert(v)
			}

			result1 := s1.Intersection(s2)
			result2 := s2.Intersection(s1)

			if !result1.Equals(result2) {
				t.Errorf("%s: intersection is not commutative", tt.name)
			}
		})
	}
}

func BenchmarkIntersection(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	overlaps := []float64{0.0, 0.5, 1.0}

	for _, size := range sizes {
		for _, overlap := range overlaps {
			name := fmt.Sprintf("size=%d,overlap=%.1f", size, overlap)
			b.Run(name, func(b *testing.B) {
				s1 := NewHashSet[int]()
				s2 := NewHashSet[int]()

				overlapSize := int(float64(size) * overlap)

				// Setup sets with specified overlap
				for i := 0; i < size; i++ {
					s1.Insert(i)
					if i < overlapSize {
						s2.Insert(i)
					} else {
						s2.Insert(i + size)
					}
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = s1.Intersection(s2)
				}
			})
		}
	}
}

func TestBasicOperations(t *testing.T) {
	tests := []struct {
		name string
		ops  []struct {
			op    string
			elem  int
			check bool
		}
		expectedCard int
		isEmpty      bool
	}{
		{
			name: "empty set operations",
			ops: []struct {
				op    string
				elem  int
				check bool
			}{
				{"contains", 1, false},
				{"insert", 1, true},
				{"contains", 1, true},
				{"remove", 1, false},
				{"contains", 1, false},
			},
			expectedCard: 0,
			isEmpty:      true,
		},
		{
			name: "multiple elements",
			ops: []struct {
				op    string
				elem  int
				check bool
			}{
				{"insert", 1, true},
				{"insert", 2, true},
				{"insert", 3, true},
				{"contains", 2, true},
				{"remove", 2, false},
				{"contains", 2, false},
			},
			expectedCard: 2,
			isEmpty:      false,
		},
		{
			name: "duplicate elements",
			ops: []struct {
				op    string
				elem  int
				check bool
			}{
				{"insert", 1, true},
				{"insert", 1, true},
				{"contains", 1, true},
			},
			expectedCard: 1,
			isEmpty:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewHashSet[int]()

			for _, op := range tt.ops {
				switch op.op {
				case "insert":
					set.Insert(op.elem)
				case "remove":
					set.Remove(op.elem)
				case "contains":
					if got := set.Contains(op.elem); got != op.check {
						t.Errorf("Contains(%v) = %v, want %v", op.elem, got, op.check)
					}
				}
			}

			if got := set.Cardinality(); got != tt.expectedCard {
				t.Errorf("Cardinality() = %v, want %v", got, tt.expectedCard)
			}

			if got := set.IsEmpty(); got != tt.isEmpty {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.isEmpty)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		want     string
	}{
		{
			name:     "empty set",
			elements: []string{},
			want:     "{}",
		},
		{
			name:     "single element",
			elements: []string{"a"},
			want:     "{a}",
		},
		{
			name:     "multiple elements",
			elements: []string{"c", "a", "b"},
			want:     "{a, b, c}", // Should be sorted
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewHashSet[string]()
			for _, elem := range tt.elements {
				set.Insert(elem)
			}

			if got := set.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetOperations(t *testing.T) {
	t.Run("book genres example", func(t *testing.T) {
		sciFi := NewHashSet[string]()
		for _, book := range []string{"Dune", "Neuromancer", "Altered Carbon", "The Caves of Steel"} {
			sciFi.Insert(book)
		}

		mystery := NewHashSet[string]()
		for _, book := range []string{"The Hound of the Baskervilles", "Altered Carbon", "Gone Girl", "The Caves of Steel"} {
			mystery.Insert(book)
		}

		// Test Union
		union := sciFi.Union(mystery)
		expectedUnion := []string{
			"Dune", "Neuromancer", "Altered Carbon", "The Caves of Steel",
			"The Hound of the Baskervilles", "Gone Girl",
		}
		for _, book := range expectedUnion {
			if !union.Contains(book) {
				t.Errorf("Union should contain %s", book)
			}
		}
		if union.Cardinality() != len(expectedUnion) {
			t.Errorf("Union should have %d elements, got %d", len(expectedUnion), union.Cardinality())
		}

		// Test Difference (sciFi \ mystery)
		diff := sciFi.Difference(mystery)
		expectedDiff := []string{"Dune", "Neuromancer"}
		for _, book := range expectedDiff {
			if !diff.Contains(book) {
				t.Errorf("Difference should contain %s", book)
			}
		}
		if diff.Cardinality() != len(expectedDiff) {
			t.Errorf("Difference should have %d elements, got %d", len(expectedDiff), diff.Cardinality())
		}

		// Test Symmetric Difference
		symDiff := sciFi.SymmetricDifference(mystery)
		expectedSymDiff := []string{
			"Dune", "Neuromancer",
			"The Hound of the Baskervilles", "Gone Girl",
		}
		for _, book := range expectedSymDiff {
			if !symDiff.Contains(book) {
				t.Errorf("Symmetric difference should contain %s", book)
			}
		}
		if symDiff.Cardinality() != len(expectedSymDiff) {
			t.Errorf("Symmetric difference should have %d elements, got %d", len(expectedSymDiff), symDiff.Cardinality())
		}
	})

	t.Run("special cases", func(t *testing.T) {
		empty := NewHashSet[int]()
		nonEmpty := NewHashSet[int]()
		nonEmpty.Insert(1)
		nonEmpty.Insert(2)

		unionWithEmpty := nonEmpty.Union(empty)
		if !unionWithEmpty.Equals(nonEmpty) {
			t.Error("Union with empty set should equal the non-empty set")
		}

		diffWithEmpty := nonEmpty.Difference(empty)
		if !diffWithEmpty.Equals(nonEmpty) {
			t.Error("Difference with empty set should equal the original set")
		}

		emptyDiff := empty.Difference(nonEmpty)
		if !emptyDiff.IsEmpty() {
			t.Error("Empty set difference should be empty")
		}

		symDiffWithEmpty := nonEmpty.SymmetricDifference(empty)
		if !symDiffWithEmpty.Equals(nonEmpty) {
			t.Error("Symmetric difference with empty set should equal the non-empty set")
		}
	})
}

func TestCartesianProduct(t *testing.T) {
	// Test using the outfits example from the book
	t.Run("outfits example", func(t *testing.T) {
		// Create set of shirts: Navy (N), Maroon (M), White (W)
		shirts := NewHashSet[string]()
		for _, shirt := range []string{"N", "M", "W"} {
			shirts.Insert(shirt)
		}

		// Create set of trousers: Black (BK), Brown (BN)
		trousers := NewHashSet[string]()
		for _, trouser := range []string{"BK", "BN"} {
			trousers.Insert(trouser)
		}

		// Calculate Cartesian product using the standalone function
		outfits := CartesianProduct(shirts, trousers)

		// Should have 6 possible outfits (3 shirts × 2 trousers)
		if outfits.Cardinality() != 6 {
			t.Errorf("Expected 6 outfits, got %d", outfits.Cardinality())
		}

		// Check specific outfits exist
		expectedOutfits := []Pair[string]{
			{First: "N", Second: "BK"},
			{First: "N", Second: "BN"},
			{First: "M", Second: "BK"},
			{First: "M", Second: "BN"},
			{First: "W", Second: "BK"},
			{First: "W", Second: "BN"},
		}

		for _, outfit := range expectedOutfits {
			if !outfits.Contains(outfit) {
				t.Errorf("Missing outfit: %v", outfit)
			}
		}
	})
}
