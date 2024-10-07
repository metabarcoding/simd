# SIMD (Single Instruction, Multiple Data)
Simd support for arithmetic operations.  
### Function Documentation
Available locally in ```simd.go```.  
Available online at [pkg.go.dev/github.com/pehringer/simd](https://pkg.go.dev/github.com/pehringer/simd).  
### AMD64 Simd Support
|          |SSE|SSE2|SSE4.1|AVX|AVX2|
|----------|---|----|------|---|----|
|AddInt32  |   |X   |⟵     |⟵  |X   |
|AddInt64  |   |X   |⟵     |⟵  |X   |
|AddFloat32|X  |⟵   |⟵     |X  |⟵   |
|AddFloat64|   |X   |⟵     |X  |⟵   |
|SubInt32  |   |X   |⟵     |⟵  |X   |
|SubInt64  |   |X   |⟵     |⟵  |X   |
|SubFloat32|X  |⟵   |⟵     |X  |⟵   |
|SubFloat64|   |X   |⟵     |X  |⟵   |
|MulInt32  |   |    |X     |⟵  |X   |
|MulInt64  |   |    |      |   |    |
|MulFloat32|X  |⟵   |⟵     |X  |⟵   |
|MulFloat64|   |X   |⟵     |X  |⟵   |
|DivInt32  |   |    |      |   |    |
|DivInt64  |   |    |      |   |    |
|DivFloat32|X  |⟵   |⟵     |X  |⟵   |
|DivFloat64|   |X   |⟵     |X  |⟵   |
