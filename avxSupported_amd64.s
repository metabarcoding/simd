// +build amd64

// func avxSupported() bool
TEXT ·avxSupported(SB), 4, $0
    MOVQ    $1, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<28), CX
    JZ      avxFalse
avxTrue:
    MOVQ    $1, bool+0(FP)
    RET
avxFalse:
    MOVQ    $0, bool+0(FP)
    RET
