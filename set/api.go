package set

import (
	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/tomove/types"
)

// Api represents a set data structure
type Api[K comparable] interface {

	// Contains checks if the set contains the given item
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	s.Contains("item1") // true
	//	s.Contains("item3") // false
	Contains(K) bool

	// String returns a string representation of the set
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	s.String() // "{item1, item2}"
	String() string

	// Add adds an item to the set
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	s.Add("item3")
	//	s.Contains("item3") // true
	Add(item K) bool

	// Remove removes an item from the set
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	s.Remove("item1")
	//	s.Contains("item1") // false
	Remove(item K)

	// Clear removes all items from the set
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	s.Clear()
	//	s.IsEmpty() // true
	Clear()

	// IsEmpty checks if the set is empty
	// Example:
	//
	//	s := set.Of[string]()
	//	s.IsEmpty() // true
	//	s.Add("item1")
	//	s.IsEmpty() // false
	IsEmpty() bool

	// Len returns the number of items in the set
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	s.Len() // 2
	Len() int

	// Values returns a list of all items in the set
	// Example:
	//
	//	s := set.Of("item1", "item2")
	//	values := s.Values()
	//	values.Contains("item1") // true
	Values() (values collection.List[K])
}

func Of[K comparable](items ...K) collection.Set[K] {
	rawSet := make(map[K]types.Empty, len(items))

	for _, item := range items {
		rawSet[item] = types.EmptyInstance
	}

	return rawSet
}

func From[K comparable, V any](rawMap map[K]V) collection.Set[K] {
	rawSet := make(map[K]types.Empty, len(rawMap))

	for item := range rawMap {
		rawSet[item] = types.EmptyInstance
	}

	return rawSet
}
