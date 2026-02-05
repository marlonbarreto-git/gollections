package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/set"
)

func BenchmarkSetAdd(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := set.Of[int]()
				for j := 0; j < size; j++ {
					s.Add(j)
				}
			}
		})
	}
}

func BenchmarkSetContains(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		s := set.Of(data...)
		target := size - 1

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s.Contains(target)
			}
		})
	}
}

func BenchmarkSetRemove(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := set.Of(data...)
				for j := 0; j < size/2; j++ {
					s.Remove(j)
				}
			}
		})
	}
}

func BenchmarkSetUnion(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data1 := make([]int, size)
		data2 := make([]int, size)
		for i := range data1 {
			data1[i] = i
			data2[i] = i + size/2
		}
		s1 := set.Of(data1...)
		s2 := set.Of(data2...)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s1.Union(s2)
			}
		})
	}
}

func BenchmarkSetIntersect(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data1 := make([]int, size)
		data2 := make([]int, size)
		for i := range data1 {
			data1[i] = i
			data2[i] = i + size/2
		}
		s1 := set.Of(data1...)
		s2 := set.Of(data2...)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s1.Intersect(s2)
			}
		})
	}
}

func BenchmarkSetSubtract(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data1 := make([]int, size)
		data2 := make([]int, size)
		for i := range data1 {
			data1[i] = i
			data2[i] = i + size/2
		}
		s1 := set.Of(data1...)
		s2 := set.Of(data2...)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s1.Subtract(s2)
			}
		})
	}
}

func BenchmarkSetFilter(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		s := set.Of(data...)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s.Filter(func(x int) bool { return x%2 == 0 })
			}
		})
	}
}

func BenchmarkSetValues(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		s := set.Of(data...)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s.Values()
			}
		})
	}
}

func BenchmarkSetToList(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		s := set.Of(data...)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s.ToList()
			}
		})
	}
}
