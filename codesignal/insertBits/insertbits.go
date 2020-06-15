package insertbits

import (
	"unsafe"
)

const nBits = unsafe.Sizeof(int(0)) * 8

type BitVector struct {
	bits []bool
}

// func (bv *BitVector) Reverse() {
// 	for i, j := 0, len(bv.bits)-1; i < j; i, j = i+1, j-1 {
// 		bv.bits[i], bv.bits[j] = bv.bits[j], bv.bits[i]
// 	}
// }

func (bv *BitVector) Insert(given *BitVector, a, b int) {
	for i, j := a, 0; i <= b && j < len(given.bits); i, j = i+1, j+1 {
		bv.bits[i] = given.bits[j]
	}
}

func (bv *BitVector) ToInt() int {
	ret := int(0)
	toggle := int(1)
	for _, v := range bv.bits {
		if v {
			ret ^= toggle
		}
		toggle = toggle << 1
	}
	return ret
}

func makeBitVector(n int) *BitVector {
	bv := BitVector{
		bits: make([]bool, 0, nBits),
	}

	for i := 0; i < int(nBits); i++ {
		bv.bits = append(bv.bits, n&1 > 0)
		n = n >> 1
	}

	return &bv
}

func insertBits(n int, a int, b int, k int) int {
	nBV := makeBitVector(n)

	kBV := makeBitVector(k)

	nBV.Insert(kBV, a, b)

	return nBV.ToInt()
}
