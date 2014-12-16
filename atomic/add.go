package atomic

import (
	"sync/atomic"
	"unsafe"
	"math"
)

func AddFloat32(addr *float32, delta float32) (new float32) {
	unsafeAddr := (*uint32)(unsafe.Pointer(addr))

	for {
		old := *addr
		new  = old + delta

		if atomic.CompareAndSwapUint32(unsafeAddr, math.Float32bits(old), math.Float32bits(new)) {
			return
		}
	}
}