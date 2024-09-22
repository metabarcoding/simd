// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool

func sse2Supported() bool

func avxSupported() bool

func avx2Supported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE SUPPORTED")
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
