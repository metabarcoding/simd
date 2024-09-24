// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool
func sseAddFloat32(s []int64) int64

func sse2Supported() bool

func avxSupported() bool

func avx2Supported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE SUPPORTED")
        res := sseAddFloat32([]int64{3, 6, 9, 12, 15, 18, 21})
        fmt.Println("len:", res)
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
