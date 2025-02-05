//go:build amd64
// +build amd64

package sse41

import (
	"testing"

	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/test"
	"golang.org/x/sys/cpu"
)

func TestSse41(t *testing.T) {
	if !cpu.X86.HasSSE41 {
		t.Skip("sse4.1 not supported")
		return
	}
	t.Run("MaxInt32", func(t *testing.T) { test.Universal(t, MaxInt32, fallback.Max) })
	t.Run("MinInt32", func(t *testing.T) { test.Universal(t, MinInt32, fallback.Min) })
	t.Run("MulInt32", func(t *testing.T) { test.Universal(t, MulInt32, fallback.Mul) })
}
