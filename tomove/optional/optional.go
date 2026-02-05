package optional

import (
	"errors"
	"reflect"
)

type (
	// Optional represents a value that may or may not be present
	Optional[T any] interface {

		// IsPresent returns true if the value is present, otherwise it returns false
		IsPresent() bool

		// IsEmpty returns true if the value is not present, otherwise it returns false
		IsEmpty() bool

		// IfPresent calls the given consumer with the value if present
		// Example:
		//
		// 	optional := optional.Of(10)
		// 	optional.IfPresent(func(value int) {
		// 	    fmt.Println(value)
		// 	})
		//
		// Output: 10
		IfPresent(consumer func(T)) *optional[T]

		// Get returns the value if present, otherwise it calls the supplier and returns the value
		// If the supplier panics, it returns the panic message as an error
		//
		// Example:
		//
		// 	optional := optional.OfGet(func() int {
		// 	    return 10
		// 	})
		// 	value, err := optional.Get()
		//
		// Output: 10, nil
		Get() (T, error)

		// GetValue returns the value if present, otherwise it calls the supplier and returns the value
		GetValue() T

		// OrElse returns the value if present, otherwise it returns the alternative
		// Example:
		//
		// 	optional := optional.Empty[int]()
		// 	optional.OrElse(10)
		//
		// Output: 10
		OrElse(T) T

		// OrElseGet returns the value if present, otherwise it calls the supplier and returns the value
		// Example:
		//
		// 	optional := optional.Empty[int]()
		// 	optional.OrElseGet(func() int {
		// 	    return 10
		// 	})
		//
		// Output: 10
		OrElseGet(supplier Supplier[T]) (result T)

		// OrElsePanic returns the value if present, otherwise it panics with the given message
		// Example:
		//
		// 	optional := optional.Empty[int]()
		// 	optional.OrElsePanic("no value present")
		//
		// Output: panic: no value present
		OrElsePanic(panicMsg any) T
	}

	// TakingArg represents an optional that takes arguments
	TakingArg[T any] interface {
		TakingArg(ArgumentNum) Optional[T]
	}

	optional[T any] struct {
		isEmpty       bool
		value         T
		supplierValue Supplier[T]
		alternative   T
	}

	optionalNumValue[T any] struct {
		optional[T]
		arguments []any
	}

	Supplier[T any] func() T

	ArgumentNum int
)

const (
	First ArgumentNum = iota
	Second
	Third
	Fourth
	Fifth
	Sixth
	Seventh
	Eighth
	Ninth
	Tenth
	Last ArgumentNum = -1
)

var NoValuePresentError = errors.New("no value present")

// Of creates a new Optional with the given value
func Of[T any](value T) Optional[T] {
	return &optional[T]{value: value}
}

// OfValues creates a new Optional with the given values
func OfValues[T any](values ...any) TakingArg[T] {
	return &optionalNumValue[T]{arguments: values}
}

// OfGet creates a new Optional with the given supplier
// The supplier is called only when the value is requested
func OfGet[T any](supplier func() T) Optional[T] {
	return &optional[T]{supplierValue: supplier}
}

// Empty creates a new empty Optional
func Empty[T any]() Optional[T] {
	return &optional[T]{isEmpty: true}
}

func (optional *optionalNumValue[T]) TakingArg(numOfArgument ArgumentNum) Optional[T] {
	if numOfArgument == Last || len(optional.arguments) < int(numOfArgument) {
		numOfArgument = ArgumentNum(len(optional.arguments) - 1)
	}

	value := optional.arguments[numOfArgument]

	optional.value, _ = value.(T)
	return optional
}

func (optional *optional[T]) IsPresent() bool {
	return !optional.isEmpty && !isEmpty(optional.value)
}

func (optional *optional[T]) IsEmpty() bool {
	return optional.isEmpty || isEmpty(optional.value)
}

func (optional *optional[T]) IfPresent(consumer func(T)) *optional[T] {
	if optional.IsPresent() {
		consumer(optional.value)
	}

	return optional
}

func (optional *optional[T]) GetValue() T {
	if optional.supplierValue != nil {
		optional.value = optional.supplierValue()
	}

	return optional.value
}

func (optional *optional[T]) Get() (result T, err error) {
	if optional.supplierValue != nil {
		defer optional.recoverGetPanicAndSetResults(&result, &err)
		optional.value = optional.supplierValue()
	}

	if optional.IsPresent() {
		return optional.value, nil
	}

	return optional.value, NoValuePresentError
}

func (optional *optional[T]) OrElse(alternative T) (result T) {
	optional.alternative = alternative

	if optional.supplierValue != nil {
		defer optional.recoverOrElsePanicAndSetResult(&result)
		optional.value = optional.supplierValue()
	}

	if optional.IsPresent() {
		return optional.value
	}

	return alternative
}

func (optional *optional[T]) OrElseGet(supplier Supplier[T]) (result T) {
	if optional.supplierValue != nil {
		defer optional.recoverOrElseGetPanicAndSetResult(supplier, &result)
		optional.value = optional.supplierValue()
	}

	if optional.IsPresent() {
		return optional.value
	}

	return supplier()
}

func (optional *optional[T]) OrElsePanic(panicMsg any) T {
	if optional.supplierValue != nil {
		optional.value = optional.supplierValue()
	}

	if !optional.IsPresent() {
		panic(panicMsg)
	}

	return optional.value
}

func (optional *optional[T]) recoverOrElsePanicAndSetResult(result *T) {
	if recoverData := recover(); recoverData != nil {
		*result = optional.alternative
	}
}

func (optional *optional[T]) recoverOrElseGetPanicAndSetResult(supplier Supplier[T], result *T) {
	if recoverData := recover(); recoverData != nil {
		optional.alternative = supplier()
		*result = optional.alternative
	}
}

func (optional *optional[T]) recoverGetPanicAndSetResults(value *T, err *error) {
	if recoverData := recover(); recoverData != nil {
		*err = NoValuePresentError
		*value = optional.value
	}
}

func isEmpty(item any) (empty bool) {
	if item == nil {
		return true
	}

	val := reflect.ValueOf(item)

	switch val.Kind() {
	case reflect.Ptr:
		return val.IsNil()
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return val.Len() == 0
	case reflect.Struct:
		return isEmptyStruct(item)
	default:
		return
	}
}

func isEmptyStruct(item any) (empty bool) {
	val := reflect.ValueOf(item)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			return false
		}
	}

	return true
}
