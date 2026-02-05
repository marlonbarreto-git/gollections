package pointer

import (
	"testing"

	assert "github.com/marlonbarreto-git/gollections/internal/testing"
)

func TestOf(t *testing.T) {
	t.Run("returns pointer to int", func(t *testing.T) {
		ptr := Of(42)
		assert.NotNil(t, ptr)
		assert.Equal(t, 42, *ptr)
	})

	t.Run("returns pointer to string", func(t *testing.T) {
		ptr := Of("hello")
		assert.NotNil(t, ptr)
		assert.Equal(t, "hello", *ptr)
	})

	t.Run("returns pointer to struct", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		p := person{name: "John", age: 30}
		ptr := Of(p)
		assert.NotNil(t, ptr)
		assert.Equal(t, "John", ptr.name)
		assert.Equal(t, 30, ptr.age)
	})

	t.Run("returns pointer to zero value", func(t *testing.T) {
		ptr := Of(0)
		assert.NotNil(t, ptr)
		assert.Equal(t, 0, *ptr)
	})

	t.Run("returns pointer to empty string", func(t *testing.T) {
		ptr := Of("")
		assert.NotNil(t, ptr)
		assert.Equal(t, "", *ptr)
	})

	t.Run("returns pointer to bool", func(t *testing.T) {
		ptrTrue := Of(true)
		ptrFalse := Of(false)
		assert.Equal(t, true, *ptrTrue)
		assert.Equal(t, false, *ptrFalse)
	})

	t.Run("modifying pointer does not affect original", func(t *testing.T) {
		original := 100
		ptr := Of(original)
		*ptr = 200
		assert.Equal(t, 100, original)
		assert.Equal(t, 200, *ptr)
	})
}
