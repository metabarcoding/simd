// +build amd64

// func sseDivFloat32(left, right, result []float32) int
TEXT ·sseDivFloat32(SB), 4, $0
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
divQuadrupleLoop:
    MOVQ    CX, BX
    SUBQ    AX, BX
    CMPQ    BX, $4
    JL      divSingleLoop
    MOVUPS  (SI)(AX*4), X0
    MOVUPS  (DX)(AX*4), X1
    DIVPS   X1, X0
    MOVUPS  X0, (DI)(AX*4)
    ADDQ    $4, AX
    JMP     divQuadrupleLoop
divSingleLoop:
    CMPQ    AX, CX
    JGE     returnLength
    MOVSS   (SI)(AX*4), X0
    MOVSS   (DX)(AX*4), X1
    DIVSS   X1, X0
    MOVSS   X0, (DI)(AX*4)
    INCQ    AX
    JMP     divSingleLoop
returnLength:
    MOVQ CX, int+72(FP)
    RET
