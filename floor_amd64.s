// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#define Big		0x4b000000 // 2**23

#include "cmd/ld/textflag.h"

// Adding the commented lines in the function body back in 
// should result in behavior exactly as in the go reference implementation
//
// func Floor(x float32) float32
TEXT ·Floor(SB),NOSPLIT,$0
	MOVL	x+0(FP), AX
//	MOVL	$0x7FFFFFFF, DX // sign bit mask
//	ANDL	AX,DX // DX = |x|
//	SUBL	$1,DX
//	MOVL    $(Big - 1), CX // if |x| >= 2**23-1 or IsNaN(x) or |x| == 0, return x
//	CMPL	DX,CX
//	JAE     isBig_floor
	MOVL	AX, X0 // X0 = x
	CVTTSS2SL	X0, AX
	CVTSL2SS	AX, X1 // X1 = float(int(x))
	CMPSS	X1, X0, 1 // compare LT; X0 = 0xFFFFFFFF or 0
	MOVSS	$(-1.0), X2
	ANDPS	X2, X0 // if x < float(int(x)) {X0 = -1} else {X0 = 0}
	ADDSS	X1, X0
	MOVL	X0, ret+8(FP)
	RET
isBig_floor: 
	MOVL    AX, ret+8(FP) // return x
	RET

// func Floor2(x, y float32) (float32, float32)
TEXT ·Floor2(SB),NOSPLIT,$0
	MOVLPS  y+0(FP), X0
	MOVQ	X0, ret+8(FP)
	RET
