// +build amd64

// func sseAddFloat32(left, right, result []float32) int
TEXT ·sseAddFloat32(SB), 4, $0
    MOVQ    leftData+0(FP), SI
    MOVQ    leftLen+8(FP), AX
    MOVQ    rightData+24(FP), DX
    MOVQ    rightLen+32(FP), BX
    MOVQ    resultData+48(FP), DI
    MOVQ    resultLen+56(FP), CX
minLenLeft:
    CMPQ    AX, CX
    JGE     minLenRight
    MOVQ    AX, CX
minLenRight:
    CMPQ    AX, CX
    JGE     minLenResult
    MOVQ    AX, CX
minLenResult:
    XORQ    AX, AX
    XORQ    BX, BX
addMany:
    MOVQ    CX, BX
    SUBQ    AX, BX
    CMPQ    BX, $4
    JL      addOne
    
    MOVUPS  (SI)(AX*4), X0
    MOVUPS  (DX)(AX*4), X1
    ADDPS   X1, X0
    MOVUPS  X0, (DI)(AX*4)
    
    ADDQ    $4, AX
    JMP     addMany
addOne:
    CMPQ    AX, CX
    JGE     ret
    MOVSS   (SI)(AX*4), X0
    MOVSS   (DX)(AX*4), X1
    ADDSS   X1, X0
    MOVSS   X0, (DI)(AX*4)
    INCQ    AX
    JMP     addOne
ret:
    MOVQ CX, int+72(FP)
    RET
