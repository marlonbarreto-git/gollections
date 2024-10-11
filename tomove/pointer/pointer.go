package pointer

// Of returns a pointer to a given item
func Of[T any](item T) *T {
	return &item
}
