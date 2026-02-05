package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	assert "github.com/marlonbarreto-git/gollections/internal/testing"
	"github.com/marlonbarreto-git/gollections/list"
)

func TestPipelineTakeUnless(t *testing.T) {
	t.Run("returns value when predicate is false", func(t *testing.T) {
		result := collection.Pipe(42).
			TakeUnless(func(x int) bool { return x < 10 })
		assert.True(t, result.IsPresent())
		assert.Equal(t, 42, result.GetValue())
	})

	t.Run("returns empty when predicate is true", func(t *testing.T) {
		result := collection.Pipe(5).
			TakeUnless(func(x int) bool { return x < 10 })
		assert.True(t, result.IsEmpty())
	})

	t.Run("works with list predicate", func(t *testing.T) {
		l := list.Of(1, 2, 3)
		result := collection.Pipe(l).
			TakeUnless(func(l collection.List[int]) bool { return l.IsEmpty() })
		assert.True(t, result.IsPresent())
	})

	t.Run("takeUnless with true predicate returns empty", func(t *testing.T) {
		l := list.Of(1, 2, 3)
		result := collection.Pipe(l).
			TakeUnless(func(l collection.List[int]) bool { return l.Len() > 0 })
		assert.True(t, result.IsEmpty())
	})
}

func TestPipeMap(t *testing.T) {
	t.Run("transforms pipeline to different type", func(t *testing.T) {
		p := collection.Pipe(42)
		result := collection.PipeMap(p, func(x int) string {
			return "transformed"
		}).Value()
		assert.Equal(t, "transformed", result)
	})

	t.Run("chains with let after type change", func(t *testing.T) {
		p := collection.Pipe(list.Of(1, 2, 3))
		result := collection.PipeMap(p, func(l collection.List[int]) int {
			return l.Len()
		}).
			Let(func(x int) int { return x * 10 }).
			Value()
		assert.Equal(t, 30, result)
	})

	t.Run("transforms list to string", func(t *testing.T) {
		p := collection.Pipe(list.Of("a", "b", "c"))
		result := collection.PipeMap(p, func(l collection.List[string]) string {
			return l.Join(",")
		}).Value()
		assert.Equal(t, "a,b,c", result)
	})
}

func TestPipelineTakeIfFalse(t *testing.T) {
	t.Run("returns empty when predicate is false", func(t *testing.T) {
		result := collection.Pipe(5).
			TakeIf(func(x int) bool { return x > 10 })
		assert.True(t, result.IsEmpty())
	})
}
