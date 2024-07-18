package iterable

import (
	"github.com/marlonbarreto-git/gollections/tomove/function"
	"github.com/marlonbarreto-git/gollections/tomove/optional"
)

type Api[T, C any] interface {

	// Filter returns a new collection with the items that pass the given function
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.Filter(func(item int) bool {
	//	    return item > 2
	//	})
	//
	// Output: [3, 4, 5]
	Filter(function.Predicate[T]) C

	// ForEach iterates over the items in the collection and applies the given function
	// Example:
	//
	//	nums := list.Of(1, 2, 3)
	//	nums.ForEach(func(item int) {
	//	    fmt.Println(item)
	//	})
	//
	// Output: 1
	//         2
	//         3
	ForEach(consumer function.Consumer[T])

	// ForEachIndexed iterates over the items in the collection and applies the given function
	// Example:
	//
	//	nums := list.Of(1, 2, 3)
	//	nums.ForEachIndexed(func(index int, item int) {
	//	    fmt.Println(index, item)
	//	})
	//
	// Output: 0 1
	//         1 2
	//         2 3
	ForEachIndexed(consumer function.IndexedConsumer[T])

	// First returns the first item in the collection
	// Example:
	//
	//	nums := list.Of(1, 2, 3)
	//	nums.First()
	//
	// Output: 1
	First() optional.Optional[T]

	// Last returns the last item in the collection
	// Example:
	//
	//	nums := list.Of(1, 2, 3)
	//	nums.Last()
	//
	// Output: 3
	Last() optional.Optional[T]
}
