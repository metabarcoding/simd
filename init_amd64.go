// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool
func sseAddFloat32(left, right, result []float32) int

func sse2Supported() bool

func avxSupported() bool

func avx2Supported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE SUPPORTED")
        AddFloat32 = sseAddFloat32
    }
    if sse2Supported() {
        fmt.Println("SSE2 SUPPORTED")
    }
    if avxSupported() {
        fmt.Println("AVX SUPPORTED")
    }
    if avx2Supported() {
        fmt.Println("AVX2 SUPPORTED")
    }
}
