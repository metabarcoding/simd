package test

import (
	"math"
	"testing"

	"github.com/pehringer/simd/internal/data"
)

func checkSlice[T data.Floating | data.Integer](t *testing.T, test, control []T) bool {
	if len(test) != len(control) {
		t.Errorf("lengths not equal")
		return false
	}
	if cap(test) != cap(control) {
		t.Errorf("capacities not equal")
		return false
	}
	for i := range len(control) {
		if test[i] != control[i] {
			t.Errorf("elements not equal")
			return false
		}
	}
	return true
}

func checkOperation[T data.Floating | data.Integer](t *testing.T, test, control data.Operation[T], left, right, result []T) bool {
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
	if !checkSlice(t, testLeft, left) {
		return false
	}
	if !checkSlice(t, testRight, right) {
		return false
	}
	if !checkSlice(t, testResult, result) {
		return false
	}
	return true
}

var (
	increment int64 = 1
	decrement int64 = math.MaxInt64
)

func large[T data.Floating | data.Integer](length int) []T {
	elements := make([]T, length)
	for i := range length {
		elements[i] = T(decrement)
		decrement--
	}
	return elements
}

func small[T data.Floating | data.Integer](length int) []T {
	elements := make([]T, length)
	for i := range length {
		elements[i] = T(increment)
		increment++
	}
	return elements
}

func Universal[T data.Floating | data.Integer](t *testing.T, test, control data.Operation[T]) {
	checkOperation(t, test, control, []T{}, []T{}, []T{})
	checkOperation(t, test, control, small[T](11), small[T](13), small[T](17))
	checkOperation(t, test, control, small[T](29), small[T](19), large[T](23))
	checkOperation(t, test, control, small[T](37), large[T](41), small[T](31))
	checkOperation(t, test, control, small[T](43), large[T](47), large[T](53))
	checkOperation(t, test, control, large[T](67), small[T](59), small[T](61))
	checkOperation(t, test, control, large[T](73), small[T](79), large[T](71))
}
