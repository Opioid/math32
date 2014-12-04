// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cmd/ld/textflag.h"

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),NOSPLIT,$0
	FMOVS   x+0(FP),F0
	FSQRT
	FMOVSP  F0,ret+8(FP)
	RET
