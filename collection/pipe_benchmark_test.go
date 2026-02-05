package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/list"
)

func BenchmarkPipeChain(b *testing.B) {
	data := make(collection.List[int], 1000)
	for i := range data {
		data[i] = i
	}

	b.Run("pipe_6_lets", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = collection.Pipe(data).
				Let(func(l collection.List[int]) collection.List[int] { return l.Filter(func(x int) bool { return x%2 == 0 }) }).
				Let(func(l collection.List[int]) collection.List[int] { return l.Take(100) }).
				Let(func(l collection.List[int]) collection.List[int] { return collection.ListMap(l, func(x int) int { return x * 2 }) }).
				Let(func(l collection.List[int]) collection.List[int] { return l.Filter(func(x int) bool { return x > 50 }) }).
				Let(func(l collection.List[int]) collection.List[int] { return l.Take(10) }).
				Let(func(l collection.List[int]) collection.List[int] { return l.Reversed() }).
				Value()
		}
	})

	b.Run("direct_calls_equivalent", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			step1 := data.Filter(func(x int) bool { return x%2 == 0 })
			step2 := step1.Take(100)
			step3 := collection.ListMap(step2, func(x int) int { return x * 2 })
			step4 := step3.Filter(func(x int) bool { return x > 50 })
			step5 := step4.Take(10)
			_ = step5.Reversed()
		}
	})
}

func BenchmarkPipeVsSequence(b *testing.B) {
	data := make(collection.List[int], 100000)
	for i := range data {
		data[i] = i
	}

	b.Run("pipe_eager", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = collection.Pipe(data).
				Let(func(l collection.List[int]) collection.List[int] {
					return l.Filter(func(x int) bool { return x%2 == 0 })
				}).
				Let(func(l collection.List[int]) collection.List[int] {
					return l.Take(10)
				}).
				Value()
		}
	})

	b.Run("direct_filter_take", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = data.Filter(func(x int) bool { return x%2 == 0 }).Take(10)
		}
	})
}

func BenchmarkPipeAlso(b *testing.B) {
	data := list.Of(1, 2, 3, 4, 5)

	b.Run("with_also", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var count int
			_ = collection.Pipe(data).
				Also(func(l collection.List[int]) { count = l.Len() }).
				Let(func(l collection.List[int]) collection.List[int] { return l.Append(6) }).
				Also(func(l collection.List[int]) { count += l.Len() }).
				Value()
			_ = count
		}
	})
}
