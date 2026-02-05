package collection

import "github.com/marlonbarreto-git/gollections/tomove/optional"

// Pipeline provides a chainable wrapper for any value, enabling
// Kotlin-style let/also/takeIf/takeUnless chaining with zero allocation overhead.
// The Pipeline type is a simple wrapper that gets optimized away by the compiler.
type Pipeline[T any] struct {
	value T
}

// Pipe creates a new Pipeline wrapping the given value.
// This is the entry point for chainable operations.
//
// Example:
//
//	result := Pipe(list.Of(1, 2, 3)).
//	    Let(func(l List[int]) List[int] { return l.Filter(...) }).
//	    Let(func(l List[int]) List[int] { return l.Map(...) }).
//	    Value()
func Pipe[T any](value T) Pipeline[T] {
	return Pipeline[T]{value: value}
}

// Value returns the wrapped value, ending the chain.
func (p Pipeline[T]) Value() T {
	return p.value
}

// Let applies a transformation function and returns a new Pipeline with the result.
// This enables chaining: .Let(...).Let(...).Let(...)
func (p Pipeline[T]) Let(fn func(T) T) Pipeline[T] {
	return Pipeline[T]{value: fn(p.value)}
}

// Also executes a side-effect function and returns the same Pipeline unchanged.
// Useful for logging, debugging, or other side effects in the middle of a chain.
func (p Pipeline[T]) Also(fn func(T)) Pipeline[T] {
	fn(p.value)
	return p
}

// TakeIf returns an Optional containing the value if the predicate returns true,
// or an empty Optional otherwise. This ends the chain.
func (p Pipeline[T]) TakeIf(predicate func(T) bool) optional.Optional[T] {
	if predicate(p.value) {
		return optional.Of(p.value)
	}
	return optional.Empty[T]()
}

// TakeUnless returns an Optional containing the value if the predicate returns false,
// or an empty Optional otherwise. This ends the chain.
func (p Pipeline[T]) TakeUnless(predicate func(T) bool) optional.Optional[T] {
	if !predicate(p.value) {
		return optional.Of(p.value)
	}
	return optional.Empty[T]()
}

// PipeTransform transforms a value to a different type.
// Use this when you need to change types in a chain.
//
// Example:
//
//	count := PipeTransform(
//	    Pipe(list.Of(1, 2, 3)).Let(...).Value(),
//	    func(l List[int]) int { return l.Len() },
//	)
func PipeTransform[T, R any](value T, fn func(T) R) R {
	return fn(value)
}

// PipeMap creates a Pipeline that transforms to a different type.
// This allows continuing the chain with a new type.
func PipeMap[T, R any](p Pipeline[T], fn func(T) R) Pipeline[R] {
	return Pipeline[R]{value: fn(p.value)}
}
