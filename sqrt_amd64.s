// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cmd/ld/textflag.h"

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),NOSPLIT,$0
	SQRTSS x+0(FP), X0
	MOVSS X0, ret+8(FP)
	RET

// This is the fast RSQRTSS approximation refined with one Newton-Raphson iteration
// 	r := Rsqrt(x)
//	return r * (1.5 - (0.5 * x * r * r))
// func Rsqrt(x float32) float32
TEXT ·Rsqrt(SB),NOSPLIT,$0
	MOVSS  	x+0(FP), X0
	RSQRTSS X0, X1 	
	MOVSS 	X1, X2	
	MULSS 	X1, X1
	MULSS 	X0, X1
	MULSS	$(0.5), X1
	MOVSS   $(1.5), X0
	SUBSS	X1, X0
	MULSS	X0, X2
	MOVSS 	X2, ret+8(FP)
	RET

// func FastSqrt(x float32) float32
TEXT ·FastSqrt(SB),NOSPLIT,$0
	MOVSS  	x+0(FP), X0
	RSQRTSS X0, X1 	
	MOVSS 	X1, X2	
	MULSS 	X1, X1
	MULSS 	X0, X1
	MULSS	$(0.5), X1
	MOVSS   $(1.5), X3
	SUBSS	X1, X3
	MULSS	X3, X2
	MULSS	X2, X0
	MOVSS 	X0, ret+8(FP)
	RET
	
