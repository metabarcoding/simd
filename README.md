# SIMD (Single Instruction, Multiple Data)
Simd support via Go assembly for arithmetic and bitwise operations.
Allowing for parallel element-wise computations.
Resulting in **200 - 470%** speedup.
Currently 64-bit x86 is supported, with partial 64-bit ARM support.
## Function Documentation
- Online at [pkg.go.dev/github.com/pehringer/simd](https://pkg.go.dev/github.com/pehringer/simd).  
- Locally in ```simd.go``` GoDoc comments.
## Simd Support
|          |64-bit x86   |64-bit ARM|
|----------|-------------|----------|
|AddFloat32|SSE / AVX    |NEON      |
|AddFloat64|SSE2 / AVX   |NEON      |
|AddInt32  |SSE2 / AVX2  |NEON      |
|AddInt64  |SSE2 / AVX2  |NEON      |
|AndInt32  |SSE2 / AVX2  |          |
|AndInt64  |SSE2 / AVX2  |          |
|DivFloat32|SSE / AVX    |          |
|DivFloat64|SSE2 / AVX   |          |
|DivInt32  |             |          |
|DivInt64  |             |          |
|MulFloat32|SSE / AVX    |NEON      |
|MulFloat64|SSE2 / AVX   |          |
|MulInt32  |SSE4.1 / AVX2|          |
|MulInt64  |             |          |
|OrInt32   |SSE2 / AVX2  |          |
|OrInt64   |SSE2 / AVX2  |          |
|SubFloat32|SSE / AVX    |NEON      |
|SubFloat64|SSE2 / AVX   |NEON      |
|SubInt32  |SSE2 / AVX2  |NEON      |
|SubInt64  |SSE2 / AVX2  |NEON      |
|XorInt32  |SSE2 / AVX2  |          |
|XorInt64  |SSE2 / AVX2  |          |
## AMD64 Simd Performance:
![Large Vectors](images/LargeVectorsFloat32Addition.png)
![Medium Vectors](images/MediumVectorsFloat32Addition.png)
![Large Vectors](images/SmallVectorsFloat32Addition.png)  

