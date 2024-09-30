// +build amd64

// func sseSubFloat32(left, right, result []float32) int
TEXT ·sseSubFloat32(SB), 4, $0
    MOVQ    leftLen+8(FP), AX
    MOVQ    rightLen+32(FP), BX
    MOVQ    resultLen+56(FP), CX
    CMPQ    AX, CX
    JGE     compareLengths
    MOVQ    AX, CX
compareLengths:
    CMPQ    BX, CX
    JGE     initializeLoops
    MOVQ    BX, CX
initializeLoops:
    XORQ    AX, AX
    MOVQ    leftData+0(FP), SI
    MOVQ    rightData+24(FP), DX
    MOVQ    resultData+48(FP), DI
multipleDataLoop:
    MOVQ    CX, BX
    SUBQ    AX, BX
    CMPQ    BX, $4
    JL      singleDataLoop
    MOVO    (SI)(AX*4), X0
    MOVO    (DX)(AX*4), X1
    SUBPS   X1, X0
    MOVO    X0, (DI)(AX*4)
    ADDQ    $4, AX
    JMP     multipleDataLoop
singleDataLoop:
    CMPQ    AX, CX
    JGE     returnLength
    MOVL    (SI)(AX*4), X0
    MOVL    (DX)(AX*4), X1
    SUBSS   X1, X0
    MOVL    X0, (DI)(AX*4)
    INCQ    AX
    JMP     singleDataLoop
returnLength:
    MOVQ    CX, int+72(FP)
    RET
