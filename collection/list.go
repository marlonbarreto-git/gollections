package collection

import (
	"cmp"
	"fmt"
	"iter"
	"math/rand"
	"reflect"
	"slices"
	"strings"

	. "github.com/marlonbarreto-git/gollections/tomove/function"
	"github.com/marlonbarreto-git/gollections/tomove/optional"
	"github.com/marlonbarreto-git/gollections/tomove/types"
)

type List[T any] []T

// ListMap creates a new List with the results of applying the given function to each item in the list
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

func (list List[T]) Reduce(initial T, fn func(acc, item T) T) T {
	acc := initial
	for _, item := range list {
		acc = fn(acc, item)
	}
	return acc
}

func Fold[T, R any](list List[T], initial R, fn func(R, T) R) R {
	acc := initial
	for _, item := range list {
		acc = fn(acc, item)
	}
	return acc
}

func (list List[T]) FlatMap(fn func(T) []T) List[T] {
	var result List[T]
	for _, item := range list {
		result = append(result, fn(item)...)
	}
	if result == nil {
		return List[T]{}
	}
	return result
}

func FlatMap[T, R any](list List[T], fn func(T) []R) List[R] {
	var result List[R]
	for _, item := range list {
		result = append(result, fn(item)...)
	}
	if result == nil {
		return List[R]{}
	}
	return result
}

func (list List[T]) Reversed() List[T] {
	if len(list) == 0 {
		return List[T]{}
	}
	result := make(List[T], len(list))
	for i, item := range list {
		result[len(list)-1-i] = item
	}
	return result
}

func (list List[T]) Sorted(cmpFn func(a, b T) int) List[T] {
	result := make(List[T], len(list))
	copy(result, list)
	slices.SortFunc(result, cmpFn)
	return result
}

func (list List[T]) Take(n int) List[T] {
	if n <= 0 {
		return List[T]{}
	}
	if n >= len(list) {
		result := make(List[T], len(list))
		copy(result, list)
		return result
	}
	result := make(List[T], n)
	copy(result, list[:n])
	return result
}

func (list List[T]) TakeLast(n int) List[T] {
	if n <= 0 {
		return List[T]{}
	}
	if n >= len(list) {
		result := make(List[T], len(list))
		copy(result, list)
		return result
	}
	result := make(List[T], n)
	copy(result, list[len(list)-n:])
	return result
}

func (list List[T]) TakeWhile(fn Predicate[T]) List[T] {
	var result List[T]
	for _, item := range list {
		if !fn(item) {
			break
		}
		result = append(result, item)
	}
	if result == nil {
		return List[T]{}
	}
	return result
}

func (list List[T]) Drop(n int) List[T] {
	if n <= 0 {
		result := make(List[T], len(list))
		copy(result, list)
		return result
	}
	if n >= len(list) {
		return List[T]{}
	}
	result := make(List[T], len(list)-n)
	copy(result, list[n:])
	return result
}

func (list List[T]) DropLast(n int) List[T] {
	if n <= 0 {
		result := make(List[T], len(list))
		copy(result, list)
		return result
	}
	if n >= len(list) {
		return List[T]{}
	}
	result := make(List[T], len(list)-n)
	copy(result, list[:len(list)-n])
	return result
}

func (list List[T]) DropWhile(fn Predicate[T]) List[T] {
	var result List[T]
	dropping := true
	for _, item := range list {
		if dropping && fn(item) {
			continue
		}
		dropping = false
		result = append(result, item)
	}
	if result == nil {
		return List[T]{}
	}
	return result
}

func (list List[T]) Chunked(size int) []List[T] {
	if size <= 0 || len(list) == 0 {
		return []List[T]{}
	}
	var result []List[T]
	for i := 0; i < len(list); i += size {
		end := i + size
		if end > len(list) {
			end = len(list)
		}
		chunk := make(List[T], end-i)
		copy(chunk, list[i:end])
		result = append(result, chunk)
	}
	return result
}

func (list List[T]) Contains(target T) bool {
	for _, item := range list {
		var targetAny any = target
		var itemAny any = item
		if targetAny == itemAny {
			return true
		}
	}
	return false
}

func (list List[T]) IndexOf(target T) int {
	for i, item := range list {
		var targetAny any = target
		var itemAny any = item
		if targetAny == itemAny {
			return i
		}
	}
	return -1
}

func (list List[T]) LastIndexOf(target T) int {
	for i := len(list) - 1; i >= 0; i-- {
		var targetAny any = target
		var itemAny any = list[i]
		if targetAny == itemAny {
			return i
		}
	}
	return -1
}

func (list List[T]) FindLast(fn Predicate[T]) optional.Optional[T] {
	for i := len(list) - 1; i >= 0; i-- {
		if fn(list[i]) {
			return optional.Of(list[i])
		}
	}
	return optional.Empty[T]()
}

func (list List[T]) None(fn Predicate[T]) bool {
	return !list.Some(fn)
}

func (list List[T]) MinBy(selector func(T) int) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	minItem := list[0]
	minVal := selector(minItem)
	for _, item := range list[1:] {
		val := selector(item)
		if val < minVal {
			minVal = val
			minItem = item
		}
	}
	return optional.Of(minItem)
}

func (list List[T]) MaxBy(selector func(T) int) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	maxItem := list[0]
	maxVal := selector(maxItem)
	for _, item := range list[1:] {
		val := selector(item)
		if val > maxVal {
			maxVal = val
			maxItem = item
		}
	}
	return optional.Of(maxItem)
}

func (list List[T]) GroupBy(keyFn func(T) string) map[string]List[T] {
	result := make(map[string]List[T])
	for _, item := range list {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func GroupBy[T any, K comparable](list List[T], keyFn func(T) K) map[K]List[T] {
	result := make(map[K]List[T])
	for _, item := range list {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func (list List[T]) Partition(fn Predicate[T]) (pass, fail List[T]) {
	for _, item := range list {
		if fn(item) {
			pass = append(pass, item)
		} else {
			fail = append(fail, item)
		}
	}
	if pass == nil {
		pass = List[T]{}
	}
	if fail == nil {
		fail = List[T]{}
	}
	return
}

type ZipPair[T, U any] struct {
	First  T
	Second U
}

func Zip[T, U any](list1 List[T], list2 List[U]) List[ZipPair[T, U]] {
	minLen := len(list1)
	if len(list2) < minLen {
		minLen = len(list2)
	}
	result := make(List[ZipPair[T, U]], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = ZipPair[T, U]{First: list1[i], Second: list2[i]}
	}
	return result
}

func Flatten[T any](list List[[]T]) List[T] {
	var result List[T]
	for _, sublist := range list {
		result = append(result, sublist...)
	}
	if result == nil {
		return List[T]{}
	}
	return result
}

func (list List[T]) Flatten() List[any] {
	var result List[any]
	for _, item := range list {
		if sublist, ok := any(item).([]any); ok {
			result = append(result, sublist...)
		} else {
			val := reflect.ValueOf(item)
			if val.Kind() == reflect.Slice {
				for i := 0; i < val.Len(); i++ {
					result = append(result, val.Index(i).Interface())
				}
			}
		}
	}
	if result == nil {
		return List[any]{}
	}
	return result
}

func (list List[T]) DistinctBy(selector func(T) any) List[T] {
	seen := make(map[any]types.Empty)
	var result List[T]
	for _, item := range list {
		key := selector(item)
		if _, exists := seen[key]; !exists {
			seen[key] = types.EmptyInstance
			result = append(result, item)
		}
	}
	if result == nil {
		return List[T]{}
	}
	return result
}

func (list List[T]) OnEach(fn Consumer[T]) List[T] {
	for _, item := range list {
		fn(item)
	}
	return list
}

type Seq[T any] struct {
	iter iter.Seq[T]
}

func (s Seq[T]) Filter(fn func(T) bool) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for item := range s.iter {
				if fn(item) {
					if !yield(item) {
						return
					}
				}
			}
		},
	}
}

func (s Seq[T]) Iter() iter.Seq[T] {
	return s.iter
}

func (s Seq[T]) ToSlice() []T {
	var result []T
	for item := range s.iter {
		result = append(result, item)
	}
	if result == nil {
		return []T{}
	}
	return result
}

func (list List[T]) AsSequence() Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for _, item := range list {
				if !yield(item) {
					return
				}
			}
		},
	}
}

func Min[T cmp.Ordered](list List[T]) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	minVal := list[0]
	for _, item := range list[1:] {
		if item < minVal {
			minVal = item
		}
	}
	return optional.Of(minVal)
}

func Max[T cmp.Ordered](list List[T]) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	maxVal := list[0]
	for _, item := range list[1:] {
		if item > maxVal {
			maxVal = item
		}
	}
	return optional.Of(maxVal)
}

func Average[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](list List[T]) float64 {
	if len(list) == 0 {
		return 0
	}
	var sum float64
	for _, item := range list {
		sum += float64(item)
	}
	return sum / float64(len(list))
}

func MapIndexed[T, R any](list List[T], fn func(int, T) R) List[R] {
	if len(list) == 0 {
		return List[R]{}
	}
	result := make(List[R], len(list))
	for i, item := range list {
		result[i] = fn(i, item)
	}
	return result
}

func (list List[T]) FilterIndexed(fn func(int, T) bool) List[T] {
	var result List[T]
	for i, item := range list {
		if fn(i, item) {
			result = append(result, item)
		}
	}
	if result == nil {
		return List[T]{}
	}
	return result
}

func (list List[T]) FilterNot(fn Predicate[T]) List[T] {
	return list.Filter(func(item T) bool { return !fn(item) })
}

func RunningFold[T, R any](list List[T], initial R, fn func(R, T) R) List[R] {
	result := make(List[R], 0, len(list)+1)
	result = append(result, initial)
	acc := initial
	for _, item := range list {
		acc = fn(acc, item)
		result = append(result, acc)
	}
	return result
}

func Scan[T, R any](list List[T], initial R, fn func(R, T) R) List[R] {
	return RunningFold(list, initial, fn)
}

func (list List[T]) RunningReduce(fn func(T, T) T) List[T] {
	if len(list) == 0 {
		return List[T]{}
	}
	result := make(List[T], 0, len(list))
	acc := list[0]
	result = append(result, acc)
	for _, item := range list[1:] {
		acc = fn(acc, item)
		result = append(result, acc)
	}
	return result
}

func (list List[T]) AssociateWith(fn Function[T, any]) map[any]any {
	result := make(map[any]any, len(list))
	for _, item := range list {
		var key any = item
		result[key] = fn(item)
	}
	return result
}

type ConsecutivePair[T any] struct {
	First  T
	Second T
}

func ZipWithNext[T any](list List[T]) []ConsecutivePair[T] {
	if len(list) < 2 {
		return []ConsecutivePair[T]{}
	}
	result := make([]ConsecutivePair[T], len(list)-1)
	for i := 0; i < len(list)-1; i++ {
		result[i] = ConsecutivePair[T]{First: list[i], Second: list[i+1]}
	}
	return result
}

func (list List[T]) Windowed(size, step int, partialWindows bool) []List[T] {
	if len(list) == 0 || size <= 0 || step <= 0 {
		return []List[T]{}
	}
	var result []List[T]
	for i := 0; i <= len(list)-size; i += step {
		window := make(List[T], size)
		copy(window, list[i:i+size])
		result = append(result, window)
	}
	if partialWindows {
		lastStart := (len(result)) * step
		if lastStart < len(list) {
			for i := lastStart; i < len(list); i += step {
				remaining := len(list) - i
				if remaining > 0 && remaining < size {
					window := make(List[T], remaining)
					copy(window, list[i:])
					result = append(result, window)
				}
			}
		}
	}
	return result
}

func (list List[T]) Single() optional.Optional[T] {
	if len(list) != 1 {
		return optional.Empty[T]()
	}
	return optional.Of(list[0])
}

func (list List[T]) ElementAt(index int) optional.Optional[T] {
	if index < 0 || index >= len(list) {
		return optional.Empty[T]()
	}
	return optional.Of(list[index])
}

func (list List[T]) Shuffled() List[T] {
	if len(list) == 0 {
		return List[T]{}
	}
	result := make(List[T], len(list))
	copy(result, list)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

func (list List[T]) Random() optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	return optional.Of(list[rand.Intn(len(list))])
}

func (list List[T]) FindIndex(fn Predicate[T]) int {
	for i, item := range list {
		if fn(item) {
			return i
		}
	}
	return -1
}

func (list List[T]) FindLastIndex(fn Predicate[T]) int {
	for i := len(list) - 1; i >= 0; i-- {
		if fn(list[i]) {
			return i
		}
	}
	return -1
}

func (list List[T]) IsNotEmpty() bool {
	return len(list) > 0
}

func (list List[T]) ContainsAll(elements List[T]) bool {
	for _, elem := range elements {
		found := false
		for _, item := range list {
			if reflect.DeepEqual(elem, item) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list List[T]) TakeLastWhile(fn Predicate[T]) List[T] {
	idx := len(list)
	for i := len(list) - 1; i >= 0; i-- {
		if !fn(list[i]) {
			break
		}
		idx = i
	}
	return list[idx:]
}

func (list List[T]) DropLastWhile(fn Predicate[T]) List[T] {
	idx := len(list)
	for i := len(list) - 1; i >= 0; i-- {
		if !fn(list[i]) {
			break
		}
		idx = i
	}
	return list[:idx]
}

func SortedDescending[T cmp.Ordered](list List[T]) List[T] {
	if len(list) == 0 {
		return List[T]{}
	}
	result := make(List[T], len(list))
	copy(result, list)
	slices.SortFunc(result, func(a, b T) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})
	return result
}

func SumOf[T any](list List[T], selector func(T) float64) float64 {
	var sum float64
	for _, item := range list {
		sum += selector(item)
	}
	return sum
}

func FoldIndexed[T, R any](list List[T], initial R, fn func(int, R, T) R) R {
	acc := initial
	for i, item := range list {
		acc = fn(i, acc, item)
	}
	return acc
}

func ReduceIndexed[T any](list List[T], fn func(int, T, T) T) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	acc := list[0]
	for i := 1; i < len(list); i++ {
		acc = fn(i, acc, list[i])
	}
	return optional.Of(acc)
}

func FoldRight[T, R any](list List[T], initial R, fn func(T, R) R) R {
	acc := initial
	for i := len(list) - 1; i >= 0; i-- {
		acc = fn(list[i], acc)
	}
	return acc
}

func ReduceRight[T any](list List[T], fn func(T, T) T) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	acc := list[len(list)-1]
	for i := len(list) - 2; i >= 0; i-- {
		acc = fn(list[i], acc)
	}
	return optional.Of(acc)
}

func RunningFoldIndexed[T, R any](list List[T], initial R, fn func(int, R, T) R) List[R] {
	result := make(List[R], len(list)+1)
	result[0] = initial
	acc := initial
	for i, item := range list {
		acc = fn(i, acc, item)
		result[i+1] = acc
	}
	return result
}

func (list List[T]) RunningReduceIndexed(fn func(int, T, T) T) List[T] {
	if len(list) == 0 {
		return List[T]{}
	}
	result := make(List[T], len(list))
	result[0] = list[0]
	for i := 1; i < len(list); i++ {
		result[i] = fn(i, result[i-1], list[i])
	}
	return result
}

func MapNotNull[T, R any](list List[T], fn func(T) *R) List[R] {
	result := make(List[R], 0)
	for _, item := range list {
		if mapped := fn(item); mapped != nil {
			result = append(result, *mapped)
		}
	}
	return result
}

func MapIndexedNotNull[T, R any](list List[T], fn func(int, T) *R) List[R] {
	result := make(List[R], 0)
	for i, item := range list {
		if mapped := fn(i, item); mapped != nil {
			result = append(result, *mapped)
		}
	}
	return result
}

func Unzip[A comparable, B any](list List[Pair[A, B]]) (List[A], List[B]) {
	first := make(List[A], len(list))
	second := make(List[B], len(list))
	for i, p := range list {
		first[i] = p.First()
		second[i] = p.Second()
	}
	return first, second
}

func FirstNotNullOf[T, R any](list List[T], fn func(T) *R) optional.Optional[R] {
	for _, item := range list {
		if result := fn(item); result != nil {
			return optional.Of(*result)
		}
	}
	return optional.Empty[R]()
}

func (list List[T]) Plus(element T) List[T] {
	return append(list, element)
}

func (list List[T]) PlusAll(elements List[T]) List[T] {
	return append(list, elements...)
}

func (list List[T]) Minus(element T) List[T] {
	result := make(List[T], 0, len(list))
	removed := false
	for _, item := range list {
		if !removed && reflect.DeepEqual(item, element) {
			removed = true
			continue
		}
		result = append(result, item)
	}
	return result
}

func (list List[T]) MinusAll(elements List[T]) List[T] {
	toRemove := make(map[int]struct{})
	for _, elem := range elements {
		for i, item := range list {
			if _, alreadyRemoved := toRemove[i]; !alreadyRemoved && reflect.DeepEqual(item, elem) {
				toRemove[i] = struct{}{}
			}
		}
	}
	result := make(List[T], 0, len(list)-len(toRemove))
	for i, item := range list {
		if _, remove := toRemove[i]; !remove {
			result = append(result, item)
		}
	}
	return result
}

func FoldRightIndexed[T, R any](list List[T], initial R, fn func(int, T, R) R) R {
	acc := initial
	for i := len(list) - 1; i >= 0; i-- {
		acc = fn(i, list[i], acc)
	}
	return acc
}

func ReduceRightIndexed[T any](list List[T], fn func(int, T, T) T) optional.Optional[T] {
	if len(list) == 0 {
		return optional.Empty[T]()
	}
	acc := list[len(list)-1]
	for i := len(list) - 2; i >= 0; i-- {
		acc = fn(i, list[i], acc)
	}
	return optional.Of(acc)
}

func ToSet[T comparable](list List[T]) Set[T] {
	result := make(Set[T], len(list))
	for _, item := range list {
		result[item] = types.EmptyInstance
	}
	return result
}

func ToMap[T any, K comparable](list List[T], keySelector func(T) K) MutableMap[K, T] {
	result := make(MutableMap[K, T], len(list))
	for _, item := range list {
		result[keySelector(item)] = item
	}
	return result
}

func ToMapWithValue[T any, K comparable, V any](list List[T], keySelector func(T) K, valueSelector func(T) V) MutableMap[K, V] {
	result := make(MutableMap[K, V], len(list))
	for _, item := range list {
		result[keySelector(item)] = valueSelector(item)
	}
	return result
}

func Let[T, R any](list List[T], fn func(List[T]) R) R {
	return fn(list)
}

func (list List[T]) Also(fn func(List[T])) List[T] {
	fn(list)
	return list
}

func (list List[T]) TakeIf(predicate func(List[T]) bool) optional.Optional[List[T]] {
	if predicate(list) {
		return optional.Of(list)
	}
	return optional.Empty[List[T]]()
}

func (list List[T]) TakeUnless(predicate func(List[T]) bool) optional.Optional[List[T]] {
	if !predicate(list) {
		return optional.Of(list)
	}
	return optional.Empty[List[T]]()
}
