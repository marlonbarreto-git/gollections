package maps

import (
	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/tomove/function"
)

// Api represents a map data structure.
type Api[K comparable, V any] interface {

	// Map applies the specified function to each key-value pair in the map and returns the mutable map with the results.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   result := m.Map(func(k string, v int) (any, any) {
	//       return k, v * 2
	//   })
	// Returns:
	//   {
	//       "1": 2,
	//       "2": 4,
	//   }
	Map(fn func(K, V) (any, any)) collection.MutableMap[any, any]

	// Reduce reduces the map to a single value by applying the specified function to each key-value pair and accumulating the results.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   sum := m.Reduce(func(acc any, key string, value int) any {
	//       return acc.(int) + value
	//   }, 0)
	// Returns:
	//    3
	Reduce(fn func(acc any, key K, value V) any, acc any) any

	// ForEach iterates over each key-value pair in the map and applies the specified consumer function.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   m.ForEach(func(k string, v int) {
	//       fmt.Println(k, v)
	//   })
	ForEach(consumer function.BiConsumer[K, V])

	// Filter returns a new mutable map containing only the key-value pairs that satisfy the specified predicate function.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   filtered := m.Filter(func(k string, v int) bool {
	//       return v > 1
	//   })
	// Returns:
	//   {
	//       "2": 2,
	//   }
	Filter(predicate function.BiPredicate[K, V]) collection.MutableMap[K, V]

	// IsEmpty returns true if the map is empty, otherwise returns false.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   empty := m.IsEmpty()
	// Returns:
	//    false
	IsEmpty() bool

	// Len returns the number of key-value pairs in the map.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   size := m.Len()
	// Returns:
	//    2
	Len() int

	// Count returns the number of key-value pairs that satisfy the specified predicate function.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   count := m.Count(func(k string, v int) bool {
	//       return v > 1
	//   })
	// Returns:
	//    1
	Count(predicate function.BiPredicate[K, V]) (count int)

	// Copy creates a shallow copy of the map.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   copy := m.Copy()
	// Returns:
	//   {
	//       "1": 1,
	//       "2": 2,
	//   }
	Copy() collection.MutableMap[K, V]

	// Keys returns a list containing all the keys in the map.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   keys := m.Keys()
	// Returns:
	//    [1, 2]
	Keys() (keys collection.List[K])

	// Values returns a list containing all the values in the map.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   values := m.Values()
	// Returns:
	//    [1, 2]
	Values() (values collection.List[V])

	// Entries returns a list containing all the key-value pairs in the map.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   entries := m.Entries()
	// Returns:
	//    [{
	//       "key": 1,
	//       "value": 1,
	//   }, {
	//       "key": 2,
	//       "value": 2,
	//   }]
	Entries() (values collection.List[collection.Pair[K, V]])

	// Remove deletes the specified key and its associated value from the map.
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	//   m.Remove(2)
	// Returns:
	//   {
	//       "1": 1,
	//   }
	Remove(key K)

	// String returns a map string with json style
	// Example:
	//   m := Map.Of(pair.Of("1",1), pair.Of("2",2))
	// Returns:
	//   {
	//       "1": 1,
	//       "2": 2
	//   }
	String() string
}

func Of[K comparable, V any](pairs ...collection.Pair[K, V]) collection.MutableMap[K, V] {
	rawMap := make(map[K]V)

	for _, pairItem := range pairs {
		rawMap[pairItem.First()] = pairItem.Second()
	}

	return rawMap
}

func From[K comparable, V any](rawMap map[K]V) collection.MutableMap[K, V] {
	return rawMap
}
