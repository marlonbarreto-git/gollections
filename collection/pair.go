package collection

type Pair[K comparable, V any] struct {
	first  K
	second V
}

func PairOf[K comparable, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{first: key, second: value}
}

func (p Pair[K, V]) First() K {
	return p.first
}

func (p Pair[K, V]) Second() V {
	return p.second
}
