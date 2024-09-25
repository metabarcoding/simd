// +build amd64

// func sseSupported() bool
TEXT ·sseSupported(SB), 4, $0
    MOVQ    $1, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<25), DX
    JZ      sseFalse
sseTrue:
    MOVQ    $1, bool+0(FP)
    RET
sseFalse:
    MOVQ    $0, bool+0(FP)
    RET
