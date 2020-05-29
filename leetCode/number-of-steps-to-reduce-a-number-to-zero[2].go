package main

import "math/bits"

func numberOfSteps(num int) int {

	u := uint(num)

	// bits.UintSize - bits.LeadingZeros(u); need these many shifts to be zero

	// bits.OnesCount(u) - 1; need these many ops for least significant bit 1
	// because only 1 are those bits which is going to be 1 at significant bit

	return bits.UintSize - bits.LeadingZeros(u) + bits.OnesCount(u) - 1
}
