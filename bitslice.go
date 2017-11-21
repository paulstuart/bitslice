package bitslice

import (
	"fmt"
)

// Bits contains the bit slice
type Bits struct {
	data []uint64
}

// NewBits returns a bitslice
func NewBits(len int64) *Bits {
	return &Bits{data: make([]uint64, (len+63)/64)}
}

// Set sets the bit at the given offset
func (b *Bits) Set(offset int) error {
	x := uint64(offset) / 64
	if x >= uint64(len(b.data)) {
		return fmt.Errorf("out of range: %d is greater than length %d", offset, len(b.data)*64)
	}
	bit := uint64(1 << (uint64(offset) % 64))
	b.data[x] |= bit
	return nil
}

// Clear clears the bit at the given offset
func (b *Bits) Clear(offset int) error {
	x := uint64(offset) / 64
	if x >= uint64(len(b.data)) {
		return fmt.Errorf("out of range: %d is greater than length %d", offset, len(b.data)*64)
	}
	bit := uint64(1 << (uint64(offset) % 64))
	b.data[x] |= ^bit
	return nil
}

// Get returns true if the bit at offset was set
func (b *Bits) Get(offset int) (bool, error) {
	x := uint64(offset) / 64
	if x >= uint64(len(b.data)) {
		return false, fmt.Errorf("out of range: %d is greater than length %d", offset, len(b.data)*64)
	}
	bit := uint64(1 << (uint64(offset) % 64))
	return (b.data[x] & bit) != 0, nil
}
