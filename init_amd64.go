// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool
func sseAddFloat32(left, right, result []float32) int
func sseSubFloat32(left, right, result []float32) int
func sseMulFloat32(left, right, result []float32) int
func sseDivFloat32(left, right, result []float32) int

func sse2Supported() bool

func avxSupported() bool

func avx2Supported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE")
        addFloat32 = sseAddFloat32
        subFloat32 = sseSubFloat32
        mulFloat32 = sseMulFloat32
        divFloat32 = sseDivFloat32
    }
    if sse2Supported() {
        fmt.Println("SSE2")
    }
    if avxSupported() {
        fmt.Println("AVX")
    }
    if avx2Supported() {
        fmt.Println("AVX2")
    }
}
