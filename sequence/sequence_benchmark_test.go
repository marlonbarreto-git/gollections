package sequence_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/sequence"
)

func BenchmarkSequenceMap(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Map(func(x int) int { return x * 2 }).ToSlice()
			}
		})
	}
}

func BenchmarkSequenceFilter(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Filter(func(x int) bool { return x%2 == 0 }).ToSlice()
			}
		})
	}
}

func BenchmarkSequenceReduce(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Reduce(0, func(acc, x int) int { return acc + x })
			}
		})
	}
}

func BenchmarkSequenceChain(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Filter(func(x int) bool { return x%2 == 0 }).
					Map(func(x int) int { return x * 2 }).
					Take(10).
					ToSlice()
			}
		})
	}
}

func BenchmarkSequenceLazyVsEager(b *testing.B) {
	size := 100000

	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	seq := sequence.From(data)

	b.Run("lazy_take_10", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = seq.Filter(func(x int) bool { return x%2 == 0 }).
				Map(func(x int) int { return x * 2 }).
				Take(10).
				ToSlice()
		}
	})

	b.Run("lazy_take_all", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = seq.Filter(func(x int) bool { return x%2 == 0 }).
				Map(func(x int) int { return x * 2 }).
				ToSlice()
		}
	})
}

func BenchmarkSequenceTake(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Take(10).ToSlice()
			}
		})
	}
}

func BenchmarkSequenceDrop(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Drop(10).ToSlice()
			}
		})
	}
}

func BenchmarkSequenceDistinct(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i % (size / 10)
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Distinct().ToSlice()
			}
		})
	}
}

func BenchmarkSequenceContains(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)
		target := size - 1

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Contains(target)
			}
		})
	}
}

func BenchmarkSequenceCount(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}
		seq := sequence.From(data)

		b.Run(intToString(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = seq.Count()
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
