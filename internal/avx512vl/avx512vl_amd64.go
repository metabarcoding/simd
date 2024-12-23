//go:build amd64
// +build amd64

package avx512vl

func AddFloat32(left, right, result []float32) int

func DivFloat32(left, right, result []float32) int

func MulFloat32(left, right, result []float32) int

func SubFloat32(left, right, result []float32) int
