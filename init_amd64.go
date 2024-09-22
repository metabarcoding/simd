// +build amd64

package simd

import (
    "fmt"
)

func sseSupported() bool

func init() {
    if sseSupported() {
        fmt.Println("SSE SUPPORTED")
    } else {
        fmt.Println("SSE NOT SUPPORTED")
    }
}
