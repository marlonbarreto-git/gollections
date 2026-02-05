package sequence

import (
	"cmp"
	"iter"
	"slices"

	"github.com/marlonbarreto-git/gollections/tomove/optional"
)

type Seq[T any] struct {
	iter iter.Seq[T]
}

type Pair[T, U any] struct {
	First  T
	Second U
}

type IndexedValue[T any] struct {
	Index int
	Value T
}

func Of[T any](items ...T) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for _, item := range items {
				if !yield(item) {
					return
				}
			}
		},
	}
}

func From[T any](slice []T) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for _, item := range slice {
				if !yield(item) {
					return
				}
			}
		},
	}
}

func FromIter[T any](it iter.Seq[T]) Seq[T] {
	return Seq[T]{iter: it}
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

func (s Seq[T]) Map(fn func(T) T) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for item := range s.iter {
				if !yield(fn(item)) {
					return
				}
			}
		},
	}
}

func Map[T, R any](s Seq[T], fn func(T) R) Seq[R] {
	return Seq[R]{
		iter: func(yield func(R) bool) {
			for item := range s.iter {
				if !yield(fn(item)) {
					return
				}
			}
		},
	}
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

func (s Seq[T]) FlatMap(fn func(T) []T) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for item := range s.iter {
				for _, subItem := range fn(item) {
					if !yield(subItem) {
						return
					}
				}
			}
		},
	}
}

func FlatMap[T, R any](s Seq[T], fn func(T) []R) Seq[R] {
	return Seq[R]{
		iter: func(yield func(R) bool) {
			for item := range s.iter {
				for _, subItem := range fn(item) {
					if !yield(subItem) {
						return
					}
				}
			}
		},
	}
}

func (s Seq[T]) Reduce(initial T, fn func(acc, item T) T) T {
	acc := initial
	for item := range s.iter {
		acc = fn(acc, item)
	}
	return acc
}

func Fold[T, R any](s Seq[T], initial R, fn func(R, T) R) R {
	acc := initial
	for item := range s.iter {
		acc = fn(acc, item)
	}
	return acc
}

func (s Seq[T]) Take(n int) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			count := 0
			for item := range s.iter {
				if count >= n {
					return
				}
				if !yield(item) {
					return
				}
				count++
			}
		},
	}
}

func (s Seq[T]) TakeWhile(fn func(T) bool) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for item := range s.iter {
				if !fn(item) {
					return
				}
				if !yield(item) {
					return
				}
			}
		},
	}
}

func (s Seq[T]) Drop(n int) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			count := 0
			for item := range s.iter {
				if count < n {
					count++
					continue
				}
				if !yield(item) {
					return
				}
			}
		},
	}
}

func (s Seq[T]) DropWhile(fn func(T) bool) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			dropping := true
			for item := range s.iter {
				if dropping && fn(item) {
					continue
				}
				dropping = false
				if !yield(item) {
					return
				}
			}
		},
	}
}

func (s Seq[T]) First() optional.Optional[T] {
	for item := range s.iter {
		return optional.Of(item)
	}
	return optional.Empty[T]()
}

func (s Seq[T]) Last() optional.Optional[T] {
	var last T
	found := false
	for item := range s.iter {
		last = item
		found = true
	}
	if !found {
		return optional.Empty[T]()
	}
	return optional.Of(last)
}

func (s Seq[T]) ForEach(fn func(T)) {
	for item := range s.iter {
		fn(item)
	}
}

func (s Seq[T]) Count() int {
	count := 0
	for range s.iter {
		count++
	}
	return count
}

func (s Seq[T]) Any(fn func(T) bool) bool {
	for item := range s.iter {
		if fn(item) {
			return true
		}
	}
	return false
}

func (s Seq[T]) All(fn func(T) bool) bool {
	for item := range s.iter {
		if !fn(item) {
			return false
		}
	}
	return true
}

func (s Seq[T]) None(fn func(T) bool) bool {
	return !s.Any(fn)
}

func (s Seq[T]) Distinct() Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			seen := make(map[any]struct{})
			for item := range s.iter {
				if _, ok := seen[item]; !ok {
					seen[item] = struct{}{}
					if !yield(item) {
						return
					}
				}
			}
		},
	}
}

func Chunked[T any](s Seq[T], size int) Seq[[]T] {
	return Seq[[]T]{
		iter: func(yield func([]T) bool) {
			chunk := make([]T, 0, size)
			for item := range s.iter {
				chunk = append(chunk, item)
				if len(chunk) == size {
					if !yield(chunk) {
						return
					}
					chunk = make([]T, 0, size)
				}
			}
			if len(chunk) > 0 {
				yield(chunk)
			}
		},
	}
}

func Zip[T, U any](s1 Seq[T], s2 Seq[U]) Seq[Pair[T, U]] {
	return Seq[Pair[T, U]]{
		iter: func(yield func(Pair[T, U]) bool) {
			next1, stop1 := iter.Pull(s1.iter)
			defer stop1()
			next2, stop2 := iter.Pull(s2.iter)
			defer stop2()

			for {
				v1, ok1 := next1()
				v2, ok2 := next2()
				if !ok1 || !ok2 {
					return
				}
				if !yield(Pair[T, U]{First: v1, Second: v2}) {
					return
				}
			}
		},
	}
}

func (s Seq[T]) Reversed() Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			items := s.ToSlice()
			for i := len(items) - 1; i >= 0; i-- {
				if !yield(items[i]) {
					return
				}
			}
		},
	}
}

func (s Seq[T]) Sorted(cmpFn func(a, b T) int) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			items := s.ToSlice()
			slices.SortFunc(items, cmpFn)
			for _, item := range items {
				if !yield(item) {
					return
				}
			}
		},
	}
}

func (s Seq[T]) Contains(target T) bool {
	return s.Any(func(item T) bool {
		var targetAny any = target
		var itemAny any = item
		return targetAny == itemAny
	})
}

func (s Seq[T]) IndexOf(target T) int {
	idx := 0
	for item := range s.iter {
		var targetAny any = target
		var itemAny any = item
		if targetAny == itemAny {
			return idx
		}
		idx++
	}
	return -1
}

func (s Seq[T]) Find(fn func(T) bool) optional.Optional[T] {
	for item := range s.iter {
		if fn(item) {
			return optional.Of(item)
		}
	}
	return optional.Empty[T]()
}

func GroupBy[T any, K comparable](s Seq[T], keyFn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for item := range s.iter {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func (s Seq[T]) Partition(fn func(T) bool) (pass, fail []T) {
	for item := range s.iter {
		if fn(item) {
			pass = append(pass, item)
		} else {
			fail = append(fail, item)
		}
	}
	if pass == nil {
		pass = []T{}
	}
	if fail == nil {
		fail = []T{}
	}
	return
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Numeric](s Seq[T]) T {
	var sum T
	for item := range s.iter {
		sum += item
	}
	return sum
}

func Average[T Numeric](s Seq[T]) float64 {
	var sum float64
	count := 0
	for item := range s.iter {
		sum += float64(item)
		count++
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func Max[T cmp.Ordered](s Seq[T]) optional.Optional[T] {
	var max T
	found := false
	for item := range s.iter {
		if !found || item > max {
			max = item
			found = true
		}
	}
	if !found {
		return optional.Empty[T]()
	}
	return optional.Of(max)
}

func Min[T cmp.Ordered](s Seq[T]) optional.Optional[T] {
	var min T
	found := false
	for item := range s.iter {
		if !found || item < min {
			min = item
			found = true
		}
	}
	if !found {
		return optional.Empty[T]()
	}
	return optional.Of(min)
}

func (s Seq[T]) OnEach(fn func(T)) Seq[T] {
	return Seq[T]{
		iter: func(yield func(T) bool) {
			for item := range s.iter {
				fn(item)
				if !yield(item) {
					return
				}
			}
		},
	}
}

func WithIndex[T any](s Seq[T]) Seq[IndexedValue[T]] {
	return Seq[IndexedValue[T]]{
		iter: func(yield func(IndexedValue[T]) bool) {
			idx := 0
			for item := range s.iter {
				if !yield(IndexedValue[T]{Index: idx, Value: item}) {
					return
				}
				idx++
			}
		},
	}
}
