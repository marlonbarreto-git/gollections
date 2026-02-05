package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	assert "github.com/marlonbarreto-git/gollections/internal/testing"
	maps "github.com/marlonbarreto-git/gollections/map"
)

func TestMapOf(t *testing.T) {
	t.Run("creates a map from pairs", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", "value1"), collection.PairOf("key2", "value2"))
		assert.Equal(t, "value1", m["key1"])
		assert.Equal(t, "value2", m["key2"])
	})
}

func TestMapFrom(t *testing.T) {
	t.Run("creates a map from raw map", func(t *testing.T) {
		rawMap := make(map[string]string)
		rawMap["key1"] = "value1"
		rawMap["key2"] = "value2"
		m := maps.From(rawMap)
		assert.Equal(t, "value1", m["key1"])
		assert.Equal(t, "value2", m["key2"])
	})
}

func TestMapMap(t *testing.T) {
	t.Run("maps keys and values", func(t *testing.T) {
		original := maps.Of(collection.PairOf("key1", "value1"), collection.PairOf("key2", "value2"))
		mapped := collection.Map(original, func(k, v string) (string, string) {
			return k + "_new", v + "_new"
		})
		assert.Equal(t, "value1_new", mapped["key1_new"])
		assert.Equal(t, "value2_new", mapped["key2_new"])
	})
}

func TestMapKeys(t *testing.T) {
	t.Run("maps keys", func(t *testing.T) {
		original := maps.Of(collection.PairOf("key1", "value1"), collection.PairOf("key2", "value2"))
		mapped := collection.MapKeys(original, func(k, v string) string {
			return k + "_new"
		})
		assert.Equal(t, "value1", mapped["key1_new"])
		assert.Equal(t, "value2", mapped["key2_new"])
	})
}

func TestMapValues(t *testing.T) {
	t.Run("maps values", func(t *testing.T) {
		original := maps.Of(collection.PairOf("key1", "value1"), collection.PairOf("key2", "value2"))
		mapped := collection.MapValues(original, func(k, v string) string {
			return v + "_new"
		})
		assert.Equal(t, "value1_new", mapped["key1"])
		assert.Equal(t, "value2_new", mapped["key2"])
	})
}

func TestMap(t *testing.T) {
	t.Run("maps keys and values", func(t *testing.T) {
		original := maps.Of(collection.PairOf("key1", "value1"), collection.PairOf("key2", "value2"))
		mapped := original.Map(func(k, v string) (any, any) {
			return k + "_new", v + "_new"
		})
		assert.Equal(t, "value1_new", mapped["key1_new"])
		assert.Equal(t, "value2_new", mapped["key2_new"])
	})
}

func TestReduce(t *testing.T) {
	t.Run("reduces map to a single value", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		sum := m.Reduce(func(acc any, key string, value int) any {
			return acc.(int) + value
		}, 0)
		assert.Equal(t, 3, sum)
	})
}

func TestForEach(t *testing.T) {
	t.Run("iterates over each key-value pair", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		sum := 0
		m.ForEach(func(k string, v int) {
			sum += v
		})
		assert.Equal(t, 3, sum)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filters map based on predicate", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		filtered := m.Filter(func(k string, v int) bool {
			return v%2 == 0
		})
		assert.Equal(t, 1, filtered.Len())
		assert.Equal(t, 2, filtered["key2"])
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("checks if map is empty", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		assert.False(t, m.IsEmpty())
		m.Remove("key1")
		m.Remove("key2")
		assert.True(t, m.IsEmpty())
	})
}

func TestLen(t *testing.T) {
	t.Run("gets length of map", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		assert.Equal(t, 2, m.Len())
	})
}

func TestCount(t *testing.T) {
	t.Run("counts elements based on predicate", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		count := m.Count(func(k string, v int) bool {
			return v%2 == 0
		})
		assert.Equal(t, 1, count)
	})
}

func TestCopy(t *testing.T) {
	t.Run("copies map", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		copied := m.Copy()
		assert.Equal(t, m, copied)
		m.Remove("key1")
		assert.NotEqual(t, m, copied)
	})
}

func TestKeys(t *testing.T) {
	t.Run("gets keys of map", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		keys := m.Keys()
		assert.Contains(t, keys, "key1")
		assert.Contains(t, keys, "key2")
	})
}

func TestValues(t *testing.T) {
	t.Run("gets values of map", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		values := m.Values()
		assert.Contains(t, values, 1)
		assert.Contains(t, values, 2)
	})
}

func TestRemove(t *testing.T) {
	t.Run("removes key-value pair from map", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		m.Remove("key1")
		assert.Equal(t, 1, m.Len())
		_, exists := m["key1"]
		assert.False(t, exists)
	})
}

func TestString(t *testing.T) {
	t.Run("converts map to string", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		assert.JSONEq(t, `{"key1":1,"key2":2}`, m.String())
	})

	t.Run("converts map with strange key to string", func(t *testing.T) {
		type customType struct {
			Key string
		}

		m := maps.Of(collection.PairOf(customType{Key: "1"}, 1), collection.PairOf(customType{Key: "2"}, 2))
		assert.JSONEq(t, `{"{1}":1,"{2}":2}`, m.String())
	})
}

func TestGetOrDefault(t *testing.T) {
	t.Run("returns value when key exists", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		result := m.GetOrDefault("key1", 99)
		assert.Equal(t, 1, result)
	})

	t.Run("returns default when key doesn't exist", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		result := m.GetOrDefault("key2", 99)
		assert.Equal(t, 99, result)
	})
}

func TestGetOrPut(t *testing.T) {
	t.Run("returns existing value", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		result := m.GetOrPut("key1", func() int { return 99 })
		assert.Equal(t, 1, result)
	})

	t.Run("puts and returns new value", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		result := m.GetOrPut("key2", func() int { return 99 })
		assert.Equal(t, 99, result)
		assert.Equal(t, 99, m["key2"])
	})
}

func TestContainsKey(t *testing.T) {
	t.Run("returns true when key exists", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		assert.True(t, m.ContainsKey("key1"))
	})

	t.Run("returns false when key doesn't exist", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		assert.False(t, m.ContainsKey("key2"))
	})
}

func TestContainsValue(t *testing.T) {
	t.Run("returns true when value exists", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		assert.True(t, m.ContainsValue(2))
	})

	t.Run("returns false when value doesn't exist", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1))
		assert.False(t, m.ContainsValue(99))
	})
}

func TestMerge(t *testing.T) {
	t.Run("merges two maps", func(t *testing.T) {
		m1 := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		m2 := maps.Of(collection.PairOf("b", 20), collection.PairOf("c", 3))
		m1.Merge(m2)
		assert.Equal(t, 1, m1["a"])
		assert.Equal(t, 20, m1["b"])
		assert.Equal(t, 3, m1["c"])
	})
}

func TestPutAll(t *testing.T) {
	t.Run("puts all key-value pairs", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1))
		m.PutAll(collection.PairOf("b", 2), collection.PairOf("c", 3))
		assert.Equal(t, 1, m["a"])
		assert.Equal(t, 2, m["b"])
		assert.Equal(t, 3, m["c"])
	})
}

func TestFilterKeys(t *testing.T) {
	t.Run("filters by key predicate", func(t *testing.T) {
		m := maps.Of(collection.PairOf("apple", 1), collection.PairOf("banana", 2), collection.PairOf("apricot", 3))
		result := m.FilterKeys(func(k string) bool { return k[0] == 'a' })
		assert.Equal(t, 2, result.Len())
		assert.Equal(t, 1, result["apple"])
		assert.Equal(t, 3, result["apricot"])
	})
}

func TestFilterValues(t *testing.T) {
	t.Run("filters by value predicate", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2), collection.PairOf("c", 3))
		result := m.FilterValues(func(v int) bool { return v > 1 })
		assert.Equal(t, 2, result.Len())
		assert.Equal(t, 2, result["b"])
		assert.Equal(t, 3, result["c"])
	})
}

func TestMapToList(t *testing.T) {
	t.Run("converts map entries to list", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.ToList()
		assert.Equal(t, 2, len(result))
	})
}

func TestMapAny(t *testing.T) {
	t.Run("returns true when any matches", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.Any(func(k string, v int) bool { return v > 1 })
		assert.True(t, result)
	})

	t.Run("returns false when none match", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.Any(func(k string, v int) bool { return v > 10 })
		assert.False(t, result)
	})
}

func TestMapAll(t *testing.T) {
	t.Run("returns true when all match", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 2), collection.PairOf("b", 4))
		result := m.All(func(k string, v int) bool { return v%2 == 0 })
		assert.True(t, result)
	})

	t.Run("returns false when one doesn't match", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.All(func(k string, v int) bool { return v%2 == 0 })
		assert.False(t, result)
	})
}

func TestMapNone(t *testing.T) {
	t.Run("returns true when none match", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 3))
		result := m.None(func(k string, v int) bool { return v%2 == 0 })
		assert.True(t, result)
	})
}

func TestEntries(t *testing.T) {
	t.Run("returns all entries as pairs", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		entries := m.Entries()
		assert.Equal(t, 2, len(entries))
	})

	t.Run("returns empty for empty map", func(t *testing.T) {
		m := maps.Of[string, int]()
		entries := m.Entries()
		assert.Equal(t, 0, len(entries))
	})
}

func TestMapToSet(t *testing.T) {
	t.Run("converts map keys to set", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2), collection.PairOf("c", 3))
		result := m.ToSet()
		assert.Equal(t, 3, result.Len())
		assert.True(t, result.Contains("a"))
		assert.True(t, result.Contains("b"))
		assert.True(t, result.Contains("c"))
	})
}

func TestMapAlso(t *testing.T) {
	t.Run("performs side effect and returns same map", func(t *testing.T) {
		var count int
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.Also(func(m collection.MutableMap[string, int]) {
			count = m.Len()
		})
		assert.Equal(t, 2, count)
		assert.Equal(t, 2, result.Len())
	})
}

func TestMapTakeIf(t *testing.T) {
	t.Run("returns map when predicate is true", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.TakeIf(func(m collection.MutableMap[string, int]) bool {
			return m.Len() > 1
		})
		assert.True(t, result.IsPresent())
	})

	t.Run("returns empty when predicate is false", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1))
		result := m.TakeIf(func(m collection.MutableMap[string, int]) bool {
			return m.Len() > 5
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestMapTakeUnless(t *testing.T) {
	t.Run("returns map when predicate is false", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.TakeUnless(func(m collection.MutableMap[string, int]) bool {
			return m.IsEmpty()
		})
		assert.True(t, result.IsPresent())
	})

	t.Run("returns empty when predicate is true", func(t *testing.T) {
		m := maps.Of(collection.PairOf("a", 1), collection.PairOf("b", 2))
		result := m.TakeUnless(func(m collection.MutableMap[string, int]) bool {
			return m.Len() > 0
		})
		assert.True(t, result.IsEmpty())
	})
}
