package fallback

import (
	"github.com/pehringer/simd/internal/shared"
)

func Add[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] + right[i]
		result[i+1] = left[i+1] + right[i+1]
		result[i+2] = left[i+2] + right[i+2]
		result[i+3] = left[i+3] + right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] + right[i]
	}
	return n
}

func And[T shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] & right[i]
		result[i+1] = left[i+1] & right[i+1]
		result[i+2] = left[i+2] & right[i+2]
		result[i+3] = left[i+3] & right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] & right[i]
	}
	return n
}

func Div[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] / right[i]
		result[i+1] = left[i+1] / right[i+1]
		result[i+2] = left[i+2] / right[i+2]
		result[i+3] = left[i+3] / right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] / right[i]
	}
	return n
}

func Mul[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] * right[i]
		result[i+1] = left[i+1] * right[i+1]
		result[i+2] = left[i+2] * right[i+2]
		result[i+3] = left[i+3] * right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] * right[i]
	}
	return n
}

func Or[T shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] | right[i]
		result[i+1] = left[i+1] | right[i+1]
		result[i+2] = left[i+2] | right[i+2]
		result[i+3] = left[i+3] | right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] | right[i]
	}
	return n
}

func Sub[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] - right[i]
		result[i+1] = left[i+1] - right[i+1]
		result[i+2] = left[i+2] - right[i+2]
		result[i+3] = left[i+3] - right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] - right[i]
	}
	return n
}

func Xor[T shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] ^ right[i]
		result[i+1] = left[i+1] ^ right[i+1]
		result[i+2] = left[i+2] ^ right[i+2]
		result[i+3] = left[i+3] ^ right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] ^ right[i]
	}
	return n
}
