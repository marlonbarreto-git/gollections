package list

import (
	"testing"

	assert "github.com/marlonbarreto-git/gollections/internal/testing"
)

func TestOf(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		expected []int
	}{
		{
			name:     "creates list from variadic args",
			items:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "creates empty list when no args",
			items:    []int{},
			expected: []int{},
		},
		{
			name:     "creates list with single item",
			items:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Of(tt.items...)
			assert.Equal(t, len(tt.expected), result.Len())
			for i, v := range tt.expected {
				assert.Equal(t, v, result.Get(i))
			}
		})
	}
}

func TestOfWithStrings(t *testing.T) {
	result := Of("a", "b", "c")
	assert.Equal(t, 3, result.Len())
	assert.Equal(t, "a", result.Get(0))
	assert.Equal(t, "b", result.Get(1))
	assert.Equal(t, "c", result.Get(2))
}

func TestFrom(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		expected []int
	}{
		{
			name:     "creates list from slice",
			items:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "creates empty list from empty slice",
			items:    []int{},
			expected: []int{},
		},
		{
			name:     "creates list from nil slice",
			items:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := From(tt.items)
			if tt.expected == nil {
				assert.Equal(t, 0, result.Len())
			} else {
				assert.Equal(t, len(tt.expected), result.Len())
				for i, v := range tt.expected {
					assert.Equal(t, v, result.Get(i))
				}
			}
		})
	}
}

func TestFromWithStrings(t *testing.T) {
	slice := []string{"x", "y", "z"}
	result := From(slice)
	assert.Equal(t, 3, result.Len())
	assert.Equal(t, "x", result.Get(0))
	assert.Equal(t, "y", result.Get(1))
	assert.Equal(t, "z", result.Get(2))
}

func TestOfReturnsWorkingList(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)

	assert.Equal(t, false, list.IsEmpty())
	assert.Equal(t, 5, list.Len())

	filtered := list.Filter(func(x int) bool { return x > 2 })
	assert.Equal(t, 3, filtered.Len())
	assert.Equal(t, 3, filtered.Get(0))
	assert.Equal(t, 4, filtered.Get(1))
	assert.Equal(t, 5, filtered.Get(2))
}

func TestFromReturnsWorkingList(t *testing.T) {
	slice := []int{10, 20, 30}
	list := From(slice)

	assert.Equal(t, false, list.IsEmpty())
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, "10, 20, 30", list.Join(", "))
}
