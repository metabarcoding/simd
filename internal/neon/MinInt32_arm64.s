//go:build arm64
// +build arm64

// func MinInt32(left, right, result []int32) int
TEXT ·MinInt32(SB), 4, $0
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
    VLD1    (R3), [V0.S4]
    VLD1    (R4), [V1.S4]
    // SMIN V2.4S, V0.4S, V1.4S - signed minimum
    WORD    $0x4ea16c02  // SMIN opcode correct
    VST1    [V2.S4], (R5)
    ADD     $16, R3, R3  // Avancer de 16 octets (4 * 4)
    ADD     $16, R4, R4
    ADD     $16, R5, R5
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
    CSEL    LT, R7, R6, R8
storeResult:
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
