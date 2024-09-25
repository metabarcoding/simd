// +build amd64

// func avx2Supported() bool
TEXT ·avx2Supported(SB), 4, $0
    MOVQ    $7, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<5), BX
    JZ      avx2False
avx2True:
    MOVQ    $1, bool+0(FP)
    RET
avx2False:
    MOVQ    $0, bool+0(FP)
    RET
