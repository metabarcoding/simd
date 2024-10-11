package shared

import (
    "testing"
)

var (
    BenchSizes []int = []int{
        1000,
        2000,
        3000,
        4000,
        5000,
        6000,
        7000,
        8000,
        9000,
       10000,
       11000,
       12000,
       13000,
       14000,
       15000,
       16000,
       17000,
       18000,
       19000,
       20000,
       21000,
       22000,
       23000,
       24000,
       25000,
       26000,
       27000,
       28000,
       29000,
       30000,
    }
)

type (
    Element interface {int32 | int64 | float32 | float64}
    Operation[E Element] func(left, right, result []E) int
)

func CheckSlice[E Element](t *testing.T, test, control []E) bool {
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

func CheckOperation[E Element](t *testing.T, test, control Operation[E], left, right, result []E) bool {
    testLeft := make([]E, len(left), cap(left))
    copy(testLeft, left)
    testRight := make([]E, len(right), cap(right))
    copy(testRight, right)
    testResult := make([]E, len(result), cap(result))
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
    counter int8
)

func Vector[E Element](length int) []E {
    elements := make([]E, length)
    for i := 0; i < length; i++ {
        counter++
        if counter == 0 {
            counter = 1
        }
        elements[i] = E(counter)
    }
    return elements
}
