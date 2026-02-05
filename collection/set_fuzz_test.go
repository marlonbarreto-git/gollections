package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/set"
)

func FuzzSetAddContains(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{1, 1, 1, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		s := set.Of[int]()

		for _, b := range data {
			s.Add(int(b))
		}

		for _, b := range data {
			if !s.Contains(int(b)) {
				t.Errorf("Set doesn't contain added element: %d", b)
			}
		}

		uniqueCount := 0
		seen := make(map[byte]bool)
		for _, b := range data {
			if !seen[b] {
				seen[b] = true
				uniqueCount++
			}
		}

		if s.Len() != uniqueCount {
			t.Errorf("Set has wrong size: got %d, want %d", s.Len(), uniqueCount)
		}
	})
}

func FuzzSetRemove(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, byte(3))
	f.Add([]byte{1}, byte(1))
	f.Add([]byte{}, byte(1))

	f.Fuzz(func(t *testing.T, data []byte, target byte) {
		s := set.Of[int]()
		for _, b := range data {
			s.Add(int(b))
		}

		originalLen := s.Len()
		containedBefore := s.Contains(int(target))

		s.Remove(int(target))

		if s.Contains(int(target)) {
			t.Errorf("Set still contains removed element: %d", target)
		}

		if containedBefore && s.Len() != originalLen-1 {
			t.Errorf("Set size didn't decrease after removing existing element")
		}

		if !containedBefore && s.Len() != originalLen {
			t.Error("Set size changed after removing non-existing element")
		}
	})
}

func FuzzSetUnion(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{3, 4, 5})
	f.Add([]byte{}, []byte{1, 2})
	f.Add([]byte{1, 2}, []byte{})

	f.Fuzz(func(t *testing.T, data1, data2 []byte) {
		s1 := set.Of[int]()
		s2 := set.Of[int]()

		for _, b := range data1 {
			s1.Add(int(b))
		}
		for _, b := range data2 {
			s2.Add(int(b))
		}

		union := s1.Union(s2)

		for _, b := range data1 {
			if !union.Contains(int(b)) {
				t.Errorf("Union missing element from s1: %d", b)
			}
		}
		for _, b := range data2 {
			if !union.Contains(int(b)) {
				t.Errorf("Union missing element from s2: %d", b)
			}
		}

		for k := range union {
			if !s1.Contains(k) && !s2.Contains(k) {
				t.Errorf("Union contains element not in either set: %d", k)
			}
		}
	})
}

func FuzzSetIntersect(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{2, 3, 4})
	f.Add([]byte{1, 2}, []byte{3, 4})
	f.Add([]byte{}, []byte{1, 2})

	f.Fuzz(func(t *testing.T, data1, data2 []byte) {
		s1 := set.Of[int]()
		s2 := set.Of[int]()

		for _, b := range data1 {
			s1.Add(int(b))
		}
		for _, b := range data2 {
			s2.Add(int(b))
		}

		intersect := s1.Intersect(s2)

		for k := range intersect {
			if !s1.Contains(k) {
				t.Errorf("Intersect contains element not in s1: %d", k)
			}
			if !s2.Contains(k) {
				t.Errorf("Intersect contains element not in s2: %d", k)
			}
		}

		for k := range s1 {
			if s2.Contains(k) && !intersect.Contains(k) {
				t.Errorf("Intersect missing common element: %d", k)
			}
		}
	})
}

func FuzzSetSubtract(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4}, []byte{2, 4})
	f.Add([]byte{1, 2}, []byte{1, 2})
	f.Add([]byte{}, []byte{1, 2})

	f.Fuzz(func(t *testing.T, data1, data2 []byte) {
		s1 := set.Of[int]()
		s2 := set.Of[int]()

		for _, b := range data1 {
			s1.Add(int(b))
		}
		for _, b := range data2 {
			s2.Add(int(b))
		}

		subtract := s1.Subtract(s2)

		for k := range subtract {
			if !s1.Contains(k) {
				t.Errorf("Subtract contains element not in s1: %d", k)
			}
			if s2.Contains(k) {
				t.Errorf("Subtract contains element that's in s2: %d", k)
			}
		}

		for k := range s1 {
			if !s2.Contains(k) && !subtract.Contains(k) {
				t.Errorf("Subtract missing element that should be included: %d", k)
			}
		}
	})
}

func FuzzSetFilter(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{2, 4, 6, 8})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		s := set.Of[int]()
		for _, b := range data {
			s.Add(int(b))
		}

		filtered := s.Filter(func(x int) bool { return x%2 == 0 })

		for k := range filtered {
			if k%2 != 0 {
				t.Errorf("Filter kept odd number: %d", k)
			}
			if !s.Contains(k) {
				t.Errorf("Filter introduced new element: %d", k)
			}
		}

		for k := range s {
			if k%2 == 0 && !filtered.Contains(k) {
				t.Errorf("Filter removed valid even number: %d", k)
			}
		}
	})
}

func FuzzSetValues(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{1, 1, 1, 2, 2})

	f.Fuzz(func(t *testing.T, data []byte) {
		s := set.Of[int]()
		for _, b := range data {
			s.Add(int(b))
		}

		values := s.Values()

		if values.Len() != s.Len() {
			t.Errorf("Values length mismatch: got %d, want %d", values.Len(), s.Len())
		}

		for _, v := range values {
			if !s.Contains(v) {
				t.Errorf("Values contains element not in set: %d", v)
			}
		}
	})
}

func FuzzSetClear(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		s := set.Of[int]()
		for _, b := range data {
			s.Add(int(b))
		}

		s.Clear()

		if !s.IsEmpty() {
			t.Error("Set is not empty after Clear")
		}

		if s.Len() != 0 {
			t.Errorf("Set length after Clear: got %d, want 0", s.Len())
		}
	})
}
