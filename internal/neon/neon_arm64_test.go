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
	t.Run("MaxFloat32", func(t *testing.T) { test.Universal(t, MaxFloat32, fallback.Max) })
	t.Run("MinFloat32", func(t *testing.T) { test.Universal(t, MinFloat32, fallback.Min) })
	t.Run("MaxInt32", func(t *testing.T) { test.UniversalGeneric(t, MaxInt32, fallback.Max) })
	t.Run("MinInt32", func(t *testing.T) { test.UniversalGeneric(t, MinInt32, fallback.Min) })
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

func TestMinInt32Simple(t *testing.T) {
	left := []int32{1, -100, 5, -200, 10, -50}
	right := []int32{-100, 1, -200, 5, -50, 10}
	expected := []int32{-100, -100, -200, -200, -50, -50}
	result := make([]int32, len(left))

	length := MinInt32(left, right, result)

	if length != len(left) {
		t.Errorf("Expected length %d, got %d", len(left), length)
	}

	for i := 0; i < len(left); i++ {
		if result[i] != expected[i] {
			t.Errorf("Index %d: left=%d, right=%d, expected min=%d, got=%d",
				i, left[i], right[i], expected[i], result[i])
		}
	}
}
