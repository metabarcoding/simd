// +build amd64

// func Supported() bool
TEXT ·Supported(SB), 4, $0
    MOVQ    $1, AX
    CPUID
    TESTQ   $(1<<26), DX
    JZ      sse2False
sse2True:
    MOVQ    $1, bool+0(FP)
    RET
sse2False:
    MOVQ    $0, bool+0(FP)
    RET
