//go:build arm64
// +build arm64

// func Supported() bool
TEXT ·Supported(SB), 4, $0
    // Check NEON supported.
    MRS     ID_AA64PFR0_EL1, R0
    TST     $(1<<0), R0
    BEQ     sseFalse
    //sseTrue:
    MOVW    $1, R0
    MOVB    R0, bool+0(FP)
    RET
sseFalse:
    MOVW    $0, R0
    MOVB    R0, bool+0(FP)
    RET
