package simd

type (
    element interface {
        int8 | int16 | int32 | int64
    }
    operation[E element] func(a, b, c []E)
)

func addInt[E element](left, right, sum []E) {
    for i := 0; i < len(sum); i++ {
        sum[i] = left[i] + right[i]
    }
}

func addInt32_x86(left, right, sum []int32) {
    for i := 0; i < len(sum); i++ {
        sum[i] = left[i] + right[i]
    }
}

var (
    AddInt8  operation[int8]  = addInt[int8]
    AddInt16 operation[int16] = addInt[int16]
    AddInt32 operation[int32] = addInt32_x86
    AddInt64 operation[int64] = addInt[int64]
)

func RunAddOp(l, r, s []int32) {
    AddInt32(l, r, s)
}
