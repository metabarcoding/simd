package simd

import (
	"testing"
)

func dup[E element](source []E) []E {
    duplicate := make([]E, len(source), cap(source))
    copy(duplicate, source)
    return duplicate
}

func checkSlice[E element](t *testing.T, test, control []E) bool {
    if len(test) != len(control) {
        t.Errorf("lengths not equal")
        return false
    }
    if cap(test) != cap(control) {
        t.Errorf("capacities not equal")
        return false
    }
    for i := 0; i < len(control); i++ {
        if test[i] != control[i] {
            t.Errorf("elements not equal")
            return false
        }
    }
    return true
}

func checkOperation[E element](t *testing.T, test, control operation[E], left, right, result []E) bool {
    testLeft := dup(left)
    testRight := dup(right)
    testResult := dup(result)
    if test(testLeft, testRight, testResult) != control(left, right, result) {
        t.Errorf("operation returned incorrect length")
	return false
    }
    if !checkSlice(t, testLeft, left) {
        return false
    }
    if !checkSlice(t, testRight, right) {
        return false
    }
    if !checkSlice(t, testResult, result) {
        return false
    }
    return true
}

var (
    counter uint8
)

func vec[E element](length int) []E {
    vector := make([]E, length)
    for i := 0; i < length; i++ {
        counter++
        if counter == 0 {
            counter = 1
        }
        vector[i] = E(counter)
    }
    return vector
}

func TestAddInt8(t *testing.T) {
    checkOperation(t, AddInt8, add, []int8{}, []int8{}, []int8{})
    checkOperation(t, AddInt8, add, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, AddInt8, add, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, AddInt8, add, vec[int8](103), vec[int8](109), vec[int8](107))
    checkOperation(t, AddInt8, add, vec[int8](1001), vec[int8](1003), vec[int8](1007))
}

func TestAddInt16(t *testing.T) {
    checkOperation(t, AddInt16, add, []int16{}, []int16{}, []int16{})
    checkOperation(t, AddInt16, add, vec[int16](1), vec[int16](3), vec[int16](9))
    checkOperation(t, AddInt16, add, vec[int16](19), vec[int16](11), vec[int16](13))
    checkOperation(t, AddInt16, add, vec[int16](103), vec[int16](109), vec[int16](107))
    checkOperation(t, AddInt16, add, vec[int16](1001), vec[int16](1003), vec[int16](1007))
}

func TestAddInt32(t *testing.T) {
    checkOperation(t, AddInt32, add, []int32{}, []int32{}, []int32{})
    checkOperation(t, AddInt32, add, vec[int32](1), vec[int32](3), vec[int32](9))
    checkOperation(t, AddInt32, add, vec[int32](19), vec[int32](11), vec[int32](13))
    checkOperation(t, AddInt32, add, vec[int32](103), vec[int32](109), vec[int32](107))
    checkOperation(t, AddInt32, add, vec[int32](1001), vec[int32](1003), vec[int32](1007))
}

func TestAddInt64(t *testing.T) {
    checkOperation(t, AddInt64, add, []int64{}, []int64{}, []int64{})
    checkOperation(t, AddInt64, add, vec[int64](1), vec[int64](3), vec[int64](9))
    checkOperation(t, AddInt64, add, vec[int64](19), vec[int64](11), vec[int64](13))
    checkOperation(t, AddInt64, add, vec[int64](103), vec[int64](109), vec[int64](107))
    checkOperation(t, AddInt64, add, vec[int64](1001), vec[int64](1003), vec[int64](1007))
}

func TestAddFloat32(t *testing.T) {
    checkOperation(t, AddFloat32, add, []float32{}, []float32{}, []float32{})
    checkOperation(t, AddFloat32, add, vec[float32](1), vec[float32](3), vec[float32](9))
    checkOperation(t, AddFloat32, add, vec[float32](19), vec[float32](11), vec[float32](13))
    checkOperation(t, AddFloat32, add, vec[float32](103), vec[float32](109), vec[float32](107))
    checkOperation(t, AddFloat32, add, vec[float32](1001), vec[float32](1003), vec[float32](1007))
}

func TestAddFloat64(t *testing.T) {
    checkOperation(t, AddFloat64, add, []float64{}, []float64{}, []float64{})
    checkOperation(t, AddFloat64, add, vec[float64](1), vec[float64](3), vec[float64](9))
    checkOperation(t, AddFloat64, add, vec[float64](19), vec[float64](11), vec[float64](13))
    checkOperation(t, AddFloat64, add, vec[float64](103), vec[float64](109), vec[float64](107))
    checkOperation(t, AddFloat64, add, vec[float64](1001), vec[float64](1003), vec[float64](1007))
}

func TestSubInt8(t *testing.T) {
    checkOperation(t, SubInt8, sub, []int8{}, []int8{}, []int8{})
    checkOperation(t, SubInt8, sub, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, SubInt8, sub, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, SubInt8, sub, vec[int8](103), vec[int8](109), vec[int8](107))
    checkOperation(t, SubInt8, sub, vec[int8](1001), vec[int8](1003), vec[int8](1007))
}

func TestSubInt16(t *testing.T) {
    checkOperation(t, SubInt16, sub, []int16{}, []int16{}, []int16{})
    checkOperation(t, SubInt16, sub, vec[int16](1), vec[int16](3), vec[int16](9))
    checkOperation(t, SubInt16, sub, vec[int16](19), vec[int16](11), vec[int16](13))
    checkOperation(t, SubInt16, sub, vec[int16](103), vec[int16](109), vec[int16](107))
    checkOperation(t, SubInt16, sub, vec[int16](1001), vec[int16](1003), vec[int16](1007))
}

func TestSubInt32(t *testing.T) {
    checkOperation(t, SubInt32, sub, []int32{}, []int32{}, []int32{})
    checkOperation(t, SubInt32, sub, vec[int32](1), vec[int32](3), vec[int32](9))
    checkOperation(t, SubInt32, sub, vec[int32](19), vec[int32](11), vec[int32](13))
    checkOperation(t, SubInt32, sub, vec[int32](103), vec[int32](109), vec[int32](107))
    checkOperation(t, SubInt32, sub, vec[int32](1001), vec[int32](1003), vec[int32](1007))
}

func TestSubInt64(t *testing.T) {
    checkOperation(t, SubInt64, sub, []int64{}, []int64{}, []int64{})
    checkOperation(t, SubInt64, sub, vec[int64](1), vec[int64](3), vec[int64](9))
    checkOperation(t, SubInt64, sub, vec[int64](19), vec[int64](11), vec[int64](13))
    checkOperation(t, SubInt64, sub, vec[int64](103), vec[int64](109), vec[int64](107))
    checkOperation(t, SubInt64, sub, vec[int64](1001), vec[int64](1003), vec[int64](1007))
}

func TestSubFloat32(t *testing.T) {
    checkOperation(t, SubFloat32, sub, []float32{}, []float32{}, []float32{})
    checkOperation(t, SubFloat32, sub, vec[float32](1), vec[float32](3), vec[float32](9))
    checkOperation(t, SubFloat32, sub, vec[float32](19), vec[float32](11), vec[float32](13))
    checkOperation(t, SubFloat32, sub, vec[float32](103), vec[float32](109), vec[float32](107))
    checkOperation(t, SubFloat32, sub, vec[float32](1001), vec[float32](1003), vec[float32](1007))
}

func TestSubFloat64(t *testing.T) {
    checkOperation(t, SubFloat64, sub, []float64{}, []float64{}, []float64{})
    checkOperation(t, SubFloat64, sub, vec[float64](1), vec[float64](3), vec[float64](9))
    checkOperation(t, SubFloat64, sub, vec[float64](19), vec[float64](11), vec[float64](13))
    checkOperation(t, SubFloat64, sub, vec[float64](103), vec[float64](109), vec[float64](107))
    checkOperation(t, SubFloat64, sub, vec[float64](1001), vec[float64](1003), vec[float64](1007))
}

func TestMulInt8(t *testing.T) {
    checkOperation(t, MulInt8, mul, []int8{}, []int8{}, []int8{})
    checkOperation(t, MulInt8, mul, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, MulInt8, mul, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, MulInt8, mul, vec[int8](103), vec[int8](109), vec[int8](107))
    checkOperation(t, MulInt8, mul, vec[int8](1001), vec[int8](1003), vec[int8](1007))
}

func TestMulInt16(t *testing.T) {
    checkOperation(t, MulInt16, mul, []int16{}, []int16{}, []int16{})
    checkOperation(t, MulInt16, mul, vec[int16](1), vec[int16](3), vec[int16](9))
    checkOperation(t, MulInt16, mul, vec[int16](19), vec[int16](11), vec[int16](13))
    checkOperation(t, MulInt16, mul, vec[int16](103), vec[int16](109), vec[int16](107))
    checkOperation(t, MulInt16, mul, vec[int16](1001), vec[int16](1003), vec[int16](1007))
}

func TestMulInt32(t *testing.T) {
    checkOperation(t, MulInt32, mul, []int32{}, []int32{}, []int32{})
    checkOperation(t, MulInt32, mul, vec[int32](1), vec[int32](3), vec[int32](9))
    checkOperation(t, MulInt32, mul, vec[int32](19), vec[int32](11), vec[int32](13))
    checkOperation(t, MulInt32, mul, vec[int32](103), vec[int32](109), vec[int32](107))
    checkOperation(t, MulInt32, mul, vec[int32](1001), vec[int32](1003), vec[int32](1007))
}

func TestMulInt64(t *testing.T) {
    checkOperation(t, MulInt64, mul, []int64{}, []int64{}, []int64{})
    checkOperation(t, MulInt64, mul, vec[int64](1), vec[int64](3), vec[int64](9))
    checkOperation(t, MulInt64, mul, vec[int64](19), vec[int64](11), vec[int64](13))
    checkOperation(t, MulInt64, mul, vec[int64](103), vec[int64](109), vec[int64](107))
    checkOperation(t, MulInt64, mul, vec[int64](1001), vec[int64](1003), vec[int64](1007))
}

func TestMulFloat32(t *testing.T) {
    checkOperation(t, MulFloat32, mul, []float32{}, []float32{}, []float32{})
    checkOperation(t, MulFloat32, mul, vec[float32](1), vec[float32](3), vec[float32](9))
    checkOperation(t, MulFloat32, mul, vec[float32](19), vec[float32](11), vec[float32](13))
    checkOperation(t, MulFloat32, mul, vec[float32](103), vec[float32](109), vec[float32](107))
    checkOperation(t, MulFloat32, mul, vec[float32](1001), vec[float32](1003), vec[float32](1007))
}

func TestMulFloat64(t *testing.T) {
    checkOperation(t, MulFloat64, mul, []float64{}, []float64{}, []float64{})
    checkOperation(t, MulFloat64, mul, vec[float64](1), vec[float64](3), vec[float64](9))
    checkOperation(t, MulFloat64, mul, vec[float64](19), vec[float64](11), vec[float64](13))
    checkOperation(t, MulFloat64, mul, vec[float64](103), vec[float64](109), vec[float64](107))
    checkOperation(t, MulFloat64, mul, vec[float64](1001), vec[float64](1003), vec[float64](1007))
}



func TestDivInt8(t *testing.T) {
    checkOperation(t, DivInt8, div, []int8{}, []int8{}, []int8{})
    checkOperation(t, DivInt8, div, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, DivInt8, div, vec[int8](19), vec[int8](11), vec[int8](13))
    checkOperation(t, DivInt8, div, vec[int8](103), vec[int8](109), vec[int8](107))
    checkOperation(t, DivInt8, div, vec[int8](1001), vec[int8](1003), vec[int8](1007))
}

func TestDivInt16(t *testing.T) {
    checkOperation(t, DivInt16, div, []int16{}, []int16{}, []int16{})
    checkOperation(t, DivInt16, div, vec[int16](1), vec[int16](3), vec[int16](9))
    checkOperation(t, DivInt16, div, vec[int16](19), vec[int16](11), vec[int16](13))
    checkOperation(t, DivInt16, div, vec[int16](103), vec[int16](109), vec[int16](107))
    checkOperation(t, DivInt16, div, vec[int16](1001), vec[int16](1003), vec[int16](1007))
}

func TestDivInt32(t *testing.T) {
    checkOperation(t, DivInt32, div, []int32{}, []int32{}, []int32{})
    checkOperation(t, DivInt32, div, vec[int32](1), vec[int32](3), vec[int32](9))
    checkOperation(t, DivInt32, div, vec[int32](19), vec[int32](11), vec[int32](13))
    checkOperation(t, DivInt32, div, vec[int32](103), vec[int32](109), vec[int32](107))
    checkOperation(t, DivInt32, div, vec[int32](1001), vec[int32](1003), vec[int32](1007))
}

func TestDivInt64(t *testing.T) {
    checkOperation(t, DivInt64, div, []int64{}, []int64{}, []int64{})
    checkOperation(t, DivInt64, div, vec[int64](1), vec[int64](3), vec[int64](9))
    checkOperation(t, DivInt64, div, vec[int64](19), vec[int64](11), vec[int64](13))
    checkOperation(t, DivInt64, div, vec[int64](103), vec[int64](109), vec[int64](107))
    checkOperation(t, DivInt64, div, vec[int64](1001), vec[int64](1003), vec[int64](1007))
}

func TestDivFloat32(t *testing.T) {
    checkOperation(t, DivFloat32, div, []float32{}, []float32{}, []float32{})
    checkOperation(t, DivFloat32, div, vec[float32](1), vec[float32](3), vec[float32](9))
    checkOperation(t, DivFloat32, div, vec[float32](19), vec[float32](11), vec[float32](13))
    checkOperation(t, DivFloat32, div, vec[float32](103), vec[float32](109), vec[float32](107))
    checkOperation(t, DivFloat32, div, vec[float32](1001), vec[float32](1003), vec[float32](1007))
}

func TestDivFloat64(t *testing.T) {
    checkOperation(t, DivFloat64, div, []float64{}, []float64{}, []float64{})
    checkOperation(t, DivFloat64, div, vec[float64](1), vec[float64](3), vec[float64](9))
    checkOperation(t, DivFloat64, div, vec[float64](19), vec[float64](11), vec[float64](13))
    checkOperation(t, DivFloat64, div, vec[float64](103), vec[float64](109), vec[float64](107))
    checkOperation(t, DivFloat64, div, vec[float64](1001), vec[float64](1003), vec[float64](1007))
}
