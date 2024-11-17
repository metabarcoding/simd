![logo](logo.png)
# SIMD (Single Instruction, Multiple Data)
Simd support via Go assembly for arithmetic and bitwise operations.
Allowing for parallel element-wise computations.
Resulting in a ***100% to 400%*** speedup.
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
|XorInt32  |SSE2 / AVX2   |     |               |
|XorInt64  |SSE2 / AVX2   |     |               |
## AMD64 AddFloat32 Performance:
|Elements      |Go ns/op|SIMD ns/op|Performance|
|--------------|--------|----------|-----------|
|Small Vectors |        |          |           |
|100           |42.5    |96.6      |0.4        |
|200           |88.4    |99.9      |0.8        |
|300           |127.6   |106.0     |1.2        |
|400           |167.8   |110.7     |1.5        |
|500           |208.8   |118.2     |1.7        |
|600           |247.2   |123.3     |2.0        |
|700           |286.5   |129.4     |2.2        |
|800           |328.5   |131.4     |2.5        |
|900           |362.7   |137.8     |2.6        |
|Medium Vectors|        |          |           |
|1000          |407.5   |139.6     |2.9        |
|2000          |818.0   |182.9     |4.4        |
|3000          |1207    |222.1     |5.4        |
|4000          |1612    |290.0     |5.5        |
|5000          |2028    |482.7     |4.2        |
|6000          |2412    |544.5     |4.4        |
|7000          |2846    |623.4     |4.5        |
|8000          |3277    |747.4     |4.3        |
|9000          |3681    |806.7     |4.5        |
|Large Vectors |        |          |           |
|10000         |4101    |858.6     |4.7        |
|20000         |8218    |1744      |4.7        |
|30000         |12188   |2587      |4.7        |
|40000         |16363   |3277      |4.9        |
|50000         |20343   |4074      |4.9        |
|60000         |24265   |5029      |4.8        |
|70000         |28435   |6210      |4.5        |
|80000         |32298   |7519      |4.2        |
|90000         |36328   |9987      |3.6        |
## ARM64 AddFloat32 Performance:
|Elements      |Go ns/op|SIMD ns/op|Performance|
|--------------|--------|----------|-----------|
|Small Vectors |        |          |           |
|100           |51.8    |13.6      |3.8        |
|200           |102.2   |24.2      |4.2        |
|300           |152.8   |35.9      |4.2        |
|400           |209.0   |47.7      |4.3        |
|500           |258.7   |64.8      |3.9        |
|600           |309.8   |73.4      |4.2        |
|700           |359.6   |89.0      |4.0        |
|800           |410.6   |101.9     |4.0        |
|900           |460.3   |112.5     |4.0        |
|Medium Vectors|        |          |           |
|1000          |511.5   |124.3     |4.1        |
|2000          |1015    |241.0     |4.2        |
|3000          |1520    |356.9     |4.2        |
|4000          |2024    |473.1     |4.2        |
|5000          |2527    |589.9     |4.2        |
|6000          |3032    |706.1     |4.2        |
|7000          |3535    |822.5     |4.2        |
|8000          |4039    |939.2     |4.3        |
|9000          |4543    |1056      |4.3        |
|Large Vectors |        |          |           |
|10000         |5046    |1172      |4.3        |
|20000         |10107   |2394      |4.2        |
|30000         |15139   |3599      |4.2        |
|40000         |20178   |4957      |4.0        |
|50000         |25218   |6190      |4.0        |
|60000         |30253   |7277      |4.1        |
|70000         |35285   |8707      |4.0        |
|80000         |40346   |9924      |4.0        |
|90000         |45378   |11189     |4.0        |
