package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/list"
	"github.com/marlonbarreto-git/gollections/tomove/function"

	"github.com/stretchr/testify/assert"
)

func TestListOf(t *testing.T) {
	t.Run("creates new list with single element", func(t *testing.T) {
		result := list.Of[int](1)
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("creates new list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("creates new list with no elements", func(t *testing.T) {
		result := list.Of[int]()
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestListFrom(t *testing.T) {
	t.Run("creates new list from single element slice", func(t *testing.T) {
		result := list.From[int]([]int{1})
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("creates new list from multiple elements slice", func(t *testing.T) {
		result := list.From[int]([]int{1, 2, 3})
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("creates new list from empty slice", func(t *testing.T) {
		result := list.From[int]([]int{})
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestListMap(t *testing.T) {
	t.Run("transforms list with single element", func(t *testing.T) {
		result := collection.ListMap[int, int](list.Of[int](1), func(i int) int { return i * 2 })
		assert.Equal(t, collection.List[int]{2}, result)
	})

	t.Run("transforms list with multiple elements", func(t *testing.T) {
		result := collection.ListMap[int, int](list.Of[int](1, 2, 3), func(i int) int { return i * 2 })
		assert.Equal(t, collection.List[int]{2, 4, 6}, result)
	})

	t.Run("transforms empty list", func(t *testing.T) {
		result := collection.ListMap[int, int](list.Of[int](), func(i int) int { return i * 2 })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestToArray(t *testing.T) {
	t.Run("converts list with single element to array", func(t *testing.T) {
		result := list.Of[int](1).ToArray()
		assert.Equal(t, []int{1}, result)
	})

	t.Run("converts list with multiple elements to array", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).ToArray()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("converts empty list to array", func(t *testing.T) {
		result := list.Of[int]().ToArray()
		assert.Equal(t, []int(nil), result)
	})
}

func TestDistinct(t *testing.T) {
	t.Run("removes duplicates from list with single element", func(t *testing.T) {
		result := list.Of[int](1, 1).Distinct()
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("removes duplicates from list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 2, 3, 3, 3).Distinct()
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("returns same list when no duplicates", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Distinct()
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("returns empty list when called on empty list", func(t *testing.T) {
		result := list.Of[int]().Distinct()
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestAssociate(t *testing.T) {
	t.Run("associates list with single element", func(t *testing.T) {
		result := list.Of[int](1).Associate(func(i int) any { return i * 2 })
		assert.Equal(t, map[any]int{2: 1}, map[any]int(result))
	})

	t.Run("associates list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Associate(func(i int) any { return i * 2 })
		assert.Equal(t, map[any]int{2: 1, 4: 2, 6: 3}, map[any]int(result))
	})

	t.Run("associates empty list", func(t *testing.T) {
		result := list.Of[int]().Associate(func(i int) any { return i * 2 })
		assert.Equal(t, map[any]int{}, map[any]int(result))
	})
}

func TestAssociateBy(t *testing.T) {
	type TestStruct struct {
		Key   int
		Value string
	}

	t.Run("associates list with single element by key", func(t *testing.T) {
		result := list.Of[TestStruct](TestStruct{Key: 1, Value: "one"}).AssociateBy("Key")
		assert.Equal(t, map[any]TestStruct{1: {Key: 1, Value: "one"}}, map[any]TestStruct(result))
	})

	t.Run("associates list with multiple elements by key", func(t *testing.T) {
		result := list.Of[TestStruct](TestStruct{Key: 1, Value: "one"}, TestStruct{Key: 2, Value: "two"}, TestStruct{Key: 3, Value: "three"}).AssociateBy("Key")
		assert.Equal(t, map[any]TestStruct{1: {Key: 1, Value: "one"}, 2: {Key: 2, Value: "two"}, 3: {Key: 3, Value: "three"}}, map[any]TestStruct(result))
	})

	t.Run("associates empty list", func(t *testing.T) {
		result := list.Of[TestStruct]().AssociateBy("Key")
		assert.Equal(t, map[any]TestStruct{}, map[any]TestStruct(result))
	})

	t.Run("associates list with struct elements by valid key", func(t *testing.T) {
		result := list.Of[TestStruct](TestStruct{Key: 1, Value: "one"}, TestStruct{Key: 2, Value: "two"}).AssociateBy("Key")
		assert.Equal(t, map[any]TestStruct{1: {Key: 1, Value: "one"}, 2: {Key: 2, Value: "two"}}, map[any]TestStruct(result))
	})

	t.Run("associates list with pointer to struct elements by valid key", func(t *testing.T) {
		result := list.Of[*TestStruct](&TestStruct{Key: 1, Value: "one"}, &TestStruct{Key: 2, Value: "two"}).AssociateBy("Key")
		assert.Equal(t, map[any]*TestStruct{1: {Key: 1, Value: "one"}, 2: {Key: 2, Value: "two"}}, map[any]*TestStruct(result))
	})

	t.Run("panics when list elements are not structs", func(t *testing.T) {
		assert.Panics(t, func() { list.Of[int](1, 2, 3).AssociateBy("Key") })
	})

	t.Run("panics when keySelector is not a valid field", func(t *testing.T) {
		assert.Panics(t, func() { list.Of[TestStruct](TestStruct{Key: 1, Value: "one"}).AssociateBy("InvalidKey") })
	})
}

func TestJoin(t *testing.T) {
	t.Run("joins list with single element", func(t *testing.T) {
		result := list.Of[int](1).Join(", ")
		assert.Equal(t, "1", result)
	})

	t.Run("joins list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Join(", ")
		assert.Equal(t, "1, 2, 3", result)
	})

	t.Run("joins empty list", func(t *testing.T) {
		result := list.Of[int]().Join(", ")
		assert.Equal(t, "", result)
	})

	t.Run("joins list with nil elements", func(t *testing.T) {
		result := list.Of[any](nil, 1, nil).Join(", ")
		assert.Equal(t, "<nil>, 1, <nil>", result)
	})
}

func TestListFilter(t *testing.T) {
	t.Run("filters list with single element", func(t *testing.T) {
		result := list.Of[int](1).Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, collection.List[int](nil), result)
	})

	t.Run("filters list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, collection.List[int]{2}, result)
	})

	t.Run("filters empty list", func(t *testing.T) {
		result := list.Of[int]().Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestAppend(t *testing.T) {
	t.Run("appends to list with single element", func(t *testing.T) {
		result := list.Of[int](1).Append(2)
		assert.Equal(t, collection.List[int]{1, 2}, result)
	})

	t.Run("appends to list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Append(4)
		assert.Equal(t, collection.List[int]{1, 2, 3, 4}, result)
	})

	t.Run("appends to empty list", func(t *testing.T) {
		result := list.Of[int]().Append(1)
		assert.Equal(t, collection.List[int]{1}, result)
	})
}

func TestListAdd(t *testing.T) {
	t.Run("adds to list with single element", func(t *testing.T) {
		myList := list.Of[int](1)
		myList.Add(2)
		assert.Equal(t, collection.List[int]{1, 2}, myList)
	})

	t.Run("adds to list with multiple elements", func(t *testing.T) {
		myList := list.Of[int](1, 2, 3)
		myList.Add(4)
		assert.Equal(t, collection.List[int]{1, 2, 3, 4}, myList)
	})

	t.Run("adds to empty list", func(t *testing.T) {
		myList := list.Of[int]()
		myList.Add(1)
		assert.Equal(t, collection.List[int]{1}, myList)
	})
}

func TestListForEach(t *testing.T) {
	t.Run("iterates over list with single element", func(t *testing.T) {
		sum := 0
		list.Of[int](1).ForEach(func(i int) { sum += i })
		assert.Equal(t, 1, sum)
	})

	t.Run("iterates over list with multiple elements", func(t *testing.T) {
		sum := 0
		list.Of[int](1, 2, 3).ForEach(func(i int) { sum += i })
		assert.Equal(t, 6, sum)
	})

	t.Run("iterates over empty list", func(t *testing.T) {
		sum := 0
		list.Of[int]().ForEach(func(i int) { sum += i })
		assert.Equal(t, 0, sum)
	})
}

func TestForEachIndexed(t *testing.T) {
	t.Run("iterates over list with single element with index", func(t *testing.T) {
		sum := 0
		list.Of[int](1).ForEachIndexed(func(index int, i int) { sum += i + index })
		assert.Equal(t, 1, sum)
	})

	t.Run("iterates over list with multiple elements with index", func(t *testing.T) {
		sum := 0
		list.Of[int](1, 2, 3).ForEachIndexed(func(index int, i int) { sum += i + index })
		assert.Equal(t, 9, sum)
	})

	t.Run("iterates over empty list with index", func(t *testing.T) {
		sum := 0
		list.Of[int]().ForEachIndexed(func(index int, i int) { sum += i + index })
		assert.Equal(t, 0, sum)
	})
}

func TestListFirst(t *testing.T) {
	t.Run("gets first element from list with single element", func(t *testing.T) {
		result := list.Of[int](1).First()
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("gets first element from list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).First()
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("gets first element from empty list", func(t *testing.T) {
		result := list.Of[int]().First()
		assert.False(t, result.IsPresent())
	})
}

func TestGet(t *testing.T) {
	t.Run("gets element by index from list with single element", func(t *testing.T) {
		result := list.Of[int](1).Get(0)
		assert.Equal(t, 1, result)
	})

	t.Run("gets element by index from list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Get(1)
		assert.Equal(t, 2, result)
	})
}

func TestFind(t *testing.T) {
	t.Run("finds element in list with single element", func(t *testing.T) {
		result := list.Of[int](1).Find(func(i int) bool { return i == 1 })
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("finds element in list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Find(func(i int) bool { return i == 2 })
		assert.Equal(t, 2, result.GetValue())
	})

	t.Run("finds element in empty list", func(t *testing.T) {
		result := list.Of[int]().Find(func(i int) bool { return i == 1 })
		assert.False(t, result.IsPresent())
	})
}

func TestLast(t *testing.T) {
	t.Run("gets last element from list with single element", func(t *testing.T) {
		result := list.Of[int](1).Last()
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("gets last element from list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Last()
		assert.Equal(t, 3, result.GetValue())
	})

	t.Run("gets last element from empty list", func(t *testing.T) {
		result := list.Of[int]().Last()
		assert.False(t, result.IsPresent())
	})
}

func TestListLen(t *testing.T) {
	t.Run("gets length of list with single element", func(t *testing.T) {
		result := list.Of[int](1).Len()
		assert.Equal(t, 1, result)
	})

	t.Run("gets length of list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Len()
		assert.Equal(t, 3, result)
	})

	t.Run("gets length of empty list", func(t *testing.T) {
		result := list.Of[int]().Len()
		assert.Equal(t, 0, result)
	})
}

func TestListIsEmpty(t *testing.T) {
	t.Run("checks if list with single element is empty", func(t *testing.T) {
		result := list.Of[int](1).IsEmpty()
		assert.False(t, result)
	})

	t.Run("checks if list with multiple elements is empty", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).IsEmpty()
		assert.False(t, result)
	})

	t.Run("checks if empty list is empty", func(t *testing.T) {
		result := list.Of[int]().IsEmpty()
		assert.True(t, result)
	})
}

func TestSome(t *testing.T) {
	t.Run("checks if some elements in list with single element satisfy the predicate", func(t *testing.T) {
		result := list.Of[int](1).Some(function.Predicate[int](func(i int) bool { return i%2 == 0 }))
		assert.False(t, result)
	})

	t.Run("checks if some elements in list with multiple elements satisfy the predicate", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Some(function.Predicate[int](func(i int) bool { return i%2 == 0 }))
		assert.True(t, result)
	})

	t.Run("checks if some elements in empty list satisfy the predicate", func(t *testing.T) {
		result := list.Of[int]().Some(function.Predicate[int](func(i int) bool { return i%2 == 0 }))
		assert.False(t, result)
	})
}

func TestEvery(t *testing.T) {
	t.Run("checks if every element in list with single element satisfy the predicate", func(t *testing.T) {
		result := list.Of[int](2).Every(func(i int) bool { return i%2 == 0 })
		assert.True(t, result)
	})

	t.Run("checks if any element in list with single element not satisfy the predicate", func(t *testing.T) {
		result := list.Of[int](2, 3).Every(func(i int) bool { return i%2 == 0 })
		assert.False(t, result)
	})

	t.Run("checks if every element in list with multiple elements satisfy the predicate", func(t *testing.T) {
		result := list.Of[int](2, 4, 6).Every(func(i int) bool { return i%2 == 0 })
		assert.True(t, result)
	})

	t.Run("checks if every element in empty list satisfy the predicate", func(t *testing.T) {
		result := list.Of[int]().Every(func(i int) bool { return i%2 == 0 })
		assert.True(t, result)
	})
}

func TestSlice(t *testing.T) {
	t.Run("slices list with single element", func(t *testing.T) {
		result := list.Of[int](1).Slice(0, 1)
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("slices list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Slice(1, 3)
		assert.Equal(t, collection.List[int]{2, 3}, result)
	})

	t.Run("slices empty list", func(t *testing.T) {
		result := list.Of[int]().Slice(0, 0)
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestListCount(t *testing.T) {
	t.Run("counts elements in list with single element satisfying the predicate", func(t *testing.T) {
		result := list.Of[int](1).Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 0, result)
	})

	t.Run("counts elements in list with multiple elements satisfying the predicate", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 1, result)
	})

	t.Run("counts elements in empty list satisfying the predicate", func(t *testing.T) {
		result := list.Of[int]().Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 0, result)
	})
}

func TestSum(t *testing.T) {
	t.Run("sums elements in list with single element", func(t *testing.T) {
		result := list.Of[int](1).Sum(func(i int) float64 { return float64(i) })
		assert.Equal(t, 1.0, result)
	})

	t.Run("sums elements in list with multiple elements", func(t *testing.T) {
		result := list.Of[int](1, 2, 3).Sum(func(i int) float64 { return float64(i) })
		assert.Equal(t, 6.0, result)
	})

	t.Run("sums elements in empty list", func(t *testing.T) {
		result := list.Of[int]().Sum(func(i int) float64 { return float64(i) })
		assert.Equal(t, 0.0, result)
	})
}
