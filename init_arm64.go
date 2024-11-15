//go:build arm64
// +build arm64

package simd

import (
	"golang.org/x/sys/cpu"
	"github.com/pehringer/simd/internal/neon"
)

func init() {
	if cpu.ARM64.HasASIMD {
		addFloat32 = neon.AddFloat32
		addFloat64 = neon.AddFloat64
		addInt32 = neon.AddInt32
		addInt64 = neon.AddInt64
		andInt32 = neon.AndInt32
		mulFloat32 = neon.MulFloat32
		mulFloat64 = neon.MulFloat64
		mulInt32 = neon.MulInt32
		subFloat32 = neon.SubFloat32
		subFloat64 = neon.SubFloat64
		subInt32 = neon.SubInt32
		subInt64 = neon.SubInt64
	}
}
