package function

type (
	// Function is a function that receives a T and returns a U
	Function[T, U any] func(T) U

	// Predicate is a function that receives a T and returns a bool
	Predicate[T any] func(T) bool

	// BiPredicate is a function that receives a T and K and returns a bool
	BiPredicate[T, V any] func(T, V) bool

	// Consumer is a function that receives a T and returns nothing
	Consumer[T any] func(T)

	// BiConsumer is a function that receives a T and V and returns nothing
	BiConsumer[T, V any] func(T, V)

	// IndexedConsumer is a function that receives a T and an int and returns nothing
	IndexedConsumer[T any] func(int, T)

	// Supplier is a function that returns a T
	Supplier[T any] func() T

	// Runnable is a function that returns nothing
	Runnable func()

	// UnaryOperator is a function that receives a T and returns a T
	UnaryOperator[T any] func(T) T

	// Comparator is a function that receives two T and returns an int
	Comparator[T any] func(T, T) int
)
