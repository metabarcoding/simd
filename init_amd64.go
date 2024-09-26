// +build amd64

package simd

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
        AddFloat32 = sseAddFloat32
        SubFloat32 = sseSubFloat32
        MulFloat32 = sseMulFloat32
        DivFloat32 = sseDivFloat32
    }
    if sse2Supported() {
    }
    if avxSupported() {
    }
    if avx2Supported() {
    }
}
