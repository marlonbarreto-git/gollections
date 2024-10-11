package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/set"

	"github.com/stretchr/testify/assert"
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
		assert.Contains(t, s.String(), "item1")
		assert.Contains(t, s.String(), "item2")
		assert.Equal(t, "{item1, item2}", s.String())
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
