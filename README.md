# 128-bit Simd Support
|          |amd64 SSE|amd64 SSE2|amd64 SSE3|amd64 SSE4.1|amd64 SSE4.2|arm64 NEON|
|----------|---------|----------|----------|------------|------------|----------|
|AddInt32  |         |X         |X         |X           |X           |          |
|AddInt64  |         |X         |X         |X           |X           |          |
|AddFloat32|X        |X         |X         |X           |X           |          |
|AddFloat64|         |X         |X         |X           |X           |          |
|SubInt32  |         |X         |X         |X           |X           |          |
|SubInt64  |         |X         |X         |X           |X           |          |
|SubFloat32|X        |X         |X         |X           |X           |          |
|SubFloat64|         |X         |X         |X           |X           |          |
|MulInt32  |         |          |          |            |            |          |
|MulInt64  |         |          |          |            |            |          |
|MulFloat32|X        |X         |X         |X           |X           |          |
|MulFloat64|         |X         |X         |X           |X           |          |
|DivInt32  |         |          |          |            |            |          |
|DivInt64  |         |          |          |            |            |          |
|DivFloat32|X        |X         |X         |X           |X           |          |
|DivFloat64|         |X         |X         |X           |X           |          |

