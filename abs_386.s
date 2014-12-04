// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cmd/ld/textflag.h"

// func Abs(x float32) float32
TEXT ·Abs(SB),NOSPLIT,$0
	FMOVF   x+0(FP), F0  // F0=x
	FABS                 // F0=|x|
	FMOVFP  F0, ret+8(FP)
	RET
