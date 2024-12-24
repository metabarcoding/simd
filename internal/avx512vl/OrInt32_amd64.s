//go:build amd64
// +build amd64

// func OrInt32(left, right, result []int32) int
TEXT ·OrInt32(SB), 4, $0
    //Load slices lengths.
    MOVQ      leftLen+8(FP), AX
    MOVQ      rightLen+32(FP), BX
    MOVQ      resultLen+56(FP), CX
    //Get minimum length.
    CMPQ      AX, CX
    CMOVQLT   AX, CX
    CMPQ      BX, CX
    CMOVQLT   BX, CX
    //Load slices data pointers.
    MOVQ      leftData+0(FP), SI
    MOVQ      rightData+24(FP), DX
    MOVQ      resultData+48(FP), DI
    //Initialize loop index.
    MOVQ      $0, AX
multipleDataLoop:
    MOVQ      CX, BX
    SUBQ      AX, BX
    CMPQ      BX, $16
    JL        singleDataLoop
    //Or sixteen int32 values.
    VMOVDQU32 (SI)(AX*4), Z0
    VMOVDQU32 (DX)(AX*4), Z1
    VPORD     Z1, Z0, Z2
    VMOVDQU32 Z2, (DI)(AX*4)
    ADDQ      $16, AX
    JMP       multipleDataLoop
singleDataLoop:
    CMPQ      AX, CX
    JGE       returnLength
    //Or one int32 value.
    MOVL      (SI)(AX*4), R8
    MOVL      (DX)(AX*4), R9
    ORL       R9, R8
    MOVL      R8, (DI)(AX*4)
    INCQ      AX
    JMP       singleDataLoop
returnLength:
    MOVQ      CX, int+72(FP)
    RET
