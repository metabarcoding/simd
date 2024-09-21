package simd

type (
    element interface {
        int8 | int16 | int32 | int64
    }
    operation[E element] func(left, right, result []E)
)

var (
    AddInt8  operation[int8]  = addInt[int8]
    AddInt16 operation[int16] = addInt[int16]
    AddInt32 operation[int32] = addInt[int32]
    AddInt64 operation[int64] = addInt[int64]

    SubInt8  operation[int8]  = subInt[int8]
    SubInt16 operation[int16] = subInt[int16]
    SubInt32 operation[int32] = subInt[int32]
    SubInt64 operation[int64] = subInt[int64]
)
