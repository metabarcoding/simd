package simd

import (
	"testing"
)

func checkResult[E element](t *testing.T, o operation[E], left, right, result, expected []E) {
    if len(expected) > len(result) {
         t.Errorf("expected must be shorter or equal in length to result")
    }
    leftCopy := make([]E, len(left))
    copy(leftCopy, left)
    rightCopy := make([]E, len(right))
    copy(rightCopy, right)
    resultCopy := make([]E, len(result))
    copy(resultCopy, result)
    if n := o(leftCopy, rightCopy, resultCopy); n != len(expected) {
        t.Errorf("operation failed returned length %v, want %v", n, len(expected))
    }
    if len(leftCopy) != len(left) {
        t.Errorf("operation failed left length changed")
    }
    for i := 0; i < len(left); i++ {
        if leftCopy[i] != left[i] {
            t.Errorf("operation failed left element changed at index %d", i)
        }
    }
    if len(rightCopy) != len(right) {
        t.Errorf("operation failed right length changed")
    }
    for i := 0; i < len(right); i++ {
        if rightCopy[i] != right[i] {
             t.Errorf("operation failed right element changed at index %d", i)
        }
    }
    if len(resultCopy) != len(result) {
        t.Errorf("operation failed result length changed")
    }
    for i := 0; i < len(expected); i++ {
        if resultCopy[i] != expected[i] {
            t.Errorf("failed result element at index %d: got %v, want %v", i, resultCopy[i], expected[i])
        }
    }
    for i := len(expected); i < len(result); i++ {
        if resultCopy[i] != result[i] {
             t.Errorf("operation failed result element changed at index %d", i)
        }
    }
}

func TestAddFloat32(t *testing.T) {
    checkResult(t, AddFloat32,
        []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0},
        []float32{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 1000.0, 1100.0},
        make([]float32, 16),
        []float32{11.0, 22.0, 33.0, 44.0, 55.0, 66.0, 77.0, 88.0, 99.0, 1010.0, 1111.0})
}

func TestSubFloat32(t *testing.T) {
    checkResult(t, SubFloat32,
        []float32{11.0, 22.0, 33.0, 44.0, 55.0, 66.0, 77.0, 88.0, 99.0, 1010.0, 1111.0},
        []float32{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 1000.0, 1100.0},
        make([]float32, 16),
        []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0})
}

func TestMulFloat32(t *testing.T) {
    checkResult(t, MulFloat32,
        []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0},
        []float32{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 100.0, 110.0},
        make([]float32, 16),
        []float32{10.0, 40.0, 90.0, 160.0, 250.0, 360.0, 490.0, 640.0, 810.0, 1000.0, 1210.0})
}

func TestDivFloat32(t *testing.T) {
    checkResult(t, DivFloat32,
        []float32{10.0, 25.0, 36.0, 49.0, 50.0, 72.0, 98.0, 120.0, 135.0, 200.0, 220.0},
        []float32{2.0, 5.0, 6.0, 7.0, 10.0, 8.0, 14.0, 15.0, 27.0, 20.0, 22.0},
        make([]float32, 16),
        []float32{5.0, 5.0, 6.0, 7.0, 5.0, 9.0, 7.0, 8.0, 5.0, 10.0, 10.0})
}
