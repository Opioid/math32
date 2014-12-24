// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

func Lerp(x, y, t float32) float32 {
	_t := 1 - t
	return _t * x + t * y
}