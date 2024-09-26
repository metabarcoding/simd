// +build amd64

// func sseMulFloat32(left, right, result []float32) int
TEXT ·sseMulFloat32(SB), 4, $0
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
mulQuadrupleLoop:
    MOVQ    CX, BX
    SUBQ    AX, BX
    CMPQ    BX, $4
    JL      mulSingleLoop
    MOVUPS  (SI)(AX*4), X0
    MOVUPS  (DX)(AX*4), X1
    MULPS   X1, X0
    MOVUPS  X0, (DI)(AX*4)
    ADDQ    $4, AX
    JMP     mulQuadrupleLoop
mulSingleLoop:
    CMPQ    AX, CX
    JGE     returnLength
    MOVSS   (SI)(AX*4), X0
    MOVSS   (DX)(AX*4), X1
    MULSS   X1, X0
    MOVSS   X0, (DI)(AX*4)
    INCQ    AX
    JMP     mulSingleLoop
returnLength:
    MOVQ CX, int+72(FP)
    RET
