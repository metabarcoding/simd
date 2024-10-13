// +build amd64

// func Supported() bool
TEXT ·Supported(SB), 4, $0
    //Check SSE supported.
    MOVQ    $1, AX
    CPUID
    TESTQ   $(1<<25), DX
    JZ      sseFalse
    //sseTrue:
    MOVB    $1, bool+0(FP)
    RET
sseFalse:
    MOVB    $0, bool+0(FP)
    RET
