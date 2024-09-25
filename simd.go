package simd

type (
    element interface {
        int8 | int16 | int32 | int64 | float32 | float64
    }
    operation[E element] func(left, right, result []E) int
)

func add[E element](left, right, result []E) int {
    n := len(result)
    if len(left) < n {
        n = len(left)
    }
    if len(right) < n {
        n = len(right)
    }
    i := 0
    for n-i >= 4 {
        result[i] = left[i] + right[i]
        result[i+1] = left[i+1] + right[i+1]
        result[i+2] = left[i+2] + right[i+2]
        result[i+3] = left[i+3] + right[i+3]
        i += 4
    }
    for i < n {
        result[i] = left[i] + right[i]
        i += 1
    }
    return n
}

func sub[E element](left, right, result []E) int {
    n := len(result)
    if len(left) < n {
        n = len(left)
    }
    if len(right) < n {
        n = len(right)
    }
    i := 0
    for n-i >= 4 {
        result[i] = left[i] - right[i]
        result[i+1] = left[i+1] - right[i+1]
        result[i+2] = left[i+2] - right[i+2]
        result[i+3] = left[i+3] - right[i+3]
        i += 4
    }
    for i < n {
        result[i] = left[i] - right[i]
        i += 1
    }
    return n
}

func mul[E element](left, right, result []E) int {
    n := len(result)
    if len(left) < n {
        n = len(left)
    }
    if len(right) < n {
        n = len(right)
    }
    i := 0
    for n-i >= 4 {
        result[i] = left[i] * right[i]
        result[i+1] = left[i+1] * right[i+1]
        result[i+2] = left[i+2] * right[i+2]
        result[i+3] = left[i+3] * right[i+3]
        i += 4
    }
    for i < n {
        result[i] = left[i] * right[i]
        i += 1
    }
    return n
}

func div[E element](left, right, result []E) int {
    n := len(result)
    if len(left) < n {
        n = len(left)
    }
    if len(right) < n {
        n = len(right)
    }
    i := 0
    for n-i >= 4 {
        result[i] = left[i] / right[i]
        result[i+1] = left[i+1] / right[i+1]
        result[i+2] = left[i+2] / right[i+2]
        result[i+3] = left[i+3] / right[i+3]
        i += 4
    }
    for i < n {
        result[i] = left[i] / right[i]
        i += 1
    }
    return n
}

var (
    AddInt8    operation[int8]    = add[int8]
    AddInt16   operation[int16]   = add[int16]
    AddInt32   operation[int32]   = add[int32]
    AddInt64   operation[int64]   = add[int64]
    AddFloat32 operation[float32] = add[float32]
    AddFloat64 operation[float64] = add[float64]

    SubInt8    operation[int8]    = sub[int8]
    SubInt16   operation[int16]   = sub[int16]
    SubInt32   operation[int32]   = sub[int32]
    SubInt64   operation[int64]   = sub[int64]
    SubFloat32 operation[float32] = sub[float32]
    SubFloat64 operation[float64] = sub[float64]

    MulInt8    operation[int8]    = mul[int8]
    MulInt16   operation[int16]   = mul[int16]
    MulInt32   operation[int32]   = mul[int32]
    MulInt64   operation[int64]   = mul[int64]
    MulFloat32 operation[float32] = mul[float32]
    MulFloat64 operation[float64] = mul[float64]

    DivInt8    operation[int8]    = div[int8]
    DivInt16   operation[int16]   = div[int16]
    DivInt32   operation[int32]   = div[int32]
    DivInt64   operation[int64]   = div[int64]
    DivFloat32 operation[float32] = div[float32]
    DivFloat64 operation[float64] = div[float64]
)

