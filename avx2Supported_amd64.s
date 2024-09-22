// +build amd64

TEXT ·avx2Supported(SB), $0
    MOVQ    $7, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<5), BX
    JZ      avx2False
avx2True:
    MOVQ    $1, AX
    RET
avx2False:
    MOVQ    $0, AX
    RET
