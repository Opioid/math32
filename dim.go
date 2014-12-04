// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

func Dim(x, y float32) float32 {
	return Max(x - y, 0)
}

func Max(x, y float32) float32 {
	if x > y {
		return x
	} 
	return y
}

func Min(x, y float32) float32 {
	if x < y {
		return x
	} 
	return y
}

func Clamp(x, mi, ma float32) float32 {
	if x < mi {
		return mi
	} else if x > ma {
		return ma
	}
	return x
}