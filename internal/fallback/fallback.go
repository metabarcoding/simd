package fallback

import (
	"github.com/pehringer/simd/internal/data"
)

func Add[T data.Floating | data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] + right[i]
		result[i+1] = left[i+1] + right[i+1]
		result[i+2] = left[i+2] + right[i+2]
		result[i+3] = left[i+3] + right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] + right[i]
	}
	return length
}

func And[T data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] & right[i]
		result[i+1] = left[i+1] & right[i+1]
		result[i+2] = left[i+2] & right[i+2]
		result[i+3] = left[i+3] & right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] & right[i]
	}
	return length
}

func Div[T data.Floating | data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] / right[i]
		result[i+1] = left[i+1] / right[i+1]
		result[i+2] = left[i+2] / right[i+2]
		result[i+3] = left[i+3] / right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] / right[i]
	}
	return length
}

func Max[T data.Floating | data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = max(left[i], right[i])
		result[i+1] = max(left[i+1], right[i+1])
		result[i+2] = max(left[i+2], right[i+2])
		result[i+3] = max(left[i+3], right[i+3])
	}
	for ; i < length; i++ {
		result[i] = max(left[i], right[i])
	}
	return length
}

func Min[T data.Floating | data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = min(left[i], right[i])
		result[i+1] = min(left[i+1], right[i+1])
		result[i+2] = min(left[i+2], right[i+2])
		result[i+3] = min(left[i+3], right[i+3])
	}
	for ; i < length; i++ {
		result[i] = min(left[i], right[i])
	}
	return length
}

func Mul[T data.Floating | data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] * right[i]
		result[i+1] = left[i+1] * right[i+1]
		result[i+2] = left[i+2] * right[i+2]
		result[i+3] = left[i+3] * right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] * right[i]
	}
	return length
}

func Or[T data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] | right[i]
		result[i+1] = left[i+1] | right[i+1]
		result[i+2] = left[i+2] | right[i+2]
		result[i+3] = left[i+3] | right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] | right[i]
	}
	return length
}

func Sub[T data.Floating | data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] - right[i]
		result[i+1] = left[i+1] - right[i+1]
		result[i+2] = left[i+2] - right[i+2]
		result[i+3] = left[i+3] - right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] - right[i]
	}
	return length
}

func Xor[T data.Integer](left, right, result []T) int {
	length := min(len(left), len(right), len(result))
	i := 0
	for ; length-i >= 4; i += 4 {
		result[i] = left[i] ^ right[i]
		result[i+1] = left[i+1] ^ right[i+1]
		result[i+2] = left[i+2] ^ right[i+2]
		result[i+3] = left[i+3] ^ right[i+3]
	}
	for ; i < length; i++ {
		result[i] = left[i] ^ right[i]
	}
	return length
}
