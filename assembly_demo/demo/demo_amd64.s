#include "textflag.h"

GLOBL ·Num(SB),NOPTR,$16

DATA ·Num+0(SB)/8,$123
DATA ·Num+8(SB)/8,$234


GLOBL ·Flag(SB),NOPTR,$1
DATA ·Flag+0(SB)/1,$1


// func Swap(a, b int) (int, int)
TEXT ·Swap(SB), NOSPLIT, $0-32
