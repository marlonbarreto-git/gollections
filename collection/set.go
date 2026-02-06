package collection

import (
	"fmt"
	"strings"

	"github.com/marlonbarreto-git/gollections/tomove/optional"
	. "github.com/marlonbarreto-git/gollections/tomove/types"
)

type Set[K comparable] map[K]Empty

func (s Set[K]) Contains(item K) bool {
	_, ok := s[item]
	return ok
}

func (s Set[K]) Add(item K) bool {
	_, exists := s[item]
	s[item] = EmptyInstance
	return !exists
}

func (s Set[K]) Remove(item K) {
	delete(s, item)
}

func (s Set[K]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[K]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[K]) Len() int {
	return len(s)
}

func (s Set[K]) String() string {
	var str strings.Builder
	str.WriteString("{")

	first := true
	for k := range s {
		if !first {
			str.WriteString(", ")
		}
		str.WriteString(fmt.Sprintf("%v", k))
		first = false
	}

	str.WriteString("}")
	return str.String()
}

func (s Set[K]) Values() (values List[K]) {
	values = make(List[K], 0, len(s))
	for k := range s {
		values = append(values, k)
	}

	return
}

func (s Set[K]) Union(other Set[K]) Set[K] {
	result := make(Set[K], len(s)+len(other))
	for k := range s {
		result[k] = EmptyInstance
	}
	for k := range other {
		result[k] = EmptyInstance
	}
	return result
}

func (s Set[K]) Intersect(other Set[K]) Set[K] {
	result := make(Set[K])
	for k := range s {
		if _, ok := other[k]; ok {
			result[k] = EmptyInstance
		}
	}
	return result
}

func (s Set[K]) Subtract(other Set[K]) Set[K] {
	result := make(Set[K])
	for k := range s {
		if _, ok := other[k]; !ok {
			result[k] = EmptyInstance
		}
	}
	return result
}

func (s Set[K]) Filter(predicate func(K) bool) Set[K] {
	result := make(Set[K])
	for k := range s {
		if predicate(k) {
			result[k] = EmptyInstance
		}
	}
	return result
}

func (s Set[K]) ForEach(fn func(K)) {
	for k := range s {
		fn(k)
	}
}

func (s Set[K]) Any(predicate func(K) bool) bool {
	for k := range s {
		if predicate(k) {
			return true
		}
	}
	return false
}

func (s Set[K]) All(predicate func(K) bool) bool {
	for k := range s {
		if !predicate(k) {
			return false
		}
	}
	return true
}

func (s Set[K]) None(predicate func(K) bool) bool {
	return !s.Any(predicate)
}

func (s Set[K]) First() optional.Optional[K] {
	for k := range s {
		return optional.Of(k)
	}
	return optional.Empty[K]()
}

func (s Set[K]) ToList() []K {
	result := make([]K, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}

func (s Set[K]) Also(fn func(Set[K])) Set[K] {
	fn(s)
	return s
}

func (s Set[K]) TakeIf(predicate func(Set[K]) bool) optional.Optional[Set[K]] {
	if predicate(s) {
		return optional.Of(s)
	}
	return optional.Empty[Set[K]]()
}

func (s Set[K]) TakeUnless(predicate func(Set[K]) bool) optional.Optional[Set[K]] {
	if !predicate(s) {
		return optional.Of(s)
	}
	return optional.Empty[Set[K]]()
}

func (s Set[K]) ToMap(valueSelector func(K) any) MutableMap[K, any] {
	result := make(MutableMap[K, any], len(s))
	for k := range s {
		result[k] = valueSelector(k)
	}
	return result
}
