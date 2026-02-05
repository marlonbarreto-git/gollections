package sequence_test

import (
	"testing"

	assert "github.com/marlonbarreto-git/gollections/internal/testing"
	"github.com/marlonbarreto-git/gollections/sequence"
)

func TestSequenceOf(t *testing.T) {
	t.Run("creates sequence from variadic args", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3)
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("creates empty sequence", func(t *testing.T) {
		seq := sequence.Of[int]()
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceFrom(t *testing.T) {
	t.Run("creates sequence from slice", func(t *testing.T) {
		seq := sequence.From([]int{1, 2, 3})
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})
}

func TestSequenceMap(t *testing.T) {
	t.Run("maps elements lazily", func(t *testing.T) {
		callCount := 0
		seq := sequence.Of(1, 2, 3).Map(func(x int) int {
			callCount++
			return x * 2
		})

		assert.Equal(t, 0, callCount)

		result := seq.ToSlice()
		assert.Equal(t, []int{2, 4, 6}, result)
		assert.Equal(t, 3, callCount)
	})

	t.Run("maps empty sequence", func(t *testing.T) {
		seq := sequence.Of[int]().Map(func(x int) int { return x * 2 })
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceFilter(t *testing.T) {
	t.Run("filters elements lazily", func(t *testing.T) {
		callCount := 0
		seq := sequence.Of(1, 2, 3, 4, 5).Filter(func(x int) bool {
			callCount++
			return x%2 == 0
		})

		assert.Equal(t, 0, callCount)

		result := seq.ToSlice()
		assert.Equal(t, []int{2, 4}, result)
		assert.Equal(t, 5, callCount)
	})

	t.Run("filters all elements", func(t *testing.T) {
		seq := sequence.Of(1, 3, 5).Filter(func(x int) bool { return x%2 == 0 })
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceFlatMap(t *testing.T) {
	t.Run("flat maps elements", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3).FlatMap(func(x int) []int {
			return []int{x, x * 10}
		})
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 10, 2, 20, 3, 30}, result)
	})

	t.Run("flat maps to empty", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3).FlatMap(func(x int) []int {
			return []int{}
		})
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceReduce(t *testing.T) {
	t.Run("reduces elements", func(t *testing.T) {
		result := sequence.Of(1, 2, 3, 4).Reduce(0, func(acc, x int) int {
			return acc + x
		})
		assert.Equal(t, 10, result)
	})

	t.Run("reduces empty sequence", func(t *testing.T) {
		result := sequence.Of[int]().Reduce(42, func(acc, x int) int {
			return acc + x
		})
		assert.Equal(t, 42, result)
	})
}

func TestSequenceFold(t *testing.T) {
	t.Run("folds elements with different type", func(t *testing.T) {
		result := sequence.Fold(sequence.Of(1, 2, 3), "", func(acc string, x int) string {
			if acc == "" {
				return string(rune('0' + x))
			}
			return acc + "," + string(rune('0'+x))
		})
		assert.Equal(t, "1,2,3", result)
	})
}

func TestSequenceTake(t *testing.T) {
	t.Run("takes n elements", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).Take(3)
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("takes more than available", func(t *testing.T) {
		seq := sequence.Of(1, 2).Take(5)
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2}, result)
	})

	t.Run("takes zero", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3).Take(0)
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceTakeWhile(t *testing.T) {
	t.Run("takes while predicate is true", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 4 })
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("takes none when first fails", func(t *testing.T) {
		seq := sequence.Of(5, 1, 2).TakeWhile(func(x int) bool { return x < 4 })
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceDrop(t *testing.T) {
	t.Run("drops n elements", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).Drop(2)
		result := seq.ToSlice()
		assert.Equal(t, []int{3, 4, 5}, result)
	})

	t.Run("drops more than available", func(t *testing.T) {
		seq := sequence.Of(1, 2).Drop(5)
		result := seq.ToSlice()
		assert.Equal(t, []int{}, result)
	})
}

func TestSequenceDropWhile(t *testing.T) {
	t.Run("drops while predicate is true", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x < 3 })
		result := seq.ToSlice()
		assert.Equal(t, []int{3, 4, 5}, result)
	})
}

func TestSequenceFirst(t *testing.T) {
	t.Run("gets first element", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).First()
		assert.True(t, result.IsPresent())
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("returns empty for empty sequence", func(t *testing.T) {
		result := sequence.Of[int]().First()
		assert.False(t, result.IsPresent())
	})
}

func TestSequenceLast(t *testing.T) {
	t.Run("gets last element", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Last()
		assert.True(t, result.IsPresent())
		assert.Equal(t, 3, result.GetValue())
	})

	t.Run("returns empty for empty sequence", func(t *testing.T) {
		result := sequence.Of[int]().Last()
		assert.False(t, result.IsPresent())
	})
}

func TestSequenceForEach(t *testing.T) {
	t.Run("iterates over all elements", func(t *testing.T) {
		sum := 0
		sequence.Of(1, 2, 3).ForEach(func(x int) {
			sum += x
		})
		assert.Equal(t, 6, sum)
	})
}

func TestSequenceCount(t *testing.T) {
	t.Run("counts all elements", func(t *testing.T) {
		count := sequence.Of(1, 2, 3, 4, 5).Count()
		assert.Equal(t, 5, count)
	})

	t.Run("counts empty sequence", func(t *testing.T) {
		count := sequence.Of[int]().Count()
		assert.Equal(t, 0, count)
	})
}

func TestSequenceAny(t *testing.T) {
	t.Run("returns true when any matches", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Any(func(x int) bool { return x > 2 })
		assert.True(t, result)
	})

	t.Run("returns false when none matches", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Any(func(x int) bool { return x > 5 })
		assert.False(t, result)
	})

	t.Run("short circuits on first match", func(t *testing.T) {
		callCount := 0
		result := sequence.Of(1, 2, 3, 4, 5).Any(func(x int) bool {
			callCount++
			return x == 2
		})
		assert.True(t, result)
		assert.Equal(t, 2, callCount)
	})
}

func TestSequenceAll(t *testing.T) {
	t.Run("returns true when all match", func(t *testing.T) {
		result := sequence.Of(2, 4, 6).All(func(x int) bool { return x%2 == 0 })
		assert.True(t, result)
	})

	t.Run("returns false when one doesn't match", func(t *testing.T) {
		result := sequence.Of(2, 3, 6).All(func(x int) bool { return x%2 == 0 })
		assert.False(t, result)
	})

	t.Run("short circuits on first non-match", func(t *testing.T) {
		callCount := 0
		result := sequence.Of(2, 3, 4, 5, 6).All(func(x int) bool {
			callCount++
			return x%2 == 0
		})
		assert.False(t, result)
		assert.Equal(t, 2, callCount)
	})
}

func TestSequenceNone(t *testing.T) {
	t.Run("returns true when none matches", func(t *testing.T) {
		result := sequence.Of(1, 3, 5).None(func(x int) bool { return x%2 == 0 })
		assert.True(t, result)
	})

	t.Run("returns false when one matches", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).None(func(x int) bool { return x%2 == 0 })
		assert.False(t, result)
	})
}

func TestSequenceDistinct(t *testing.T) {
	t.Run("removes duplicates", func(t *testing.T) {
		seq := sequence.Of(1, 2, 2, 3, 1, 3, 4).Distinct()
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3, 4}, result)
	})
}

func TestSequenceChunked(t *testing.T) {
	t.Run("chunks into groups", func(t *testing.T) {
		seq := sequence.Chunked(sequence.Of(1, 2, 3, 4, 5), 2)
		result := seq.ToSlice()
		assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, result)
	})

	t.Run("chunks with exact fit", func(t *testing.T) {
		seq := sequence.Chunked(sequence.Of(1, 2, 3, 4), 2)
		result := seq.ToSlice()
		assert.Equal(t, [][]int{{1, 2}, {3, 4}}, result)
	})
}

func TestSequenceZip(t *testing.T) {
	t.Run("zips two sequences", func(t *testing.T) {
		seq1 := sequence.Of(1, 2, 3)
		seq2 := sequence.Of("a", "b", "c")
		result := sequence.Zip(seq1, seq2).ToSlice()
		assert.Equal(t, 3, len(result))
		assert.Equal(t, 1, result[0].First)
		assert.Equal(t, "a", result[0].Second)
	})

	t.Run("zips different lengths stops at shorter", func(t *testing.T) {
		seq1 := sequence.Of(1, 2)
		seq2 := sequence.Of("a", "b", "c")
		result := sequence.Zip(seq1, seq2).ToSlice()
		assert.Equal(t, 2, len(result))
	})
}

func TestChaining(t *testing.T) {
	t.Run("chains multiple operations lazily", func(t *testing.T) {
		callCount := 0
		seq := sequence.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
			Filter(func(x int) bool {
				callCount++
				return x%2 == 0
			}).
			Map(func(x int) int {
				callCount++
				return x * 2
			}).
			Take(2)

		assert.Equal(t, 0, callCount)

		result := seq.ToSlice()
		assert.Equal(t, []int{4, 8}, result)
		assert.Less(t, callCount, 20)
	})
}

func TestSequenceReversed(t *testing.T) {
	t.Run("reverses sequence", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3).Reversed()
		result := seq.ToSlice()
		assert.Equal(t, []int{3, 2, 1}, result)
	})
}

func TestSequenceSorted(t *testing.T) {
	t.Run("sorts sequence", func(t *testing.T) {
		seq := sequence.Of(3, 1, 4, 1, 5, 9, 2, 6).Sorted(func(a, b int) int {
			return a - b
		})
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, result)
	})
}

func TestSequenceContains(t *testing.T) {
	t.Run("returns true when contains", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Contains(2)
		assert.True(t, result)
	})

	t.Run("returns false when not contains", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Contains(5)
		assert.False(t, result)
	})
}

func TestSequenceIndexOf(t *testing.T) {
	t.Run("finds index of element", func(t *testing.T) {
		result := sequence.Of(1, 2, 3, 2).IndexOf(2)
		assert.Equal(t, 1, result)
	})

	t.Run("returns -1 when not found", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).IndexOf(5)
		assert.Equal(t, -1, result)
	})
}

func TestSequenceFind(t *testing.T) {
	t.Run("finds element matching predicate", func(t *testing.T) {
		result := sequence.Of(1, 2, 3, 4).Find(func(x int) bool { return x > 2 })
		assert.True(t, result.IsPresent())
		assert.Equal(t, 3, result.GetValue())
	})

	t.Run("returns empty when not found", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Find(func(x int) bool { return x > 5 })
		assert.False(t, result.IsPresent())
	})
}

func TestSequenceGroupBy(t *testing.T) {
	t.Run("groups by key", func(t *testing.T) {
		groups := sequence.GroupBy(sequence.Of(1, 2, 3, 4, 5, 6), func(x int) string {
			if x%2 == 0 {
				return "even"
			}
			return "odd"
		})
		assert.Equal(t, []int{1, 3, 5}, groups["odd"])
		assert.Equal(t, []int{2, 4, 6}, groups["even"])
	})
}

func TestSequencePartition(t *testing.T) {
	t.Run("partitions by predicate", func(t *testing.T) {
		pass, fail := sequence.Of(1, 2, 3, 4, 5).Partition(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, []int{2, 4}, pass)
		assert.Equal(t, []int{1, 3, 5}, fail)
	})
}

func TestSequenceSum(t *testing.T) {
	t.Run("sums numeric sequence", func(t *testing.T) {
		result := sequence.Sum(sequence.Of(1, 2, 3, 4, 5))
		assert.Equal(t, 15, result)
	})
}

func TestSequenceAverage(t *testing.T) {
	t.Run("averages numeric sequence", func(t *testing.T) {
		result := sequence.Average(sequence.Of(1.0, 2.0, 3.0, 4.0, 5.0))
		assert.Equal(t, 3.0, result)
	})
}

func TestSequenceMax(t *testing.T) {
	t.Run("finds max", func(t *testing.T) {
		result := sequence.Max(sequence.Of(3, 1, 4, 1, 5, 9, 2, 6))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 9, result.GetValue())
	})

	t.Run("returns empty for empty sequence", func(t *testing.T) {
		result := sequence.Max(sequence.Of[int]())
		assert.False(t, result.IsPresent())
	})
}

func TestSequenceMin(t *testing.T) {
	t.Run("finds min", func(t *testing.T) {
		result := sequence.Min(sequence.Of(3, 1, 4, 1, 5, 9, 2, 6))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("returns empty for empty sequence", func(t *testing.T) {
		result := sequence.Min(sequence.Of[int]())
		assert.False(t, result.IsPresent())
	})
}

func TestSequenceOnEach(t *testing.T) {
	t.Run("calls function for each without consuming", func(t *testing.T) {
		var collected []int
		seq := sequence.Of(1, 2, 3).OnEach(func(x int) {
			collected = append(collected, x)
		})

		assert.Equal(t, 0, len(collected))

		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
		assert.Equal(t, []int{1, 2, 3}, collected)
	})
}

func TestSequenceWithIndex(t *testing.T) {
	t.Run("pairs elements with indices", func(t *testing.T) {
		result := sequence.WithIndex(sequence.Of("a", "b", "c")).ToSlice()
		assert.Equal(t, 3, len(result))
		assert.Equal(t, 0, result[0].Index)
		assert.Equal(t, "a", result[0].Value)
		assert.Equal(t, 2, result[2].Index)
		assert.Equal(t, "c", result[2].Value)
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.WithIndex(sequence.Of("a", "b", "c"))
		count := 0
		for item := range seq.Iter() {
			count++
			if item.Index >= 1 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestFromIter(t *testing.T) {
	t.Run("creates sequence from iter.Seq", func(t *testing.T) {
		iter := func(yield func(int) bool) {
			for i := 1; i <= 3; i++ {
				if !yield(i) {
					return
				}
			}
		}
		seq := sequence.FromIter(iter)
		result := seq.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})
}

func TestIter(t *testing.T) {
	t.Run("returns iter.Seq", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3)
		iter := seq.Iter()
		var result []int
		for v := range iter {
			result = append(result, v)
		}
		assert.Equal(t, []int{1, 2, 3}, result)
	})
}

func TestMapFreeFunction(t *testing.T) {
	t.Run("maps to different type", func(t *testing.T) {
		result := sequence.Map(sequence.Of(1, 2, 3), func(x int) string {
			return string(rune('a' + x - 1))
		}).ToSlice()
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("early termination", func(t *testing.T) {
		mapped := sequence.Map(sequence.Of(1, 2, 3), func(x int) int { return x * 2 })
		count := 0
		for v := range mapped.Iter() {
			count++
			_ = v
			break
		}
		assert.Equal(t, 1, count)
	})
}

func TestFlatMapFreeFunction(t *testing.T) {
	t.Run("flat maps to different type", func(t *testing.T) {
		result := sequence.FlatMap(sequence.Of(1, 2), func(x int) []string {
			return []string{string(rune('a' + x - 1)), string(rune('A' + x - 1))}
		}).ToSlice()
		assert.Equal(t, []string{"a", "A", "b", "B"}, result)
	})

	t.Run("early termination", func(t *testing.T) {
		flatMapped := sequence.FlatMap(sequence.Of(1, 2, 3), func(x int) []int { return []int{x, x * 10} })
		count := 0
		for v := range flatMapped.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestTakeEdgeCases(t *testing.T) {
	t.Run("take zero", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Take(0).ToSlice()
		assert.Equal(t, []int{}, result)
	})

	t.Run("take negative", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Take(-1).ToSlice()
		assert.Equal(t, []int{}, result)
	})

	t.Run("early termination in take", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).Take(3)
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestTakeWhileEdgeCases(t *testing.T) {
	t.Run("takes none when first fails", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).TakeWhile(func(x int) bool { return x > 10 }).ToSlice()
		assert.Equal(t, []int{}, result)
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 10 })
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestDropEdgeCases(t *testing.T) {
	t.Run("drop zero", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Drop(0).ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("drop negative", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).Drop(-1).ToSlice()
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("early termination in drop", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).Drop(2)
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 1 {
				break
			}
		}
		assert.Equal(t, 1, count)
	})
}

func TestDropWhileEdgeCases(t *testing.T) {
	t.Run("drops all when all match", func(t *testing.T) {
		result := sequence.Of(1, 2, 3).DropWhile(func(x int) bool { return x < 10 }).ToSlice()
		assert.Equal(t, []int{}, result)
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x < 3 })
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 1 {
				break
			}
		}
		assert.Equal(t, 1, count)
	})
}

func TestDistinctEdgeCases(t *testing.T) {
	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Of(1, 2, 2, 3, 3, 3).Distinct()
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestChunkedEdgeCases(t *testing.T) {
	t.Run("chunk size zero returns single chunk with all elements", func(t *testing.T) {
		result := sequence.Chunked(sequence.Of(1, 2, 3), 0).ToSlice()
		assert.Equal(t, 1, len(result))
		assert.Equal(t, []int{1, 2, 3}, result[0])
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Chunked(sequence.Of(1, 2, 3, 4, 5, 6), 2)
		count := 0
		for chunk := range seq.Iter() {
			count++
			_ = chunk
			if count == 1 {
				break
			}
		}
		assert.Equal(t, 1, count)
	})
}

func TestZipEdgeCases(t *testing.T) {
	t.Run("first shorter", func(t *testing.T) {
		result := sequence.Zip(sequence.Of(1, 2), sequence.Of("a", "b", "c")).ToSlice()
		assert.Equal(t, 2, len(result))
	})

	t.Run("second shorter", func(t *testing.T) {
		result := sequence.Zip(sequence.Of(1, 2, 3), sequence.Of("a")).ToSlice()
		assert.Equal(t, 1, len(result))
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Zip(sequence.Of(1, 2, 3), sequence.Of("a", "b", "c"))
		count := 0
		for pair := range seq.Iter() {
			count++
			_ = pair
			if count == 1 {
				break
			}
		}
		assert.Equal(t, 1, count)
	})
}

func TestReversedEdgeCases(t *testing.T) {
	t.Run("reverses empty", func(t *testing.T) {
		result := sequence.Of[int]().Reversed().ToSlice()
		assert.Equal(t, []int{}, result)
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3, 4, 5).Reversed()
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestSortedEdgeCases(t *testing.T) {
	t.Run("sorts empty", func(t *testing.T) {
		result := sequence.Of[int]().Sorted(func(a, b int) int { return a - b }).ToSlice()
		assert.Equal(t, []int{}, result)
	})

	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Of(3, 1, 4, 1, 5).Sorted(func(a, b int) int { return a - b })
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
	})
}

func TestPartitionEdgeCases(t *testing.T) {
	t.Run("all pass", func(t *testing.T) {
		pass, fail := sequence.Of(2, 4, 6).Partition(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, []int{2, 4, 6}, pass)
		assert.Equal(t, []int{}, fail)
	})

	t.Run("all fail", func(t *testing.T) {
		pass, fail := sequence.Of(1, 3, 5).Partition(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, []int{}, pass)
		assert.Equal(t, []int{1, 3, 5}, fail)
	})

	t.Run("empty sequence", func(t *testing.T) {
		pass, fail := sequence.Of[int]().Partition(func(x int) bool { return x%2 == 0 })
		assert.Equal(t, []int{}, pass)
		assert.Equal(t, []int{}, fail)
	})
}

func TestAverageEdgeCases(t *testing.T) {
	t.Run("empty sequence returns 0", func(t *testing.T) {
		result := sequence.Average(sequence.Of[float64]())
		assert.Equal(t, 0.0, result)
	})
}

func TestMinEdgeCases(t *testing.T) {
	t.Run("single element", func(t *testing.T) {
		result := sequence.Min(sequence.Of(42))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 42, result.GetValue())
	})

	t.Run("min not first element", func(t *testing.T) {
		result := sequence.Min(sequence.Of(5, 3, 1, 4, 2))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 1, result.GetValue())
	})

	t.Run("min at end", func(t *testing.T) {
		result := sequence.Min(sequence.Of(5, 4, 3, 2, 1))
		assert.True(t, result.IsPresent())
		assert.Equal(t, 1, result.GetValue())
	})
}

func TestOnEachEdgeCases(t *testing.T) {
	t.Run("early termination", func(t *testing.T) {
		var collected []int
		seq := sequence.Of(1, 2, 3, 4, 5).OnEach(func(x int) {
			collected = append(collected, x)
		})
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 2 {
				break
			}
		}
		assert.Equal(t, 2, count)
		assert.Equal(t, []int{1, 2}, collected)
	})
}

func TestFlatMapMethodEdgeCases(t *testing.T) {
	t.Run("early termination", func(t *testing.T) {
		seq := sequence.Of(1, 2, 3).FlatMap(func(x int) []int { return []int{x, x * 10} })
		count := 0
		for v := range seq.Iter() {
			count++
			_ = v
			if count == 3 {
				break
			}
		}
		assert.Equal(t, 3, count)
	})
}
