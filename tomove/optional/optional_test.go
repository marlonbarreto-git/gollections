package optional_test

import (
	"testing"

	assert "github.com/marlonbarreto-git/gollections/internal/testing"
	"github.com/marlonbarreto-git/gollections/tomove/optional"
)

func TestOf(t *testing.T) {
	t.Run("creates optional with non-empty value", func(t *testing.T) {
		opt := optional.Of(10)
		assert.True(t, opt.IsPresent())
		value, err := opt.Get()
		assert.NoError(t, err)
		assert.Equal(t, 10, value)
	})

	t.Run("creates optional with zero value", func(t *testing.T) {
		opt := optional.Of(0)
		assert.True(t, opt.IsPresent())
		value, err := opt.Get()
		assert.NoError(t, err)
		assert.Equal(t, 0, value)
	})

	t.Run("creates optional with nil value", func(t *testing.T) {
		opt := optional.Of[*int](nil)
		assert.False(t, opt.IsPresent())
		_, err := opt.Get()
		assert.Error(t, err)
	})
}

func TestOfValues(t *testing.T) {
	t.Run("creates optional with non-empty values", func(t *testing.T) {
		opt := optional.OfValues[int](10, 20, 30)
		assert.True(t, opt.TakingArg(optional.First).IsPresent())
		value, err := opt.TakingArg(optional.First).Get()
		assert.NoError(t, err)
		assert.Equal(t, 10, value)
	})

	t.Run("creates optional with zero values", func(t *testing.T) {
		opt := optional.OfValues[int](0, 0, 0)
		assert.True(t, opt.TakingArg(optional.Second).IsPresent())
		value, err := opt.TakingArg(optional.Second).Get()
		assert.NoError(t, err)
		assert.Equal(t, 0, value)
	})

	t.Run("creates optional with nil values", func(t *testing.T) {
		opt := optional.OfValues[*int](nil, nil, nil)
		assert.False(t, opt.TakingArg(optional.Third).IsPresent())
		_, err := opt.TakingArg(optional.Third).Get()
		assert.Error(t, err)
	})
}

func TestOfGet(t *testing.T) {
	t.Run("creates optional with non-empty supplier", func(t *testing.T) {
		opt := optional.OfGet(func() int { return 10 })
		assert.True(t, opt.IsPresent())
		value, err := opt.Get()
		assert.NoError(t, err)
		assert.Equal(t, 10, value)
	})

	t.Run("creates optional with zero supplier", func(t *testing.T) {
		opt := optional.OfGet(func() int { return 0 })
		assert.True(t, opt.IsPresent())
		value, err := opt.Get()
		assert.NoError(t, err)
		assert.Equal(t, 0, value)
	})

	t.Run("creates optional with nil supplier", func(t *testing.T) {
		opt := optional.OfGet(func() *int { return nil })
		assert.False(t, opt.IsPresent())
		_, err := opt.Get()
		assert.Error(t, err)
	})
}

func TestTakingArg(t *testing.T) {
	t.Run("gets value from optional with multiple values", func(t *testing.T) {
		opt := optional.OfValues[int](10, 20, 30)
		assert.Equal(t, 20, opt.TakingArg(optional.Second).GetValue())
	})

	t.Run("gets last value from optional with multiple values", func(t *testing.T) {
		opt := optional.OfValues[int](10, 20, 30)
		assert.Equal(t, 30, opt.TakingArg(optional.Last).GetValue())
	})

	t.Run("gets last value from optional when argument number is greater than length", func(t *testing.T) {
		opt := optional.OfValues[int](10, 20, 30)
		assert.Equal(t, 30, opt.TakingArg(optional.ArgumentNum(5)).GetValue())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("checks if struct is empty", func(t *testing.T) {
		type TestStruct struct {
			Field1 int
			Field2 string
		}
		assert.True(t, optional.Of(TestStruct{}).IsEmpty())
	})

	t.Run("checks if struct is not empty", func(t *testing.T) {
		type TestStruct struct {
			Field1 int
			Field2 string
		}
		assert.False(t, optional.Of(TestStruct{Field1: 10, Field2: "test"}).IsEmpty())
	})

	t.Run("checks if map is empty", func(t *testing.T) {
		testMap := make(map[string]int)
		assert.True(t, optional.Of(testMap).IsEmpty())
	})

	t.Run("checks if map is not empty", func(t *testing.T) {
		testMap := map[string]int{"test": 10}
		assert.False(t, optional.Of(testMap).IsEmpty())
	})

	t.Run("checks if slice is empty", func(t *testing.T) {
		testSlice := make([]int, 0)
		assert.True(t, optional.Of(testSlice).IsEmpty())
	})

	t.Run("checks if slice is not empty", func(t *testing.T) {
		testSlice := []int{10}
		assert.False(t, optional.Of(testSlice).IsEmpty())
	})

	t.Run("checks if item is nil", func(t *testing.T) {
		assert.True(t, optional.Of[any](nil).IsEmpty())
	})
}

func TestIsPresent(t *testing.T) {
	t.Run("checks if value is present in non-empty optional", func(t *testing.T) {
		opt := optional.Of(10)
		assert.True(t, opt.IsPresent())
	})

	t.Run("checks if value is present in empty optional", func(t *testing.T) {
		opt := optional.Empty[int]()
		assert.False(t, opt.IsPresent())
	})
}

func TestIfPresent(t *testing.T) {
	t.Run("calls consumer if value is present", func(t *testing.T) {
		opt := optional.Of(10)
		var result int
		opt.IfPresent(func(value int) { result = value })
		assert.Equal(t, 10, result)
	})

	t.Run("does not call consumer if value is not present", func(t *testing.T) {
		opt := optional.Empty[int]()
		var result int
		opt.IfPresent(func(value int) { result = value })
		assert.Equal(t, 0, result)
	})
}

func TestGetValue(t *testing.T) {
	t.Run("gets value from non-empty optional", func(t *testing.T) {
		opt := optional.Of(10)
		assert.Equal(t, 10, opt.GetValue())
	})

	t.Run("gets zero value from empty optional", func(t *testing.T) {
		opt := optional.Empty[int]()
		assert.Equal(t, 0, opt.GetValue())
	})

	t.Run("gets value from optional with supplier", func(t *testing.T) {
		opt := optional.OfGet(func() int { return 20 })
		assert.Equal(t, 20, opt.GetValue())
	})
}

func TestGet(t *testing.T) {
	t.Run("gets value from non-empty optional", func(t *testing.T) {
		opt := optional.Of(10)
		value, err := opt.Get()
		assert.NoError(t, err)
		assert.Equal(t, 10, value)
	})

	t.Run("returns error from empty optional", func(t *testing.T) {
		opt := optional.Empty[int]()
		_, err := opt.Get()
		assert.ErrorIs(t, err, optional.NoValuePresentError)
	})

	t.Run("creates optional with panic supplier", func(t *testing.T) {
		opt := optional.OfGet(func() *int { panic("test error") })
		assert.False(t, opt.IsPresent())
		_, err := opt.Get()
		assert.Error(t, err)
	})
}

func TestOrElse(t *testing.T) {
	t.Run("gets value from non-empty optional", func(t *testing.T) {
		opt := optional.Of(10)
		assert.Equal(t, 10, opt.OrElse(20))
	})

	t.Run("gets alternative from empty optional", func(t *testing.T) {
		opt := optional.Empty[int]()
		assert.Equal(t, 20, opt.OrElse(20))
	})

	t.Run("panics from supplier for OrElsePanic function", func(t *testing.T) {
		opt := optional.OfGet(func() int { panic("test panic") })
		assert.Equal(t, 20, opt.OrElse(20))
	})
}

func TestOrElseGet(t *testing.T) {
	t.Run("gets value from non-empty optional", func(t *testing.T) {
		opt := optional.Of(10)
		assert.Equal(t, 10, opt.OrElseGet(func() int { return 20 }))
	})

	t.Run("gets value from supplier for empty optional", func(t *testing.T) {
		opt := optional.Empty[int]()
		assert.Equal(t, 20, opt.OrElseGet(func() int { return 20 }))
	})

	t.Run("recovers from panic in supplier for OrElseGet function", func(t *testing.T) {
		opt := optional.OfGet(func() int { panic("test panic") })
		assert.Equal(t, 20, opt.OrElseGet(func() int { return 20 }))
	})
}

func TestOrElsePanic(t *testing.T) {
	t.Run("gets value from non-empty optional", func(t *testing.T) {
		opt := optional.Of(10)
		assert.Equal(t, 10, opt.OrElsePanic("no value present"))
	})

	t.Run("panics for empty optional", func(t *testing.T) {
		opt := optional.Empty[int]()
		assert.PanicsWithValue(t, "no value present", func() { opt.OrElsePanic("no value present") })
	})

	t.Run("panics from supplier for OrElsePanic function", func(t *testing.T) {
		opt := optional.OfGet(func() *int { return nil })
		assert.PanicsWithValue(t, "no value present", func() { opt.OrElsePanic("no value present") })
	})
}
