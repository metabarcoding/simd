//go:build arm64
// +build arm64

// func SubFloat32(left, right, result []float32) int
TEXT ·SubFloat32(SB), 4, $0
    //Load slices lengths.
    MOVD    leftLen+8(FP), R0
    MOVD    rightLen+32(FP), R1
    MOVD    resultLen+56(FP), R2
    //Get minimum length.
    CMP     R0, R2
    CSEL    LT, R2, R0, R2
    CMP     R1, R2
    CSEL    LT, R2, R1, R2
    //Load slices data pointers.
    MOVD    leftData+0(FP), R3
    MOVD    rightData+24(FP), R4
    MOVD    resultData+48(FP), R5
    //Initialize loop index.
    MOVD    $0, R0
multipleDataLoop:
    SUB     R0, R2, R1
    CMP     $4, R1
    BLT     singleDataLoop
    //Sub four float32 values.
    VLD1.P  16(R3), [V0.S4]
    VLD1.P  16(R4), [V1.S4]
    WORD    $0x4EA1D402 //VFSUB V1.S4, V0.S4, V2.S4
    VST1.P  [V2.S4], 16(R5)
    ADD     $4, R0, R0
    B       multipleDataLoop
singleDataLoop:
    CMP     R2, R0
    BGE     returnLength
    //Sub one float32 value.
    FMOVS.P 4(R3), F0
    FMOVS.P 4(R4), F1
    FSUBS   F1, F0, F2
    FMOVS.P F2, 4(R5)
    ADD     $1, R0, R0
    B       singleDataLoop
returnLength:
    MOVD    R2, int+72(FP)
    RET
