// +build amd64

package avx2

func Supported() bool

func AddInt32(left, right, result []int32) int

func AddInt64(left, right, result []int64) int

func SubInt32(left, right, result []int32) int

func SubInt64(left, right, result []int64) int
