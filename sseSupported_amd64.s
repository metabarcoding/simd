// +build amd64

TEXT ·sseSupported(SB), $0
    MOVQ    $1, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<25), DX
    JZ      sseFalse
sseTrue:
    MOVQ    $1, AX
    RET
sseFalse:
    MOVQ    $0, AX
    RET
