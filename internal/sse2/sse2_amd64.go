// +build amd64

package sse2

func Supported() bool

func AddInt32(left, right, result []int32) int

func AddInt64(left, right, result []int64) int

func AddFloat64(left, right, result []float64) int

func SubInt32(left, right, result []int32) int

func SubInt64(left, right, result []int64) int

func SubFloat64(left, right, result []float64) int

func MulFloat64(left, right, result []float64) int

func DivFloat64(left, right, result []float64) int
