package utils

import "math"

type BitArray struct {
	size int
	bits []uint8
}

func NewBitArray(size int) (ba *BitArray) {
	ba = new(BitArray)
	ba.size = size + 1
	ba.bits = make([]uint8, int(math.Ceil(float64(ba.size)/8.0)))

	return
}

func (ba *BitArray) SetBit(n int) {
	i, j := uint(n/8), uint(n%8)
	ba.bits[i] |= 1 << j
}

func (ba *BitArray) IsSet(n int) bool {
	i, j := uint(n/8), uint(n%8)
	return ba.bits[i]&(1<<j) != 0
}
