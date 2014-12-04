// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

func Sqrt(x float32) float32

func sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}