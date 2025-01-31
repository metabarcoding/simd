![SIMD](logo/300x300.png)
# SIMD (Single Instruction, Multiple Data)
SIMD support via Go assembly for arithmetic, bitwise, maximum, and minimum operations.
Allowing for parallel element-wise computations.
Resulting in a ***100% to 400%*** speedup.
Currently AMD64 (x86_64) and ARM64 processors are supported.
## Function Documentation
- Online at [pkg.go.dev/github.com/pehringer/simd](https://pkg.go.dev/github.com/pehringer/simd).
- Locally in ```simd.go``` GoDoc comments.
## SIMD Support
|          |AMD64 (x86_64)          |ARM64|
|----------|------------------------|-----|
|AddFloat32|SSE / AVX / AVX512VL    |NEON |
|AddFloat64|SSE2 / AVX / AVX512VL   |NEON |
|AddInt32  |SSE2 / AVX2 / AVX512VL  |NEON |
|AddInt64  |SSE2 / AVX2 / AVX512VL  |NEON |
|AndInt32  |SSE2 / AVX2 / AVX512VL  |NEON |
|AndInt64  |SSE2 / AVX2 / AVX512VL  |NEON |
|DivFloat32|SSE / AVX / AVX512VL    |     |
|DivFloat64|SSE2 / AVX / AVX512VL   |     |
|DivInt32  |                        |     |
|DivInt64  |                        |     |
|MaxFloat32|SSE                     |     |
|MaxFloat64|SSE2                    |     |
|MaxInt32  |SSE4.1                  |     |
|MaxInt64  |                        |     |
|MinFloat32|SSE                     |     |
|MinFloat64|SSE2                    |     |
|MinInt32  |SSE4.1                  |     |
|MinInt64  |                        |     |
|MulFloat32|SSE / AVX / AVX512VL    |NEON |
|MulFloat64|SSE2 / AVX / AVX512VL   |NEON |
|MulInt32  |SSE4.1 / AVX2 / AVX512VL|NEON |
|MulInt64  |AVX512VL                |     |
|OrInt32   |SSE2 / AVX2 / AVX512VL  |NEON |
|OrInt64   |SSE2 / AVX2 / AVX512VL  |NEON |
|SubFloat32|SSE / AVX / AVX512VL    |NEON |
|SubFloat64|SSE2 / AVX / AVX512VL   |NEON |
|SubInt32  |SSE2 / AVX2 / AVX512VL  |NEON |
|SubInt64  |SSE2 / AVX2 / AVX512VL  |NEON |
|XorInt32  |SSE2 / AVX2 / AVX512VL  |     |
|XorInt64  |SSE2 / AVX2 / AVX512VL  |     |
## Make Targets
#### Tests
|Command              |Description                                                           |
|---------------------|----------------------------------------------------------------------|
|```make test```      |Compiles and runs tests natively on hardware.                         |
|```make test_amd64```|Cross compiles for amd64 and runs tests via QEMU (```qemu-x86_64```). |
|```make test_arm64```|Cross compiles for arm64 and runs tests via QEMU (```qemu-aarch64```).|
#### Benchmarks
|Command               |Description                                                                |
|----------------------|---------------------------------------------------------------------------|
|```make bench```      |Compiles and runs benchmarks natively on hardware.                         |
|```make bench_amd64```|Cross compiles for amd64 and runs benchmarks via QEMU (```qemu-x86_64```). |
|```make bench_arm64```|Cross compiles for arm64 and runs benchmarks via QEMU (```qemu-aarch64```).|
## AMD64 Performance (AMD Ryzen 7 7840U / DDR5 SO-DIMM)
|Elements      |Go ns/op|SIMD ns/op|Performance x|
|--------------|--------|----------|-------------|
|Small Vectors |        |          |             |
|100           |38.33   |7.580     |5.056        |
|200           |79.59   |12.80     |6.217        |
|300           |117.0   |18.45     |9.593        |
|400           |154.5   |16.20     |9.537        |
|500           |191.5   |20.38     |9.396        |
|600           |228.6   |26.37     |8.668        |
|700           |265.6   |33.70     |7.881        |
|800           |303.1   |29.38     |10.31        |
|900           |340.3   |33.54     |10.14        |
|Medium Vectors|        |          |             |
|1000          |377.4   |39.60     |9.530        |
|2000          |751.2   |69.45     |10.81        |
|3000          |1153    |148.3     |7.774        |
|4000          |1499    |325.1     |4.610        |
|5000          |1871    |431.6     |4.335        |
|6000          |2243    |523.6     |4.283        |
|7000          |2614    |614.1     |4.256        |
|8000          |2987    |701.6     |4.257        |
|9000          |3360    |792.5     |4.239        |
|Large Vectors |        |          |             |
|10000         |3725    |878.5     |4.240        |
|20000         |7458    |1754      |4.251        |
|30000         |11187   |2631      |4.251        |
|40000         |14908   |3509      |4.248        |
|50000         |18677   |4373      |4.270        |
|60000         |22363   |5276      |4.238        |
|70000         |26107   |6319      |4.131        |
|80000         |29854   |7820      |3.817        |
|90000         |33613   |9222      |3.644        |
## ARM64 Performance (Apple M1 Pro / LPDDR5 SDRAM)
|Elements      |Go ns/op|SIMD ns/op|Performance x|
|--------------|--------|----------|-------------|
|Small Vectors |        |          |             |
|100           |51.81   |13.68     |3.787        |
|200           |102.2   |24.24     |4.216        |
|300           |152.8   |35.93     |4.252        |
|400           |209.0   |47.71     |4.380        |
|500           |258.7   |64.88     |3.987        |
|600           |309.8   |73.42     |4.219        |
|700           |359.6   |89.01     |4.039        |
|800           |410.6   |101.9     |4.029        |
|900           |460.3   |112.5     |4.091        |
|Medium Vectors|        |          |             |
|1000          |511.5   |124.3     |4.115        |
|2000          |1015    |241.0     |4.211        |
|3000          |1520    |356.9     |4.258        |
|4000          |2024    |473.1     |4.278        |
|5000          |2527    |589.9     |4.283        |
|6000          |3032    |706.1     |4.294        |
|7000          |3535    |822.5     |4.297        |
|8000          |4039    |939.2     |4.300        |
|9000          |4543    |1056      |4.302        |
|Large Vectors |        |          |             |
|10000         |5046    |1172      |4.305        |
|20000         |10107   |2394      |4.221        |
|30000         |15139   |3599      |4.206        |
|40000         |20178   |4957      |4.070        |
|50000         |25218   |6190      |4.073        |
|60000         |30253   |7277      |4.157        |
|70000         |35285   |8707      |4.052        |
|80000         |40346   |9924      |4.065        |
|90000         |45378   |11189     |4.055        |
