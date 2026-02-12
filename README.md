<p align="center">
  <img src="https://github.com/user-attachments/assets/2d54a7dd-8ce9-4867-af87-ccff111986c9" height="400">
  <h1 align="center">
    Gollections
    <br>
    <img alt="build-status" src="https://img.shields.io/badge/build-passing-brightgreen.svg?style=flat-square" />
    <img alt="license" src="https://img.shields.io/badge/license-MIT-E91E63.svg?style=flat-square" />
    <img alt="go-version" src="https://img.shields.io/badge/go-1.23%2B-00ADD8.svg?style=flat-square" />
  </h1>
</p>

A type-safe collections framework for Go using generics, inspired by Kotlin's collections API. Provides `List`, `Set`, `MutableMap`, `Sequence`, and `Pipeline` with a rich set of functional operations.

## Installation

```bash
go get github.com/marlonbarreto-git/gollections
```

Requires Go 1.23+.

## Collections

### List

A generic ordered collection built on top of Go slices. Supports functional operations like `Filter`, `Map`, `Reduce`, `FlatMap`, `Sorted`, `GroupBy`, and many more.

```go
import "github.com/marlonbarreto-git/gollections/list"

nums := list.Of(1, 2, 3, 4, 5)

// Filter
evens := nums.Filter(func(n int) bool { return n%2 == 0 })
// [2, 4]

// Map (type-changing requires free function)
doubled := collection.ListMap(nums, func(n int) string {
    return fmt.Sprintf("%d!", n)
})
// ["1!", "2!", "3!", "4!", "5!"]

// Find
found := nums.Find(func(n int) bool { return n > 3 })
// Optional[4]

// Reduce
sum := nums.Reduce(0, func(acc, n int) int { return acc + n })
// 15

// Chain operations
result := nums.
    Filter(func(n int) bool { return n > 2 }).
    Sorted(func(a, b int) int { return b - a }).
    Take(2)
// [5, 4]
```

**Key methods**: `Filter`, `Find`, `FindLast`, `First`, `Last`, `Get`, `Append`, `ForEach`, `ForEachIndexed`, `Some`, `Every`, `None`, `Count`, `Sum`, `Reduce`, `Sorted`, `Reversed`, `Distinct`, `DistinctBy`, `Take`, `TakeLast`, `TakeWhile`, `Drop`, `DropLast`, `DropWhile`, `Chunked`, `Contains`, `ContainsAll`, `IndexOf`, `LastIndexOf`, `Slice`, `FlatMap`, `Partition`, `GroupBy`, `MinBy`, `MaxBy`, `Join`, `Associate`, `AssociateBy`, `Windowed`, `Single`, `ElementAt`, `Shuffled`, `Random`, `Plus`, `Minus`, `OnEach`, `Also`, `TakeIf`, `TakeUnless`, `IsEmpty`, `IsNotEmpty`, `Len`, `AsSequence`.

**Free functions**: `ListMap`, `Fold`, `FlatMap`, `GroupBy`, `Zip`, `Flatten`, `Min`, `Max`, `Average`, `MapIndexed`, `MapNotNull`, `MapIndexedNotNull`, `RunningFold`, `Scan`, `FoldIndexed`, `ReduceIndexed`, `FoldRight`, `ReduceRight`, `FoldRightIndexed`, `ReduceRightIndexed`, `RunningFoldIndexed`, `SortedDescending`, `SumOf`, `ToSet`, `ToMap`, `ToMapWithValue`, `Let`, `Unzip`, `FirstNotNullOf`, `ZipWithNext`.

### Set

A generic set backed by a `map[K]Empty`. Supports standard set operations.

```go
import "github.com/marlonbarreto-git/gollections/set"

s := set.Of("go", "rust", "zig")

s.Contains("go")    // true
s.Add("python")     // true (was new)
s.Len()             // 4

// Set operations
other := set.Of("go", "python", "java")
s.Union(other)      // {go, rust, zig, python, java}
s.Intersect(other)  // {go, python}
s.Subtract(other)   // {rust, zig}

// Functional operations
s.Filter(func(k string) bool { return len(k) <= 3 })
s.Any(func(k string) bool { return k == "go" })
s.All(func(k string) bool { return len(k) > 0 })
```

**Key methods**: `Contains`, `Add`, `Remove`, `Clear`, `IsEmpty`, `Len`, `Values`, `Union`, `Intersect`, `Subtract`, `Filter`, `ForEach`, `Any`, `All`, `None`, `First`, `ToList`, `ToMap`, `Also`, `TakeIf`, `TakeUnless`, `String`.

### MutableMap

A generic map wrapper providing functional operations on top of Go maps.

```go
import (
    maps "github.com/marlonbarreto-git/gollections/map"
    "github.com/marlonbarreto-git/gollections/collection"
)

m := maps.Of(
    collection.PairOf("name", "Alice"),
    collection.PairOf("city", "Singapore"),
)

m.ContainsKey("name")  // true
m.Keys()               // ["name", "city"]
m.Values()             // ["Alice", "Singapore"]

// Filter
m.Filter(func(k, v string) bool { return len(v) > 4 })
// {"city": "Singapore"}

// GetOrDefault / GetOrPut
m.GetOrDefault("age", "unknown")  // "unknown"
m.GetOrPut("age", func() string { return "30" })

// Transform
collection.MapKeys(m, func(k, v string) string {
    return strings.ToUpper(k)
})
```

**Key methods**: `Map`, `Reduce`, `ForEach`, `Filter`, `FilterKeys`, `FilterValues`, `IsEmpty`, `Len`, `Count`, `Copy`, `Keys`, `Values`, `Entries`, `Remove`, `GetOrDefault`, `GetOrPut`, `ContainsKey`, `ContainsValue`, `Merge`, `PutAll`, `ToList`, `ToSet`, `Any`, `All`, `None`, `Also`, `TakeIf`, `TakeUnless`, `String`.

**Free functions**: `Map`, `MapKeys`, `MapValues`.

### Sequence

Lazy evaluation sequences built on Go 1.23+ iterators (`iter.Seq`). Operations are deferred until terminal operations like `ToSlice()`, `Count()`, or `ForEach()` are called.

```go
import "github.com/marlonbarreto-git/gollections/sequence"

seq := sequence.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

// Lazy chain - nothing computed until ToSlice()
result := seq.
    Filter(func(n int) bool { return n%2 == 0 }).
    Map(func(n int) int { return n * 10 }).
    Take(3).
    ToSlice()
// [20, 40, 60]

// From existing slices or iterators
sequence.From(existingSlice)
sequence.FromIter(existingIterator)
```

**Key methods**: `Filter`, `Map`, `FlatMap`, `Reduce`, `Take`, `TakeWhile`, `Drop`, `DropWhile`, `First`, `Last`, `ForEach`, `Count`, `Any`, `All`, `None`, `Distinct`, `Reversed`, `Sorted`, `Contains`, `IndexOf`, `Find`, `Partition`, `OnEach`, `ToSlice`, `Iter`.

**Free functions**: `Map`, `FlatMap`, `Fold`, `Chunked`, `Zip`, `Sum`, `Average`, `Max`, `Min`, `GroupBy`, `WithIndex`.

### Pipeline

A chainable wrapper for any value, enabling Kotlin-style `let`/`also`/`takeIf`/`takeUnless` chaining.

```go
import "github.com/marlonbarreto-git/gollections/collection"

result := collection.Pipe(list.Of(1, 2, 3, 4, 5)).
    Let(func(l collection.List[int]) collection.List[int] {
        return l.Filter(func(n int) bool { return n > 2 })
    }).
    Also(func(l collection.List[int]) {
        fmt.Println("After filter:", l)
    }).
    Value()
// [3, 4, 5]

// TakeIf / TakeUnless return Optional
opt := collection.Pipe(42).
    TakeIf(func(n int) bool { return n > 0 })
// Optional[42]
```

**Key methods**: `Let`, `Also`, `TakeIf`, `TakeUnless`, `Value`.

**Free functions**: `Pipe`, `PipeTransform`, `PipeMap`.

## Project Structure

```
gollections/
  collection/     # Core types: List, Set, MutableMap, Pair, Pipeline
  list/           # List factory functions (Of, From)
  set/            # Set factory functions (Of, From)
  map/            # MutableMap factory functions (Of, From)
  sequence/       # Lazy sequence type and operations
  iterable/       # Shared collection interface
  internal/       # Internal utilities
```

## Testing

The project includes comprehensive tests, benchmarks, and fuzz tests:

```bash
go test ./...
go test -bench=. -benchmem ./collection/...
go test -fuzz=FuzzList ./collection/...
```

## License

[MIT](LICENSE)
