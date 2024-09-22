// +build amd64

TEXT ·sseSupported(SB), $0
    MOVQ    $1, AX
    CPUID
    TESTB   $1, DX
    JZ      sseFalse
sseTrue:
    MOVQ    $1, AX
    RET
sseFalse:
    MOVQ    $0, AX
    RET
