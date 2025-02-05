package simd

import (
	"fmt"
)

func ExampleAddFloat32() {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := AddFloat32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [4 16 6 14 0 0]
}

func ExampleAddFloat64() {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := AddFloat64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [4 16 6 14 0 0]
}

func ExampleAddInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := AddInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [4 16 6 14 0 0]
}

func ExampleAddInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := AddInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [4 16 6 14 0 0]
}

func ExampleAndInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := AndInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [1 1 0 0 0 0]
}

func ExampleAndInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := AndInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [1 1 0 0 0 0]
}

func ExampleDivFloat32() {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := DivFloat32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [0.33333334 1.2857143 0.5 1.3333334 0 0]
}

func ExampleDivFloat64() {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := DivFloat64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [0.3333333333333333 1.2857142857142858 0.5 1.3333333333333333 0 0]
}

func ExampleDivInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := DivInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [0 1 0 1 0 0]
}

func ExampleDivInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := DivInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [0 1 0 1 0 0]
}

func ExampleMaxFloat32() {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := MaxFloat32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 9 4 8 0 0]
}

func ExampleMaxFloat64() {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := MaxFloat64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 9 4 8 0 0]
}

func ExampleMaxInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := MaxInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 9 4 8 0 0]
}

func ExampleMaxInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := MaxInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 9 4 8 0 0]
}

func ExampleMinFloat32() {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := MinFloat32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [1 7 2 6 0 0]
}

func ExampleMinFloat64() {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := MinFloat64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [1 7 2 6 0 0]
}

func ExampleMinInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := MinInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [1 7 2 6 0 0]
}

func ExampleMinInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := MinInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [1 7 2 6 0 0]
}

func ExampleMulFloat32() {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := MulFloat32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 63 8 48 0 0]
}

func ExampleMulFloat64() {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := MulFloat64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 63 8 48 0 0]
}

func ExampleMulInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := MulInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 63 8 48 0 0]
}

func ExampleMulInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := MulInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 63 8 48 0 0]
}

func ExampleOrInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := OrInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 15 6 14 0 0]
}

func ExampleOrInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := OrInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [3 15 6 14 0 0]
}

func ExampleSubFloat32() {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := SubFloat32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [-2 2 -2 2 0 0]
}

func ExampleSubFloat64() {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := SubFloat64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [-2 2 -2 2 0 0]
}

func ExampleSubInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := SubInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [-2 2 -2 2 0 0]
}

func ExampleSubInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := SubInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [-2 2 -2 2 0 0]
}

func ExampleXorInt32() {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := XorInt32(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [2 14 6 14 0 0]
}

func ExampleXorInt64() {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := XorInt64(left, right, result)
	fmt.Print(length, result)
	// Output: 4 [2 14 6 14 0 0]
}
