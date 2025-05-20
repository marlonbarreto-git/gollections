package collection

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/marlonbarreto-git/gollections/tomove/function"
	"github.com/marlonbarreto-git/gollections/tomove/optional"
	"github.com/marlonbarreto-git/gollections/tomove/types"
)

type List[T any] []T

// Map creates a new List with the results of applying the given function to each item in the list
func ListMap[T, R any](list List[T], fn Function[T, R]) List[R] {
	mappedItems := make([]R, list.Len())

	list.ForEachIndexed(func(index int, item T) {
		mappedItems[index] = fn(item)
	})

	return mappedItems
}

func (list List[T]) ToArray() []T {
	return list
}

func (list List[T]) Distinct() List[T] {
	if len(list) == 0 {
		return nil
	}

	result := make(List[T], 0, len(list))
	set := make(map[any]types.Empty, len(list))

	if !reflect.TypeOf(list[0]).Comparable() {
		panic("collection: Distinct requires list elements to be comparable")
	}

	for _, item := range list {
		if _, exists := set[item]; !exists {
			set[item] = types.EmptyInstance
			result = append(result, item)
		}
	}

	return result
}

func (list List[T]) Associate(fn Function[T, any]) (result MutableMap[any, T]) {
	result = make(map[any]T, list.Len())

	for _, item := range list {
		result[fn(item)] = item
	}

	return
}

func (list List[T]) AssociateBy(keySelector string) (result MutableMap[any, T]) {
	if list.Len() == 0 {
		return map[any]T{}
	}

	typeOf := reflect.TypeOf(list).Elem()
	kind := typeOf.Kind()
	isPointer := kind == reflect.Pointer
	if isPointer {
		typeOf = typeOf.Elem()
		kind = typeOf.Kind()
	}

	if kind != reflect.Struct {
		panic("list elements must be a structure to use AssociateBy")
	}

	if _, ok := typeOf.FieldByName(keySelector); !ok {
		panic("keySelector not found")
	}

	return list.Associate(func(item T) any {
		valueOf := reflect.ValueOf(item)
		if isPointer {
			valueOf = reflect.Indirect(valueOf)
		}

		return valueOf.FieldByName(keySelector).Interface()
	})
}

func (list List[T]) Join(separator string, toString ...func(item T) string) string {
	if len(toString) == 0 {
		toString = append(toString, func(item T) string {
			return fmt.Sprintf("%v", item)
		})
	}

	return strings.Join(ListMap[T, string](list, toString[0]).ToArray(), separator)
}

func (list List[T]) Filter(fn Predicate[T]) (filteredItems List[T]) {
	for _, item := range list {
		if fn(item) {
			filteredItems = append(filteredItems, item)
		}
	}

	return filteredItems

}

func (list List[T]) Append(item T) List[T] {
	return append(list, item)
}

func (list *List[T]) Add(item T) *List[T] {
	*list = append(*list, item)
	return list
}

func (list List[T]) ForEach(fn Consumer[T]) {
	for _, item := range list {
		fn(item)
	}
}

func (list List[T]) ForEachIndexed(fn IndexedConsumer[T]) {
	for i, item := range list {
		fn(i, item)
	}
}

func (list List[T]) First() optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}

	return optional.Of(list[0])
}

func (list List[T]) Get(index int) T {
	return list[index]
}

func (list List[T]) Find(fn Predicate[T]) optional.Optional[T] {
	for _, item := range list {
		if fn(item) {
			return optional.Of(item)
		}
	}

	return optional.Empty[T]()
}

func (list List[T]) Last() optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}

	return optional.Of(list[len(list)-1])
}

func (list List[T]) Len() int {
	return len(list)
}

func (list List[T]) IsEmpty() bool {
	return list.Len() == 0
}

func (list List[T]) Some(fn Predicate[T]) bool {
	for _, item := range list {
		if fn(item) {
			return true
		}
	}

	return false
}

func (list List[T]) Every(fn Predicate[T]) bool {
	for _, item := range list {
		if !fn(item) {
			return false
		}
	}

	return true
}

func (list List[T]) Slice(start, end int) List[T] {
	return list[start:end]
}

func (list List[T]) Count(fn Predicate[T]) (count int) {
	for _, item := range list {
		if fn(item) {
			count++
		}
	}

	return count
}

func (list List[T]) Sum(fn Function[T, float64]) (sum float64) {
	for _, item := range list {
		sum += fn(item)
	}

	return sum
}
