// +build amd64

TEXT ·sseAddFloat32(SB), $0
    MOVQ a+8(FP), AX
    MOVQ AX, c+24(FP)
    // Return from the function
    RET
