#include "textflag.h"
#include "funcdata.h"

TEXT ·asmfn(SB), NOSPLIT, $0
    RET

TEXT ·issue53(SB), NOSPLIT, $0
    ADDL R1, $0x8, R1
    ADDL R1, $0xff, R1
    ADDL R1, $0x100, R1
    ADDL R1, $0x7fff, R1
    ADDL R1, $-8, R1
    ADDL R1, $-0x8000, R1
    ADDL R1, $0x8000, R1
    ADDL R1, $0x7fffffffffffffff, R1
    ADDL R1, $-0x8001, R1
    ADDL R1, $-0x8000000000000000, R1
    ADDL R1, $0x8000000000000002, R1
    RET

TEXT ·issue55(SB), NOSPLIT, $0
        LDI  R0, $4096
l255:   ADDL R0, $-8, R0
        BNE  R0, l255
l256:   ADDL R1, $0x8000, R1
        BEQ  R1, l256
        RET

TEXT ·issue5502(SB), NOFRAME|NOSPLIT, $0
l255:   LDL R0, $capn-0x8001(SP)
        SUBL R0, R0, R0
        BNE  R0, l255
        RET

TEXT ·issue22(SB), NOSPLIT, $0
    JMP  2(PC)
    LDL R0, (ZERO)
    RET

