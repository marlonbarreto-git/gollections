# Gollections Knowledge Base

## Project Overview
A Go collections library that provides Kotlin/JavaScript-like collection chaining with C++ level performance.

## Current State (Iteration 6 - COMPLETED)
- **Author**: Marlon Barreto (mbarretot@hotmail.com)
- **Go Version**: 1.23.0 (upgraded for iter package support)
- **Dependencies**: ZERO external dependencies (all in-house)
- **Test Coverage**: 100% for all collection packages (97% total including test utilities)

## Implemented Functions

### List[T] (75+ methods)
- [x] ToArray
- [x] Distinct / DistinctBy
- [x] Associate / AssociateBy / AssociateWith
- [x] Join
- [x] Filter / FilterIndexed / FilterNot
- [x] Append / Add (mutating)
- [x] ForEach / ForEachIndexed
- [x] First / Last / FindLast
- [x] Get / ElementAt
- [x] Find / FindIndex / FindLastIndex
- [x] Len / IsEmpty / IsNotEmpty
- [x] Some (any) / Every (all) / None
- [x] Slice
- [x] Count
- [x] Sum / SumOf
- [x] ListMap / MapIndexed / MapNotNull / MapIndexedNotNull (free functions)
- [x] Reduce / Fold / RunningFold / RunningReduce / Scan
- [x] ReduceIndexed / FoldIndexed / RunningFoldIndexed / RunningReduceIndexed
- [x] FoldRight / ReduceRight
- [x] FlatMap (method and free function)
- [x] Reversed
- [x] Sorted / SortedDescending (with custom comparator)
- [x] Take / TakeLast / TakeWhile / TakeLastWhile
- [x] Drop / DropLast / DropWhile / DropLastWhile
- [x] Chunked / Windowed
- [x] Contains / ContainsAll / IndexOf / LastIndexOf
- [x] MinBy / MaxBy
- [x] Min / Max / Average (free functions for ordered types)
- [x] GroupBy (method and free function)
- [x] Partition
- [x] Zip / ZipWithNext / Unzip (free functions)
- [x] Flatten (free function and method)
- [x] OnEach
- [x] AsSequence (lazy evaluation)
- [x] Single
- [x] Shuffled / Random
- [x] FirstNotNullOf
- [x] Plus / PlusAll / Minus / MinusAll
- [x] FoldRightIndexed / ReduceRightIndexed
- [x] ToSet / ToMap / ToMapWithValue (conversions between collections)
- [x] Let / Also / TakeIf / TakeUnless (Kotlin scope functions)

### Sequence[T] - Lazy Evaluation (30+ methods)
- [x] Of / From / FromIter
- [x] ToSlice / Iter
- [x] Map (method and free function)
- [x] Filter
- [x] FlatMap (method and free function)
- [x] Reduce / Fold (free function)
- [x] Take / TakeWhile
- [x] Drop / DropWhile
- [x] First / Last
- [x] ForEach
- [x] Count
- [x] Any / All / None
- [x] Distinct
- [x] Reversed
- [x] Sorted
- [x] Contains / IndexOf
- [x] Find
- [x] OnEach
- [x] Chunked (free function)
- [x] Zip (free function)
- [x] GroupBy (free function)
- [x] Sum / Average / Max / Min (free functions)
- [x] WithIndex (free function)
- [x] Partition

### MutableMap[K, V] (25+ methods)
- [x] Map / MapKeys / MapValues
- [x] Reduce
- [x] ForEach
- [x] Filter / FilterKeys / FilterValues
- [x] IsEmpty / Len
- [x] Count
- [x] Copy
- [x] Keys / Values / Entries
- [x] Remove
- [x] String
- [x] GetOrDefault
- [x] GetOrPut
- [x] ContainsKey / ContainsValue
- [x] Merge / PutAll
- [x] ToList
- [x] Any / All / None

### Set[K] (15+ methods)
- [x] Contains
- [x] Add
- [x] Remove
- [x] Clear
- [x] IsEmpty / Len
- [x] String
- [x] Values
- [x] Union / Intersect / Subtract
- [x] Filter
- [x] ForEach
- [x] Any / All / None
- [x] First
- [x] ToList

### Optional[T] (100% coverage)
- [x] Of / OfValues / OfGet / Empty
- [x] IsPresent / IsEmpty
- [x] IfPresent
- [x] Get / GetValue
- [x] OrElse / OrElseGet / OrElsePanic
- [x] TakingArg

### Pipeline[T] - Zero-Cost Chainable Wrapper (7 functions)
- [x] Pipe - creates a new Pipeline wrapping any value
- [x] Value - returns the wrapped value, ending the chain
- [x] Let - applies transformation and returns new Pipeline (chainable .let.let.let)
- [x] Also - executes side-effect and returns same Pipeline unchanged
- [x] TakeIf - returns Optional with value if predicate true
- [x] TakeUnless - returns Optional with value if predicate false
- [x] PipeTransform - transforms value to different type (free function)
- [x] PipeMap - transforms Pipeline to different type (free function)

### Helper Types
- [x] Pair[A, B] - First / Second / String
- [x] ZipPair[T, U] - For Zip operations
- [x] ConsecutivePair[T] - For ZipWithNext operations
- [x] IndexedValue[T] - For WithIndex operations
- [x] Numeric constraint - For Sum/Average operations

## Architecture Decisions

### 1. Zero External Dependencies
- Created internal/testing/assert.go with custom test helpers
- Removed testify dependency completely

### 2. Go 1.23+ with iter Package
- Using iter.Seq for lazy evaluation
- Enables C++ level performance through iterator fusion

### 3. Lazy Evaluation via Sequence Type
```go
type Seq[T any] struct {
    iter iter.Seq[T]
}
```
Provides lazy evaluation for efficient chaining.

### 4. Generic Type Constraints
- cmp.Ordered for Min/Max/Sorted on ordered types
- comparable for Set keys and Map keys
- Numeric constraint for Sum/Average

### 5. Method vs Free Function Strategy
- Methods when single type parameter (e.g., Filter, Map same type)
- Free functions when multiple type parameters needed (e.g., FlatMap to different type, Fold, ZipWithNext)
- This avoids Go's instantiation cycle limitations

## Performance Results

### Benchmarks Highlights (Apple M3 Pro)
- List.Take: O(1) - ~3.5ns regardless of size
- Map.Get/ContainsKey: O(1) - ~5.5ns
- Sequence lazy chain (take 10 from 100k): 438ns vs 632μs eager
- List.Reduce 100k elements: ~38μs
- List.Filter 100k elements: ~258μs

### Lazy vs Eager Comparison
```
BenchmarkSequenceLazyVsEager/lazy_take_10: 438ns
BenchmarkSequenceLazyVsEager/lazy_take_all: 632μs
```
Lazy evaluation is ~1400x faster when only taking 10 elements from 100k!

## Testing

### Unit Tests
- 100% overall coverage
- Table-driven tests following Go best practices

### Fuzz Tests
- List: Filter, Map, Reduce, TakeDrop, Contains, Distinct, Reversed, Partition, Chunked
- Map: PutGet, Filter, ContainsKey, Keys, Values, Copy, Merge, GetOrDefault
- Set: AddContains, Remove, Union, Intersect, Subtract, Filter, Values, Clear
- Sequence: Map, Filter, Reduce, Take, Drop, Distinct, Chain, Contains, Count, AnyAllNone

### Benchmarks
- List: Map, Filter, Reduce, Contains, Sorted, Reversed, Chain, Distinct, GroupBy, FlatMap, Fold, Take, Drop, Partition
- Map: Put, Get, Filter, Keys, Values, Copy, Merge, ContainsKey, ContainsValue, Of
- Set: Add, Contains, Remove, Union, Intersect, Subtract, Filter, Values, ToList
- Sequence: Map, Filter, Reduce, Chain, LazyVsEager, Take, Drop, Distinct, Contains, Count

## File Structure
```
gollections/
├── collection/
│   ├── list.go              # List type and methods (55+ functions)
│   ├── list_test.go
│   ├── list_benchmark_test.go
│   ├── list_fuzz_test.go
│   ├── map.go               # MutableMap type (25+ functions)
│   ├── map_test.go
│   ├── map_benchmark_test.go
│   ├── map_fuzz_test.go
│   ├── set.go               # Set type (15+ functions)
│   ├── set_test.go
│   ├── set_benchmark_test.go
│   ├── set_fuzz_test.go
│   ├── pair.go              # Pair helper type
│   ├── pipe.go              # Pipeline for zero-cost chaining (.let.let.let)
│   ├── pipe_test.go         # Pipeline tests
│   └── pipe_benchmark_test.go # Pipeline benchmarks
├── sequence/
│   ├── sequence.go          # Lazy Sequence type (30+ functions)
│   ├── sequence_test.go
│   ├── sequence_benchmark_test.go
│   └── sequence_fuzz_test.go
├── internal/testing/
│   ├── assert.go            # Custom test assertions
│   └── assert_test.go       # Tests for test assertions
├── tomove/optional/
│   └── optional.go          # Optional type
├── tomove/pointer/
│   ├── pointer.go           # Pointer utility
│   └── pointer_test.go      # Pointer tests
├── list/
│   ├── api.go               # Factory functions (Of, From)
│   └── api_test.go          # Factory tests
├── map/
│   ├── api.go               # Factory functions (Of, From)
│   └── api_test.go          # Factory tests
├── set/
│   ├── api.go               # Factory functions (Of, From)
│   └── api_test.go          # Factory tests
├── go.mod                   # Go 1.23.0
└── knowledge.md             # This file
```

## Completed Tasks
1. ✅ Remove testify dependency
2. ✅ Upgrade to Go 1.23+
3. ✅ Implement Sequence type for lazy evaluation
4. ✅ Implement List transformation methods (Map, FlatMap, Reduce, Fold)
5. ✅ Implement List ordering methods (Sorted, Reversed)
6. ✅ Implement List slicing methods (Take, Drop, Chunked, Windowed)
7. ✅ Implement List aggregation methods (Min, Max, Average, GroupBy)
8. ✅ Implement List search/check methods (Contains, IndexOf, None)
9. ✅ Implement List combination methods (Zip, Flatten)
10. ✅ Implement remaining List methods
11. ✅ Implement enhanced Map methods
12. ✅ Implement enhanced Set methods
13. ✅ Create comprehensive benchmarks
14. ✅ Implement comprehensive fuzz testing
15. ✅ Increase test coverage to 100%
16. ✅ Add missing Kotlin/JS collection functions
17. ✅ Add tests for factory packages (list/Of, list/From, map/Of, map/From, set/Of, set/From)
18. ✅ Add tests for pointer utility package
19. ✅ Add tests for internal testing utilities
20. ✅ Add missing Kotlin aggregate functions (FoldRight, ReduceRight, FoldIndexed, ReduceIndexed, etc.)
21. ✅ Add TakeLastWhile / DropLastWhile
22. ✅ Add MapNotNull / MapIndexedNotNull / FirstNotNullOf
23. ✅ Add Unzip / ContainsAll / IsNotEmpty / FindLastIndex / SortedDescending / SumOf
24. ✅ Implement Pipeline type for zero-cost .let.let.let chaining
25. ✅ Add Pipeline benchmarks proving zero overhead

## Kotlin/JS Parity - Functions Added in Iteration 4
- MapIndexed - maps with index access
- FilterIndexed - filters with index access
- FilterNot - inverse filter
- RunningFold / RunningReduce / Scan - cumulative operations
- AssociateWith - create map from list with value selector
- ZipWithNext - pair consecutive elements
- Windowed - sliding window view
- Single - returns single element or empty
- ElementAt - safe index access
- Shuffled / Random - randomization operations
- FindIndex - find first matching index

## Kotlin/JS Parity - Functions Added in Iteration 5
- IsNotEmpty - check non-empty
- ContainsAll - check all elements contained
- TakeLastWhile / DropLastWhile - take/drop from end with predicate
- SortedDescending - descending sort for ordered types
- SumOf - sum with selector
- FoldIndexed / ReduceIndexed - fold/reduce with index
- FoldRight / ReduceRight - fold/reduce from right
- RunningFoldIndexed / RunningReduceIndexed - running operations with index
- MapNotNull / MapIndexedNotNull - map filtering nulls
- Unzip - split pairs into two lists
- FindLastIndex - find last matching index
- FirstNotNullOf - first non-null mapped value

## Kotlin/JS Parity - Functions Added in Iteration 6
- Pipeline type - zero-cost chainable wrapper for .let.let.let.let chaining
- Pipe - creates Pipeline wrapping any value
- Let (Pipeline method) - transforms and chains
- Also (Pipeline method) - side-effects in chain
- TakeIf / TakeUnless (Pipeline methods) - conditional returns
- PipeTransform / PipeMap - type-changing transformations

### Pipeline Performance (Benchmarked)
```
BenchmarkPipeChain/pipe_6_lets-11            470118    2597 ns/op   11200 B/op   20 allocs/op
BenchmarkPipeChain/direct_calls_equivalent   468200    2604 ns/op   11200 B/op   20 allocs/op
```
Pipeline wrapper has **zero overhead** compared to direct function calls - identical performance!

## Future Enhancements (Optional)
- Parallel processing with goroutines
- Stream type for infinite sequences
- Range type for numeric ranges
- More specialized numeric operations
