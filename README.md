# SIMD (Single Instruction, Multiple Data)
Simd support via Go assembly for arithmetic and bitwise operations.
Allowing for parallel element-wise computations.
Resulting in **200 - 470%** speedup.
Currently amd64 (x86_64) and arm64 processors are supported.
## Function Documentation
- Online at [pkg.go.dev/github.com/pehringer/simd](https://pkg.go.dev/github.com/pehringer/simd).  
- Locally in ```simd.go``` GoDoc comments.
## Simd Support
|          |amd64 (x86_64)|arm64|ppc64 / ppc64le|
|----------|--------------|-----|---------------|
|AddFloat32|SSE2 / AVX    |NEON |               |
|AddFloat64|SSE2 / AVX    |NEON |               |
|AddInt32  |SSE2 / AVX2   |NEON |               |
|AddInt64  |SSE2 / AVX2   |NEON |               |
|AndInt32  |SSE2 / AVX2   |NEON |               |
|AndInt64  |SSE2 / AVX2   |NEON |               |
|DivFloat32|SSE2 / AVX    |     |               |
|DivFloat64|SSE2 / AVX    |     |               |
|DivInt32  |              |     |               |
|DivInt64  |              |     |               |
|MulFloat32|SSE2 / AVX    |NEON |               |
|MulFloat64|SSE2 / AVX    |NEON |               |
|MulInt32  |SSE4.1 / AVX2 |NEON |               |
|MulInt64  |              |     |               |
|OrInt32   |SSE2 / AVX2   |NEON |               |
|OrInt64   |SSE2 / AVX2   |NEON |               |
|SubFloat32|SSE2 / AVX    |NEON |               |
|SubFloat64|SSE2 / AVX    |NEON |               |
|SubInt32  |SSE2 / AVX2   |NEON |               |
|SubInt64  |SSE2 / AVX2   |NEON |               |
|XorInt32  |SSE2 / AVX2   |NEON |               |
|XorInt64  |SSE2 / AVX2   |NEON |               |
## AMD64 Simd Performance:
![Large Vectors](images/LargeVectorsFloat32Addition.png)
![Medium Vectors](images/MediumVectorsFloat32Addition.png)
![Large Vectors](images/SmallVectorsFloat32Addition.png)  

