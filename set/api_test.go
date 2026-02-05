package set

import (
	"testing"

	assert "github.com/marlonbarreto-git/gollections/internal/testing"
)

func TestOf(t *testing.T) {
	t.Run("creates set from variadic args", func(t *testing.T) {
		s := Of("a", "b", "c")

		assert.Equal(t, 3, s.Len())
		assert.Equal(t, true, s.Contains("a"))
		assert.Equal(t, true, s.Contains("b"))
		assert.Equal(t, true, s.Contains("c"))
	})

	t.Run("creates empty set when no args", func(t *testing.T) {
		s := Of[string]()

		assert.Equal(t, 0, s.Len())
		assert.Equal(t, true, s.IsEmpty())
	})

	t.Run("creates set with single item", func(t *testing.T) {
		s := Of(42)

		assert.Equal(t, 1, s.Len())
		assert.Equal(t, true, s.Contains(42))
	})

	t.Run("deduplicates items", func(t *testing.T) {
		s := Of(1, 2, 2, 3, 3, 3)

		assert.Equal(t, 3, s.Len())
		assert.Equal(t, true, s.Contains(1))
		assert.Equal(t, true, s.Contains(2))
		assert.Equal(t, true, s.Contains(3))
	})
}

func TestOfWithStrings(t *testing.T) {
	s := Of("x", "y", "z")
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, true, s.Contains("x"))
	assert.Equal(t, true, s.Contains("y"))
	assert.Equal(t, true, s.Contains("z"))
	assert.Equal(t, false, s.Contains("w"))
}

func TestFrom(t *testing.T) {
	t.Run("creates set from map keys", func(t *testing.T) {
		raw := map[string]int{"a": 1, "b": 2, "c": 3}
		s := From(raw)

		assert.Equal(t, 3, s.Len())
		assert.Equal(t, true, s.Contains("a"))
		assert.Equal(t, true, s.Contains("b"))
		assert.Equal(t, true, s.Contains("c"))
	})

	t.Run("creates empty set from empty map", func(t *testing.T) {
		raw := map[string]int{}
		s := From(raw)

		assert.Equal(t, 0, s.Len())
		assert.Equal(t, true, s.IsEmpty())
	})

	t.Run("creates set from nil map", func(t *testing.T) {
		var raw map[string]int
		s := From(raw)

		assert.Equal(t, 0, s.Len())
	})
}

func TestFromWithIntKeys(t *testing.T) {
	raw := map[int]string{1: "one", 2: "two", 3: "three"}
	s := From(raw)

	assert.Equal(t, 3, s.Len())
	assert.Equal(t, true, s.Contains(1))
	assert.Equal(t, true, s.Contains(2))
	assert.Equal(t, true, s.Contains(3))
	assert.Equal(t, false, s.Contains(4))
}

func TestOfReturnsWorkingSet(t *testing.T) {
	s := Of(1, 2, 3, 4, 5)

	assert.Equal(t, false, s.IsEmpty())
	assert.Equal(t, 5, s.Len())

	s.Add(6)
	assert.Equal(t, 6, s.Len())
	assert.Equal(t, true, s.Contains(6))

	s.Remove(1)
	assert.Equal(t, 5, s.Len())
	assert.Equal(t, false, s.Contains(1))
}

func TestFromReturnsWorkingSet(t *testing.T) {
	raw := map[string]bool{"x": true, "y": false, "z": true}
	s := From(raw)

	assert.Equal(t, false, s.IsEmpty())
	assert.Equal(t, 3, s.Len())

	values := s.Values()
	assert.Equal(t, 3, values.Len())
}

func TestSetOperationsAfterOf(t *testing.T) {
	s1 := Of(1, 2, 3)
	s2 := Of(2, 3, 4)

	union := s1.Union(s2)
	assert.Equal(t, 4, union.Len())

	intersect := s1.Intersect(s2)
	assert.Equal(t, 2, intersect.Len())
	assert.Equal(t, true, intersect.Contains(2))
	assert.Equal(t, true, intersect.Contains(3))
}
