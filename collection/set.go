package collection

import (
	"fmt"
	"strings"

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
	for k := range s {
		values = append(values, k)
	}

	return
}
