package simd

import (
	"testing"
        "github.com/pehringer/simd/internal/fallback"
        "github.com/pehringer/simd/internal/shared"
)

func TestAddInt32Zero(t *testing.T) {
    shared.CheckOperation(t, AddInt32, fallback.Add, []int32{}, []int32{}, []int32{})
}

func TestAddInt32Prime(t *testing.T) {
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](11), shared.Vector[int32](13), shared.Vector[int32](17))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](29), shared.Vector[int32](19), shared.Vector[int32](23))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](37), shared.Vector[int32](41), shared.Vector[int32](31))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](43), shared.Vector[int32](47), shared.Vector[int32](53))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](67), shared.Vector[int32](59), shared.Vector[int32](61))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](73), shared.Vector[int32](79), shared.Vector[int32](71))
}

func TestAddInt32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](1024), shared.Vector[int32](1024), shared.Vector[int32](1024))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](2048), shared.Vector[int32](2048), shared.Vector[int32](2048))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](4096), shared.Vector[int32](4096), shared.Vector[int32](4096))
    shared.CheckOperation(t, AddInt32, fallback.Add, shared.Vector[int32](8192), shared.Vector[int32](8192), shared.Vector[int32](8192))
}

func TestAddInt64Zero(t *testing.T) {
    shared.CheckOperation(t, AddInt64, fallback.Add, []int64{}, []int64{}, []int64{})
}

func TestAddInt64Prime(t *testing.T) {
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](11), shared.Vector[int64](13), shared.Vector[int64](17))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](29), shared.Vector[int64](19), shared.Vector[int64](23))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](37), shared.Vector[int64](41), shared.Vector[int64](31))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](43), shared.Vector[int64](47), shared.Vector[int64](53))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](67), shared.Vector[int64](59), shared.Vector[int64](61))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](73), shared.Vector[int64](79), shared.Vector[int64](71))
}

func TestAddInt64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](1024), shared.Vector[int64](1024), shared.Vector[int64](1024))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](2048), shared.Vector[int64](2048), shared.Vector[int64](2048))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](4096), shared.Vector[int64](4096), shared.Vector[int64](4096))
    shared.CheckOperation(t, AddInt64, fallback.Add, shared.Vector[int64](8192), shared.Vector[int64](8192), shared.Vector[int64](8192))
}

func TestAddFloat32Zero(t *testing.T) {
    shared.CheckOperation(t, AddFloat32, fallback.Add, []float32{}, []float32{}, []float32{})
}

func TestAddFloat32Prime(t *testing.T) {
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](11), shared.Vector[float32](13), shared.Vector[float32](17))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](29), shared.Vector[float32](19), shared.Vector[float32](23))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](37), shared.Vector[float32](41), shared.Vector[float32](31))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](43), shared.Vector[float32](47), shared.Vector[float32](53))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](67), shared.Vector[float32](59), shared.Vector[float32](61))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](73), shared.Vector[float32](79), shared.Vector[float32](71))
}

func TestAddFloat32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](1024), shared.Vector[float32](1024), shared.Vector[float32](1024))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](2048), shared.Vector[float32](2048), shared.Vector[float32](2048))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](4096), shared.Vector[float32](4096), shared.Vector[float32](4096))
    shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Vector[float32](8192), shared.Vector[float32](8192), shared.Vector[float32](8192))
}

func TestAddFloat64Zero(t *testing.T) {
    shared.CheckOperation(t, AddFloat64, fallback.Add, []float64{}, []float64{}, []float64{})
}

func TestAddFloat64Prime(t *testing.T) {
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](11), shared.Vector[float64](13), shared.Vector[float64](17))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](29), shared.Vector[float64](19), shared.Vector[float64](23))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](37), shared.Vector[float64](41), shared.Vector[float64](31))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](43), shared.Vector[float64](47), shared.Vector[float64](53))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](67), shared.Vector[float64](59), shared.Vector[float64](61))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](73), shared.Vector[float64](79), shared.Vector[float64](71))
}

func TestAddFloat64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](1024), shared.Vector[float64](1024), shared.Vector[float64](1024))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](2048), shared.Vector[float64](2048), shared.Vector[float64](2048))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](4096), shared.Vector[float64](4096), shared.Vector[float64](4096))
    shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Vector[float64](8192), shared.Vector[float64](8192), shared.Vector[float64](8192))
}

func TestSubInt32Zero(t *testing.T) {
    shared.CheckOperation(t, SubInt32, fallback.Sub, []int32{}, []int32{}, []int32{})
}

func TestSubInt32Prime(t *testing.T) {
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](11), shared.Vector[int32](13), shared.Vector[int32](17))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](29), shared.Vector[int32](19), shared.Vector[int32](23))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](37), shared.Vector[int32](41), shared.Vector[int32](31))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](43), shared.Vector[int32](47), shared.Vector[int32](53))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](67), shared.Vector[int32](59), shared.Vector[int32](61))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](73), shared.Vector[int32](79), shared.Vector[int32](71))
}

func TestSubInt32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](1024), shared.Vector[int32](1024), shared.Vector[int32](1024))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](2048), shared.Vector[int32](2048), shared.Vector[int32](2048))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](4096), shared.Vector[int32](4096), shared.Vector[int32](4096))
    shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Vector[int32](8192), shared.Vector[int32](8192), shared.Vector[int32](8192))
}

func TestSubInt64Zero(t *testing.T) {
    shared.CheckOperation(t, SubInt64, fallback.Sub, []int64{}, []int64{}, []int64{})
}

func TestSubInt64Prime(t *testing.T) {
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](11), shared.Vector[int64](13), shared.Vector[int64](17))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](29), shared.Vector[int64](19), shared.Vector[int64](23))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](37), shared.Vector[int64](41), shared.Vector[int64](31))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](43), shared.Vector[int64](47), shared.Vector[int64](53))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](67), shared.Vector[int64](59), shared.Vector[int64](61))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](73), shared.Vector[int64](79), shared.Vector[int64](71))
}

func TestSubInt64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](1024), shared.Vector[int64](1024), shared.Vector[int64](1024))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](2048), shared.Vector[int64](2048), shared.Vector[int64](2048))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](4096), shared.Vector[int64](4096), shared.Vector[int64](4096))
    shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Vector[int64](8192), shared.Vector[int64](8192), shared.Vector[int64](8192))
}

func TestSubFloat32Zero(t *testing.T) {
    shared.CheckOperation(t, SubFloat32, fallback.Sub, []float32{}, []float32{}, []float32{})
}

func TestSubFloat32Prime(t *testing.T) {
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](11), shared.Vector[float32](13), shared.Vector[float32](17))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](29), shared.Vector[float32](19), shared.Vector[float32](23))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](37), shared.Vector[float32](41), shared.Vector[float32](31))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](43), shared.Vector[float32](47), shared.Vector[float32](53))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](67), shared.Vector[float32](59), shared.Vector[float32](61))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](73), shared.Vector[float32](79), shared.Vector[float32](71))
}

func TestSubFloat32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](1024), shared.Vector[float32](1024), shared.Vector[float32](1024))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](2048), shared.Vector[float32](2048), shared.Vector[float32](2048))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](4096), shared.Vector[float32](4096), shared.Vector[float32](4096))
    shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Vector[float32](8192), shared.Vector[float32](8192), shared.Vector[float32](8192))
}

func TestSubFloat64Zero(t *testing.T) {
    shared.CheckOperation(t, SubFloat64, fallback.Sub, []float64{}, []float64{}, []float64{})
}

func TestSubFloat64Prime(t *testing.T) {
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](11), shared.Vector[float64](13), shared.Vector[float64](17))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](29), shared.Vector[float64](19), shared.Vector[float64](23))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](37), shared.Vector[float64](41), shared.Vector[float64](31))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](43), shared.Vector[float64](47), shared.Vector[float64](53))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](67), shared.Vector[float64](59), shared.Vector[float64](61))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](73), shared.Vector[float64](79), shared.Vector[float64](71))
}

func TestSubFloat64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](1024), shared.Vector[float64](1024), shared.Vector[float64](1024))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](2048), shared.Vector[float64](2048), shared.Vector[float64](2048))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](4096), shared.Vector[float64](4096), shared.Vector[float64](4096))
    shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Vector[float64](8192), shared.Vector[float64](8192), shared.Vector[float64](8192))
}

func TestMulInt32Zero(t *testing.T) {
    shared.CheckOperation(t, MulInt32, fallback.Mul, []int32{}, []int32{}, []int32{})
}

func TestMulInt32Prime(t *testing.T) {
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](11), shared.Vector[int32](13), shared.Vector[int32](17))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](29), shared.Vector[int32](19), shared.Vector[int32](23))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](37), shared.Vector[int32](41), shared.Vector[int32](31))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](43), shared.Vector[int32](47), shared.Vector[int32](53))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](67), shared.Vector[int32](59), shared.Vector[int32](61))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](73), shared.Vector[int32](79), shared.Vector[int32](71))
}

func TestMulInt32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](1024), shared.Vector[int32](1024), shared.Vector[int32](1024))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](2048), shared.Vector[int32](2048), shared.Vector[int32](2048))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](4096), shared.Vector[int32](4096), shared.Vector[int32](4096))
    shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Vector[int32](8192), shared.Vector[int32](8192), shared.Vector[int32](8192))
}

func TestMulInt64Zero(t *testing.T) {
    shared.CheckOperation(t, MulInt64, fallback.Mul, []int64{}, []int64{}, []int64{})
}

func TestMulInt64Prime(t *testing.T) {
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](11), shared.Vector[int64](13), shared.Vector[int64](17))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](29), shared.Vector[int64](19), shared.Vector[int64](23))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](37), shared.Vector[int64](41), shared.Vector[int64](31))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](43), shared.Vector[int64](47), shared.Vector[int64](53))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](67), shared.Vector[int64](59), shared.Vector[int64](61))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](73), shared.Vector[int64](79), shared.Vector[int64](71))
}

func TestMulInt64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](1024), shared.Vector[int64](1024), shared.Vector[int64](1024))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](2048), shared.Vector[int64](2048), shared.Vector[int64](2048))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](4096), shared.Vector[int64](4096), shared.Vector[int64](4096))
    shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Vector[int64](8192), shared.Vector[int64](8192), shared.Vector[int64](8192))
}

func TestMulFloat32Zero(t *testing.T) {
    shared.CheckOperation(t, MulFloat32, fallback.Mul, []float32{}, []float32{}, []float32{})
}

func TestMulFloat32Prime(t *testing.T) {
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](11), shared.Vector[float32](13), shared.Vector[float32](17))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](29), shared.Vector[float32](19), shared.Vector[float32](23))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](37), shared.Vector[float32](41), shared.Vector[float32](31))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](43), shared.Vector[float32](47), shared.Vector[float32](53))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](67), shared.Vector[float32](59), shared.Vector[float32](61))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](73), shared.Vector[float32](79), shared.Vector[float32](71))
}

func TestMulFloat32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](1024), shared.Vector[float32](1024), shared.Vector[float32](1024))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](2048), shared.Vector[float32](2048), shared.Vector[float32](2048))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](4096), shared.Vector[float32](4096), shared.Vector[float32](4096))
    shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Vector[float32](8192), shared.Vector[float32](8192), shared.Vector[float32](8192))
}

func TestMulFloat64Zero(t *testing.T) {
    shared.CheckOperation(t, MulFloat64, fallback.Mul, []float64{}, []float64{}, []float64{})
}

func TestMulFloat64Prime(t *testing.T) {
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](11), shared.Vector[float64](13), shared.Vector[float64](17))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](29), shared.Vector[float64](19), shared.Vector[float64](23))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](37), shared.Vector[float64](41), shared.Vector[float64](31))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](43), shared.Vector[float64](47), shared.Vector[float64](53))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](67), shared.Vector[float64](59), shared.Vector[float64](61))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](73), shared.Vector[float64](79), shared.Vector[float64](71))
}

func TestMulFloat64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](1024), shared.Vector[float64](1024), shared.Vector[float64](1024))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](2048), shared.Vector[float64](2048), shared.Vector[float64](2048))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](4096), shared.Vector[float64](4096), shared.Vector[float64](4096))
    shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Vector[float64](8192), shared.Vector[float64](8192), shared.Vector[float64](8192))
}

func TestDivInt32Zero(t *testing.T) {
    shared.CheckOperation(t, DivInt32, fallback.Div, []int32{}, []int32{}, []int32{})
}

func TestDivInt32Prime(t *testing.T) {
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](11), shared.Vector[int32](13), shared.Vector[int32](17))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](29), shared.Vector[int32](19), shared.Vector[int32](23))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](37), shared.Vector[int32](41), shared.Vector[int32](31))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](43), shared.Vector[int32](47), shared.Vector[int32](53))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](67), shared.Vector[int32](59), shared.Vector[int32](61))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](73), shared.Vector[int32](79), shared.Vector[int32](71))
}

func TestDivInt32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](1024), shared.Vector[int32](1024), shared.Vector[int32](1024))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](2048), shared.Vector[int32](2048), shared.Vector[int32](2048))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](4096), shared.Vector[int32](4096), shared.Vector[int32](4096))
    shared.CheckOperation(t, DivInt32, fallback.Div, shared.Vector[int32](8192), shared.Vector[int32](8192), shared.Vector[int32](8192))
}

func TestDivInt64Zero(t *testing.T) {
    shared.CheckOperation(t, DivInt64, fallback.Div, []int64{}, []int64{}, []int64{})
}

func TestDivInt64Prime(t *testing.T) {
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](11), shared.Vector[int64](13), shared.Vector[int64](17))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](29), shared.Vector[int64](19), shared.Vector[int64](23))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](37), shared.Vector[int64](41), shared.Vector[int64](31))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](43), shared.Vector[int64](47), shared.Vector[int64](53))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](67), shared.Vector[int64](59), shared.Vector[int64](61))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](73), shared.Vector[int64](79), shared.Vector[int64](71))
}

func TestDivInt64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](1024), shared.Vector[int64](1024), shared.Vector[int64](1024))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](2048), shared.Vector[int64](2048), shared.Vector[int64](2048))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](4096), shared.Vector[int64](4096), shared.Vector[int64](4096))
    shared.CheckOperation(t, DivInt64, fallback.Div, shared.Vector[int64](8192), shared.Vector[int64](8192), shared.Vector[int64](8192))
}

func TestDivFloat32Zero(t *testing.T) {
    shared.CheckOperation(t, DivFloat32, fallback.Div, []float32{}, []float32{}, []float32{})
}

func TestDivFloat32Prime(t *testing.T) {
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](11), shared.Vector[float32](13), shared.Vector[float32](17))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](29), shared.Vector[float32](19), shared.Vector[float32](23))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](37), shared.Vector[float32](41), shared.Vector[float32](31))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](43), shared.Vector[float32](47), shared.Vector[float32](53))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](67), shared.Vector[float32](59), shared.Vector[float32](61))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](73), shared.Vector[float32](79), shared.Vector[float32](71))
}

func TestDivFloat32PowerTwo(t *testing.T) {
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](1024), shared.Vector[float32](1024), shared.Vector[float32](1024))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](2048), shared.Vector[float32](2048), shared.Vector[float32](2048))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](4096), shared.Vector[float32](4096), shared.Vector[float32](4096))
    shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Vector[float32](8192), shared.Vector[float32](8192), shared.Vector[float32](8192))
}

func TestDivFloat64Zero(t *testing.T) {
    shared.CheckOperation(t, DivFloat64, fallback.Div, []float64{}, []float64{}, []float64{})
}

func TestDivFloat64Prime(t *testing.T) {
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](11), shared.Vector[float64](13), shared.Vector[float64](17))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](29), shared.Vector[float64](19), shared.Vector[float64](23))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](37), shared.Vector[float64](41), shared.Vector[float64](31))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](43), shared.Vector[float64](47), shared.Vector[float64](53))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](67), shared.Vector[float64](59), shared.Vector[float64](61))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](73), shared.Vector[float64](79), shared.Vector[float64](71))
}

func TestDivFloat64PowerTwo(t *testing.T) {
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](1024), shared.Vector[float64](1024), shared.Vector[float64](1024))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](2048), shared.Vector[float64](2048), shared.Vector[float64](2048))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](4096), shared.Vector[float64](4096), shared.Vector[float64](4096))
    shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Vector[float64](8192), shared.Vector[float64](8192), shared.Vector[float64](8192))
}
