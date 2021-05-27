#include "../../../../../runtime/textflag.h"
TEXT foo(SB), NOSPLIT, $0
	ADDW	R17, $0x4d, R0 //00 a0 29 4a
	S4ADDW	R17, R3, R1 // 41 00 23 42
	S4SUBL	R3, R9, R11 // 6b 01 69 40
	S8SUBL	R3, $0, R11 // 6b 01 60 48
	STB	R3, 4(R0)     // 04 00 60 a0
	SYS_CALL_B	$1      // 01 00 00 02
	BSR	R2, $52        // 0d 00 40 14
