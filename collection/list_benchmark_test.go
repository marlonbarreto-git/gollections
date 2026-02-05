package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/list"
)

func BenchmarkListMap(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = collection.ListMap(l, func(x int) int { return x * 2 })
			}
		})
	}
}

func BenchmarkListFilter(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Filter(func(x int) bool { return x%2 == 0 })
			}
		})
	}
}

func BenchmarkListReduce(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Reduce(0, func(acc, x int) int { return acc + x })
			}
		})
	}
}

func BenchmarkListContains(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)
		target := size - 1

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Contains(target)
			}
		})
	}
}

func BenchmarkListSorted(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = size - i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Sorted(func(a, b int) int { return a - b })
			}
		})
	}
}

func BenchmarkListReversed(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Reversed()
			}
		})
	}
}

func BenchmarkListChain(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Filter(func(x int) bool { return x%2 == 0 }).Take(10)
			}
		})
	}
}

func BenchmarkListDistinct(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i % (size / 10)
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Distinct()
			}
		})
	}
}

func BenchmarkListGroupBy(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = collection.GroupBy(l, func(x int) int { return x % 10 })
			}
		})
	}
}

func BenchmarkListFlatMap(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = collection.FlatMap(l, func(x int) []int {
					return []int{x, x + 1}
				})
			}
		})
	}
}

func BenchmarkListFold(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = collection.Fold(l, 0, func(acc, x int) int { return acc + x })
			}
		})
	}
}

func BenchmarkListTake(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Take(10)
			}
		})
	}
}

func BenchmarkListDrop(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = l.Drop(10)
			}
		})
	}
}

func BenchmarkListPartition(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		l := list.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = l.Partition(func(x int) bool { return x%2 == 0 })
			}
		})
	}
}

func intToString(n int) string {
	switch n {
	case 100:
		return "100"
	case 1000:
		return "1000"
	case 10000:
		return "10000"
	case 100000:
		return "100000"
	default:
		return "unknown"
	}
}
