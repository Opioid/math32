// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

// Beware that this returns an uint32 and not a bool like the go float64 version!
func Signbit(x float32) uint32 {
	if x < 0.0 {
		return 1
	} else {
		return 0
	}
}