// +build amd64

TEXT ·avxSupported(SB), $0
    MOVQ    $1, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<28), CX
    JZ      avxFalse
avxTrue:
    MOVQ    $1, AX
    RET
avxFalse:
    MOVQ    $0, AX
    RET
