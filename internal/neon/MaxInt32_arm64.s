//go:build arm64
// +build arm64

// func MaxInt32(left, right, result []int32) int
TEXT ·MaxInt32(SB), 4, $0
    // Load slices lengths.
    MOVD    leftLen+8(FP), R0
    MOVD    rightLen+32(FP), R1
    MOVD    resultLen+56(FP), R2
    // Get minimum length.
    CMP     R0, R2
    CSEL    LT, R2, R0, R2
    CMP     R1, R2
    CSEL    LT, R2, R1, R2
    // Load slices data pointers.
    MOVD    leftData+0(FP), R3
    MOVD    rightData+24(FP), R4
    MOVD    resultData+48(FP), R5
    // Initialize loop index.
    MOVD    $0, R0

multipleDataLoop:
    SUB     R0, R2, R1
    CMP     $4, R1
    BLT     singleDataLoop
    // Load four int32 values into V0 and V1.
    VLD1.P  16(R3), [V0.S4]
    VLD1.P  16(R4), [V1.S4]
    // SMAX V2.4S, V0.4S, V1.4S - signed maximum (corrigé)
    WORD    $0x4E206C02  // Correct opcode for SMAX
    VST1.P  [V2.S4], 16(R5)
    ADD     $4, R0, R0
    B       multipleDataLoop

singleDataLoop:
    CMP     R2, R0
    BGE     returnLength
    // Load one int32 value.
    MOVW    (R3), R6
    MOVW    (R4), R7
    // Compare and select minimum (signed)
    CMP     R6, R7
    CSEL    LE, R6, R7, R8  // ✅ LE pour le minimum
    MOVW    R8, (R5)
    // Advance pointers
    ADD     $4, R3, R3
    ADD     $4, R4, R4
    ADD     $4, R5, R5
    ADD     $1, R0, R0
    B       singleDataLoop
    
returnLength:
    MOVD    R2, int+72(FP)
    RET
