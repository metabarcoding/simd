// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool

func sseAddFloat32(left, right, result []float32) int

func sseSubFloat32(left, right, result []float32) int

func sseMulFloat32(left, right, result []float32) int

func sseDivFloat32(left, right, result []float32) int

func sse2Supported() bool

//func sse2AddInt8(left, right, result []int8) int

//func sse2AddInt16(left, right, result []int16) int

//func sse2AddInt32(left, right, result []int32) int

//func sse2AddInt64(left, right, result []int64) int

func sse2AddFloat64(left, right, result []float64) int

//func sse2SubInt8(left, right, result []int8) int

//func sse2SubInt16(left, right, result []int16) int

//func sse2SubInt32(left, right, result []int32) int

//func sse2SubInt64(left, right, result []int64) int

func sse2SubFloat64(left, right, result []float64) int

func sse2MulFloat64(left, right, result []float64) int

func sse2DivFloat64(left, right, result []float64) int

func avxSupported() bool

func avx2Supported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE")
        addFloat32 = sseAddFloat32
        subFloat32 = sseSubFloat32
        mulFloat32 = sseMulFloat32
        divFloat32 = sseDivFloat32
    }
    if sse2Supported() {
        fmt.Println("SSE2")
//        addInt8    = sse2AddInt8
//        addInt16   = sse2AddInt16
//        addInt32   = sse2AddInt32
//        addInt64   = sse2AddInt64
        addFloat64 = sse2AddFloat64
//        subInt8    = sse2SubInt8
//        subInt16   = sse2SubInt16
//        subInt32   = sse2SubInt32
//        subInt64   = sse2SubInt64
        subFloat64 = sse2SubFloat64
        mulFloat64 = sse2MulFloat64
        divFloat64 = sse2DivFloat64
    }
    if avxSupported() {
        fmt.Println("AVX")
    }
    if avx2Supported() {
        fmt.Println("AVX2")
    }
}
