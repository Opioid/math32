// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cmd/ld/textflag.h"

// func Abs(x float32) float32
TEXT ·Abs(SB),NOSPLIT,$0
	MOVL   $(1<<31), BX
	MOVL   BX, X0 // movsd $(-0.0), x0
	MOVSS  x+0(FP), X1
	ANDNPS X1, X0
	MOVSS  X0, ret+8(FP)
	RET

//	FMOVF   x+0(FP), F0  // F0=x
//	FABS                 // F0=|x|
//	FMOVFP  F0, ret+8(FP)
//	RET
