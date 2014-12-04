// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import (
	
)

// http://http.developer.nvidia.com/Cg/atan2.html
func Atan2(x, y float32) float32 {
	ax := Abs(x)
	ay := Abs(y)  

	t0 := Max(ax, ay)
	t1 := Min(ax, ay)
	t2 := 1 / t0
	t2 = t1 * t2

	t1 = t2 * t2
	t0 =         - 0.013480470
	t0 = t0 * t1 + 0.057477314
	t0 = t0 * t1 - 0.121239071
	t0 = t0 * t1 + 0.195635925
	t0 = t0 * t1 - 0.332994597
	t0 = t0 * t1 + 0.999995630
	t2 = t0 * t2

	if ax > ay {
		t2 = 1.570796327 - t2
	}

	if y < 0 {
		t2 = 3.141592654 - t2
	}

	if x < 0 {
		t2 = -t2
	}

	return t2
}