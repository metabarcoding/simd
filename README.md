# Simd Hardware Support
|          |SSE (128-Bit)|SSE2 (128-Bit)|AVX (256-Bit)|AVX2 (256-Bit)|
|----------|-------------|--------------|-------------|--------------|
|AddInt32  |             |X             |⟵           |X             |
|AddInt64  |             |X             |⟵           |X             |
|AddFloat32|X            |⟵            |X            |⟵            |
|AddFloat64|             |X             |X            |⟵            |
|SubInt32  |             |X             |⟵           |X             |
|SubInt64  |             |X             |⟵           |X             |
|SubFloat32|X            |⟵            |X            |⟵            |
|SubFloat64|             |X             |X            |⟵            |
|MulInt32  |             |              |             |              |
|MulInt64  |             |              |             |              |
|MulFloat32|X            |⟵            |X            |⟵            |
|MulFloat64|             |X             |X            |⟵            |
|DivInt32  |             |              |             |              |
|DivInt64  |             |              |             |              |
|DivFloat32|X            |⟵            |X            |⟵            |
|DivFloat64|             |X             |X            |⟵            |
# Library Functions
---
### ```func AddInt32(left, right, result []int32) int```
### ```func AddInt64(left, right, result []int64) int```
### ```func AddFloat32(left, right, result []float32) int```
### ```func AddFloat64(left, right, result []float64) int```
Performs element-wise addition on left and right, storing the sums in result.
The operation is performed up to the shortest length of left, right, and result.
Returns the number of operations performed.
- ```left``` left-hand operands.
- ```right``` right-hand operands.
- ```result``` operation results.
---
### ```func SubInt32(left, right, result []int32) int```
### ```func SubInt64(left, right, result []int64) int```
### ```func SubFloat32(left, right, result []float32) int```
### ```func SubFloat64(left, right, result []float64) int```
Performs element-wise subtraction on left and right, storing the differences in result.
The operation is performed up to the shortest length of left, right, and result.
Returns the number of operations performed.
- ```left``` left-hand operands.
- ```right``` right-hand operands.
- ```result``` operation results.
---
### ```func MulInt32(left, right, result []int32) int```
### ```func MulInt64(left, right, result []int64) int```
### ```func MulFloat32(left, right, result []float32) int```
### ```func MulFloat64(left, right, result []float64) int```
Performs element-wise multiplication on left and right, storing the products in result.
The operation is performed up to the shortest length of left, right, and result.
Returns the number of operations performed.
- ```left``` left-hand operands.
- ```right``` right-hand operands.
- ```result``` operation results.
---
### ```func DivInt32(left, right, result []int32) int```
### ```func DivInt64(left, right, result []int64) int```
### ```func DivFloat32(left, right, result []float32) int```
### ```func DivFloat64(left, right, result []float64) int```
Performs element-wise division on left and right, storing the quotients in result.
The operation is performed up to the shortest length of left, right, and result.
Returns the number of operations performed.
- ```left``` left-hand operands.
- ```right``` right-hand operands.
- ```result``` operation results.
---
