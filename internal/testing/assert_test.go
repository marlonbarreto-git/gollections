package testing

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
	Equal(t, "hello", "hello")
	Equal(t, []int{1, 2, 3}, []int{1, 2, 3})
	Equal(t, map[string]int{"a": 1}, map[string]int{"a": 1})
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 1, 2)
	NotEqual(t, "hello", "world")
	NotEqual(t, []int{1, 2, 3}, []int{1, 2})
}

func TestTrue(t *testing.T) {
	True(t, true)
	True(t, 1 == 1)
	True(t, "a" < "b")
}

func TestFalse(t *testing.T) {
	False(t, false)
	False(t, 1 == 2)
	False(t, "b" < "a")
}

func TestNil(t *testing.T) {
	Nil(t, nil)
	var ptr *int
	Nil(t, ptr)
	var slice []int
	Nil(t, slice)
}

func TestNotNil(t *testing.T) {
	x := 5
	NotNil(t, &x)
	NotNil(t, []int{1, 2, 3})
	NotNil(t, map[string]int{"a": 1})
}

func TestPanics(t *testing.T) {
	Panics(t, func() {
		panic("expected panic")
	})
}

func TestNoPanic(t *testing.T) {
	NoPanic(t, func() {
		_ = 1 + 1
	})
}

func TestLen(t *testing.T) {
	Len(t, []int{1, 2, 3}, 3)
	Len(t, []string{}, 0)
	Len(t, []int{1}, 1)
}

func TestEmpty(t *testing.T) {
	Empty(t, []int{})
	Empty(t, []string{})
}

func TestNotEmpty(t *testing.T) {
	NotEmpty(t, []int{1})
	NotEmpty(t, []string{"a", "b"})
}

func TestContains(t *testing.T) {
	Contains(t, []int{1, 2, 3}, 2)
	Contains(t, []string{"a", "b", "c"}, "b")
}

func TestNotContains(t *testing.T) {
	NotContains(t, []int{1, 2, 3}, 5)
	NotContains(t, []string{"a", "b", "c"}, "d")
}

func TestMapEqual(t *testing.T) {
	MapEqual(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	MapEqual(t, map[int]string{1: "a"}, map[int]string{1: "a"})
}

func TestError(t *testing.T) {
	Error(t, errors.New("some error"))
}

func TestNoError(t *testing.T) {
	NoError(t, nil)
}

func TestErrorIs(t *testing.T) {
	err := errors.New("target error")
	ErrorIs(t, err, err)
}

func TestGreater(t *testing.T) {
	Greater(t, 5, 3)
	Greater(t, 10.5, 10.0)
	Greater(t, int64(100), int64(50))
}

func TestGreaterOrEqual(t *testing.T) {
	GreaterOrEqual(t, 5, 3)
	GreaterOrEqual(t, 5, 5)
	GreaterOrEqual(t, 10.5, 10.0)
	GreaterOrEqual(t, 10.0, 10.0)
}

func TestLess(t *testing.T) {
	Less(t, 3, 5)
	Less(t, 10.0, 10.5)
	Less(t, int64(50), int64(100))
}

func TestLessOrEqual(t *testing.T) {
	LessOrEqual(t, 3, 5)
	LessOrEqual(t, 5, 5)
	LessOrEqual(t, 10.0, 10.5)
	LessOrEqual(t, 10.0, 10.0)
}

func TestPanicsWithValue(t *testing.T) {
	PanicsWithValue(t, "expected value", func() {
		panic("expected value")
	})
}

func TestJSONEq(t *testing.T) {
	JSONEq(t, `{"a":1,"b":2}`, `{"b":2,"a":1}`)
	JSONEq(t, `[1,2,3]`, `[1,2,3]`)
	JSONEq(t, `"hello"`, `"hello"`)
}
