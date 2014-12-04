// Copyright 2014 Benjamin Savs.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

// sin coefficients
var _sin = [...]float32{
	1.58962301576546568060E-10, // 0x3de5d8fd1fd19ccd
	-2.50507477628578072866E-8, // 0xbe5ae5e5a9291f5d
	2.75573136213857245213E-6,  // 0x3ec71de3567d48a1
	-1.98412698295895385996E-4, // 0xbf2a01a019bfdf03
	8.33333333332211858878E-3,  // 0x3f8111111110f7d0
	-1.66666666666666307295E-1, // 0xbfc5555555555548
}

// cos coefficients
var _cos = [...]float32{
	-1.13585365213876817300E-11, // 0xbda8fa49a0861a9b
	2.08757008419747316778E-9,   // 0x3e21ee9d7b4e3f05
	-2.75573141792967388112E-7,  // 0xbe927e4f7eac4bc6
	2.48015872888517045348E-5,   // 0x3efa01a019c844f5
	-1.38888888888730564116E-3,  // 0xbf56c16c16c14f91
	4.16666666666665929218E-2,   // 0x3fa555555555554b
}

func Sin(x float32) float32 {
	const (
		PI4A = 7.85398125648498535156E-1                             // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668E-8                             // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645E-15                            // 0x3ce8469898cc5170,
		M4PI = 1.273239544735162542821171882678754627704620361328125 // 4/pi
	)
/*	// special cases
	switch {
	case x == 0 || IsNaN(x):
		return x // return ±0 || NaN()
	case IsInf(x, 0):
		return NaN()
	}
*/
	// make argument positive but save the sign
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	j := int32(x * M4PI) // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := float32(j)      // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j += 1
		y += 1
	}
	j &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}

	z := ((x - y*PI4A) - y*PI4B) - y*PI4C // Extended precision modular arithmetic
	zz := z * z
	if j == 1 || j == 2 {
		y = 1.0 - 0.5*zz + zz*zz*((((((_cos[0]*zz)+_cos[1])*zz+_cos[2])*zz+_cos[3])*zz+_cos[4])*zz+_cos[5])
	} else {
		y = z + z*zz*((((((_sin[0]*zz)+_sin[1])*zz+_sin[2])*zz+_sin[3])*zz+_sin[4])*zz+_sin[5])
	}
	if sign {
		y = -y
	}
	return y
}