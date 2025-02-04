//go:build amd64
// +build amd64

package simd

import (
	"golang.org/x/sys/cpu"
	"github.com/pehringer/simd/internal/avx"
	"github.com/pehringer/simd/internal/avx2"
	"github.com/pehringer/simd/internal/avx512vl"
	"github.com/pehringer/simd/internal/sse"
	"github.com/pehringer/simd/internal/sse2"
	"github.com/pehringer/simd/internal/sse41"
)

func init() {
	if cpu.X86.HasSSE2 {
		addFloat32 = sse.AddFloat32
		divFloat32 = sse.DivFloat32
		maxFloat32 = sse.MaxFloat32
		minFloat32 = sse.MinFloat32
		mulFloat32 = sse.MulFloat32
		subFloat32 = sse.SubFloat32
		addFloat64 = sse2.AddFloat64
		addInt32 = sse2.AddInt32
		addInt64 = sse2.AddInt64
		andInt32 = sse2.AndInt32
		andInt64 = sse2.AndInt64
		divFloat64 = sse2.DivFloat64
		maxFloat64 = sse2.MaxFloat64
		minFloat64 = sse2.MinFloat64
		mulFloat64 = sse2.MulFloat64
		orInt32 = sse2.OrInt32
		orInt64 = sse2.OrInt64
		subFloat64 = sse2.SubFloat64
		subInt32 = sse2.SubInt32
		subInt64 = sse2.SubInt64
		xorInt32 = sse2.XorInt32
		xorInt64 = sse2.XorInt64
	}
	if cpu.X86.HasSSE41 {
		maxInt32 = sse41.MaxInt32
		minInt32 = sse41.MinInt32
		mulInt32 = sse41.MulInt32
	}
	if cpu.X86.HasAVX {
		addFloat32 = avx.AddFloat32
		addFloat64 = avx.AddFloat64
		divFloat32 = avx.DivFloat32
		divFloat64 = avx.DivFloat64
		maxFloat32 = avx.MaxFloat32
		maxFloat64 = avx.MaxFloat64
		minFloat32 = avx.MinFloat32
		minFloat64 = avx.MinFloat64
		mulFloat32 = avx.MulFloat32
		mulFloat64 = avx.MulFloat64
		subFloat32 = avx.SubFloat32
		subFloat64 = avx.SubFloat64
	}
	if cpu.X86.HasAVX2 {
		addInt32 = avx2.AddInt32
		addInt64 = avx2.AddInt64
		andInt32 = avx2.AndInt32
		andInt64 = avx2.AndInt64
		maxInt32 = avx2.MaxInt32
		minInt32 = avx2.MinInt32
		mulInt32 = avx2.MulInt32
		orInt32 = avx2.OrInt32
		orInt64 = avx2.OrInt64
		subInt32 = avx2.SubInt32
		subInt64 = sse2.SubInt64
		xorInt32 = avx2.XorInt32
		xorInt64 = avx2.XorInt64
	}
	if cpu.X86.HasAVX512VL {
		addFloat32 = avx512vl.AddFloat32
		addFloat64 = avx512vl.AddFloat64
		addInt32 = avx512vl.AddInt32
		addInt64 = avx512vl.AddInt64
		andInt32 = avx512vl.AndInt32
		andInt64 = avx512vl.AndInt64
		divFloat32 = avx512vl.DivFloat32
		divFloat64 = avx512vl.DivFloat64
		maxFloat32 = avx512vl.MaxFloat32
		maxFloat64 = avx512vl.MaxFloat64
		maxInt32 = avx512vl.MaxInt32
		maxInt64 = avx512vl.MaxInt64
		minFloat32 = avx512vl.MinFloat32
		minFloat64 = avx512vl.MinFloat64
		minInt32 = avx512vl.MinInt32
		minInt64 = avx512vl.MinInt64
		mulFloat32 = avx512vl.MulFloat32
		mulFloat64 = avx512vl.MulFloat64
		mulInt32 = avx512vl.MulInt32
		mulInt64 = avx512vl.MulInt64
		orInt32 = avx512vl.OrInt32
		orInt64 = avx512vl.OrInt64
		subFloat32 = avx512vl.SubFloat32
		subFloat64 = avx512vl.SubFloat64
		subInt32 = avx512vl.SubInt32
		subInt64 = avx512vl.SubInt64
		xorInt32 = avx512vl.XorInt32
		xorInt64 = avx512vl.XorInt64
	}
}
