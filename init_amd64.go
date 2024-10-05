// +build amd64

package simd

import (
    "fmt"
    "github.com/pehringer/simd/internal/sse"
    "github.com/pehringer/simd/internal/sse2"
    "github.com/pehringer/simd/internal/avx"
    "github.com/pehringer/simd/internal/avx2"
)

func init() {
    if sse.Supported() {
        fmt.Println("SSE")
        addFloat32 = sse.AddFloat32
        subFloat32 = sse.SubFloat32
        mulFloat32 = sse.MulFloat32
        divFloat32 = sse.DivFloat32
    }
    if sse2.Supported() {
        fmt.Println("SSE2")
        addInt32   = sse2.AddInt32
        addInt64   = sse2.AddInt64
        addFloat64 = sse2.AddFloat64
        subInt32   = sse2.SubInt32
        subInt64   = sse2.SubInt64
        subFloat64 = sse2.SubFloat64
        mulFloat64 = sse2.MulFloat64
        divFloat64 = sse2.DivFloat64
    }
    if avx.Supported() {
        fmt.Println("AVX")
        addFloat32 = avx.AddFloat32
        addFloat64 = avx.AddFloat64
        subFloat32 = avx.SubFloat32
        subFloat64 = avx.SubFloat64
        mulFloat32 = avx.MulFloat32
        mulFloat64 = avx.MulFloat64
        divFloat32 = avx.DivFloat32
        divFloat64 = avx.DivFloat64
    }
    if avx2.Supported() {
        fmt.Println("AVX2")
    }
}
