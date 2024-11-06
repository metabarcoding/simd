package shared

import (
	"math"
	"testing"
)

func CheckSlice[T Floating | Integer](t *testing.T, test, control []T) bool {
	if len(test) != len(control) {
		t.Errorf("lengths not equal")
		return false
	}
	if cap(test) != cap(control) {
		t.Errorf("capacities not equal")
		return false
	}
	for i := 0; i < len(control); i++ {
		if test[i] != control[i] {
			t.Errorf("elements not equal")
			return false
		}
	}
	return true
}

func CheckOperation[T Floating | Integer](t *testing.T, test, control Operation[T], left, right, result []T) bool {
	testLeft := make([]T, len(left), cap(left))
	copy(testLeft, left)
	testRight := make([]T, len(right), cap(right))
	copy(testRight, right)
	testResult := make([]T, len(result), cap(result))
	copy(testResult, result)
	if test(testLeft, testRight, testResult) != control(left, right, result) {
		t.Errorf("operation returned incorrect length")
		return false
	}
	if !CheckSlice(t, testLeft, left) {
		return false
	}
	if !CheckSlice(t, testRight, right) {
		return false
	}
	if !CheckSlice(t, testResult, result) {
		return false
	}
	return true
}

var (
	increment int64 = 1
	decrement int64 = math.MaxInt64
)

func Large[T Floating | Integer](length int) []T {
	elements := make([]T, length)
	for i := 0; i < length; i++ {
		elements[i] = T(decrement)
		decrement--
	}
	return elements
}

func Small[T Floating | Integer](length int) []T {
	elements := make([]T, length)
	for i := 0; i < length; i++ {
		elements[i] = T(increment)
		increment++
	}
	return elements
}
