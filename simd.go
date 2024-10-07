// Simd support for arithmetic operations. Allowing for parallel element-wise computations.
package simd

type (
    element interface {int32 | int64 | float32 | float64}
    operation[E element] func(left, right, result []E) int
)

func goAdd[E element](left, right, result []E) int {
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    i := 0
    for ; n-i >= 4; i += 4 {
        result[i] = left[i] + right[i]
        result[i+1] = left[i+1] + right[i+1]
        result[i+2] = left[i+2] + right[i+2]
        result[i+3] = left[i+3] + right[i+3]
    }
    for ; i < n; i++ {
        result[i] = left[i] + right[i]
    }
    return n
}

func goSub[E element](left, right, result []E) int {
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    i := 0
    for ; n-i >= 4; i += 4 {
        result[i] = left[i] - right[i]
        result[i+1] = left[i+1] - right[i+1]
        result[i+2] = left[i+2] - right[i+2]
        result[i+3] = left[i+3] - right[i+3]
    }
    for ; i < n; i++ {
        result[i] = left[i] - right[i]
    }
    return n
}

func goMul[E element](left, right, result []E) int {
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    i := 0
    for ; n-i >= 4; i += 4 {
        result[i] = left[i] * right[i]
        result[i+1] = left[i+1] * right[i+1]
        result[i+2] = left[i+2] * right[i+2]
        result[i+3] = left[i+3] * right[i+3]
    }
    for ; i < n; i++ {
        result[i] = left[i] * right[i]
    }
    return n
}

func goDiv[E element](left, right, result []E) int {
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    i := 0
    for ; n-i >= 4; i += 4 {
        result[i] = left[i] / right[i]
        result[i+1] = left[i+1] / right[i+1]
        result[i+2] = left[i+2] / right[i+2]
        result[i+3] = left[i+3] / right[i+3]
    }
    for ; i < n; i++ {
        result[i] = left[i] / right[i]
    }
    return n
}

var (
    addInt32   operation[int32]   = goAdd[int32]
    addInt64   operation[int64]   = goAdd[int64]
    addFloat32 operation[float32] = goAdd[float32]
    addFloat64 operation[float64] = goAdd[float64]
    subInt32   operation[int32]   = goSub[int32]
    subInt64   operation[int64]   = goSub[int64]
    subFloat32 operation[float32] = goSub[float32]
    subFloat64 operation[float64] = goSub[float64]
    mulInt32   operation[int32]   = goMul[int32]
    mulInt64   operation[int64]   = goMul[int64]
    mulFloat32 operation[float32] = goMul[float32]
    mulFloat64 operation[float64] = goMul[float64]
    divInt32   operation[int32]   = goDiv[int32]
    divInt64   operation[int64]   = goDiv[int64]
    divFloat32 operation[float32] = goDiv[float32]
    divFloat64 operation[float64] = goDiv[float64]
)

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
