//go:build amd64
// +build amd64

package sse

import (
	"testing"

	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/test"
	"golang.org/x/sys/cpu"
)

func TestSse(t *testing.T) {
	if !cpu.X86.HasSSE2 {
		t.Skip("sse not supported")
		return
	}
	t.Run("AddFloat32", func(t *testing.T) { test.Universal(t, AddFloat32, fallback.Add) })
	t.Run("DivFloat32", func(t *testing.T) { test.Universal(t, DivFloat32, fallback.Div) })
	t.Run("MaxFloat32", func(t *testing.T) { test.Universal(t, MaxFloat32, fallback.Max) })
	t.Run("MinFloat32", func(t *testing.T) { test.Universal(t, MinFloat32, fallback.Min) })
	t.Run("MulFloat32", func(t *testing.T) { test.Universal(t, MulFloat32, fallback.Mul) })
	t.Run("SubFloat32", func(t *testing.T) { test.Universal(t, SubFloat32, fallback.Sub) })
}
