// +build amd64

TEXT ·avx2Supported(SB), $0
    MOVQ    $7, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<5), BX
    JZ      avx2False
avx2True:
    MOVQ    $1, returnValue+0(FP)
    RET
avx2False:
    MOVQ    $0, returnValue+0(FP)
    RET
