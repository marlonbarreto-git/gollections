package sequence_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/sequence"
)

func FuzzSequenceMap(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{127, 255})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)

		result := seq.Map(func(x int) int { return x * 2 }).ToSlice()

		if len(result) != len(ints) {
			t.Errorf("Map changed length: got %d, want %d", len(result), len(ints))
		}

		for i, item := range result {
			expected := ints[i] * 2
			if item != expected {
				t.Errorf("Map returned wrong value at %d: got %d, want %d", i, item, expected)
			}
		}
	})
}

func FuzzSequenceFilter(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{2, 4, 6, 8})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)

		result := seq.Filter(func(x int) bool { return x%2 == 0 }).ToSlice()

		for _, item := range result {
			if item%2 != 0 {
				t.Errorf("Filter returned odd number: %d", item)
			}
		}

		if len(result) > len(ints) {
			t.Error("Filter result is larger than original")
		}
	})
}

func FuzzSequenceReduce(f *testing.F) {
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
		seq := sequence.From(ints)

		result := seq.Reduce(0, func(acc, x int) int { return acc + x })

		if result != expectedSum {
			t.Errorf("Reduce sum incorrect: got %d, want %d", result, expectedSum)
		}
	})
}

func FuzzSequenceTake(f *testing.F) {
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
		seq := sequence.From(ints)

		result := seq.Take(n).ToSlice()

		expectedLen := n
		if expectedLen > len(ints) {
			expectedLen = len(ints)
		}
		if len(result) != expectedLen {
			t.Errorf("Take(%d) returned wrong length: got %d, want %d", n, len(result), expectedLen)
		}

		for i, item := range result {
			if item != ints[i] {
				t.Errorf("Take returned wrong element at %d: got %d, want %d", i, item, ints[i])
			}
		}
	})
}

func FuzzSequenceDrop(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, 2)
	f.Add([]byte{}, 0)
	f.Add([]byte{1, 2}, 10)

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
		seq := sequence.From(ints)

		result := seq.Drop(n).ToSlice()

		expectedLen := len(ints) - n
		if expectedLen < 0 {
			expectedLen = 0
		}
		if len(result) != expectedLen {
			t.Errorf("Drop(%d) returned wrong length: got %d, want %d", n, len(result), expectedLen)
		}

		for i, item := range result {
			if item != ints[n+i] {
				t.Errorf("Drop returned wrong element at %d: got %d, want %d", i, item, ints[n+i])
			}
		}
	})
}

func FuzzSequenceDistinct(f *testing.F) {
	f.Add([]byte{1, 2, 2, 3, 3, 3})
	f.Add([]byte{})
	f.Add([]byte{1, 1, 1, 1, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)

		result := seq.Distinct().ToSlice()

		if len(result) > len(ints) {
			t.Error("Distinct result is larger than original")
		}

		seen := make(map[int]bool)
		for _, item := range result {
			if seen[item] {
				t.Errorf("Distinct returned duplicate: %d", item)
			}
			seen[item] = true
		}
	})
}

func FuzzSequenceChain(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.Add([]byte{})
	f.Add([]byte{2, 4, 6, 8})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)

		result := seq.
			Filter(func(x int) bool { return x%2 == 0 }).
			Map(func(x int) int { return x * 2 }).
			Take(5).
			ToSlice()

		for _, item := range result {
			if item%4 != 0 {
				t.Errorf("Chain produced invalid element: %d (should be divisible by 4)", item)
			}
		}

		if len(result) > 5 {
			t.Errorf("Chain Take(5) returned more than 5 elements: %d", len(result))
		}
	})
}

func FuzzSequenceContains(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, byte(3))
	f.Add([]byte{}, byte(1))
	f.Add([]byte{1, 1, 1}, byte(1))

	f.Fuzz(func(t *testing.T, data []byte, target byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)

		contains := seq.Contains(int(target))

		expected := false
		for _, b := range data {
			if b == target {
				expected = true
				break
			}
		}

		if contains != expected {
			t.Errorf("Contains(%d) = %v, want %v", target, contains, expected)
		}
	})
}

func FuzzSequenceCount(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{1, 1, 1, 1, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)

		count := seq.Count()

		if count != len(data) {
			t.Errorf("Count() = %d, want %d", count, len(data))
		}
	})
}

func FuzzSequenceAnyAllNone(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{2, 4, 6, 8})
	f.Add([]byte{1, 3, 5, 7})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		ints := make([]int, len(data))
		for i, b := range data {
			ints[i] = int(b)
		}
		seq := sequence.From(ints)
		predicate := func(x int) bool { return x%2 == 0 }

		anyResult := seq.Any(predicate)
		allResult := seq.All(predicate)
		noneResult := seq.None(predicate)

		hasEven := false
		allEven := len(data) > 0
		for _, b := range data {
			if int(b)%2 == 0 {
				hasEven = true
			} else {
				allEven = false
			}
		}
		if len(data) == 0 {
			allEven = true
		}

		if anyResult != hasEven {
			t.Errorf("Any() = %v, want %v", anyResult, hasEven)
		}
		if allResult != allEven {
			t.Errorf("All() = %v, want %v", allResult, allEven)
		}
		if noneResult != !hasEven {
			t.Errorf("None() = %v, want %v", noneResult, !hasEven)
		}
	})
}
