// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

func Floor(x float32) float32

func floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

func Floor2(x, y float32) (float32, float32)

func PlainFloor2(x, y float32) (float32, float32) {
	return Floor(x), Floor(y)
}

func Ceil(x float32) float32 {
	return float32(math.Ceil(float64(x)))
}