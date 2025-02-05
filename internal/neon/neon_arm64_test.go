//go:build arm64
// +build arm64

package neon

import (
	"testing"

	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/test"
	"golang.org/x/sys/cpu"
)

func TestNeon(t *testing.T) {
	if !cpu.ARM64.HasASIMD {
		t.Skip("neon not supported")
		return
	}
	t.Run("AddFloat32", func(t *testing.T) { test.Universal(t, AddFloat32, fallback.Add) })
	t.Run("AddFloat64", func(t *testing.T) { test.Universal(t, AddFloat64, fallback.Add) })
	t.Run("AddInt32", func(t *testing.T) { test.Universal(t, AddInt32, fallback.Add) })
	t.Run("AddInt64", func(t *testing.T) { test.Universal(t, AddInt64, fallback.Add) })
	t.Run("AndInt32", func(t *testing.T) { test.Universal(t, AndInt32, fallback.And) })
	t.Run("AndInt64", func(t *testing.T) { test.Universal(t, AndInt64, fallback.And) })
	t.Run("MulFloat32", func(t *testing.T) { test.Universal(t, MulFloat32, fallback.Mul) })
	t.Run("MulFloat64", func(t *testing.T) { test.Universal(t, MulFloat64, fallback.Mul) })
	t.Run("MulInt32", func(t *testing.T) { test.Universal(t, MulInt32, fallback.Mul) })
	t.Run("OrInt32", func(t *testing.T) { test.Universal(t, OrInt32, fallback.Or) })
	t.Run("OrInt64", func(t *testing.T) { test.Universal(t, OrInt64, fallback.Or) })
	t.Run("SubFloat32", func(t *testing.T) { test.Universal(t, SubFloat32, fallback.Sub) })
	t.Run("SubFloat64", func(t *testing.T) { test.Universal(t, SubFloat64, fallback.Sub) })
	t.Run("SubInt32", func(t *testing.T) { test.Universal(t, SubInt32, fallback.Sub) })
	t.Run("SubInt64", func(t *testing.T) { test.Universal(t, SubInt64, fallback.Sub) })
}
