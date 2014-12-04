// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import (
	"math"
)

// http://http.developer.nvidia.com/Cg/acos.html
// Absolute error <= 6.7e-5
func Acos(x float32) float32 {
	negate := float32(0)
	if x < 0 {
		negate = float32(1)
	}

	x = Abs(x)
	ret := float32(-0.0187293)
	ret *= x
	ret += 0.0742610
	ret *= x
	ret -= 0.2121144
	ret *= x
	ret += 1.5707288
	ret *= Sqrt(1 - x)
	ret -= 2 * negate * ret
	return negate * math.Pi + ret
}