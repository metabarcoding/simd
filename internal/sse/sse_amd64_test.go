// +build amd64

package sse

import (
    "fmt"
    "testing"
    "github.com/pehringer/simd/internal/shared"
)

func BenchmarkAddFloat32(b *testing.B) {
    for _, size := range shared.BenchSizes {
        b.Run(fmt.Sprint(size), func(b *testing.B) {
            left := shared.Vector[float32](size)
            right := shared.Vector[float32](size)
            result := shared.Vector[float32](size)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                AddFloat32(left, right, result)
            }
	})
    }
}
