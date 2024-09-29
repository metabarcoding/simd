package simd

import (
	"testing"
)

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
    testLeft := make([]E, len(left), cap(left))
    copy(testLeft, left)
    testRight := make([]E, len(right), cap(right))
    copy(testRight, right)
    testResult := make([]E, len(result), cap(result))
    copy(testResult, result)
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

func vector[E element](length int) []E {
    elements := make([]E, length)
    for i := 0; i < length; i++ {
        counter++
        if counter == 0 {
            counter = 1
        }
        elements[i] = E(counter)
    }
    return elements
}

func TestAddInt8Zero(t *testing.T) {
    checkOperation(t, AddInt8, goAdd, []int8{}, []int8{}, []int8{})
}

func TestAddInt8Prime(t *testing.T) {
    checkOperation(t, AddInt8, goAdd, vector[int8](11), vector[int8](13), vector[int8](17))
    checkOperation(t, AddInt8, goAdd, vector[int8](29), vector[int8](19), vector[int8](23))
    checkOperation(t, AddInt8, goAdd, vector[int8](37), vector[int8](41), vector[int8](31))
    checkOperation(t, AddInt8, goAdd, vector[int8](43), vector[int8](47), vector[int8](53))
    checkOperation(t, AddInt8, goAdd, vector[int8](67), vector[int8](59), vector[int8](61))
    checkOperation(t, AddInt8, goAdd, vector[int8](73), vector[int8](79), vector[int8](71))
}

func TestAddInt8PowerTwo(t *testing.T) {
    checkOperation(t, AddInt8, goAdd, vector[int8](1024), vector[int8](1024), vector[int8](1024))
    checkOperation(t, AddInt8, goAdd, vector[int8](2048), vector[int8](2048), vector[int8](2048))
    checkOperation(t, AddInt8, goAdd, vector[int8](4096), vector[int8](4096), vector[int8](4096))
    checkOperation(t, AddInt8, goAdd, vector[int8](8192), vector[int8](8192), vector[int8](8192))
}

func TestAddInt16Zero(t *testing.T) {
    checkOperation(t, AddInt16, goAdd, []int16{}, []int16{}, []int16{})
}

func TestAddInt16Prime(t *testing.T) {
    checkOperation(t, AddInt16, goAdd, vector[int16](11), vector[int16](13), vector[int16](17))
    checkOperation(t, AddInt16, goAdd, vector[int16](29), vector[int16](19), vector[int16](23))
    checkOperation(t, AddInt16, goAdd, vector[int16](37), vector[int16](41), vector[int16](31))
    checkOperation(t, AddInt16, goAdd, vector[int16](43), vector[int16](47), vector[int16](53))
    checkOperation(t, AddInt16, goAdd, vector[int16](67), vector[int16](59), vector[int16](61))
    checkOperation(t, AddInt16, goAdd, vector[int16](73), vector[int16](79), vector[int16](71))
}

func TestAddInt16PowerTwo(t *testing.T) {
    checkOperation(t, AddInt16, goAdd, vector[int16](1024), vector[int16](1024), vector[int16](1024))
    checkOperation(t, AddInt16, goAdd, vector[int16](2048), vector[int16](2048), vector[int16](2048))
    checkOperation(t, AddInt16, goAdd, vector[int16](4096), vector[int16](4096), vector[int16](4096))
    checkOperation(t, AddInt16, goAdd, vector[int16](8192), vector[int16](8192), vector[int16](8192))
}

func TestAddInt32Zero(t *testing.T) {
    checkOperation(t, AddInt32, goAdd, []int32{}, []int32{}, []int32{})
}

func TestAddInt32Prime(t *testing.T) {
    checkOperation(t, AddInt32, goAdd, vector[int32](11), vector[int32](13), vector[int32](17))
    checkOperation(t, AddInt32, goAdd, vector[int32](29), vector[int32](19), vector[int32](23))
    checkOperation(t, AddInt32, goAdd, vector[int32](37), vector[int32](41), vector[int32](31))
    checkOperation(t, AddInt32, goAdd, vector[int32](43), vector[int32](47), vector[int32](53))
    checkOperation(t, AddInt32, goAdd, vector[int32](67), vector[int32](59), vector[int32](61))
    checkOperation(t, AddInt32, goAdd, vector[int32](73), vector[int32](79), vector[int32](71))
}

func TestAddInt32PowerTwo(t *testing.T) {
    checkOperation(t, AddInt32, goAdd, vector[int32](1024), vector[int32](1024), vector[int32](1024))
    checkOperation(t, AddInt32, goAdd, vector[int32](2048), vector[int32](2048), vector[int32](2048))
    checkOperation(t, AddInt32, goAdd, vector[int32](4096), vector[int32](4096), vector[int32](4096))
    checkOperation(t, AddInt32, goAdd, vector[int32](8192), vector[int32](8192), vector[int32](8192))
}

func TestAddInt64Zero(t *testing.T) {
    checkOperation(t, AddInt64, goAdd, []int64{}, []int64{}, []int64{})
}

func TestAddInt64Prime(t *testing.T) {
    checkOperation(t, AddInt64, goAdd, vector[int64](11), vector[int64](13), vector[int64](17))
    checkOperation(t, AddInt64, goAdd, vector[int64](29), vector[int64](19), vector[int64](23))
    checkOperation(t, AddInt64, goAdd, vector[int64](37), vector[int64](41), vector[int64](31))
    checkOperation(t, AddInt64, goAdd, vector[int64](43), vector[int64](47), vector[int64](53))
    checkOperation(t, AddInt64, goAdd, vector[int64](67), vector[int64](59), vector[int64](61))
    checkOperation(t, AddInt64, goAdd, vector[int64](73), vector[int64](79), vector[int64](71))
}

func TestAddInt64PowerTwo(t *testing.T) {
    checkOperation(t, AddInt64, goAdd, vector[int64](1024), vector[int64](1024), vector[int64](1024))
    checkOperation(t, AddInt64, goAdd, vector[int64](2048), vector[int64](2048), vector[int64](2048))
    checkOperation(t, AddInt64, goAdd, vector[int64](4096), vector[int64](4096), vector[int64](4096))
    checkOperation(t, AddInt64, goAdd, vector[int64](8192), vector[int64](8192), vector[int64](8192))
}

func TestAddFloat32Zero(t *testing.T) {
    checkOperation(t, AddFloat32, goAdd, []float32{}, []float32{}, []float32{})
}

func TestAddFloat32Prime(t *testing.T) {
    checkOperation(t, AddFloat32, goAdd, vector[float32](11), vector[float32](13), vector[float32](17))
    checkOperation(t, AddFloat32, goAdd, vector[float32](29), vector[float32](19), vector[float32](23))
    checkOperation(t, AddFloat32, goAdd, vector[float32](37), vector[float32](41), vector[float32](31))
    checkOperation(t, AddFloat32, goAdd, vector[float32](43), vector[float32](47), vector[float32](53))
    checkOperation(t, AddFloat32, goAdd, vector[float32](67), vector[float32](59), vector[float32](61))
    checkOperation(t, AddFloat32, goAdd, vector[float32](73), vector[float32](79), vector[float32](71))
}

func TestAddFloat32PowerTwo(t *testing.T) {
    checkOperation(t, AddFloat32, goAdd, vector[float32](1024), vector[float32](1024), vector[float32](1024))
    checkOperation(t, AddFloat32, goAdd, vector[float32](2048), vector[float32](2048), vector[float32](2048))
    checkOperation(t, AddFloat32, goAdd, vector[float32](4096), vector[float32](4096), vector[float32](4096))
    checkOperation(t, AddFloat32, goAdd, vector[float32](8192), vector[float32](8192), vector[float32](8192))
}

func TestAddFloat64Zero(t *testing.T) {
    checkOperation(t, AddFloat64, goAdd, []float64{}, []float64{}, []float64{})
}

func TestAddFloat64Prime(t *testing.T) {
    checkOperation(t, AddFloat64, goAdd, vector[float64](11), vector[float64](13), vector[float64](17))
    checkOperation(t, AddFloat64, goAdd, vector[float64](29), vector[float64](19), vector[float64](23))
    checkOperation(t, AddFloat64, goAdd, vector[float64](37), vector[float64](41), vector[float64](31))
    checkOperation(t, AddFloat64, goAdd, vector[float64](43), vector[float64](47), vector[float64](53))
    checkOperation(t, AddFloat64, goAdd, vector[float64](67), vector[float64](59), vector[float64](61))
    checkOperation(t, AddFloat64, goAdd, vector[float64](73), vector[float64](79), vector[float64](71))
}

func TestAddFloat64PowerTwo(t *testing.T) {
    checkOperation(t, AddFloat64, goAdd, vector[float64](1024), vector[float64](1024), vector[float64](1024))
    checkOperation(t, AddFloat64, goAdd, vector[float64](2048), vector[float64](2048), vector[float64](2048))
    checkOperation(t, AddFloat64, goAdd, vector[float64](4096), vector[float64](4096), vector[float64](4096))
    checkOperation(t, AddFloat64, goAdd, vector[float64](8192), vector[float64](8192), vector[float64](8192))
}

func TestSubInt8Zero(t *testing.T) {
    checkOperation(t, SubInt8, goSub, []int8{}, []int8{}, []int8{})
}

func TestSubInt8Prime(t *testing.T) {
    checkOperation(t, SubInt8, goSub, vector[int8](11), vector[int8](13), vector[int8](17))
    checkOperation(t, SubInt8, goSub, vector[int8](29), vector[int8](19), vector[int8](23))
    checkOperation(t, SubInt8, goSub, vector[int8](37), vector[int8](41), vector[int8](31))
    checkOperation(t, SubInt8, goSub, vector[int8](43), vector[int8](47), vector[int8](53))
    checkOperation(t, SubInt8, goSub, vector[int8](67), vector[int8](59), vector[int8](61))
    checkOperation(t, SubInt8, goSub, vector[int8](73), vector[int8](79), vector[int8](71))
}

func TestSubInt8PowerTwo(t *testing.T) {
    checkOperation(t, SubInt8, goSub, vector[int8](1024), vector[int8](1024), vector[int8](1024))
    checkOperation(t, SubInt8, goSub, vector[int8](2048), vector[int8](2048), vector[int8](2048))
    checkOperation(t, SubInt8, goSub, vector[int8](4096), vector[int8](4096), vector[int8](4096))
    checkOperation(t, SubInt8, goSub, vector[int8](8192), vector[int8](8192), vector[int8](8192))
}

func TestSubInt16Zero(t *testing.T) {
    checkOperation(t, SubInt16, goSub, []int16{}, []int16{}, []int16{})
}

func TestSubInt16Prime(t *testing.T) {
    checkOperation(t, SubInt16, goSub, vector[int16](11), vector[int16](13), vector[int16](17))
    checkOperation(t, SubInt16, goSub, vector[int16](29), vector[int16](19), vector[int16](23))
    checkOperation(t, SubInt16, goSub, vector[int16](37), vector[int16](41), vector[int16](31))
    checkOperation(t, SubInt16, goSub, vector[int16](43), vector[int16](47), vector[int16](53))
    checkOperation(t, SubInt16, goSub, vector[int16](67), vector[int16](59), vector[int16](61))
    checkOperation(t, SubInt16, goSub, vector[int16](73), vector[int16](79), vector[int16](71))
}

func TestSubInt16PowerTwo(t *testing.T) {
    checkOperation(t, SubInt16, goSub, vector[int16](1024), vector[int16](1024), vector[int16](1024))
    checkOperation(t, SubInt16, goSub, vector[int16](2048), vector[int16](2048), vector[int16](2048))
    checkOperation(t, SubInt16, goSub, vector[int16](4096), vector[int16](4096), vector[int16](4096))
    checkOperation(t, SubInt16, goSub, vector[int16](8192), vector[int16](8192), vector[int16](8192))
}

func TestSubInt32Zero(t *testing.T) {
    checkOperation(t, SubInt32, goSub, []int32{}, []int32{}, []int32{})
}

func TestSubInt32Prime(t *testing.T) {
    checkOperation(t, SubInt32, goSub, vector[int32](11), vector[int32](13), vector[int32](17))
    checkOperation(t, SubInt32, goSub, vector[int32](29), vector[int32](19), vector[int32](23))
    checkOperation(t, SubInt32, goSub, vector[int32](37), vector[int32](41), vector[int32](31))
    checkOperation(t, SubInt32, goSub, vector[int32](43), vector[int32](47), vector[int32](53))
    checkOperation(t, SubInt32, goSub, vector[int32](67), vector[int32](59), vector[int32](61))
    checkOperation(t, SubInt32, goSub, vector[int32](73), vector[int32](79), vector[int32](71))
}

func TestSubInt32PowerTwo(t *testing.T) {
    checkOperation(t, SubInt32, goSub, vector[int32](1024), vector[int32](1024), vector[int32](1024))
    checkOperation(t, SubInt32, goSub, vector[int32](2048), vector[int32](2048), vector[int32](2048))
    checkOperation(t, SubInt32, goSub, vector[int32](4096), vector[int32](4096), vector[int32](4096))
    checkOperation(t, SubInt32, goSub, vector[int32](8192), vector[int32](8192), vector[int32](8192))
}

func TestSubInt64Zero(t *testing.T) {
    checkOperation(t, SubInt64, goSub, []int64{}, []int64{}, []int64{})
}

func TestSubInt64Prime(t *testing.T) {
    checkOperation(t, SubInt64, goSub, vector[int64](11), vector[int64](13), vector[int64](17))
    checkOperation(t, SubInt64, goSub, vector[int64](29), vector[int64](19), vector[int64](23))
    checkOperation(t, SubInt64, goSub, vector[int64](37), vector[int64](41), vector[int64](31))
    checkOperation(t, SubInt64, goSub, vector[int64](43), vector[int64](47), vector[int64](53))
    checkOperation(t, SubInt64, goSub, vector[int64](67), vector[int64](59), vector[int64](61))
    checkOperation(t, SubInt64, goSub, vector[int64](73), vector[int64](79), vector[int64](71))
}

func TestSubInt64PowerTwo(t *testing.T) {
    checkOperation(t, SubInt64, goSub, vector[int64](1024), vector[int64](1024), vector[int64](1024))
    checkOperation(t, SubInt64, goSub, vector[int64](2048), vector[int64](2048), vector[int64](2048))
    checkOperation(t, SubInt64, goSub, vector[int64](4096), vector[int64](4096), vector[int64](4096))
    checkOperation(t, SubInt64, goSub, vector[int64](8192), vector[int64](8192), vector[int64](8192))
}

func TestSubFloat32Zero(t *testing.T) {
    checkOperation(t, SubFloat32, goSub, []float32{}, []float32{}, []float32{})
}

func TestSubFloat32Prime(t *testing.T) {
    checkOperation(t, SubFloat32, goSub, vector[float32](11), vector[float32](13), vector[float32](17))
    checkOperation(t, SubFloat32, goSub, vector[float32](29), vector[float32](19), vector[float32](23))
    checkOperation(t, SubFloat32, goSub, vector[float32](37), vector[float32](41), vector[float32](31))
    checkOperation(t, SubFloat32, goSub, vector[float32](43), vector[float32](47), vector[float32](53))
    checkOperation(t, SubFloat32, goSub, vector[float32](67), vector[float32](59), vector[float32](61))
    checkOperation(t, SubFloat32, goSub, vector[float32](73), vector[float32](79), vector[float32](71))
}

func TestSubFloat32PowerTwo(t *testing.T) {
    checkOperation(t, SubFloat32, goSub, vector[float32](1024), vector[float32](1024), vector[float32](1024))
    checkOperation(t, SubFloat32, goSub, vector[float32](2048), vector[float32](2048), vector[float32](2048))
    checkOperation(t, SubFloat32, goSub, vector[float32](4096), vector[float32](4096), vector[float32](4096))
    checkOperation(t, SubFloat32, goSub, vector[float32](8192), vector[float32](8192), vector[float32](8192))
}

func TestSubFloat64Zero(t *testing.T) {
    checkOperation(t, SubFloat64, goSub, []float64{}, []float64{}, []float64{})
}

func TestSubFloat64Prime(t *testing.T) {
    checkOperation(t, SubFloat64, goSub, vector[float64](11), vector[float64](13), vector[float64](17))
    checkOperation(t, SubFloat64, goSub, vector[float64](29), vector[float64](19), vector[float64](23))
    checkOperation(t, SubFloat64, goSub, vector[float64](37), vector[float64](41), vector[float64](31))
    checkOperation(t, SubFloat64, goSub, vector[float64](43), vector[float64](47), vector[float64](53))
    checkOperation(t, SubFloat64, goSub, vector[float64](67), vector[float64](59), vector[float64](61))
    checkOperation(t, SubFloat64, goSub, vector[float64](73), vector[float64](79), vector[float64](71))
}

func TestSubFloat64PowerTwo(t *testing.T) {
    checkOperation(t, SubFloat64, goSub, vector[float64](1024), vector[float64](1024), vector[float64](1024))
    checkOperation(t, SubFloat64, goSub, vector[float64](2048), vector[float64](2048), vector[float64](2048))
    checkOperation(t, SubFloat64, goSub, vector[float64](4096), vector[float64](4096), vector[float64](4096))
    checkOperation(t, SubFloat64, goSub, vector[float64](8192), vector[float64](8192), vector[float64](8192))
}

func TestMulInt8Zero(t *testing.T) {
    checkOperation(t, MulInt8, goMul, []int8{}, []int8{}, []int8{})
}

func TestMulInt8Prime(t *testing.T) {
    checkOperation(t, MulInt8, goMul, vector[int8](11), vector[int8](13), vector[int8](17))
    checkOperation(t, MulInt8, goMul, vector[int8](29), vector[int8](19), vector[int8](23))
    checkOperation(t, MulInt8, goMul, vector[int8](37), vector[int8](41), vector[int8](31))
    checkOperation(t, MulInt8, goMul, vector[int8](43), vector[int8](47), vector[int8](53))
    checkOperation(t, MulInt8, goMul, vector[int8](67), vector[int8](59), vector[int8](61))
    checkOperation(t, MulInt8, goMul, vector[int8](73), vector[int8](79), vector[int8](71))
}

func TestMulInt8PowerTwo(t *testing.T) {
    checkOperation(t, MulInt8, goMul, vector[int8](1024), vector[int8](1024), vector[int8](1024))
    checkOperation(t, MulInt8, goMul, vector[int8](2048), vector[int8](2048), vector[int8](2048))
    checkOperation(t, MulInt8, goMul, vector[int8](4096), vector[int8](4096), vector[int8](4096))
    checkOperation(t, MulInt8, goMul, vector[int8](8192), vector[int8](8192), vector[int8](8192))
}

func TestMulInt16Zero(t *testing.T) {
    checkOperation(t, MulInt16, goMul, []int16{}, []int16{}, []int16{})
}

func TestMulInt16Prime(t *testing.T) {
    checkOperation(t, MulInt16, goMul, vector[int16](11), vector[int16](13), vector[int16](17))
    checkOperation(t, MulInt16, goMul, vector[int16](29), vector[int16](19), vector[int16](23))
    checkOperation(t, MulInt16, goMul, vector[int16](37), vector[int16](41), vector[int16](31))
    checkOperation(t, MulInt16, goMul, vector[int16](43), vector[int16](47), vector[int16](53))
    checkOperation(t, MulInt16, goMul, vector[int16](67), vector[int16](59), vector[int16](61))
    checkOperation(t, MulInt16, goMul, vector[int16](73), vector[int16](79), vector[int16](71))
}

func TestMulInt16PowerTwo(t *testing.T) {
    checkOperation(t, MulInt16, goMul, vector[int16](1024), vector[int16](1024), vector[int16](1024))
    checkOperation(t, MulInt16, goMul, vector[int16](2048), vector[int16](2048), vector[int16](2048))
    checkOperation(t, MulInt16, goMul, vector[int16](4096), vector[int16](4096), vector[int16](4096))
    checkOperation(t, MulInt16, goMul, vector[int16](8192), vector[int16](8192), vector[int16](8192))
}

func TestMulInt32Zero(t *testing.T) {
    checkOperation(t, MulInt32, goMul, []int32{}, []int32{}, []int32{})
}

func TestMulInt32Prime(t *testing.T) {
    checkOperation(t, MulInt32, goMul, vector[int32](11), vector[int32](13), vector[int32](17))
    checkOperation(t, MulInt32, goMul, vector[int32](29), vector[int32](19), vector[int32](23))
    checkOperation(t, MulInt32, goMul, vector[int32](37), vector[int32](41), vector[int32](31))
    checkOperation(t, MulInt32, goMul, vector[int32](43), vector[int32](47), vector[int32](53))
    checkOperation(t, MulInt32, goMul, vector[int32](67), vector[int32](59), vector[int32](61))
    checkOperation(t, MulInt32, goMul, vector[int32](73), vector[int32](79), vector[int32](71))
}

func TestMulInt32PowerTwo(t *testing.T) {
    checkOperation(t, MulInt32, goMul, vector[int32](1024), vector[int32](1024), vector[int32](1024))
    checkOperation(t, MulInt32, goMul, vector[int32](2048), vector[int32](2048), vector[int32](2048))
    checkOperation(t, MulInt32, goMul, vector[int32](4096), vector[int32](4096), vector[int32](4096))
    checkOperation(t, MulInt32, goMul, vector[int32](8192), vector[int32](8192), vector[int32](8192))
}

func TestMulInt64Zero(t *testing.T) {
    checkOperation(t, MulInt64, goMul, []int64{}, []int64{}, []int64{})
}

func TestMulInt64Prime(t *testing.T) {
    checkOperation(t, MulInt64, goMul, vector[int64](11), vector[int64](13), vector[int64](17))
    checkOperation(t, MulInt64, goMul, vector[int64](29), vector[int64](19), vector[int64](23))
    checkOperation(t, MulInt64, goMul, vector[int64](37), vector[int64](41), vector[int64](31))
    checkOperation(t, MulInt64, goMul, vector[int64](43), vector[int64](47), vector[int64](53))
    checkOperation(t, MulInt64, goMul, vector[int64](67), vector[int64](59), vector[int64](61))
    checkOperation(t, MulInt64, goMul, vector[int64](73), vector[int64](79), vector[int64](71))
}

func TestMulInt64PowerTwo(t *testing.T) {
    checkOperation(t, MulInt64, goMul, vector[int64](1024), vector[int64](1024), vector[int64](1024))
    checkOperation(t, MulInt64, goMul, vector[int64](2048), vector[int64](2048), vector[int64](2048))
    checkOperation(t, MulInt64, goMul, vector[int64](4096), vector[int64](4096), vector[int64](4096))
    checkOperation(t, MulInt64, goMul, vector[int64](8192), vector[int64](8192), vector[int64](8192))
}

func TestMulFloat32Zero(t *testing.T) {
    checkOperation(t, MulFloat32, goMul, []float32{}, []float32{}, []float32{})
}

func TestMulFloat32Prime(t *testing.T) {
    checkOperation(t, MulFloat32, goMul, vector[float32](11), vector[float32](13), vector[float32](17))
    checkOperation(t, MulFloat32, goMul, vector[float32](29), vector[float32](19), vector[float32](23))
    checkOperation(t, MulFloat32, goMul, vector[float32](37), vector[float32](41), vector[float32](31))
    checkOperation(t, MulFloat32, goMul, vector[float32](43), vector[float32](47), vector[float32](53))
    checkOperation(t, MulFloat32, goMul, vector[float32](67), vector[float32](59), vector[float32](61))
    checkOperation(t, MulFloat32, goMul, vector[float32](73), vector[float32](79), vector[float32](71))
}

func TestMulFloat32PowerTwo(t *testing.T) {
    checkOperation(t, MulFloat32, goMul, vector[float32](1024), vector[float32](1024), vector[float32](1024))
    checkOperation(t, MulFloat32, goMul, vector[float32](2048), vector[float32](2048), vector[float32](2048))
    checkOperation(t, MulFloat32, goMul, vector[float32](4096), vector[float32](4096), vector[float32](4096))
    checkOperation(t, MulFloat32, goMul, vector[float32](8192), vector[float32](8192), vector[float32](8192))
}

func TestMulFloat64Zero(t *testing.T) {
    checkOperation(t, MulFloat64, goMul, []float64{}, []float64{}, []float64{})
}

func TestMulFloat64Prime(t *testing.T) {
    checkOperation(t, MulFloat64, goMul, vector[float64](11), vector[float64](13), vector[float64](17))
    checkOperation(t, MulFloat64, goMul, vector[float64](29), vector[float64](19), vector[float64](23))
    checkOperation(t, MulFloat64, goMul, vector[float64](37), vector[float64](41), vector[float64](31))
    checkOperation(t, MulFloat64, goMul, vector[float64](43), vector[float64](47), vector[float64](53))
    checkOperation(t, MulFloat64, goMul, vector[float64](67), vector[float64](59), vector[float64](61))
    checkOperation(t, MulFloat64, goMul, vector[float64](73), vector[float64](79), vector[float64](71))
}

func TestMulFloat64PowerTwo(t *testing.T) {
    checkOperation(t, MulFloat64, goMul, vector[float64](1024), vector[float64](1024), vector[float64](1024))
    checkOperation(t, MulFloat64, goMul, vector[float64](2048), vector[float64](2048), vector[float64](2048))
    checkOperation(t, MulFloat64, goMul, vector[float64](4096), vector[float64](4096), vector[float64](4096))
    checkOperation(t, MulFloat64, goMul, vector[float64](8192), vector[float64](8192), vector[float64](8192))
}

func TestDivInt8Zero(t *testing.T) {
    checkOperation(t, DivInt8, goDiv, []int8{}, []int8{}, []int8{})
}

func TestDivInt8Prime(t *testing.T) {
    checkOperation(t, DivInt8, goDiv, vector[int8](11), vector[int8](13), vector[int8](17))
    checkOperation(t, DivInt8, goDiv, vector[int8](29), vector[int8](19), vector[int8](23))
    checkOperation(t, DivInt8, goDiv, vector[int8](37), vector[int8](41), vector[int8](31))
    checkOperation(t, DivInt8, goDiv, vector[int8](43), vector[int8](47), vector[int8](53))
    checkOperation(t, DivInt8, goDiv, vector[int8](67), vector[int8](59), vector[int8](61))
    checkOperation(t, DivInt8, goDiv, vector[int8](73), vector[int8](79), vector[int8](71))
}

func TestDivInt8PowerTwo(t *testing.T) {
    checkOperation(t, DivInt8, goDiv, vector[int8](1024), vector[int8](1024), vector[int8](1024))
    checkOperation(t, DivInt8, goDiv, vector[int8](2048), vector[int8](2048), vector[int8](2048))
    checkOperation(t, DivInt8, goDiv, vector[int8](4096), vector[int8](4096), vector[int8](4096))
    checkOperation(t, DivInt8, goDiv, vector[int8](8192), vector[int8](8192), vector[int8](8192))
}

func TestDivInt16Zero(t *testing.T) {
    checkOperation(t, DivInt16, goDiv, []int16{}, []int16{}, []int16{})
}

func TestDivInt16Prime(t *testing.T) {
    checkOperation(t, DivInt16, goDiv, vector[int16](11), vector[int16](13), vector[int16](17))
    checkOperation(t, DivInt16, goDiv, vector[int16](29), vector[int16](19), vector[int16](23))
    checkOperation(t, DivInt16, goDiv, vector[int16](37), vector[int16](41), vector[int16](31))
    checkOperation(t, DivInt16, goDiv, vector[int16](43), vector[int16](47), vector[int16](53))
    checkOperation(t, DivInt16, goDiv, vector[int16](67), vector[int16](59), vector[int16](61))
    checkOperation(t, DivInt16, goDiv, vector[int16](73), vector[int16](79), vector[int16](71))
}

func TestDivInt16PowerTwo(t *testing.T) {
    checkOperation(t, DivInt16, goDiv, vector[int16](1024), vector[int16](1024), vector[int16](1024))
    checkOperation(t, DivInt16, goDiv, vector[int16](2048), vector[int16](2048), vector[int16](2048))
    checkOperation(t, DivInt16, goDiv, vector[int16](4096), vector[int16](4096), vector[int16](4096))
    checkOperation(t, DivInt16, goDiv, vector[int16](8192), vector[int16](8192), vector[int16](8192))
}

func TestDivInt32Zero(t *testing.T) {
    checkOperation(t, DivInt32, goDiv, []int32{}, []int32{}, []int32{})
}

func TestDivInt32Prime(t *testing.T) {
    checkOperation(t, DivInt32, goDiv, vector[int32](11), vector[int32](13), vector[int32](17))
    checkOperation(t, DivInt32, goDiv, vector[int32](29), vector[int32](19), vector[int32](23))
    checkOperation(t, DivInt32, goDiv, vector[int32](37), vector[int32](41), vector[int32](31))
    checkOperation(t, DivInt32, goDiv, vector[int32](43), vector[int32](47), vector[int32](53))
    checkOperation(t, DivInt32, goDiv, vector[int32](67), vector[int32](59), vector[int32](61))
    checkOperation(t, DivInt32, goDiv, vector[int32](73), vector[int32](79), vector[int32](71))
}

func TestDivInt32PowerTwo(t *testing.T) {
    checkOperation(t, DivInt32, goDiv, vector[int32](1024), vector[int32](1024), vector[int32](1024))
    checkOperation(t, DivInt32, goDiv, vector[int32](2048), vector[int32](2048), vector[int32](2048))
    checkOperation(t, DivInt32, goDiv, vector[int32](4096), vector[int32](4096), vector[int32](4096))
    checkOperation(t, DivInt32, goDiv, vector[int32](8192), vector[int32](8192), vector[int32](8192))
}

func TestDivInt64Zero(t *testing.T) {
    checkOperation(t, DivInt64, goDiv, []int64{}, []int64{}, []int64{})
}

func TestDivInt64Prime(t *testing.T) {
    checkOperation(t, DivInt64, goDiv, vector[int64](11), vector[int64](13), vector[int64](17))
    checkOperation(t, DivInt64, goDiv, vector[int64](29), vector[int64](19), vector[int64](23))
    checkOperation(t, DivInt64, goDiv, vector[int64](37), vector[int64](41), vector[int64](31))
    checkOperation(t, DivInt64, goDiv, vector[int64](43), vector[int64](47), vector[int64](53))
    checkOperation(t, DivInt64, goDiv, vector[int64](67), vector[int64](59), vector[int64](61))
    checkOperation(t, DivInt64, goDiv, vector[int64](73), vector[int64](79), vector[int64](71))
}

func TestDivInt64PowerTwo(t *testing.T) {
    checkOperation(t, DivInt64, goDiv, vector[int64](1024), vector[int64](1024), vector[int64](1024))
    checkOperation(t, DivInt64, goDiv, vector[int64](2048), vector[int64](2048), vector[int64](2048))
    checkOperation(t, DivInt64, goDiv, vector[int64](4096), vector[int64](4096), vector[int64](4096))
    checkOperation(t, DivInt64, goDiv, vector[int64](8192), vector[int64](8192), vector[int64](8192))
}

func TestDivFloat32Zero(t *testing.T) {
    checkOperation(t, DivFloat32, goDiv, []float32{}, []float32{}, []float32{})
}

func TestDivFloat32Prime(t *testing.T) {
    checkOperation(t, DivFloat32, goDiv, vector[float32](11), vector[float32](13), vector[float32](17))
    checkOperation(t, DivFloat32, goDiv, vector[float32](29), vector[float32](19), vector[float32](23))
    checkOperation(t, DivFloat32, goDiv, vector[float32](37), vector[float32](41), vector[float32](31))
    checkOperation(t, DivFloat32, goDiv, vector[float32](43), vector[float32](47), vector[float32](53))
    checkOperation(t, DivFloat32, goDiv, vector[float32](67), vector[float32](59), vector[float32](61))
    checkOperation(t, DivFloat32, goDiv, vector[float32](73), vector[float32](79), vector[float32](71))
}

func TestDivFloat32PowerTwo(t *testing.T) {
    checkOperation(t, DivFloat32, goDiv, vector[float32](1024), vector[float32](1024), vector[float32](1024))
    checkOperation(t, DivFloat32, goDiv, vector[float32](2048), vector[float32](2048), vector[float32](2048))
    checkOperation(t, DivFloat32, goDiv, vector[float32](4096), vector[float32](4096), vector[float32](4096))
    checkOperation(t, DivFloat32, goDiv, vector[float32](8192), vector[float32](8192), vector[float32](8192))
}

func TestDivFloat64Zero(t *testing.T) {
    checkOperation(t, DivFloat64, goDiv, []float64{}, []float64{}, []float64{})
}

func TestDivFloat64Prime(t *testing.T) {
    checkOperation(t, DivFloat64, goDiv, vector[float64](11), vector[float64](13), vector[float64](17))
    checkOperation(t, DivFloat64, goDiv, vector[float64](29), vector[float64](19), vector[float64](23))
    checkOperation(t, DivFloat64, goDiv, vector[float64](37), vector[float64](41), vector[float64](31))
    checkOperation(t, DivFloat64, goDiv, vector[float64](43), vector[float64](47), vector[float64](53))
    checkOperation(t, DivFloat64, goDiv, vector[float64](67), vector[float64](59), vector[float64](61))
    checkOperation(t, DivFloat64, goDiv, vector[float64](73), vector[float64](79), vector[float64](71))
}

func TestDivFloat64PowerTwo(t *testing.T) {
    checkOperation(t, DivFloat64, goDiv, vector[float64](1024), vector[float64](1024), vector[float64](1024))
    checkOperation(t, DivFloat64, goDiv, vector[float64](2048), vector[float64](2048), vector[float64](2048))
    checkOperation(t, DivFloat64, goDiv, vector[float64](4096), vector[float64](4096), vector[float64](4096))
    checkOperation(t, DivFloat64, goDiv, vector[float64](8192), vector[float64](8192), vector[float64](8192))
}

