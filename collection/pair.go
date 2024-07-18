package collection

type Pair[K comparable, V any] [2]any

func PairOf[K comparable, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{key, value}
}

func (p Pair[K, V]) First() K {
	if first, ok := p[0].(K); ok {
		return first
	}

	panic("invalid type of first element")
}

func (p Pair[K, V]) Second() V {
	if second, ok := p[1].(V); ok {
		return second
	}

	panic("invalid type of second element")
}
