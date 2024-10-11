// Simd support for arithmetic operations. Allowing for parallel element-wise computations.
package simd

import (
    "github.com/pehringer/simd/internal/fallback"
)

var (
    addFloat32 func(left, right, result []float32) int = fallback.Add[float32]
    addFloat64 func(left, right, result []float64) int = fallback.Add[float64]
    addInt32   func(left, right, result []int32) int   = fallback.Add[int32]
    addInt64   func(left, right, result []int64) int   = fallback.Add[int64]
    divFloat32 func(left, right, result []float32) int = fallback.Div[float32]
    divFloat64 func(left, right, result []float64) int = fallback.Div[float64]
    divInt32   func(left, right, result []int32) int   = fallback.Div[int32]
    divInt64   func(left, right, result []int64) int   = fallback.Div[int64]
    mulFloat32 func(left, right, result []float32) int = fallback.Mul[float32]
    mulFloat64 func(left, right, result []float64) int = fallback.Mul[float64]
    mulInt32   func(left, right, result []int32) int   = fallback.Mul[int32]
    mulInt64   func(left, right, result []int64) int   = fallback.Mul[int64]
    subFloat32 func(left, right, result []float32) int = fallback.Sub[float32]
    subFloat64 func(left, right, result []float64) int = fallback.Sub[float64]
    subInt32   func(left, right, result []int32) int   = fallback.Sub[int32]
    subInt64   func(left, right, result []int64) int   = fallback.Sub[int64]
)

// AddFloat32 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddFloat32(left, right, result []float32) int {
    return addFloat32(left, right, result)
}

// AddFloat64 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddFloat64(left, right, result []float64) int {
    return addFloat64(left, right, result)
}

// AddInt32 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddInt32(left, right, result []int32) int {
    return addInt32(left, right, result)
}

// AddInt64 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddInt64(left, right, result []int64) int {
    return addInt64(left, right, result)
}

// DivFloat32 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivFloat32(left, right, result []float32) int {
    return divFloat32(left, right, result)
}

// DivFloat64 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivFloat64(left, right, result []float64) int {
    return divFloat64(left, right, result)
}

// DivInt32 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivInt32(left, right, result []int32) int {
    return divInt32(left, right, result)
}
// DivInt64 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivInt64(left, right, result []int64) int {
    return divInt64(left, right, result)
}

// MulFloat32 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulFloat32(left, right, result []float32) int {
    return mulFloat32(left, right, result)
}

// MulFloat64 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulFloat64(left, right, result []float64) int {
    return mulFloat64(left, right, result)
}

// MulInt32 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulInt32(left, right, result []int32) int {
    return mulInt32(left, right, result)
}

// MulInt64 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulInt64(left, right, result []int64) int {
    return mulInt64(left, right, result)
}

// SubFloat32 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubFloat32(left, right, result []float32) int {
    return subFloat32(left, right, result)
}

// SubFloat64 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubFloat64(left, right, result []float64) int {
    return subFloat64(left, right, result)
}

// SubInt32 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubInt32(left, right, result []int32) int {
    return subInt32(left, right, result)
}

// SubInt64 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubInt64(left, right, result []int64) int {
    return subInt64(left, right, result)
}
