// +build amd64

package simd

import (
    "github.com/pehringer/simd/internal/avx"
    "github.com/pehringer/simd/internal/avx2"
    "github.com/pehringer/simd/internal/sse"
    "github.com/pehringer/simd/internal/sse2"
    "github.com/pehringer/simd/internal/sse41"
)

func init() {
    if sse.Supported() {
        addFloat32 = sse.AddFloat32
        divFloat32 = sse.DivFloat32
        mulFloat32 = sse.MulFloat32
        subFloat32 = sse.SubFloat32
    }
    if sse2.Supported() {
        addFloat64 = sse2.AddFloat64
        addInt32   = sse2.AddInt32
        addInt64   = sse2.AddInt64
        divFloat64 = sse2.DivFloat64
        mulFloat64 = sse2.MulFloat64
        subFloat64 = sse2.SubFloat64
        subInt32   = sse2.SubInt32
        subInt64   = sse2.SubInt64
    }
    if sse41.Supported() {
        mulInt32   = sse41.MulInt32
    }
    if avx.Supported() {
        addFloat32 = avx.AddFloat32
        addFloat64 = avx.AddFloat64
        divFloat32 = avx.DivFloat32
        divFloat64 = avx.DivFloat64
        mulFloat32 = avx.MulFloat32
        mulFloat64 = avx.MulFloat64
        subFloat32 = avx.SubFloat32
        subFloat64 = avx.SubFloat64
    }
    if avx2.Supported() {
        addInt32   = avx2.AddInt32
        addInt64   = avx2.AddInt64
        mulInt32   = avx2.MulInt32
        subInt32   = avx2.SubInt32
        subInt64   = sse2.SubInt64
    }
}
