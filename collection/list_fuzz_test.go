package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/list"
)

func FuzzListFilter(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{0, 0, 0})
	f.Add([]byte{255, 128, 64, 32, 16})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		result := l.Filter(func(x int) bool { return x%2 == 0 })

		for _, item := range result {
			if item%2 != 0 {
				t.Errorf("Filter returned odd number: %d", item)
			}
		}

		if result.Len() > l.Len() {
			t.Error("Filter result is larger than original")
		}
	})
}

func FuzzListMap(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{127, 255})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		result := collection.ListMap(l, func(x int) int { return x * 2 })

		if result.Len() != l.Len() {
			t.Errorf("Map changed list length: got %d, want %d", result.Len(), l.Len())
		}

		for i, item := range result {
			expected := ints[i] * 2
			if item != expected {
				t.Errorf("Map returned wrong value at %d: got %d, want %d", i, item, expected)
			}
		}
	})
}

func FuzzListReduce(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{100, 200})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		expectedSum := 0
		for i, b := range data {
			ints[i] = int(b)
			expectedSum += int(b)
		}
		l := list.From(ints)

		result := l.Reduce(0, func(acc, x int) int { return acc + x })

		if result != expectedSum {
			t.Errorf("Reduce sum incorrect: got %d, want %d", result, expectedSum)
		}
	})
}

func FuzzListTakeDrop(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, 3)
	f.Add([]byte{}, 0)
	f.Add([]byte{1}, 10)

	f.Fuzz(func(t *testing.T, data []byte, n int) {
		if n < 0 {
			n = -n
		}
		if n > 1000 {
			n = n % 1000
		}

		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		taken := l.Take(n)
		dropped := l.Drop(n)

		expectedTakeLen := n
		if expectedTakeLen > l.Len() {
			expectedTakeLen = l.Len()
		}
		if taken.Len() != expectedTakeLen {
			t.Errorf("Take(%d) returned wrong length: got %d, want %d", n, taken.Len(), expectedTakeLen)
		}

		expectedDropLen := l.Len() - n
		if expectedDropLen < 0 {
			expectedDropLen = 0
		}
		if dropped.Len() != expectedDropLen {
			t.Errorf("Drop(%d) returned wrong length: got %d, want %d", n, dropped.Len(), expectedDropLen)
		}
	})
}

func FuzzListContains(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, byte(3))
	f.Add([]byte{}, byte(1))
	f.Add([]byte{1, 1, 1}, byte(1))

	f.Fuzz(func(t *testing.T, data []byte, target byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		contains := l.Contains(int(target))

		foundInSlice := false
		for _, b := range data {
			if b == target {
				foundInSlice = true
				break
			}
		}

		if contains != foundInSlice {
			t.Errorf("Contains(%d) returned %v, but should be %v", target, contains, foundInSlice)
		}
	})
}

func FuzzListDistinct(f *testing.F) {
	f.Add([]byte{1, 2, 2, 3, 3, 3})
	f.Add([]byte{})
	f.Add([]byte{1, 1, 1, 1, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		result := l.Distinct()

		if result.Len() > l.Len() {
			t.Error("Distinct result is larger than original")
		}

		seen := make(map[int]bool)
		for _, item := range result {
			if seen[item] {
				t.Errorf("Distinct returned duplicate: %d", item)
			}
			seen[item] = true

			if !l.Contains(item) {
				t.Errorf("Distinct introduced new element: %d", item)
			}
		}

		for _, item := range l {
			if !result.Contains(item) {
				t.Errorf("Distinct lost element: %d", item)
			}
		}
	})
}

func FuzzListReversed(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{1})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		reversed := l.Reversed()

		if reversed.Len() != l.Len() {
			t.Errorf("Reversed changed length: got %d, want %d", reversed.Len(), l.Len())
		}

		for i := range l {
			if reversed[i] != l[l.Len()-1-i] {
				t.Errorf("Reversed element mismatch at %d", i)
			}
		}

		doubleReversed := reversed.Reversed()
		for i := range l {
			if doubleReversed[i] != l[i] {
				t.Errorf("Double reversed should equal original at %d", i)
			}
		}
	})
}

func FuzzListPartition(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{2, 4, 6, 8})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		pass, fail := l.Partition(func(x int) bool { return x%2 == 0 })

		if pass.Len()+fail.Len() != l.Len() {
			t.Error("Partition lost or duplicated elements")
		}

		for _, item := range pass {
			if item%2 != 0 {
				t.Errorf("Partition pass contains odd: %d", item)
			}
		}

		for _, item := range fail {
			if item%2 == 0 {
				t.Errorf("Partition fail contains even: %d", item)
			}
		}
	})
}

func FuzzListChunked(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, 2)
	f.Add([]byte{1, 2, 3}, 5)
	f.Add([]byte{}, 3)

	f.Fuzz(func(t *testing.T, data []byte, size int) {
		if size <= 0 || size > 100 {
			size = 1
		}

		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		l := list.From(ints)

		chunks := l.Chunked(size)

		if len(data) == 0 {
			if len(chunks) != 0 {
				t.Error("Empty list should produce no chunks")
			}
			return
		}

		totalElements := 0
		for i, chunk := range chunks {
			if chunk.Len() == 0 {
				t.Errorf("Chunk %d is empty", i)
			}
			if i < len(chunks)-1 && chunk.Len() != size {
				t.Errorf("Non-last chunk has wrong size: got %d, want %d", chunk.Len(), size)
			}
			totalElements += chunk.Len()
		}

		if totalElements != l.Len() {
			t.Errorf("Chunks total elements mismatch: got %d, want %d", totalElements, l.Len())
		}
	})
}
