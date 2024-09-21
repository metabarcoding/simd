package simd

func minThree(first, second, third int) int {
    min := first
    if second < min {
        min = second
    }
    if third < min {
        min = third
    }
    return min
}

func addInt[E element](left, right, result []E) {
    n := minThree(len(left), len(right), len(result))
    i := 0
    for n-i >= 4 {
        result[i] = left[i] + right[i]
        result[i+1] = left[i+1] + right[i+1]
        result[i+2] = left[i+2] + right[i+2]
        result[i+3] = left[i+3] + right[i+3]
        i += 4
    }
    for i < n {
        result[i] = left[i] + right[i]
        i += 1
    }
}

func subInt[E element](left, right, result []E) {
    n := minThree(len(left), len(right), len(result))
    i := 0
    for n-i >= 4 {
        result[i] = left[i] - right[i]
        result[i+1] = left[i+1] - right[i+1]
        result[i+2] = left[i+2] - right[i+2]
        result[i+3] = left[i+3] - right[i+3]
        i += 4
    }
    for i < n {
        result[i] = left[i] - right[i]
        i += 1
    }
}
