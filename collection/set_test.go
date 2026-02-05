package collection_test

import (
	"strings"
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	assert "github.com/marlonbarreto-git/gollections/internal/testing"
	"github.com/marlonbarreto-git/gollections/set"
)

func TestSetOf(t *testing.T) {
	t.Run("creates a set from items", func(t *testing.T) {
		s := set.Of("item1", "item2")
		assert.True(t, s.Contains("item1"))
		assert.True(t, s.Contains("item2"))
	})
}

func TestSetFrom(t *testing.T) {
	t.Run("creates a set from raw map", func(t *testing.T) {
		rawMap := make(map[string]int)
		rawMap["item1"] = 1
		rawMap["item2"] = 2
		s := set.From(rawMap)
		assert.True(t, s.Contains("item1"))
		assert.True(t, s.Contains("item2"))
	})
}

func TestContains(t *testing.T) {
	t.Run("checks if item is in set", func(t *testing.T) {
		s := set.Of("item1", "item2")
		assert.True(t, s.Contains("item1"))
		assert.False(t, s.Contains("item3"))
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds item to set", func(t *testing.T) {
		s := set.Of("item1", "item2")
		s.Add("item3")
		assert.True(t, s.Contains("item3"))
	})
}

func TestSetIsEmpty(t *testing.T) {
	t.Run("checks if set is empty", func(t *testing.T) {
		s := set.Of[string]()
		assert.True(t, s.IsEmpty())
		s.Add("item1")
		assert.False(t, s.IsEmpty())
	})
}

func TestSetLen(t *testing.T) {
	t.Run("gets length of set", func(t *testing.T) {
		s := set.Of("item1", "item2")
		assert.Equal(t, 2, s.Len())
	})
}

func TestSetValues(t *testing.T) {
	t.Run("gets values of set", func(t *testing.T) {
		s := set.Of("item1", "item2")
		values := s.Values()
		assert.Contains(t, values, "item1")
		assert.Contains(t, values, "item2")
	})
}

func TestSetString(t *testing.T) {
	t.Run("converts set to string", func(t *testing.T) {
		s := set.Of("item1", "item2")
		str := s.String()
		assert.True(t, strings.Contains(str, "item1"))
		assert.True(t, strings.Contains(str, "item2"))
	})
}

func TestSetRemove(t *testing.T) {
	t.Run("removes item from set", func(t *testing.T) {
		s := set.Of("item1", "item2")
		s.Remove("item1")
		assert.False(t, s.Contains("item1"))
	})
}

func TestClear(t *testing.T) {
	t.Run("clears all items from set", func(t *testing.T) {
		s := set.Of("item1", "item2")
		s.Clear()
		assert.True(t, s.IsEmpty())
	})
}

func TestSetUnion(t *testing.T) {
	t.Run("returns union of two sets", func(t *testing.T) {
		s1 := set.Of(1, 2, 3)
		s2 := set.Of(3, 4, 5)
		result := s1.Union(s2)
		assert.Equal(t, 5, result.Len())
		assert.True(t, result.Contains(1))
		assert.True(t, result.Contains(5))
	})
}

func TestSetIntersect(t *testing.T) {
	t.Run("returns intersection of two sets", func(t *testing.T) {
		s1 := set.Of(1, 2, 3)
		s2 := set.Of(2, 3, 4)
		result := s1.Intersect(s2)
		assert.Equal(t, 2, result.Len())
		assert.True(t, result.Contains(2))
		assert.True(t, result.Contains(3))
		assert.False(t, result.Contains(1))
	})
}

func TestSetSubtract(t *testing.T) {
	t.Run("returns difference of two sets", func(t *testing.T) {
		s1 := set.Of(1, 2, 3)
		s2 := set.Of(2, 3, 4)
		result := s1.Subtract(s2)
		assert.Equal(t, 1, result.Len())
		assert.True(t, result.Contains(1))
		assert.False(t, result.Contains(2))
	})
}

func TestSetFilter(t *testing.T) {
	t.Run("filters set by predicate", func(t *testing.T) {
		s := set.Of(1, 2, 3, 4, 5)
		result := s.Filter(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, 2, result.Len())
		assert.True(t, result.Contains(2))
		assert.True(t, result.Contains(4))
	})
}

func TestSetForEach(t *testing.T) {
	t.Run("iterates over all items", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		sum := 0
		s.ForEach(func(x int) { sum += x })
		assert.Equal(t, 6, sum)
	})
}

func TestSetAny(t *testing.T) {
	t.Run("returns true when any matches", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		assert.True(t, s.Any(func(x int) bool { return x > 2 }))
	})

	t.Run("returns false when none match", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		assert.False(t, s.Any(func(x int) bool { return x > 10 }))
	})
}

func TestSetAll(t *testing.T) {
	t.Run("returns true when all match", func(t *testing.T) {
		s := set.Of(2, 4, 6)
		assert.True(t, s.All(func(x int) bool { return x%2 == 0 }))
	})

	t.Run("returns false when one doesn't match", func(t *testing.T) {
		s := set.Of(1, 2, 4)
		assert.False(t, s.All(func(x int) bool { return x%2 == 0 }))
	})
}

func TestSetNone(t *testing.T) {
	t.Run("returns true when none match", func(t *testing.T) {
		s := set.Of(1, 3, 5)
		assert.True(t, s.None(func(x int) bool { return x%2 == 0 }))
	})
}

func TestSetFirst(t *testing.T) {
	t.Run("returns first element", func(t *testing.T) {
		s := set.Of(1)
		result := s.First()
		assert.True(t, result.IsPresent())
	})

	t.Run("returns empty for empty set", func(t *testing.T) {
		s := set.Of[int]()
		result := s.First()
		assert.False(t, result.IsPresent())
	})
}

func TestSetToList(t *testing.T) {
	t.Run("converts set to list", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		result := s.ToList()
		assert.Equal(t, 3, len(result))
	})
}

func TestSetAlso(t *testing.T) {
	t.Run("performs side effect and returns same set", func(t *testing.T) {
		var count int
		s := set.Of(1, 2, 3)
		result := s.Also(func(s collection.Set[int]) {
			count = s.Len()
		})
		assert.Equal(t, 3, count)
		assert.Equal(t, 3, result.Len())
	})
}

func TestSetTakeIf(t *testing.T) {
	t.Run("returns set when predicate is true", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		result := s.TakeIf(func(s collection.Set[int]) bool {
			return s.Len() > 2
		})
		assert.True(t, result.IsPresent())
	})

	t.Run("returns empty when predicate is false", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		result := s.TakeIf(func(s collection.Set[int]) bool {
			return s.Len() > 10
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestSetTakeUnless(t *testing.T) {
	t.Run("returns set when predicate is false", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		result := s.TakeUnless(func(s collection.Set[int]) bool {
			return s.IsEmpty()
		})
		assert.True(t, result.IsPresent())
	})

	t.Run("returns empty when predicate is true", func(t *testing.T) {
		s := set.Of(1, 2, 3)
		result := s.TakeUnless(func(s collection.Set[int]) bool {
			return s.Len() > 0
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestSetToMap(t *testing.T) {
	t.Run("converts set to map with value selector", func(t *testing.T) {
		s := set.Of("a", "bb", "ccc")
		result := s.ToMap(func(k string) any {
			return len(k)
		})
		assert.Equal(t, 3, result.Len())
	})
}
