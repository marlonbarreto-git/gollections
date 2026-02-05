package maps

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	assert "github.com/marlonbarreto-git/gollections/internal/testing"
)

func TestOf(t *testing.T) {
	t.Run("creates map from pairs", func(t *testing.T) {
		m := Of(
			collection.PairOf("a", 1),
			collection.PairOf("b", 2),
			collection.PairOf("c", 3),
		)

		assert.Equal(t, 3, m.Len())
		assert.Equal(t, 1, m["a"])
		assert.Equal(t, 2, m["b"])
		assert.Equal(t, 3, m["c"])
	})

	t.Run("creates empty map when no pairs", func(t *testing.T) {
		m := Of[string, int]()
		assert.Equal(t, 0, m.Len())
		assert.Equal(t, true, m.IsEmpty())
	})

	t.Run("creates map with single pair", func(t *testing.T) {
		m := Of(collection.PairOf("key", "value"))
		assert.Equal(t, 1, m.Len())
		assert.Equal(t, "value", m["key"])
	})

	t.Run("last pair wins on duplicate keys", func(t *testing.T) {
		m := Of(
			collection.PairOf("x", 1),
			collection.PairOf("x", 2),
		)
		assert.Equal(t, 1, m.Len())
		assert.Equal(t, 2, m["x"])
	})
}

func TestOfWithIntKeys(t *testing.T) {
	m := Of(
		collection.PairOf(1, "one"),
		collection.PairOf(2, "two"),
	)
	assert.Equal(t, 2, m.Len())
	assert.Equal(t, "one", m[1])
	assert.Equal(t, "two", m[2])
}

func TestFrom(t *testing.T) {
	t.Run("creates mutable map from raw map", func(t *testing.T) {
		raw := map[string]int{"a": 1, "b": 2}
		m := From(raw)

		assert.Equal(t, 2, m.Len())
		assert.Equal(t, 1, m["a"])
		assert.Equal(t, 2, m["b"])
	})

	t.Run("creates empty mutable map from empty raw map", func(t *testing.T) {
		raw := map[string]int{}
		m := From(raw)

		assert.Equal(t, 0, m.Len())
		assert.Equal(t, true, m.IsEmpty())
	})

	t.Run("creates mutable map from nil", func(t *testing.T) {
		var raw map[string]int
		m := From(raw)

		assert.Equal(t, 0, m.Len())
	})
}

func TestOfReturnsWorkingMap(t *testing.T) {
	m := Of(
		collection.PairOf("one", 1),
		collection.PairOf("two", 2),
		collection.PairOf("three", 3),
	)

	assert.Equal(t, false, m.IsEmpty())
	assert.Equal(t, 3, m.Len())

	filtered := m.Filter(func(k string, v int) bool { return v > 1 })
	assert.Equal(t, 2, filtered.Len())
	assert.Equal(t, true, filtered.ContainsKey("two"))
	assert.Equal(t, true, filtered.ContainsKey("three"))
	assert.Equal(t, false, filtered.ContainsKey("one"))
}

func TestFromReturnsWorkingMap(t *testing.T) {
	raw := map[int]string{1: "a", 2: "b", 3: "c"}
	m := From(raw)

	assert.Equal(t, false, m.IsEmpty())
	assert.Equal(t, 3, m.Len())

	keys := m.Keys()
	assert.Equal(t, 3, keys.Len())

	values := m.Values()
	assert.Equal(t, 3, values.Len())
}
