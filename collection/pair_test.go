package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	assert "github.com/marlonbarreto-git/gollections/internal/testing"
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
}

func TestPairSecond(t *testing.T) {
	t.Run("gets second element of pair", func(t *testing.T) {
		p := collection.PairOf("key", "value")
		assert.Equal(t, "value", p.Second())
	})
}
