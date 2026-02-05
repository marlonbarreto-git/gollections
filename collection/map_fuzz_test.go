package collection_test

import (
	"testing"

	"github.com/marlonbarreto-git/gollections/collection"
)

func FuzzMapPutGet(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{10, 20, 30})
	f.Add([]byte{}, []byte{})
	f.Add([]byte{1, 1, 1}, []byte{10, 20, 30})

	f.Fuzz(func(t *testing.T, keys, values []byte) {
		m := make(collection.MutableMap[int, int])

		minLen := len(keys)
		if len(values) < minLen {
			minLen = len(values)
		}

		expected := make(map[int]int)
		for i := 0; i < minLen; i++ {
			key := int(keys[i])
			value := int(values[i])
			m[key] = value
			expected[key] = value
		}

		for k, v := range expected {
			if got := m[k]; got != v {
				t.Errorf("Get(%d) = %d, want %d", k, got, v)
			}
		}

		if m.Len() != len(expected) {
			t.Errorf("Len() = %d, want %d", m.Len(), len(expected))
		}
	})
}

func FuzzMapFilter(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{2, 4, 6, 8})

	f.Fuzz(func(t *testing.T, data []byte) {
		m := make(collection.MutableMap[int, int])
		for i, b := range data {
			m[i] = int(b)
		}

		filtered := m.Filter(func(k, v int) bool { return v%2 == 0 })

		for k, v := range filtered {
			if v%2 != 0 {
				t.Errorf("Filter kept odd value: %d at key %d", v, k)
			}
		}

		if filtered.Len() > m.Len() {
			t.Error("Filter result is larger than original")
		}
	})
}

func FuzzMapContainsKey(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, byte(3))
	f.Add([]byte{}, byte(1))
	f.Add([]byte{0, 0, 0}, byte(0))

	f.Fuzz(func(t *testing.T, keys []byte, target byte) {
		m := make(collection.MutableMap[int, int])
		for i, k := range keys {
			m[int(k)] = i
		}

		contains := m.ContainsKey(int(target))

		expected := false
		for _, k := range keys {
			if k == target {
				expected = true
				break
			}
		}

		if contains != expected {
			t.Errorf("ContainsKey(%d) = %v, want %v", target, contains, expected)
		}
	})
}

func FuzzMapKeys(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{1, 1, 2, 2, 3})

	f.Fuzz(func(t *testing.T, data []byte) {
		m := make(collection.MutableMap[int, int])
		for i, b := range data {
			m[int(b)] = i
		}

		keys := m.Keys()

		if keys.Len() != m.Len() {
			t.Errorf("Keys length mismatch: got %d, want %d", keys.Len(), m.Len())
		}

		for _, k := range keys {
			if !m.ContainsKey(k) {
				t.Errorf("Keys returned non-existent key: %d", k)
			}
		}
	})
}

func FuzzMapValues(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})
	f.Add([]byte{10, 10, 10})

	f.Fuzz(func(t *testing.T, data []byte) {
		m := make(collection.MutableMap[int, int])
		for i, b := range data {
			m[i] = int(b)
		}

		values := m.Values()

		if values.Len() != m.Len() {
			t.Errorf("Values length mismatch: got %d, want %d", values.Len(), m.Len())
		}
	})
}

func FuzzMapCopy(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		m := make(collection.MutableMap[int, int])
		for i, b := range data {
			m[i] = int(b)
		}

		copied := m.Copy()

		if copied.Len() != m.Len() {
			t.Errorf("Copy length mismatch: got %d, want %d", copied.Len(), m.Len())
		}

		for k, v := range m {
			if copied[k] != v {
				t.Errorf("Copy value mismatch at key %d", k)
			}
		}

		if m.Len() > 0 {
			for k := range m {
				m[k] = -1
				break
			}

			mutationAffectedCopy := false
			for k, v := range copied {
				if m[k] == -1 && v == -1 {
					mutationAffectedCopy = true
					break
				}
			}

			if mutationAffectedCopy {
				t.Error("Copy was affected by mutation of original")
			}
		}
	})
}

func FuzzMapMerge(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{4, 5, 6})
	f.Add([]byte{}, []byte{1, 2})
	f.Add([]byte{1, 2}, []byte{})

	f.Fuzz(func(t *testing.T, data1, data2 []byte) {
		m1 := make(collection.MutableMap[int, int])
		m2 := make(collection.MutableMap[int, int])

		for i, b := range data1 {
			m1[int(b)] = i
		}
		for i, b := range data2 {
			m2[int(b)] = i + 100
		}

		m1Copy := m1.Copy()
		m1Copy.Merge(m2)

		for k, v := range m2 {
			if m1Copy[k] != v {
				t.Errorf("Merge didn't include key %d from m2", k)
			}
		}

		for k := range m1 {
			if !m1Copy.ContainsKey(k) && !m2.ContainsKey(k) {
				t.Errorf("Merge lost key %d from m1", k)
			}
		}
	})
}

func FuzzMapGetOrDefault(f *testing.F) {
	f.Add([]byte{1, 2, 3}, byte(2), 99)
	f.Add([]byte{}, byte(1), 42)

	f.Fuzz(func(t *testing.T, data []byte, key byte, defaultVal int) {
		m := make(collection.MutableMap[int, int])
		for i, b := range data {
			m[int(b)] = i
		}

		result := m.GetOrDefault(int(key), defaultVal)

		if m.ContainsKey(int(key)) {
			if result != m[int(key)] {
				t.Errorf("GetOrDefault returned %d instead of existing value %d", result, m[int(key)])
			}
		} else {
			if result != defaultVal {
				t.Errorf("GetOrDefault returned %d instead of default %d", result, defaultVal)
			}
		}
	})
}
