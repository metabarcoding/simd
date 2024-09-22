package simd

import (
    "testing"
    "fmt"
)

func TestType(t *testing.T) {
    a := []int32{1, 2, 3, 4}
    b := []int32{3, 4, 5, 6}
    s := []int32{0, 0, 0, 0}
    AddInt32(a, b, s)
    fmt.Println(s)
}
