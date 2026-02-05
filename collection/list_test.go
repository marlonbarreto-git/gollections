package collection_test

import (
	"fmt"
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	assert "github.com/marlonbarreto-git/gollections/internal/testing"
	"github.com/marlonbarreto-git/gollections/list"
	"github.com/marlonbarreto-git/gollections/tomove/function"
)

func TestListOf(t *testing.T) {
	t.Run("creates new list with single element", func(t *testing.T) {
		result := list.Of(1)
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("creates new list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("creates new list with no elements", func(t *testing.T) {
		result := list.Of[int]()
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestListFrom(t *testing.T) {
	t.Run("creates new list from single element slice", func(t *testing.T) {
		result := list.From([]int{1})
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("creates new list from multiple elements slice", func(t *testing.T) {
		result := list.From([]int{1, 2, 3})
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("creates new list from empty slice", func(t *testing.T) {
		result := list.From([]int{})
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestListMap(t *testing.T) {
	t.Run("transforms list with single element", func(t *testing.T) {
		result := collection.ListMap(list.Of(1), func(i int) int { return i * 2 })
		assert.Equal(t, collection.List[int]{2}, result)
	})

	t.Run("transforms list with multiple elements", func(t *testing.T) {
		result := collection.ListMap(list.Of(1, 2, 3), func(i int) int { return i * 2 })
		assert.Equal(t, collection.List[int]{2, 4, 6}, result)
	})

	t.Run("transforms empty list", func(t *testing.T) {
		result := collection.ListMap(list.Of[int](), func(i int) int { return i * 2 })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestToArray(t *testing.T) {
	t.Run("converts list with single element to array", func(t *testing.T) {
		result := list.Of(1).ToArray()
		assert.Equal(t, []int{1}, result)
	})

	t.Run("converts list with multiple elements to array", func(t *testing.T) {
		result := list.Of(1, 2, 3).ToArray()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("converts empty list to array", func(t *testing.T) {
		result := list.Of[int]().ToArray()
		assert.Equal(t, []int(nil), result)
	})
}

func TestDistinct(t *testing.T) {
	t.Run("removes duplicates from list with single element", func(t *testing.T) {
		result := list.Of(1, 1).Distinct()
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("removes duplicates from list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 2, 3, 3, 3).Distinct()
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("returns same list when no duplicates", func(t *testing.T) {
		result := list.Of(1, 2, 3).Distinct()
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("returns empty list when called on empty list", func(t *testing.T) {
		result := list.Of[int]().Distinct()
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestAssociate(t *testing.T) {
	t.Run("associates list with single element", func(t *testing.T) {
		result := list.Of(1).Associate(func(i int) any { return i * 2 })
		assert.MapEqual(t, map[any]int{2: 1}, map[any]int(result))
	})

	t.Run("associates list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Associate(func(i int) any { return i * 2 })
		assert.MapEqual(t, map[any]int{2: 1, 4: 2, 6: 3}, map[any]int(result))
	})

	t.Run("associates empty list", func(t *testing.T) {
		result := list.Of[int]().Associate(func(i int) any { return i * 2 })
		assert.MapEqual(t, map[any]int{}, map[any]int(result))
	})
}

func TestAssociateBy(t *testing.T) {
	type TestStruct struct {
		Key   int
		Value string
	}

	t.Run("associates list with single element by key", func(t *testing.T) {
		result := list.Of(TestStruct{Key: 1, Value: "one"}).AssociateBy("Key")
		assert.MapEqual(t, map[any]TestStruct{1: {Key: 1, Value: "one"}}, map[any]TestStruct(result))
	})

	t.Run("associates list with multiple elements by key", func(t *testing.T) {
		result := list.Of(TestStruct{Key: 1, Value: "one"}, TestStruct{Key: 2, Value: "two"}, TestStruct{Key: 3, Value: "three"}).AssociateBy("Key")
		assert.MapEqual(t, map[any]TestStruct{1: {Key: 1, Value: "one"}, 2: {Key: 2, Value: "two"}, 3: {Key: 3, Value: "three"}}, map[any]TestStruct(result))
	})

	t.Run("associates empty list", func(t *testing.T) {
		result := list.Of[TestStruct]().AssociateBy("Key")
		assert.MapEqual(t, map[any]TestStruct{}, map[any]TestStruct(result))
	})

	t.Run("associates list with struct elements by valid key", func(t *testing.T) {
		result := list.Of(TestStruct{Key: 1, Value: "one"}, TestStruct{Key: 2, Value: "two"}).AssociateBy("Key")
		assert.MapEqual(t, map[any]TestStruct{1: {Key: 1, Value: "one"}, 2: {Key: 2, Value: "two"}}, map[any]TestStruct(result))
	})

	t.Run("associates list with pointer to struct elements by valid key", func(t *testing.T) {
		result := list.Of(&TestStruct{Key: 1, Value: "one"}, &TestStruct{Key: 2, Value: "two"}).AssociateBy("Key")
		assert.MapEqual(t, map[any]*TestStruct{1: {Key: 1, Value: "one"}, 2: {Key: 2, Value: "two"}}, map[any]*TestStruct(result))
	})

	t.Run("panics when list elements are not structs", func(t *testing.T) {
		assert.Panics(t, func() { list.Of(1, 2, 3).AssociateBy("Key") })
	})

	t.Run("panics when keySelector is not a valid field", func(t *testing.T) {
		assert.Panics(t, func() { list.Of(TestStruct{Key: 1, Value: "one"}).AssociateBy("InvalidKey") })
	})
}

func TestJoin(t *testing.T) {
	t.Run("joins list with single element", func(t *testing.T) {
		result := list.Of(1).Join(", ")
		assert.Equal(t, "1", result)
	})

	t.Run("joins list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Join(", ")
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
		result := list.Of(1).Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, collection.List[int](nil), result)
	})

	t.Run("filters list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, collection.List[int]{2}, result)
	})

	t.Run("filters empty list", func(t *testing.T) {
		result := list.Of[int]().Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestAppend(t *testing.T) {
	t.Run("appends to list with single element", func(t *testing.T) {
		result := list.Of(1).Append(2)
		assert.Equal(t, collection.List[int]{1, 2}, result)
	})

	t.Run("appends to list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Append(4)
		assert.Equal(t, collection.List[int]{1, 2, 3, 4}, result)
	})

	t.Run("appends to empty list", func(t *testing.T) {
		result := list.Of[int]().Append(1)
		assert.Equal(t, collection.List[int]{1}, result)
	})
}

func TestListAdd(t *testing.T) {
	t.Run("adds to list with single element", func(t *testing.T) {
		myList := list.Of(1)
		myList.Add(2)
		assert.Equal(t, collection.List[int]{1, 2}, myList)
	})

	t.Run("adds to list with multiple elements", func(t *testing.T) {
		myList := list.Of(1, 2, 3)
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
		list.Of(1).ForEach(func(i int) { sum += i })
		assert.Equal(t, 1, sum)
	})

	t.Run("iterates over list with multiple elements", func(t *testing.T) {
		sum := 0
		list.Of(1, 2, 3).ForEach(func(i int) { sum += i })
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
		list.Of(1).ForEachIndexed(func(index int, i int) { sum += i + index })
		assert.Equal(t, 1, sum)
	})

	t.Run("iterates over list with multiple elements with index", func(t *testing.T) {
		sum := 0
		list.Of(1, 2, 3).ForEachIndexed(func(index int, i int) { sum += i + index })
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
		result := list.Of(1).First()
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("gets first element from list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).First()
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("gets first element from empty list", func(t *testing.T) {
		result := list.Of[int]().First()
		assert.False(t, result.IsPresent())
	})
}

func TestGet(t *testing.T) {
	t.Run("gets element by index from list with single element", func(t *testing.T) {
		result := list.Of(1).Get(0)
		assert.Equal(t, 1, result)
	})

	t.Run("gets element by index from list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Get(1)
		assert.Equal(t, 2, result)
	})
}

func TestFind(t *testing.T) {
	t.Run("finds element in list with single element", func(t *testing.T) {
		result := list.Of(1).Find(func(i int) bool { return i == 1 })
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("finds element in list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Find(func(i int) bool { return i == 2 })
		assert.Equal(t, 2, result.GetValue())
	})

	t.Run("finds element in empty list", func(t *testing.T) {
		result := list.Of[int]().Find(func(i int) bool { return i == 1 })
		assert.False(t, result.IsPresent())
	})
}

func TestLast(t *testing.T) {
	t.Run("gets last element from list with single element", func(t *testing.T) {
		result := list.Of(1).Last()
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("gets last element from list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Last()
		assert.Equal(t, 3, result.GetValue())
	})

	t.Run("gets last element from empty list", func(t *testing.T) {
		result := list.Of[int]().Last()
		assert.False(t, result.IsPresent())
	})
}

func TestListLen(t *testing.T) {
	t.Run("gets length of list with single element", func(t *testing.T) {
		result := list.Of(1).Len()
		assert.Equal(t, 1, result)
	})

	t.Run("gets length of list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Len()
		assert.Equal(t, 3, result)
	})

	t.Run("gets length of empty list", func(t *testing.T) {
		result := list.Of[int]().Len()
		assert.Equal(t, 0, result)
	})
}

func TestListIsEmpty(t *testing.T) {
	t.Run("checks if list with single element is empty", func(t *testing.T) {
		result := list.Of(1).IsEmpty()
		assert.False(t, result)
	})

	t.Run("checks if list with multiple elements is empty", func(t *testing.T) {
		result := list.Of(1, 2, 3).IsEmpty()
		assert.False(t, result)
	})

	t.Run("checks if empty list is empty", func(t *testing.T) {
		result := list.Of[int]().IsEmpty()
		assert.True(t, result)
	})
}

func TestSome(t *testing.T) {
	t.Run("checks if some elements in list with single element satisfy the predicate", func(t *testing.T) {
		result := list.Of(1).Some(function.Predicate[int](func(i int) bool { return i%2 == 0 }))
		assert.False(t, result)
	})

	t.Run("checks if some elements in list with multiple elements satisfy the predicate", func(t *testing.T) {
		result := list.Of(1, 2, 3).Some(function.Predicate[int](func(i int) bool { return i%2 == 0 }))
		assert.True(t, result)
	})

	t.Run("checks if some elements in empty list satisfy the predicate", func(t *testing.T) {
		result := list.Of[int]().Some(function.Predicate[int](func(i int) bool { return i%2 == 0 }))
		assert.False(t, result)
	})
}

func TestEvery(t *testing.T) {
	t.Run("checks if every element in list with single element satisfy the predicate", func(t *testing.T) {
		result := list.Of(2).Every(func(i int) bool { return i%2 == 0 })
		assert.True(t, result)
	})

	t.Run("checks if any element in list with single element not satisfy the predicate", func(t *testing.T) {
		result := list.Of(2, 3).Every(func(i int) bool { return i%2 == 0 })
		assert.False(t, result)
	})

	t.Run("checks if every element in list with multiple elements satisfy the predicate", func(t *testing.T) {
		result := list.Of(2, 4, 6).Every(func(i int) bool { return i%2 == 0 })
		assert.True(t, result)
	})

	t.Run("checks if every element in empty list satisfy the predicate", func(t *testing.T) {
		result := list.Of[int]().Every(func(i int) bool { return i%2 == 0 })
		assert.True(t, result)
	})
}

func TestSlice(t *testing.T) {
	t.Run("slices list with single element", func(t *testing.T) {
		result := list.Of(1).Slice(0, 1)
		assert.Equal(t, collection.List[int]{1}, result)
	})

	t.Run("slices list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Slice(1, 3)
		assert.Equal(t, collection.List[int]{2, 3}, result)
	})

	t.Run("slices empty list", func(t *testing.T) {
		result := list.Of[int]().Slice(0, 0)
		assert.Equal(t, collection.List[int](nil), result)
	})
}

func TestListCount(t *testing.T) {
	t.Run("counts elements in list with single element satisfying the predicate", func(t *testing.T) {
		result := list.Of(1).Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 0, result)
	})

	t.Run("counts elements in list with multiple elements satisfying the predicate", func(t *testing.T) {
		result := list.Of(1, 2, 3).Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 1, result)
	})

	t.Run("counts elements in empty list satisfying the predicate", func(t *testing.T) {
		result := list.Of[int]().Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 0, result)
	})
}

func TestSum(t *testing.T) {
	t.Run("sums elements in list with single element", func(t *testing.T) {
		result := list.Of(1).Sum(func(i int) float64 { return float64(i) })
		assert.Equal(t, 1.0, result)
	})

	t.Run("sums elements in list with multiple elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Sum(func(i int) float64 { return float64(i) })
		assert.Equal(t, 6.0, result)
	})

	t.Run("sums elements in empty list", func(t *testing.T) {
		result := list.Of[int]().Sum(func(i int) float64 { return float64(i) })
		assert.Equal(t, 0.0, result)
	})
}

func TestListReduce(t *testing.T) {
	t.Run("reduces list to single value", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4).Reduce(0, func(acc, x int) int { return acc + x })
		assert.Equal(t, 10, result)
	})

	t.Run("reduces empty list returns initial", func(t *testing.T) {
		result := list.Of[int]().Reduce(42, func(acc, x int) int { return acc + x })
		assert.Equal(t, 42, result)
	})

	t.Run("reduces strings", func(t *testing.T) {
		result := list.Of("a", "b", "c").Reduce("", func(acc, x string) string { return acc + x })
		assert.Equal(t, "abc", result)
	})
}

func TestFold(t *testing.T) {
	t.Run("folds to different type", func(t *testing.T) {
		result := collection.Fold(list.Of(1, 2, 3), "", func(acc string, x int) string {
			if acc == "" {
				return string(rune('0' + x))
			}
			return acc + "," + string(rune('0'+x))
		})
		assert.Equal(t, "1,2,3", result)
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("flat maps list", func(t *testing.T) {
		result := list.Of(1, 2, 3).FlatMap(func(x int) []int { return []int{x, x * 10} })
		assert.Equal(t, collection.List[int]{1, 10, 2, 20, 3, 30}, result)
	})

	t.Run("flat maps to empty", func(t *testing.T) {
		result := list.Of(1, 2, 3).FlatMap(func(x int) []int { return []int{} })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestReversed(t *testing.T) {
	t.Run("reverses list", func(t *testing.T) {
		result := list.Of(1, 2, 3).Reversed()
		assert.Equal(t, collection.List[int]{3, 2, 1}, result)
	})

	t.Run("reverses empty list", func(t *testing.T) {
		result := list.Of[int]().Reversed()
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestSorted(t *testing.T) {
	t.Run("sorts list ascending", func(t *testing.T) {
		result := list.Of(3, 1, 4, 1, 5).Sorted(func(a, b int) int { return a - b })
		assert.Equal(t, collection.List[int]{1, 1, 3, 4, 5}, result)
	})

	t.Run("sorts list descending", func(t *testing.T) {
		result := list.Of(3, 1, 4, 1, 5).Sorted(func(a, b int) int { return b - a })
		assert.Equal(t, collection.List[int]{5, 4, 3, 1, 1}, result)
	})
}

func TestTake(t *testing.T) {
	t.Run("takes first n elements", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).Take(3)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("takes more than available", func(t *testing.T) {
		result := list.Of(1, 2).Take(5)
		assert.Equal(t, collection.List[int]{1, 2}, result)
	})

	t.Run("takes zero", func(t *testing.T) {
		result := list.Of(1, 2, 3).Take(0)
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestTakeLast(t *testing.T) {
	t.Run("takes last n elements", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).TakeLast(3)
		assert.Equal(t, collection.List[int]{3, 4, 5}, result)
	})

	t.Run("takes more than available", func(t *testing.T) {
		result := list.Of(1, 2).TakeLast(5)
		assert.Equal(t, collection.List[int]{1, 2}, result)
	})
}

func TestTakeWhile(t *testing.T) {
	t.Run("takes while predicate true", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 4 })
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestDrop(t *testing.T) {
	t.Run("drops first n elements", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).Drop(2)
		assert.Equal(t, collection.List[int]{3, 4, 5}, result)
	})

	t.Run("drops more than available", func(t *testing.T) {
		result := list.Of(1, 2).Drop(5)
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestDropLast(t *testing.T) {
	t.Run("drops last n elements", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).DropLast(2)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestDropWhile(t *testing.T) {
	t.Run("drops while predicate true", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x < 3 })
		assert.Equal(t, collection.List[int]{3, 4, 5}, result)
	})
}

func TestChunked(t *testing.T) {
	t.Run("chunks list into groups", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).Chunked(2)
		assert.Equal(t, []collection.List[int]{{1, 2}, {3, 4}, {5}}, result)
	})
}

func TestListContains(t *testing.T) {
	t.Run("returns true when contains", func(t *testing.T) {
		result := list.Of(1, 2, 3).Contains(2)
		assert.True(t, result)
	})

	t.Run("returns false when not contains", func(t *testing.T) {
		result := list.Of(1, 2, 3).Contains(5)
		assert.False(t, result)
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("finds index", func(t *testing.T) {
		result := list.Of(1, 2, 3, 2).IndexOf(2)
		assert.Equal(t, 1, result)
	})

	t.Run("returns -1 when not found", func(t *testing.T) {
		result := list.Of(1, 2, 3).IndexOf(5)
		assert.Equal(t, -1, result)
	})
}

func TestLastIndexOf(t *testing.T) {
	t.Run("finds last index", func(t *testing.T) {
		result := list.Of(1, 2, 3, 2).LastIndexOf(2)
		assert.Equal(t, 3, result)
	})

	t.Run("returns -1 when not found", func(t *testing.T) {
		result := list.Of(1, 2, 3).LastIndexOf(5)
		assert.Equal(t, -1, result)
	})
}

func TestFindLast(t *testing.T) {
	t.Run("finds last matching", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4).FindLast(func(x int) bool { return x%2 == 0 })
		assert.True(t, result.IsPresent())
		assert.Equal(t, 4, result.GetValue())
	})

	t.Run("returns empty when not found", func(t *testing.T) {
		result := list.Of(1, 3, 5).FindLast(func(x int) bool { return x%2 == 0 })
		assert.False(t, result.IsPresent())
	})
}

func TestListNone(t *testing.T) {
	t.Run("returns true when none match", func(t *testing.T) {
		result := list.Of(1, 3, 5).None(func(x int) bool { return x%2 == 0 })
		assert.True(t, result)
	})

	t.Run("returns false when one matches", func(t *testing.T) {
		result := list.Of(1, 2, 3).None(func(x int) bool { return x%2 == 0 })
		assert.False(t, result)
	})
}

func TestMinBy(t *testing.T) {
	t.Run("finds min by selector", func(t *testing.T) {
		result := list.Of("apple", "pie", "a").MinBy(func(s string) int { return len(s) })
		assert.True(t, result.IsPresent())
		assert.Equal(t, "a", result.GetValue())
	})
}

func TestMaxBy(t *testing.T) {
	t.Run("finds max by selector", func(t *testing.T) {
		result := list.Of("apple", "pie", "a").MaxBy(func(s string) int { return len(s) })
		assert.True(t, result.IsPresent())
		assert.Equal(t, "apple", result.GetValue())
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("groups by key", func(t *testing.T) {
		groups := list.Of(1, 2, 3, 4, 5, 6).GroupBy(func(x int) string {
			if x%2 == 0 {
				return "even"
			}
			return "odd"
		})
		assert.Equal(t, collection.List[int]{1, 3, 5}, groups["odd"])
		assert.Equal(t, collection.List[int]{2, 4, 6}, groups["even"])
	})
}

func TestPartition(t *testing.T) {
	t.Run("partitions by predicate", func(t *testing.T) {
		pass, fail := list.Of(1, 2, 3, 4, 5).Partition(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, collection.List[int]{2, 4}, pass)
		assert.Equal(t, collection.List[int]{1, 3, 5}, fail)
	})
}

func TestZip(t *testing.T) {
	t.Run("zips two lists", func(t *testing.T) {
		l1 := list.Of(1, 2, 3)
		l2 := list.Of("a", "b", "c")
		result := collection.Zip(l1, l2)
		assert.Equal(t, 3, result.Len())
		assert.Equal(t, 1, result[0].First)
		assert.Equal(t, "a", result[0].Second)
	})
}

func TestFlatten(t *testing.T) {
	t.Run("flattens nested lists", func(t *testing.T) {
		nested := collection.List[[]int]{{1, 2}, {3, 4}, {5}}
		result := collection.Flatten(nested)
		assert.Equal(t, collection.List[int]{1, 2, 3, 4, 5}, result)
	})
}

func TestDistinctBy(t *testing.T) {
	t.Run("removes duplicates by selector", func(t *testing.T) {
		result := list.Of("apple", "apricot", "banana", "blueberry").DistinctBy(func(s string) any { return s[0] })
		assert.Equal(t, collection.List[string]{"apple", "banana"}, result)
	})
}

func TestOnEach(t *testing.T) {
	t.Run("calls function for each then returns list", func(t *testing.T) {
		var collected []int
		result := list.Of(1, 2, 3).OnEach(func(x int) {
			collected = append(collected, x)
		})
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
		assert.Equal(t, []int{1, 2, 3}, collected)
	})
}

func TestAsSequence(t *testing.T) {
	t.Run("converts to lazy sequence", func(t *testing.T) {
		seq := list.Of(1, 2, 3).AsSequence()
		result := seq.Filter(func(x int) bool { return x > 1 }).ToSlice()
		assert.Equal(t, []int{2, 3}, result)
	})

	t.Run("early termination in sequence", func(t *testing.T) {
		seq := list.Of(1, 2, 3, 4, 5).AsSequence()
		result := seq.Filter(func(x int) bool { return x <= 2 }).ToSlice()
		assert.Equal(t, []int{1, 2}, result)
	})
}

func TestListFlatMapMethod(t *testing.T) {
	t.Run("flat maps list same type", func(t *testing.T) {
		result := list.Of(1, 2, 3).FlatMap(func(x int) []int { return []int{x, x * 10} })
		assert.Equal(t, collection.List[int]{1, 10, 2, 20, 3, 30}, result)
	})

	t.Run("flat maps empty list", func(t *testing.T) {
		result := list.Of[int]().FlatMap(func(x int) []int { return []int{x} })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestGroupByFreeFunction(t *testing.T) {
	t.Run("groups by int key", func(t *testing.T) {
		groups := collection.GroupBy(list.Of(1, 2, 3, 4, 5), func(x int) int { return x % 2 })
		assert.Equal(t, collection.List[int]{1, 3, 5}, groups[1])
		assert.Equal(t, collection.List[int]{2, 4}, groups[0])
	})
}

func TestFlattenMethod(t *testing.T) {
	t.Run("flattens using method with reflection", func(t *testing.T) {
		nested := collection.List[[]int]{{1, 2}, {3, 4}}
		result := nested.Flatten()
		assert.Equal(t, 4, result.Len())
	})
}

func TestMin(t *testing.T) {
	t.Run("finds minimum", func(t *testing.T) {
		result := collection.Min(list.Of(3, 1, 4, 1, 5))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := collection.Min(list.Of[int]())
		assert.False(t, result.IsPresent())
	})
}

func TestMax(t *testing.T) {
	t.Run("finds maximum", func(t *testing.T) {
		result := collection.Max(list.Of(3, 1, 4, 1, 5))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 5, result.GetValue())
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := collection.Max(list.Of[int]())
		assert.False(t, result.IsPresent())
	})
}

func TestAverage(t *testing.T) {
	t.Run("calculates average", func(t *testing.T) {
		result := collection.Average(list.Of(2, 4, 6, 8))
		assert.Equal(t, 5.0, result)
	})

	t.Run("returns 0 for empty list", func(t *testing.T) {
		result := collection.Average(list.Of[int]())
		assert.Equal(t, 0.0, result)
	})
}

func TestTakeLastEdgeCases(t *testing.T) {
	t.Run("takes last zero elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeLast(0)
		assert.Equal(t, collection.List[int]{}, result)
	})

	t.Run("takes negative becomes zero", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeLast(-1)
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestTakeWhileEdgeCases(t *testing.T) {
	t.Run("takes all when all match", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeWhile(func(x int) bool { return x < 10 })
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("takes none when first fails", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeWhile(func(x int) bool { return x > 10 })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestDropLastEdgeCases(t *testing.T) {
	t.Run("drops last zero elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).DropLast(0)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("drops negative becomes zero", func(t *testing.T) {
		result := list.Of(1, 2, 3).DropLast(-1)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("drops more than available", func(t *testing.T) {
		result := list.Of(1, 2).DropLast(5)
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestDropWhileEdgeCases(t *testing.T) {
	t.Run("drops all when all match", func(t *testing.T) {
		result := list.Of(1, 2, 3).DropWhile(func(x int) bool { return x < 10 })
		assert.Equal(t, collection.List[int]{}, result)
	})

	t.Run("drops none when first fails", func(t *testing.T) {
		result := list.Of(1, 2, 3).DropWhile(func(x int) bool { return x > 10 })
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestMinByEdgeCases(t *testing.T) {
	t.Run("returns empty for empty list", func(t *testing.T) {
		result := list.Of[string]().MinBy(func(s string) int { return len(s) })
		assert.False(t, result.IsPresent())
	})
}

func TestMaxByEdgeCases(t *testing.T) {
	t.Run("returns empty for empty list", func(t *testing.T) {
		result := list.Of[string]().MaxBy(func(s string) int { return len(s) })
		assert.False(t, result.IsPresent())
	})

	t.Run("finds max with single element", func(t *testing.T) {
		result := list.Of("hello").MaxBy(func(s string) int { return len(s) })
		assert.True(t, result.IsPresent())
		assert.Equal(t, "hello", result.GetValue())
	})
}

func TestZipEdgeCases(t *testing.T) {
	t.Run("zips lists of different lengths", func(t *testing.T) {
		l1 := list.Of(1, 2, 3, 4, 5)
		l2 := list.Of("a", "b")
		result := collection.Zip(l1, l2)
		assert.Equal(t, 2, result.Len())
	})

	t.Run("second list shorter", func(t *testing.T) {
		l1 := list.Of(1, 2)
		l2 := list.Of("a", "b", "c", "d")
		result := collection.Zip(l1, l2)
		assert.Equal(t, 2, result.Len())
	})
}

func TestFlattenEdgeCases(t *testing.T) {
	t.Run("flattens empty nested list", func(t *testing.T) {
		nested := collection.List[[]int]{}
		result := collection.Flatten(nested)
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestDistinctByEdgeCases(t *testing.T) {
	t.Run("returns empty for empty list", func(t *testing.T) {
		result := list.Of[string]().DistinctBy(func(s string) any { return s[0] })
		assert.Equal(t, collection.List[string]{}, result)
	})
}

func TestDropEdgeCases(t *testing.T) {
	t.Run("drops zero elements", func(t *testing.T) {
		result := list.Of(1, 2, 3).Drop(0)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("drops negative becomes zero", func(t *testing.T) {
		result := list.Of(1, 2, 3).Drop(-1)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestFlatMapFreeFunction(t *testing.T) {
	t.Run("flat maps to different type", func(t *testing.T) {
		result := collection.FlatMap(list.Of(1, 2, 3), func(x int) []string {
			return []string{string(rune('a' + x - 1))}
		})
		assert.Equal(t, collection.List[string]{"a", "b", "c"}, result)
	})

	t.Run("flat maps empty list returns empty", func(t *testing.T) {
		result := collection.FlatMap(list.Of[int](), func(x int) []string {
			return []string{string(rune('a' + x))}
		})
		assert.Equal(t, collection.List[string]{}, result)
	})
}

func TestSeqFilterEarlyTermination(t *testing.T) {
	t.Run("early termination in filter", func(t *testing.T) {
		seq := list.Of(1, 2, 3, 4, 5).AsSequence()
		count := 0
		filtered := seq.Filter(func(x int) bool { return x%2 == 0 })

		for item := range filtered.ToSlice() {
			count++
			_ = item
		}
		assert.Equal(t, 2, count)
	})

	t.Run("filter with early break via iter", func(t *testing.T) {
		seq := list.Of(2, 4, 6, 8, 10).AsSequence()
		filtered := seq.Filter(func(x int) bool { return x > 0 })

		count := 0
		for item := range filtered.Iter() {
			count++
			_ = item
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestSeqIter(t *testing.T) {
	t.Run("returns iter.Seq for direct iteration", func(t *testing.T) {
		seq := list.Of(1, 2, 3).AsSequence()
		var result []int
		for item := range seq.Iter() {
			result = append(result, item)
		}
		assert.Equal(t, []int{1, 2, 3}, result)
	})
}

func TestSeqToSliceEmpty(t *testing.T) {
	t.Run("empty sequence to slice", func(t *testing.T) {
		seq := list.Of[int]().AsSequence()
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestAsSequenceEarlyTermination(t *testing.T) {
	t.Run("early termination stops iteration", func(t *testing.T) {
		l := list.Of(1, 2, 3, 4, 5)
		seq := l.AsSequence()
		count := 0
		for item := range seq.Filter(func(x int) bool { return true }).ToSlice() {
			count++
			_ = item
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestFlattenMethodReflection(t *testing.T) {
	t.Run("flattens with reflection for nested slices", func(t *testing.T) {
		nested := collection.List[[]string]{{"a", "b"}, {"c"}}
		result := nested.Flatten()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("handles []any elements directly", func(t *testing.T) {
		single := collection.List[[]any]{{1, 2}, {3}}
		result := single.Flatten()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("empty nested list returns empty", func(t *testing.T) {
		nested := collection.List[[]int]{}
		result := nested.Flatten()
		assert.Equal(t, collection.List[any]{}, result)
	})
}

func TestMaxByIterations(t *testing.T) {
	t.Run("iterates through all elements", func(t *testing.T) {
		result := list.Of("a", "bb", "ccc", "dd").MaxBy(func(s string) int { return len(s) })
		assert.True(t, result.IsPresent())
		assert.Equal(t, "ccc", result.GetValue())
	})
}

func TestMapIndexed(t *testing.T) {
	t.Run("maps with index", func(t *testing.T) {
		result := collection.MapIndexed(list.Of("a", "b", "c"), func(index int, item string) string {
			return string(rune('0'+index)) + item
		})
		assert.Equal(t, collection.List[string]{"0a", "1b", "2c"}, result)
	})

	t.Run("maps empty list", func(t *testing.T) {
		result := collection.MapIndexed(list.Of[string](), func(index int, item string) string {
			return item
		})
		assert.Equal(t, collection.List[string]{}, result)
	})
}

func TestFilterIndexed(t *testing.T) {
	t.Run("filters with index", func(t *testing.T) {
		result := list.Of("a", "b", "c", "d").FilterIndexed(func(index int, item string) bool {
			return index%2 == 0
		})
		assert.Equal(t, collection.List[string]{"a", "c"}, result)
	})

	t.Run("filters empty list", func(t *testing.T) {
		result := list.Of[string]().FilterIndexed(func(index int, item string) bool {
			return true
		})
		assert.Equal(t, collection.List[string]{}, result)
	})
}

func TestFilterNot(t *testing.T) {
	t.Run("filters negated predicate", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).FilterNot(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, collection.List[int]{1, 3, 5}, result)
	})
}

func TestRunningFold(t *testing.T) {
	t.Run("running fold accumulates", func(t *testing.T) {
		result := collection.RunningFold(list.Of(1, 2, 3, 4), 0, func(acc, x int) int { return acc + x })
		assert.Equal(t, collection.List[int]{0, 1, 3, 6, 10}, result)
	})

	t.Run("empty list returns initial only", func(t *testing.T) {
		result := collection.RunningFold(list.Of[int](), 10, func(acc, x int) int { return acc + x })
		assert.Equal(t, collection.List[int]{10}, result)
	})
}

func TestRunningReduce(t *testing.T) {
	t.Run("running reduce accumulates", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4).RunningReduce(func(acc, x int) int { return acc + x })
		assert.Equal(t, collection.List[int]{1, 3, 6, 10}, result)
	})

	t.Run("empty list returns empty", func(t *testing.T) {
		result := list.Of[int]().RunningReduce(func(acc, x int) int { return acc + x })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestScan(t *testing.T) {
	t.Run("scan is alias for runningFold", func(t *testing.T) {
		result := collection.Scan(list.Of(1, 2, 3), 0, func(acc, x int) int { return acc + x })
		assert.Equal(t, collection.List[int]{0, 1, 3, 6}, result)
	})
}

func TestAssociateWith(t *testing.T) {
	t.Run("associates elements with values", func(t *testing.T) {
		result := list.Of("a", "bb", "ccc").AssociateWith(func(s string) any { return len(s) })
		assert.Equal(t, 1, result["a"])
		assert.Equal(t, 2, result["bb"])
		assert.Equal(t, 3, result["ccc"])
	})
}

func TestZipWithNext(t *testing.T) {
	t.Run("zips consecutive pairs", func(t *testing.T) {
		l := list.Of(1, 2, 3, 4)
		result := collection.ZipWithNext(l)
		assert.Equal(t, 3, len(result))
		assert.Equal(t, 1, result[0].First)
		assert.Equal(t, 2, result[0].Second)
		assert.Equal(t, 3, result[2].First)
		assert.Equal(t, 4, result[2].Second)
	})

	t.Run("single element returns empty", func(t *testing.T) {
		l := list.Of(1)
		result := collection.ZipWithNext(l)
		assert.Equal(t, 0, len(result))
	})

	t.Run("empty returns empty", func(t *testing.T) {
		l := list.Of[int]()
		result := collection.ZipWithNext(l)
		assert.Equal(t, 0, len(result))
	})
}

func TestWindowed(t *testing.T) {
	t.Run("creates sliding windows", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).Windowed(3, 1, false)
		assert.Equal(t, 3, len(result))
		assert.Equal(t, collection.List[int]{1, 2, 3}, result[0])
		assert.Equal(t, collection.List[int]{2, 3, 4}, result[1])
		assert.Equal(t, collection.List[int]{3, 4, 5}, result[2])
	})

	t.Run("with step 2", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).Windowed(2, 2, false)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, collection.List[int]{1, 2}, result[0])
		assert.Equal(t, collection.List[int]{3, 4}, result[1])
	})

	t.Run("partial windows when allowed", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).Windowed(3, 2, true)
		assert.Equal(t, 3, len(result))
		assert.Equal(t, collection.List[int]{5}, result[2])
	})

	t.Run("empty list returns empty", func(t *testing.T) {
		result := list.Of[int]().Windowed(3, 1, false)
		assert.Equal(t, 0, len(result))
	})

	t.Run("size zero returns empty", func(t *testing.T) {
		result := list.Of(1, 2, 3).Windowed(0, 1, false)
		assert.Equal(t, 0, len(result))
	})

	t.Run("step zero returns empty", func(t *testing.T) {
		result := list.Of(1, 2, 3).Windowed(2, 0, false)
		assert.Equal(t, 0, len(result))
	})

	t.Run("partial windows with larger step", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5, 6, 7).Windowed(3, 3, true)
		assert.Equal(t, 3, len(result))
		assert.Equal(t, collection.List[int]{1, 2, 3}, result[0])
		assert.Equal(t, collection.List[int]{4, 5, 6}, result[1])
		assert.Equal(t, collection.List[int]{7}, result[2])
	})
}

func TestSingle(t *testing.T) {
	t.Run("returns single element", func(t *testing.T) {
		result := list.Of(42).Single()
		assert.True(t, result.IsPresent())
		assert.Equal(t, 42, result.GetValue())
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := list.Of[int]().Single()
		assert.False(t, result.IsPresent())
	})

	t.Run("returns empty for multiple elements", func(t *testing.T) {
		result := list.Of(1, 2).Single()
		assert.False(t, result.IsPresent())
	})
}

func TestElementAt(t *testing.T) {
	t.Run("returns element at index", func(t *testing.T) {
		result := list.Of(10, 20, 30).ElementAt(1)
		assert.True(t, result.IsPresent())
		assert.Equal(t, 20, result.GetValue())
	})

	t.Run("returns empty for negative index", func(t *testing.T) {
		result := list.Of(10, 20, 30).ElementAt(-1)
		assert.False(t, result.IsPresent())
	})

	t.Run("returns empty for out of bounds", func(t *testing.T) {
		result := list.Of(10, 20, 30).ElementAt(5)
		assert.False(t, result.IsPresent())
	})
}

func TestShuffled(t *testing.T) {
	t.Run("returns shuffled copy", func(t *testing.T) {
		original := list.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		result := original.Shuffled()
		assert.Equal(t, 10, result.Len())
		assert.NotEqual(t, original, result)
		for _, v := range original {
			assert.True(t, result.Contains(v))
		}
	})

	t.Run("empty list returns empty", func(t *testing.T) {
		result := list.Of[int]().Shuffled()
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestRandom(t *testing.T) {
	t.Run("returns random element", func(t *testing.T) {
		l := list.Of(1, 2, 3, 4, 5)
		result := l.Random()
		assert.True(t, result.IsPresent())
		assert.True(t, l.Contains(result.GetValue()))
	})

	t.Run("empty list returns empty optional", func(t *testing.T) {
		result := list.Of[int]().Random()
		assert.False(t, result.IsPresent())
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("finds first matching index", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).FindIndex(func(x int) bool { return x > 2 })
		assert.Equal(t, 2, result)
	})

	t.Run("returns -1 when not found", func(t *testing.T) {
		result := list.Of(1, 2, 3).FindIndex(func(x int) bool { return x > 10 })
		assert.Equal(t, -1, result)
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("returns true for non-empty list", func(t *testing.T) {
		result := list.Of(1, 2, 3).IsNotEmpty()
		assert.True(t, result)
	})

	t.Run("returns false for empty list", func(t *testing.T) {
		result := list.Of[int]().IsNotEmpty()
		assert.False(t, result)
	})
}

func TestContainsAll(t *testing.T) {
	t.Run("returns true when all elements are contained", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).ContainsAll(list.Of(2, 3, 4))
		assert.True(t, result)
	})

	t.Run("returns false when some elements are not contained", func(t *testing.T) {
		result := list.Of(1, 2, 3).ContainsAll(list.Of(2, 3, 10))
		assert.False(t, result)
	})

	t.Run("returns true for empty elements list", func(t *testing.T) {
		result := list.Of(1, 2, 3).ContainsAll(list.Of[int]())
		assert.True(t, result)
	})
}

func TestTakeLastWhile(t *testing.T) {
	t.Run("takes elements from end while predicate is true", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).TakeLastWhile(func(x int) bool { return x > 3 })
		assert.Equal(t, collection.List[int]{4, 5}, result)
	})

	t.Run("returns empty when last element doesn't match", func(t *testing.T) {
		result := list.Of(5, 4, 3, 2, 1).TakeLastWhile(func(x int) bool { return x > 3 })
		assert.Equal(t, collection.List[int]{}, result)
	})

	t.Run("returns all when all match from end", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeLastWhile(func(x int) bool { return x > 0 })
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestDropLastWhile(t *testing.T) {
	t.Run("drops elements from end while predicate is true", func(t *testing.T) {
		result := list.Of(1, 2, 3, 4, 5).DropLastWhile(func(x int) bool { return x > 3 })
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})

	t.Run("returns all when last element doesn't match", func(t *testing.T) {
		result := list.Of(5, 4, 3, 2, 1).DropLastWhile(func(x int) bool { return x > 3 })
		assert.Equal(t, collection.List[int]{5, 4, 3, 2, 1}, result)
	})

	t.Run("returns empty when all match", func(t *testing.T) {
		result := list.Of(1, 2, 3).DropLastWhile(func(x int) bool { return x > 0 })
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestSortedDescending(t *testing.T) {
	t.Run("sorts in descending order", func(t *testing.T) {
		result := collection.SortedDescending(list.Of(3, 1, 4, 1, 5, 9, 2, 6))
		assert.Equal(t, collection.List[int]{9, 6, 5, 4, 3, 2, 1, 1}, result)
	})

	t.Run("handles empty list", func(t *testing.T) {
		result := collection.SortedDescending(list.Of[int]())
		assert.Equal(t, collection.List[int]{}, result)
	})

	t.Run("handles single element", func(t *testing.T) {
		result := collection.SortedDescending(list.Of(42))
		assert.Equal(t, collection.List[int]{42}, result)
	})
}

func TestSumOf(t *testing.T) {
	t.Run("sums values using selector", func(t *testing.T) {
		type item struct{ value int }
		items := list.Of(item{1}, item{2}, item{3})
		result := collection.SumOf(items, func(i item) float64 { return float64(i.value) })
		assert.Equal(t, 6.0, result)
	})

	t.Run("returns 0 for empty list", func(t *testing.T) {
		result := collection.SumOf(list.Of[int](), func(x int) float64 { return float64(x) })
		assert.Equal(t, 0.0, result)
	})
}

func TestFoldIndexed(t *testing.T) {
	t.Run("folds with index", func(t *testing.T) {
		result := collection.FoldIndexed(list.Of("a", "b", "c"), "", func(idx int, acc string, s string) string {
			return acc + fmt.Sprintf("%d:%s,", idx, s)
		})
		assert.Equal(t, "0:a,1:b,2:c,", result)
	})

	t.Run("returns initial for empty list", func(t *testing.T) {
		result := collection.FoldIndexed(list.Of[int](), 100, func(idx int, acc int, x int) int {
			return acc + x
		})
		assert.Equal(t, 100, result)
	})
}

func TestReduceIndexed(t *testing.T) {
	t.Run("reduces with index", func(t *testing.T) {
		result := collection.ReduceIndexed(list.Of(1, 2, 3, 4), func(idx int, acc int, x int) int {
			return acc + x*idx
		})
		assert.Equal(t, 1+2*1+3*2+4*3, result.GetValue())
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := collection.ReduceIndexed(list.Of[int](), func(idx int, acc int, x int) int {
			return acc + x
		})
		assert.True(t, result.IsEmpty())
	})

	t.Run("returns single element for single element list", func(t *testing.T) {
		result := collection.ReduceIndexed(list.Of(42), func(idx int, acc int, x int) int {
			return acc + x
		})
		assert.Equal(t, 42, result.GetValue())
	})
}

func TestFoldRight(t *testing.T) {
	t.Run("folds from right to left", func(t *testing.T) {
		result := collection.FoldRight(list.Of("a", "b", "c"), "", func(s string, acc string) string {
			return acc + s
		})
		assert.Equal(t, "cba", result)
	})

	t.Run("returns initial for empty list", func(t *testing.T) {
		result := collection.FoldRight(list.Of[int](), 100, func(x int, acc int) int {
			return acc + x
		})
		assert.Equal(t, 100, result)
	})
}

func TestReduceRight(t *testing.T) {
	t.Run("reduces from right to left", func(t *testing.T) {
		result := collection.ReduceRight(list.Of("a", "b", "c"), func(s string, acc string) string {
			return acc + s
		})
		assert.Equal(t, "cba", result.GetValue())
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := collection.ReduceRight(list.Of[int](), func(x int, acc int) int {
			return acc + x
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestRunningFoldIndexed(t *testing.T) {
	t.Run("creates running fold with index", func(t *testing.T) {
		result := collection.RunningFoldIndexed(list.Of(1, 2, 3), 0, func(idx int, acc int, x int) int {
			return acc + x*(idx+1)
		})
		assert.Equal(t, collection.List[int]{0, 1, 1 + 4, 1 + 4 + 9}, result)
	})
}

func TestRunningReduceIndexed(t *testing.T) {
	t.Run("creates running reduce with index", func(t *testing.T) {
		result := list.Of(10, 20, 30).RunningReduceIndexed(func(idx int, acc int, x int) int {
			return acc + x + idx
		})
		assert.Equal(t, collection.List[int]{10, 10 + 20 + 1, 10 + 20 + 1 + 30 + 2}, result)
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := list.Of[int]().RunningReduceIndexed(func(idx int, acc int, x int) int {
			return acc + x
		})
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestMapNotNull(t *testing.T) {
	t.Run("maps and filters out nils", func(t *testing.T) {
		result := collection.MapNotNull(list.Of(1, 2, 3, 4, 5), func(x int) *int {
			if x%2 == 0 {
				v := x * 10
				return &v
			}
			return nil
		})
		assert.Equal(t, collection.List[int]{20, 40}, result)
	})

	t.Run("returns empty when all nil", func(t *testing.T) {
		result := collection.MapNotNull(list.Of(1, 3, 5), func(x int) *int {
			return nil
		})
		assert.Equal(t, collection.List[int]{}, result)
	})
}

func TestMapIndexedNotNull(t *testing.T) {
	t.Run("maps with index and filters out nils", func(t *testing.T) {
		result := collection.MapIndexedNotNull(list.Of("a", "b", "c"), func(idx int, s string) *string {
			if idx%2 == 0 {
				v := fmt.Sprintf("%d:%s", idx, s)
				return &v
			}
			return nil
		})
		assert.Equal(t, collection.List[string]{"0:a", "2:c"}, result)
	})
}

func TestUnzip(t *testing.T) {
	t.Run("unzips pairs into two lists", func(t *testing.T) {
		pairs := list.Of(
			collection.PairOf(1, "a"),
			collection.PairOf(2, "b"),
			collection.PairOf(3, "c"),
		)
		first, second := collection.Unzip(pairs)
		assert.Equal(t, collection.List[int]{1, 2, 3}, first)
		assert.Equal(t, collection.List[string]{"a", "b", "c"}, second)
	})

	t.Run("returns empty lists for empty input", func(t *testing.T) {
		first, second := collection.Unzip(list.Of[collection.Pair[int, string]]())
		assert.Equal(t, collection.List[int]{}, first)
		assert.Equal(t, collection.List[string]{}, second)
	})
}

func TestFindLastIndex(t *testing.T) {
	t.Run("finds last matching index", func(t *testing.T) {
		result := list.Of(1, 2, 3, 2, 1).FindLastIndex(func(x int) bool { return x == 2 })
		assert.Equal(t, 3, result)
	})

	t.Run("returns -1 when not found", func(t *testing.T) {
		result := list.Of(1, 2, 3).FindLastIndex(func(x int) bool { return x > 10 })
		assert.Equal(t, -1, result)
	})
}

func TestFirstNotNullOf(t *testing.T) {
	t.Run("returns first non-null result", func(t *testing.T) {
		result := collection.FirstNotNullOf(list.Of(1, 2, 3, 4), func(x int) *int {
			if x > 2 {
				v := x * 10
				return &v
			}
			return nil
		})
		assert.Equal(t, 30, result.GetValue())
	})

	t.Run("returns empty when all null", func(t *testing.T) {
		result := collection.FirstNotNullOf(list.Of(1, 2, 3), func(x int) *int {
			return nil
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestPlus(t *testing.T) {
	t.Run("adds single element", func(t *testing.T) {
		result := list.Of(1, 2, 3).Plus(4)
		assert.Equal(t, collection.List[int]{1, 2, 3, 4}, result)
	})

	t.Run("adds to empty list", func(t *testing.T) {
		result := list.Of[int]().Plus(1)
		assert.Equal(t, collection.List[int]{1}, result)
	})
}

func TestPlusAll(t *testing.T) {
	t.Run("adds all elements from another list", func(t *testing.T) {
		result := list.Of(1, 2, 3).PlusAll(list.Of(4, 5, 6))
		assert.Equal(t, collection.List[int]{1, 2, 3, 4, 5, 6}, result)
	})

	t.Run("adds empty list", func(t *testing.T) {
		result := list.Of(1, 2, 3).PlusAll(list.Of[int]())
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestMinus(t *testing.T) {
	t.Run("removes first occurrence of element", func(t *testing.T) {
		result := list.Of(1, 2, 3, 2, 4).Minus(2)
		assert.Equal(t, collection.List[int]{1, 3, 2, 4}, result)
	})

	t.Run("returns same list when element not found", func(t *testing.T) {
		result := list.Of(1, 2, 3).Minus(10)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestMinusAll(t *testing.T) {
	t.Run("removes all occurrences of elements", func(t *testing.T) {
		result := list.Of(1, 2, 3, 2, 4, 3).MinusAll(list.Of(2, 3))
		assert.Equal(t, collection.List[int]{1, 4}, result)
	})

	t.Run("removes nothing when empty list", func(t *testing.T) {
		result := list.Of(1, 2, 3).MinusAll(list.Of[int]())
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestFoldRightIndexed(t *testing.T) {
	t.Run("folds from right with index", func(t *testing.T) {
		result := collection.FoldRightIndexed(list.Of("a", "b", "c"), "", func(idx int, s string, acc string) string {
			return acc + fmt.Sprintf("%d:%s,", idx, s)
		})
		assert.Equal(t, "2:c,1:b,0:a,", result)
	})
}

func TestReduceRightIndexed(t *testing.T) {
	t.Run("reduces from right with index", func(t *testing.T) {
		result := collection.ReduceRightIndexed(list.Of(1, 2, 3, 4), func(idx int, x int, acc int) int {
			return acc + x*idx
		})
		assert.Equal(t, 4+3*2+2*1+1*0, result.GetValue())
	})

	t.Run("returns empty for empty list", func(t *testing.T) {
		result := collection.ReduceRightIndexed(list.Of[int](), func(idx int, x int, acc int) int {
			return acc + x
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestToSet(t *testing.T) {
	t.Run("converts list to set", func(t *testing.T) {
		result := collection.ToSet(list.Of(1, 2, 3, 2, 1))
		assert.Equal(t, 3, result.Len())
		assert.True(t, result.Contains(1))
		assert.True(t, result.Contains(2))
		assert.True(t, result.Contains(3))
	})

	t.Run("empty list to empty set", func(t *testing.T) {
		result := collection.ToSet(list.Of[int]())
		assert.True(t, result.IsEmpty())
	})
}

func TestToMap(t *testing.T) {
	t.Run("converts list to map using key selector", func(t *testing.T) {
		result := collection.ToMap(list.Of("a", "bb", "ccc"), func(s string) int {
			return len(s)
		})
		assert.Equal(t, 3, result.Len())
		assert.Equal(t, "a", result[1])
		assert.Equal(t, "bb", result[2])
		assert.Equal(t, "ccc", result[3])
	})
}

func TestToMapWithValue(t *testing.T) {
	t.Run("converts list to map using key and value selectors", func(t *testing.T) {
		result := collection.ToMapWithValue(list.Of("a", "bb", "ccc"),
			func(s string) int { return len(s) },
			func(s string) string { return s + "!" },
		)
		assert.Equal(t, 3, result.Len())
		assert.Equal(t, "a!", result[1])
		assert.Equal(t, "bb!", result[2])
		assert.Equal(t, "ccc!", result[3])
	})
}

func TestLet(t *testing.T) {
	t.Run("transforms list with let", func(t *testing.T) {
		result := collection.Let(list.Of(1, 2, 3), func(l collection.List[int]) string {
			return fmt.Sprintf("count: %d", l.Len())
		})
		assert.Equal(t, "count: 3", result)
	})

	t.Run("chains transformations", func(t *testing.T) {
		filtered := list.Of(1, 2, 3, 4, 5).Filter(func(x int) bool { return x > 2 })
		result := collection.Let(filtered, func(l collection.List[int]) int {
			return l.Len() * 10
		})
		assert.Equal(t, 30, result)
	})
}

func TestAlso(t *testing.T) {
	t.Run("performs side effect and returns same list", func(t *testing.T) {
		var sideEffect int
		result := list.Of(1, 2, 3).Also(func(l collection.List[int]) {
			sideEffect = l.Len()
		})
		assert.Equal(t, 3, sideEffect)
		assert.Equal(t, collection.List[int]{1, 2, 3}, result)
	})
}

func TestTakeIf(t *testing.T) {
	t.Run("returns list when predicate is true", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeIf(func(l collection.List[int]) bool {
			return l.Len() > 2
		})
		assert.True(t, result.IsPresent())
		assert.Equal(t, collection.List[int]{1, 2, 3}, result.GetValue())
	})

	t.Run("returns empty when predicate is false", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeIf(func(l collection.List[int]) bool {
			return l.Len() > 10
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestTakeUnless(t *testing.T) {
	t.Run("returns list when predicate is false", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeUnless(func(l collection.List[int]) bool {
			return l.IsEmpty()
		})
		assert.True(t, result.IsPresent())
	})

	t.Run("returns empty when predicate is true", func(t *testing.T) {
		result := list.Of(1, 2, 3).TakeUnless(func(l collection.List[int]) bool {
			return l.Len() > 0
		})
		assert.True(t, result.IsEmpty())
	})
}

func TestChainListToMapToSetToList(t *testing.T) {
	t.Run("full chain: list -> map -> set -> list", func(t *testing.T) {
		// Start with list of strings
		initial := list.Of("apple", "banana", "cherry", "apricot")

		// Convert to map (first letter -> word)
		asMap := collection.ToMap(initial, func(s string) byte {
			return s[0]
		})

		// Get keys as set
		keysSet := collection.ToSet(asMap.Keys())

		// Convert back to list
		finalList := keysSet.ToList()

		// We should have unique first letters: a, b, c
		assert.Equal(t, 3, len(finalList))
	})
}

func TestPipe(t *testing.T) {
	t.Run("chains multiple let operations", func(t *testing.T) {
		result := collection.Pipe(list.Of(1, 2, 3, 4, 5)).
			Let(func(l collection.List[int]) collection.List[int] {
				return l.Filter(func(x int) bool { return x > 2 })
			}).
			Let(func(l collection.List[int]) collection.List[int] {
				return collection.ListMap(l, func(x int) int { return x * 10 })
			}).
			Let(func(l collection.List[int]) collection.List[int] {
				return l.Take(2)
			}).
			Value()

		assert.Equal(t, collection.List[int]{30, 40}, result)
	})

	t.Run("deep chaining .let.let.let.let.let.let", func(t *testing.T) {
		result := collection.Pipe(list.Of(1, 2, 3)).
			Let(func(l collection.List[int]) collection.List[int] { return l.Append(4) }).
			Let(func(l collection.List[int]) collection.List[int] { return l.Append(5) }).
			Let(func(l collection.List[int]) collection.List[int] { return l.Append(6) }).
			Let(func(l collection.List[int]) collection.List[int] { return l.Reversed() }).
			Let(func(l collection.List[int]) collection.List[int] { return l.Take(3) }).
			Let(func(l collection.List[int]) collection.List[int] { return l.Reversed() }).
			Value()

		assert.Equal(t, collection.List[int]{4, 5, 6}, result)
	})

	t.Run("also in chain", func(t *testing.T) {
		var sideEffects []int
		result := collection.Pipe(list.Of(1, 2, 3)).
			Also(func(l collection.List[int]) { sideEffects = append(sideEffects, l.Len()) }).
			Let(func(l collection.List[int]) collection.List[int] { return l.Append(4) }).
			Also(func(l collection.List[int]) { sideEffects = append(sideEffects, l.Len()) }).
			Value()

		assert.Equal(t, collection.List[int]{1, 2, 3, 4}, result)
		assert.Equal(t, []int{3, 4}, sideEffects)
	})

	t.Run("takeIf in chain", func(t *testing.T) {
		result := collection.Pipe(list.Of(1, 2, 3, 4, 5)).
			Let(func(l collection.List[int]) collection.List[int] { return l.Filter(func(x int) bool { return x > 2 }) }).
			TakeIf(func(l collection.List[int]) bool { return l.Len() > 0 })

		assert.True(t, result.IsPresent())
		assert.Equal(t, collection.List[int]{3, 4, 5}, result.GetValue())
	})
}

func TestPipeTransform(t *testing.T) {
	t.Run("transforms to different type", func(t *testing.T) {
		result := collection.PipeTransform(
			collection.Pipe(list.Of(1, 2, 3, 4, 5)).
				Let(func(l collection.List[int]) collection.List[int] {
					return l.Filter(func(x int) bool { return x > 2 })
				}).
				Value(),
			func(l collection.List[int]) string {
				return fmt.Sprintf("count=%d", l.Len())
			},
		)
		assert.Equal(t, "count=3", result)
	})

	t.Run("chain list to set to list", func(t *testing.T) {
		initialList := list.Of(1, 2, 2, 3, 3, 3)
		asSet := collection.ToSet(initialList)
		backToList := asSet.ToList()
		assert.Equal(t, 3, len(backToList))
	})
}
