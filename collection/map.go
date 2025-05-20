package collection

import (
	"encoding/json"
	"fmt"

	"github.com/marlonbarreto-git/gollections/tomove/function"
)

type MutableMap[K comparable, V any] map[K]V

func Map[K, NK comparable, V, NV any](original MutableMap[K, V], fn func(K, V) (NK, NV)) MutableMap[NK, NV] {
	newMap := make(map[NK]NV, len(original))

	for key, value := range original {
		newKey, newValue := fn(key, value)
		newMap[newKey] = newValue
	}

	return newMap
}

func MapKeys[K, NK comparable, V any](original MutableMap[K, V], fn func(K, V) NK) MutableMap[NK, V] {
	newMap := make(map[NK]V, len(original))

	for key, value := range original {
		newMap[fn(key, value)] = value
	}

	return newMap
}

func MapValues[K comparable, V, NV any](original MutableMap[K, V], fn func(K, V) NV) MutableMap[K, NV] {
	newMap := make(map[K]NV, len(original))

	for key, value := range original {
		newMap[key] = fn(key, value)
	}

	return newMap
}

func (m MutableMap[K, V]) Map(fn func(K, V) (any, any)) MutableMap[any, any] {
	newMap := make(map[any]any, len(m))

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
	filteredMap := make(map[K]V, len(m))

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
	keys = make(List[K], 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return
}

func (m MutableMap[K, V]) Values() (values List[V]) {
	values = make(List[V], 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}

	return
}

func (m MutableMap[K, V]) Entries() (values []Pair[K, V]) {
	values = make([]Pair[K, V], 0, len(m))
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
