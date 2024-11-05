//go:build arm64
// +build arm64

package neon

func Supported() bool

func AddFloat32(left, right, result []float32) int

func AddInt32(left, right, result []int32) int

func MulFloat32(left, right, result []float32) int

func SubFloat32(left, right, result []float32) int

func SubInt32(left, right, result []int32) int
