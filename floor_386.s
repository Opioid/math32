// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cmd/ld/textflag.h"

// func Floor(x float32) float32
TEXT ·Floor(SB),NOSPLIT,$0
	JMP ·floor(SB)
	