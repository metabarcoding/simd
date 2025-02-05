//go:build amd64
// +build amd64

package avx512vl

import (
	"testing"

	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/test"
	"golang.org/x/sys/cpu"
)

func TestAvx512vl(t *testing.T) {
	if !cpu.X86.HasAVX512VL {
		t.Skip("avx512vl not supported")
		return
	}
	t.Run("AddFloat32", func(t *testing.T) { test.Universal(t, AddFloat32, fallback.Add) })
	t.Run("AddFloat64", func(t *testing.T) { test.Universal(t, AddFloat64, fallback.Add) })
	t.Run("AddInt32", func(t *testing.T) { test.Universal(t, AddInt32, fallback.Add) })
	t.Run("AddInt64", func(t *testing.T) { test.Universal(t, AddInt64, fallback.Add) })
	t.Run("AndInt32", func(t *testing.T) { test.Universal(t, AndInt32, fallback.And) })
	t.Run("AndInt64", func(t *testing.T) { test.Universal(t, AndInt64, fallback.And) })
	t.Run("DivFloat32", func(t *testing.T) { test.Universal(t, DivFloat32, fallback.Div) })
	t.Run("DivFloat64", func(t *testing.T) { test.Universal(t, DivFloat64, fallback.Div) })
	t.Run("MaxFloat32", func(t *testing.T) { test.Universal(t, MaxFloat32, fallback.Max) })
	t.Run("MaxFloat64", func(t *testing.T) { test.Universal(t, MaxFloat64, fallback.Max) })
	t.Run("MaxInt32", func(t *testing.T) { test.Universal(t, MaxInt32, fallback.Max) })
	t.Run("MaxInt64", func(t *testing.T) { test.Universal(t, MaxInt64, fallback.Max) })
	t.Run("MinFloat32", func(t *testing.T) { test.Universal(t, MinFloat32, fallback.Min) })
	t.Run("MinFloat64", func(t *testing.T) { test.Universal(t, MinFloat64, fallback.Min) })
	t.Run("MinInt32", func(t *testing.T) { test.Universal(t, MinInt32, fallback.Min) })
	t.Run("MinInt64", func(t *testing.T) { test.Universal(t, MinInt64, fallback.Min) })
	t.Run("MulFloat32", func(t *testing.T) { test.Universal(t, MulFloat32, fallback.Mul) })
	t.Run("MulFloat64", func(t *testing.T) { test.Universal(t, MulFloat64, fallback.Mul) })
	t.Run("MulInt32", func(t *testing.T) { test.Universal(t, MulInt32, fallback.Mul) })
	t.Run("MulInt64", func(t *testing.T) { test.Universal(t, MulInt64, fallback.Mul) })
	t.Run("OrInt32", func(t *testing.T) { test.Universal(t, OrInt32, fallback.Or) })
	t.Run("OrInt64", func(t *testing.T) { test.Universal(t, OrInt64, fallback.Or) })
	t.Run("SubFloat32", func(t *testing.T) { test.Universal(t, SubFloat32, fallback.Sub) })
	t.Run("SubFloat64", func(t *testing.T) { test.Universal(t, SubFloat64, fallback.Sub) })
	t.Run("SubInt32", func(t *testing.T) { test.Universal(t, SubInt32, fallback.Sub) })
	t.Run("SubInt64", func(t *testing.T) { test.Universal(t, SubInt64, fallback.Sub) })
	t.Run("XorInt32", func(t *testing.T) { test.Universal(t, XorInt32, fallback.Xor) })
	t.Run("XorInt64", func(t *testing.T) { test.Universal(t, XorInt64, fallback.Xor) })
}
