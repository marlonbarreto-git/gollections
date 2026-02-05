package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	maps "github.com/marlonbarreto-git/gollections/map"
)

func BenchmarkMapPut(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := make(collection.MutableMap[int, int], size)
				for j := 0; j < size; j++ {
					m[j] = j
				}
			}
		})
	}
}

func BenchmarkMapGet(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}
		target := size - 1

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m[target]
			}
		})
	}
}

func BenchmarkMapFilter(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m.Filter(func(k, v int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkMapKeys(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m.Keys()
			}
		})
	}
}

func BenchmarkMapValues(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m.Values()
			}
		})
	}
}

func BenchmarkMapCopy(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m.Copy()
			}
		})
	}
}

func BenchmarkMapMerge(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		m1 := make(collection.MutableMap[int, int], size)
		m2 := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m1[j] = j
			m2[j+size] = j + size
		}

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				clone := m1.Copy()
				clone.Merge(m2)
			}
		})
	}
}

func BenchmarkMapContainsKey(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}
		target := size - 1

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m.ContainsKey(target)
			}
		})
	}
}

func BenchmarkMapContainsValue(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		m := make(collection.MutableMap[int, int], size)
		for j := 0; j < size; j++ {
			m[j] = j
		}
		target := size - 1

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = m.ContainsValue(target)
			}
		})
	}
}

func BenchmarkMapOf(b *testing.B) {
	b.Run("10pairs", func(b *testing.B) {
		pairs := make([]collection.Pair[int, int], 10)
		for i := range pairs {
			pairs[i] = collection.PairOf(i, i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = maps.Of(pairs...)
		}
	})

	b.Run("100pairs", func(b *testing.B) {
		pairs := make([]collection.Pair[int, int], 100)
		for i := range pairs {
			pairs[i] = collection.PairOf(i, i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = maps.Of(pairs...)
		}
	})
}
