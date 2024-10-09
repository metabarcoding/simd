# SIMD (Single Instruction, Multiple Data)
Simd support for arithmetic operations. Allowing for parallel element-wise computations.
## Function Documentation
- Online at [pkg.go.dev/github.com/pehringer/simd](https://pkg.go.dev/github.com/pehringer/simd).  
- Locally in ```simd.go``` GoDoc comments.  
## AMD64 Simd Support
|          |SSE|SSE2|AVX|AVX2|
|----------|---|----|---|----|
|AddFloat32|X  |    |X  |    |
|AddFloat64|   |X   |X  |    |
|AddInt32  |   |X   |   |X   |
|AddInt64  |   |X   |   |X   |
|DivFloat32|X  |    |X  |    |
|DivFloat64|   |X   |X  |    |
|MulFloat32|X  |    |X  |    |
|MulFloat64|   |X   |X  |    |
|SubFloat32|X  |    |X  |    |
|SubFloat64|   |X   |X  |    |
|SubInt32  |   |X   |   |X   |
|SubInt64  |   |X   |   |X   |
