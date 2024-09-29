package simd

//import (
//    "fmt"
//)

type (
    element interface {int8 | int16 | int32 | int64 | float32 | float64}
    operation[E element] func(left, right, result []E) int
)

func goAdd[E element](left, right, result []E) int {
//    fmt.Println("goAdd")
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    for i := 0; i < n; i++ {
        result[i] = left[i] + right[i]
    }
    return n
}

func goSub[E element](left, right, result []E) int {
//    fmt.Println("goSub")
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    for i := 0; i < n; i++ {
        result[i] = left[i] - right[i]
    }
    return n
}

func goMul[E element](left, right, result []E) int {
//    fmt.Println("goMul")
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    for i := 0; i < n; i++ {
        result[i] = left[i] * right[i]
    }
    return n
}

func goDiv[E element](left, right, result []E) int {
//    fmt.Println("goDiv")
    n := len(result)
    if m := len(left); m < n {
        n = m
    }
    if m := len(right); m < n {
        n = m
    }
    for i := 0; i < n; i++ {
        result[i] = left[i] / right[i]
    }
    return n
}

var (
    addInt8    operation[int8]    = goAdd[int8]
    addInt16   operation[int16]   = goAdd[int16]
    addInt32   operation[int32]   = goAdd[int32]
    addInt64   operation[int64]   = goAdd[int64]
    addFloat32 operation[float32] = goAdd[float32]
    addFloat64 operation[float64] = goAdd[float64]
    subInt8    operation[int8]    = goSub[int8]
    subInt16   operation[int16]   = goSub[int16]
    subInt32   operation[int32]   = goSub[int32]
    subInt64   operation[int64]   = goSub[int64]
    subFloat32 operation[float32] = goSub[float32]
    subFloat64 operation[float64] = goSub[float64]
    mulInt8    operation[int8]    = goMul[int8]
    mulInt16   operation[int16]   = goMul[int16]
    mulInt32   operation[int32]   = goMul[int32]
    mulInt64   operation[int64]   = goMul[int64]
    mulFloat32 operation[float32] = goMul[float32]
    mulFloat64 operation[float64] = goMul[float64]
    divInt8    operation[int8]    = goDiv[int8]
    divInt16   operation[int16]   = goDiv[int16]
    divInt32   operation[int32]   = goDiv[int32]
    divInt64   operation[int64]   = goDiv[int64]
    divFloat32 operation[float32] = goDiv[float32]
    divFloat64 operation[float64] = goDiv[float64]
)

func AddInt8(left, right, result []int8) int {
    return addInt8(left, right, result)
}

func AddInt16(left, right, result []int16) int {
    return addInt16(left, right, result)
}

func AddInt32(left, right, result []int32) int {
    return addInt32(left, right, result)
}

func AddInt64(left, right, result []int64) int {
    return addInt64(left, right, result)
}

func AddFloat32(left, right, result []float32) int {
    return addFloat32(left, right, result)
}

func AddFloat64(left, right, result []float64) int {
    return addFloat64(left, right, result)
}

func SubInt8(left, right, result []int8) int {
    return subInt8(left, right, result)
}

func SubInt16(left, right, result []int16) int {
    return subInt16(left, right, result)
}

func SubInt32(left, right, result []int32) int {
    return subInt32(left, right, result)
}

func SubInt64(left, right, result []int64) int {
    return subInt64(left, right, result)
}

func SubFloat32(left, right, result []float32) int {
    return subFloat32(left, right, result)
}

func SubFloat64(left, right, result []float64) int {
    return subFloat64(left, right, result)
}

func MulInt8(left, right, result []int8) int {
    return mulInt8(left, right, result)
}

func MulInt16(left, right, result []int16) int {
    return mulInt16(left, right, result)
}

func MulInt32(left, right, result []int32) int {
    return mulInt32(left, right, result)
}

func MulInt64(left, right, result []int64) int {
    return mulInt64(left, right, result)
}

func MulFloat32(left, right, result []float32) int {
    return mulFloat32(left, right, result)
}

func MulFloat64(left, right, result []float64) int {
    return mulFloat64(left, right, result)
}

func DivInt8(left, right, result []int8) int {
    return divInt8(left, right, result)
}

func DivInt16(left, right, result []int16) int {
    return divInt16(left, right, result)
}

func DivInt32(left, right, result []int32) int {
    return divInt32(left, right, result)
}

func DivInt64(left, right, result []int64) int {
    return divInt64(left, right, result)
}

func DivFloat32(left, right, result []float32) int {
    return divFloat32(left, right, result)
}

func DivFloat64(left, right, result []float64) int {
    return divFloat64(left, right, result)
}
