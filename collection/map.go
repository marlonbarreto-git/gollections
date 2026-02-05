package collection

import (
	"encoding/json"
	"fmt"

	"github.com/marlonbarreto-git/gollections/tomove/function"
	"github.com/marlonbarreto-git/gollections/tomove/optional"
	"github.com/marlonbarreto-git/gollections/tomove/types"
)

type MutableMap[K comparable, V any] map[K]V

func Map[K, NK comparable, V, NV any](original MutableMap[K, V], fn func(K, V) (NK, NV)) MutableMap[NK, NV] {
	newMap := map[NK]NV{}

	for key, value := range original {
		newKey, newValue := fn(key, value)
		newMap[newKey] = newValue
	}

	return newMap
}

func MapKeys[K, NK comparable, V any](original MutableMap[K, V], fn func(K, V) NK) MutableMap[NK, V] {
	newMap := map[NK]V{}

	for key, value := range original {
		newMap[fn(key, value)] = value
	}

	return newMap
}

func MapValues[K comparable, V, NV any](original MutableMap[K, V], fn func(K, V) NV) MutableMap[K, NV] {
	newMap := map[K]NV{}

	for key, value := range original {
		newMap[key] = fn(key, value)
	}

	return newMap
}

func (m MutableMap[K, V]) Map(fn func(K, V) (any, any)) MutableMap[any, any] {
	newMap := map[any]any{}

	for key, value := range m {
		newKey, newValue := fn(key, value)
		newMap[newKey] = newValue
	}

	return newMap
}

func (m MutableMap[K, V]) Reduce(fn func(acc any, key K, value V) any, acc any) any {
	for key, value := range m {
		acc = fn(acc, key, value)
	}

	return acc
}

func (m MutableMap[K, V]) ForEach(consumer function.BiConsumer[K, V]) {
	for k, v := range m {
		consumer(k, v)
	}
}

func (m MutableMap[K, V]) Filter(predicate function.BiPredicate[K, V]) MutableMap[K, V] {
	filteredMap := map[K]V{}

	for k, v := range m {
		if predicate(k, v) {
			filteredMap[k] = v
		}
	}

	return filteredMap
}

func (m MutableMap[K, V]) IsEmpty() bool {
	return m.Len() == 0
}

func (m MutableMap[K, V]) Len() int {
	return len(m)
}

func (m MutableMap[K, V]) Count(predicate function.BiPredicate[K, V]) (count int) {
	for k, v := range m {
		if predicate(k, v) {
			count++
		}
	}

	return
}

func (m MutableMap[K, V]) Copy() MutableMap[K, V] {
	copiedMap := make(map[K]V, len(m))

	for k, v := range m {
		copiedMap[k] = v
	}

	return copiedMap
}

func (m MutableMap[K, V]) Keys() (keys List[K]) {
	for key := range m {
		keys = append(keys, key)
	}

	return
}

func (m MutableMap[K, V]) Values() (values List[V]) {
	for _, value := range m {
		values = append(values, value)
	}

	return
}

func (m MutableMap[K, V]) Entries() (values []Pair[K, V]) {
	for k, v := range m {
		values = append(values, PairOf(k, v))
	}

	return
}

func (m MutableMap[K, V]) Remove(key K) {
	delete(m, key)
}

func (m MutableMap[K, V]) String() string {
	bytes, err := json.Marshal(map[K]V(m))
	if err != nil {
		return m.stringFallthrough()
	}

	return string(bytes)
}

func (m MutableMap[K, V]) stringFallthrough() string {
	return Map(m, func(k K, v V) (string, any) {
		return fmt.Sprintf("%v", k), v
	}).String()
}

func (m MutableMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

func (m MutableMap[K, V]) GetOrPut(key K, defaultFn func() V) V {
	if value, ok := m[key]; ok {
		return value
	}
	value := defaultFn()
	m[key] = value
	return value
}

func (m MutableMap[K, V]) ContainsKey(key K) bool {
	_, ok := m[key]
	return ok
}

func (m MutableMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m {
		var valAny any = value
		var vAny any = v
		if valAny == vAny {
			return true
		}
	}
	return false
}

func (m MutableMap[K, V]) Merge(other MutableMap[K, V]) {
	for k, v := range other {
		m[k] = v
	}
}

func (m MutableMap[K, V]) PutAll(pairs ...Pair[K, V]) {
	for _, pair := range pairs {
		m[pair.First()] = pair.Second()
	}
}

func (m MutableMap[K, V]) FilterKeys(predicate func(K) bool) MutableMap[K, V] {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k) {
			result[k] = v
		}
	}
	return result
}

func (m MutableMap[K, V]) FilterValues(predicate func(V) bool) MutableMap[K, V] {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(v) {
			result[k] = v
		}
	}
	return result
}

func (m MutableMap[K, V]) ToList() []Pair[K, V] {
	result := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, PairOf(k, v))
	}
	return result
}

func (m MutableMap[K, V]) Any(predicate function.BiPredicate[K, V]) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

func (m MutableMap[K, V]) All(predicate function.BiPredicate[K, V]) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func (m MutableMap[K, V]) None(predicate function.BiPredicate[K, V]) bool {
	return !m.Any(predicate)
}

func (m MutableMap[K, V]) ToSet() Set[K] {
	result := make(Set[K], len(m))
	for k := range m {
		result[k] = types.EmptyInstance
	}
	return result
}

func (m MutableMap[K, V]) Also(fn func(MutableMap[K, V])) MutableMap[K, V] {
	fn(m)
	return m
}

func (m MutableMap[K, V]) TakeIf(predicate func(MutableMap[K, V]) bool) optional.Optional[MutableMap[K, V]] {
	if predicate(m) {
		return optional.Of(m)
	}
	return optional.Empty[MutableMap[K, V]]()
}

func (m MutableMap[K, V]) TakeUnless(predicate func(MutableMap[K, V]) bool) optional.Optional[MutableMap[K, V]] {
	if !predicate(m) {
		return optional.Of(m)
	}
	return optional.Empty[MutableMap[K, V]]()
}
