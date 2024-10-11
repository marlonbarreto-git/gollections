package list

import (
	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/iterable"
	"github.com/marlonbarreto-git/gollections/tomove/function"
	"github.com/marlonbarreto-git/gollections/tomove/optional"
)

type Api[T any] interface {
	iterable.Api[T, collection.List[T]]

	// ToArray returns the array representation of the list
	// Example:
	//
	// 	nums := list.Of(1, 2, 3, 4, 5)
	// 	nums.ToArray()
	//
	// Output: [1, 2, 3, 4, 5]
	ToArray() []T

	// Append adds an item to the list
	// Example:
	//
	// 	nums := list.Of(1, 2, 3, 4, 5)
	// 	nums.Append(6)
	//
	// Output: [1, 2, 3, 4, 5, 6]
	Append(item T) collection.List[T]

	// Associate returns a map using the result of the function as a key
	// Example:
	//
	//	cars := list.Of(Car{Brand: "bmw"}, Car{Brand: "audi"})
	// 	cars.Associate(func(car Car) (key any) {
	//		return fmt.Sprintf("Brand %v", car.Brand)
	// 	})
	//
	// Output map[Brand bmw:{"bmw"} Brand audi:{"audi"}]
	Associate(fn function.Function[T, any]) (result map[any]T)

	// AssociateBy returns a map using a keySelector as a key
	// Example:
	//
	//	cars := list.Of(Car{Brand: "bmw"}, Car{Brand: "audi"})
	// 	cars.AssociateBy("Brand")
	//
	// Output map[bmw:{"bmw"} audi:{"audi"}]
	AssociateBy(keySelector string) (result map[any]T)

	// Join returns a string representation of the list with the given separator
	// Example:
	//
	// 	nums := list.Of(1, 2, 3, 4, 5)
	// 	nums.Join(", ")
	//
	// Output: "1, 2, 3, 4, 5"
	Join(separator string, toString ...func(item T) string) string

	// Get returns the item at the given index
	// Example:
	//
	// 	nums := list.Of(1, 2, 3, 4, 5)
	// 	nums.Get(2)
	//
	// Output: 3
	Get(index int) T

	// Find returns the first item in the list that satisfies the given predicate
	// Example:
	//
	// 	nums := list.Of(1, 2, 3, 4, 5)
	// 	nums.Find(func(item int) bool {
	// 	    return item > 2
	// 	})
	//
	// Output: Optional(3)
	Find(fn function.Predicate[T]) optional.Optional[T]

	// Some returns true if some item in the list pass the given function
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.Some(func(item int) bool {
	//	    return item > 10
	//	})
	//
	// Output: false
	Some(fn function.Predicate[T]) bool

	// Every returns true if all the items in the list pass the given function
	// Example:
	//
	// 	nums := list.Of(1, 2, 3, 4, 5)
	// 	nums.Every(func(item int) bool {
	// 	    return item > 0
	// 	})
	//
	// Output: true
	Every(fn function.Predicate[T]) bool

	// Slice returns a slice between start position and end position
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.Slice(1, 3)
	//
	// Output: [2, 3]
	Slice(start, end int) collection.List[T]

	// Count returns the number of items in the list that satisfy the given predicate
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.Count(func(item int) bool {
	//	    return item > 2
	//	})
	//
	// Output: 3
	Count(fn function.Predicate[T]) (count int)

	// Len returns the length of the list
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.Len()
	//
	// Output: 5
	Len() int

	// IsEmpty returns if the list is empty
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.IsEmpty()
	//
	// Output: false
	IsEmpty() bool

	// Distinct returns a list with distinct elements
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	//	nums.Distinct()
	//
	// Output: [1, 2, 3, 4, 5]
	Distinct() collection.List[T]

	// Sum returns the sum of the elements in the list
	// Example:
	//
	//	nums := list.Of(1, 2, 3, 4, 5)
	//	nums.Sum(func(item int) float64 {
	//	    return float64(item)
	//	})
	//
	// Output: 15
	Sum(fn function.Function[T, float64]) float64
}

// Interface verification
//var _ Api[int] = collection.List[int](nil)

// Of creates a new List with the given items
func Of[T any](items ...T) collection.List[T] {
	return items
}

// From creates a new List with the given items
func From[T any](items []T) collection.List[T] {
	return items
}
