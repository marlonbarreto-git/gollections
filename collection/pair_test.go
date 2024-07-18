package collection_test

import (
	"github.com/marlonbarreto-git/gollections/collection"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairOf(t *testing.T) {
	t.Run("creates a pair", func(t *testing.T) {
		p := collection.PairOf("key", "value")
		assert.Equal(t, "key", p.First())
		assert.Equal(t, "value", p.Second())
	})
}

func TestPairFirst(t *testing.T) {
	t.Run("gets first element of pair", func(t *testing.T) {
		p := collection.PairOf("key", "value")
		assert.Equal(t, "key", p.First())
	})

	t.Run("panics on invalid type", func(t *testing.T) {
		p := collection.Pair[string, string]{1, "value"}
		assert.Panics(t, func() { p.First() })
	})
}

func TestPairSecond(t *testing.T) {
	t.Run("gets second element of pair", func(t *testing.T) {
		p := collection.PairOf("key", "value")
		assert.Equal(t, "value", p.Second())
	})

	t.Run("panics on invalid type", func(t *testing.T) {
		p := collection.Pair[string, int]{"key", "value"}
		assert.Panics(t, func() { p.Second() })
	})
}
