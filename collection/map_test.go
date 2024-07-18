package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	maps "github.com/marlonbarreto-git/gollections/map"

	"github.com/stretchr/testify/assert"
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

//func TestEntries(t *testing.T) {
//	t.Run("gets entries of map", func(t *testing.T) {
//		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
//		entries := m.Entries()
//		assert.Contains(t, entries, collection.PairOf("key1", 1))
//		assert.Contains(t, entries, collection.PairOf("key2", 2))
//	})
//}

func TestRemove(t *testing.T) {
	t.Run("removes key-value pair from map", func(t *testing.T) {
		m := maps.Of(collection.PairOf("key1", 1), collection.PairOf("key2", 2))
		m.Remove("key1")
		assert.Equal(t, 1, m.Len())
		assert.NotContains(t, m, "key1")
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
