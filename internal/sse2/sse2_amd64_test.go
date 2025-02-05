//go:build amd64
// +build amd64

package sse2

import (
	"testing"

	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/test"
	"golang.org/x/sys/cpu"
)

func TestSse2(t *testing.T) {
	if !cpu.X86.HasSSE2 {
		t.Skip("sse2 not supported")
		return
	}
	t.Run("AddFloat64", func(t *testing.T) { test.Universal(t, AddFloat64, fallback.Add) })
	t.Run("AddInt32", func(t *testing.T) { test.Universal(t, AddInt32, fallback.Add) })
	t.Run("AddInt64", func(t *testing.T) { test.Universal(t, AddInt64, fallback.Add) })
	t.Run("AndInt32", func(t *testing.T) { test.Universal(t, AndInt32, fallback.And) })
	t.Run("AndInt64", func(t *testing.T) { test.Universal(t, AndInt64, fallback.And) })
	t.Run("DivFloat64", func(t *testing.T) { test.Universal(t, DivFloat64, fallback.Div) })
	t.Run("MaxFloat64", func(t *testing.T) { test.Universal(t, MaxFloat64, fallback.Max) })
	t.Run("MinFloat64", func(t *testing.T) { test.Universal(t, MinFloat64, fallback.Min) })
	t.Run("MulFloat64", func(t *testing.T) { test.Universal(t, MulFloat64, fallback.Mul) })
	t.Run("OrInt32", func(t *testing.T) { test.Universal(t, OrInt32, fallback.Or) })
	t.Run("OrInt64", func(t *testing.T) { test.Universal(t, OrInt64, fallback.Or) })
	t.Run("SubFloat64", func(t *testing.T) { test.Universal(t, SubFloat64, fallback.Sub) })
	t.Run("SubInt32", func(t *testing.T) { test.Universal(t, SubInt32, fallback.Sub) })
	t.Run("SubInt64", func(t *testing.T) { test.Universal(t, SubInt64, fallback.Sub) })
	t.Run("XorInt32", func(t *testing.T) { test.Universal(t, XorInt32, fallback.Xor) })
	t.Run("XorInt64", func(t *testing.T) { test.Universal(t, XorInt64, fallback.Xor) })
}
