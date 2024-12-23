//go:build amd64
// +build amd64

// func DivFloat64(left, right, result []float64) int
TEXT ·DivFloat64(SB), 4, $0
    //Load slices lengths.
    MOVQ    leftLen+8(FP), AX
    MOVQ    rightLen+32(FP), BX
    MOVQ    resultLen+56(FP), CX
    //Get minimum length.
    CMPQ    AX, CX
    CMOVQLT AX, CX
    CMPQ    BX, CX
    CMOVQLT BX, CX
    //Load slices data pointers.
    MOVQ    leftData+0(FP), SI
    MOVQ    rightData+24(FP), DX
    MOVQ    resultData+48(FP), DI
    //Initialize loop index.
    MOVQ    $0, AX
multipleDataLoop:
    MOVQ    CX, BX
    SUBQ    AX, BX
    CMPQ    BX, $8
    JL      singleDataLoop
    //Div eight float64 values.
    VMOVUPD (SI)(AX*8), Z0
    VMOVUPD (DX)(AX*8), Z1
    VDIVPD  Z1, Z0, Z2
    VMOVUPD Z2, (DI)(AX*8)
    ADDQ    $8, AX
    JMP     multipleDataLoop
singleDataLoop:
    CMPQ    AX, CX
    JGE     returnLength
    //Div one float64 value.
    MOVSD   (SI)(AX*8), X0
    MOVSD   (DX)(AX*8), X1
    DIVSD   X1, X0
    MOVSD   X0, (DI)(AX*8)
    INCQ    AX
    JMP     singleDataLoop
returnLength:
    MOVQ    CX, int+72(FP)
    RET
