// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool

func sse2Supported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE SUPPORTED")
    }
    if sse2Supported() {
        fmt.Println("SSE2 SUPPORTED")
    }
}
