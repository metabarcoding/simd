package test

import (
	"log"
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
			t.Errorf("elements not equal %d : %v!=%v", i, test[i], control[i])
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
	log.Printf("checkOperation 1")
	checkOperation(t, test, control, []T{}, []T{}, []T{})
	log.Printf("checkOperation 2")
	checkOperation(t, test, control, small[T](11), small[T](13), small[T](17))
	log.Printf("checkOperation 3")
	checkOperation(t, test, control, small[T](29), small[T](19), large[T](23))
	log.Printf("checkOperation 4")
	checkOperation(t, test, control, small[T](37), large[T](41), small[T](31))
	log.Printf("checkOperation 5")
	checkOperation(t, test, control, small[T](43), large[T](47), large[T](53))
	log.Printf("checkOperation 6")
	checkOperation(t, test, control, large[T](67), small[T](59), small[T](61))
	log.Printf("checkOperation 7")
	checkOperation(t, test, control, large[T](73), small[T](79), large[T](71))
}

// Fonctions génériques pour tous les types
func generateSmall[T data.Integer](length int) []T {
	elements := make([]T, length)
	for i := range length {
		elements[i] = T(i + 1)
	}
	return elements
}

func generateLarge[T data.Integer](length int) []T {
	elements := make([]T, length)
	for i := range length {
		elements[i] = T(-i - 100) // Valeurs négatives
	}
	return elements
}

func generateMixed[T data.Integer](length int) []T {
	elements := make([]T, length)
	for i := range length {
		if i%2 == 0 {
			elements[i] = T(i)
		} else {
			elements[i] = T(-i)
		}
	}
	return elements
}

func UniversalGeneric[T data.Integer](t *testing.T, test, control data.Operation[T]) {
	log.Printf("Test 1 - slices vides")
	checkOperation(t, test, control, []T{}, []T{}, []T{})

	log.Printf("Test 2 - petites valeurs")
	checkOperation(t, test, control,
		generateSmall[T](11), generateSmall[T](13), make([]T, 17))

	log.Printf("Test 3 - grandes vs petites valeurs")
	checkOperation(t, test, control,
		generateSmall[T](29), generateLarge[T](19), make([]T, 23))

	log.Printf("Test 4 - valeurs mélangées")
	checkOperation(t, test, control,
		generateMixed[T](37), generateMixed[T](41), make([]T, 31))
}
