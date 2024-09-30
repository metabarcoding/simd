// +build amd64

// func sse2AddInt64(left, right, result []int64) int
TEXT ·sse2AddInt64(SB), 4, $0
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
    CMPQ    BX, $2
    JL      singleDataLoop
    MOVOU   (SI)(AX*8), X0
    MOVOU   (DX)(AX*8), X1
    PADDQ   X1, X0
    MOVOU   X0, (DI)(AX*8)
    ADDQ    $2, AX
    JMP     multipleDataLoop
singleDataLoop:
    CMPQ    AX, CX
    JGE     returnLength
    MOVQ    (SI)(AX*8), R8
    MOVQ    (DX)(AX*8), R9
    ADDQ    R9, R8
    MOVQ    R8, (DI)(AX*8)
    INCQ    AX
    JMP     singleDataLoop
returnLength:
    MOVQ    CX, int+72(FP)
    RET
