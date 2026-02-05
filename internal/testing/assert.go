package testing

import (
	"encoding/json"
	"reflect"
	"slices"
	"testing"
)

func Equal[T any](t *testing.T, expected, actual T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func NotEqual[T any](t *testing.T, expected, actual T) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v to not equal %v", expected, actual)
	}
}

func True(t *testing.T, value bool) {
	t.Helper()
	if !value {
		t.Error("expected true, got false")
	}
}

func False(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Error("expected false, got true")
	}
}

func Nil(t *testing.T, value any) {
	t.Helper()
	if value != nil && !reflect.ValueOf(value).IsNil() {
		t.Errorf("expected nil, got %v", value)
	}
}

func NotNil(t *testing.T, value any) {
	t.Helper()
	if value == nil || reflect.ValueOf(value).IsNil() {
		t.Error("expected non-nil value")
	}
}

func Panics(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, but none occurred")
		}
	}()
	fn()
}

func NoPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	fn()
}

func Len[T any](t *testing.T, collection []T, expectedLen int) {
	t.Helper()
	if len(collection) != expectedLen {
		t.Errorf("expected length %d, got %d", expectedLen, len(collection))
	}
}

func Empty[T any](t *testing.T, collection []T) {
	t.Helper()
	if len(collection) != 0 {
		t.Errorf("expected empty collection, got %d elements", len(collection))
	}
}

func NotEmpty[T any](t *testing.T, collection []T) {
	t.Helper()
	if len(collection) == 0 {
		t.Error("expected non-empty collection")
	}
}

func Contains[T comparable](t *testing.T, collection []T, element T) {
	t.Helper()
	if !slices.Contains(collection, element) {
		t.Errorf("expected collection to contain %v", element)
	}
}

func NotContains[T comparable](t *testing.T, collection []T, element T) {
	t.Helper()
	if slices.Contains(collection, element) {
		t.Errorf("expected collection to not contain %v", element)
	}
}

func MapEqual[K comparable, V any](t *testing.T, expected, actual map[K]V) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected map %v, got %v", expected, actual)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func ErrorIs(t *testing.T, err, target error) {
	t.Helper()
	if err != target {
		t.Errorf("expected error %v, got %v", target, err)
	}
}

func Greater[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](t *testing.T, a, b T) {
	t.Helper()
	if a <= b {
		t.Errorf("expected %v > %v", a, b)
	}
}

func GreaterOrEqual[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](t *testing.T, a, b T) {
	t.Helper()
	if a < b {
		t.Errorf("expected %v >= %v", a, b)
	}
}

func Less[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](t *testing.T, a, b T) {
	t.Helper()
	if a >= b {
		t.Errorf("expected %v < %v", a, b)
	}
}

func LessOrEqual[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](t *testing.T, a, b T) {
	t.Helper()
	if a > b {
		t.Errorf("expected %v <= %v", a, b)
	}
}

func PanicsWithValue(t *testing.T, expected any, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, but none occurred")
		} else if !reflect.DeepEqual(r, expected) {
			t.Errorf("expected panic value %v, got %v", expected, r)
		}
	}()
	fn()
}

func JSONEq(t *testing.T, expected, actual string) {
	t.Helper()
	var expectedMap, actualMap any
	if err := json.Unmarshal([]byte(expected), &expectedMap); err != nil {
		t.Errorf("failed to unmarshal expected JSON: %v", err)
		return
	}
	if err := json.Unmarshal([]byte(actual), &actualMap); err != nil {
		t.Errorf("failed to unmarshal actual JSON: %v", err)
		return
	}
	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("expected JSON %s, got %s", expected, actual)
	}
}
