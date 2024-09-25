// +build amd64

// func sse2Supported() bool
TEXT ·sse2Supported(SB), 4, $0
    MOVQ    $1, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<26), DX
    JZ      sse2False
sse2True:
    MOVQ    $1, bool+0(FP)
    RET
sse2False:
    MOVQ    $0, bool+0(FP)
    RET
