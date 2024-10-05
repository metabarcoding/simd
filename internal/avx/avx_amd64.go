// +build amd64

package avx

func Supported() bool

func AddFloat32(left, right, result []float32) int

func SubFloat32(left, right, result []float32) int

func MulFloat32(left, right, result []float32) int

func DivFloat32(left, right, result []float32) int

