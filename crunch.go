package main

import (
	"fmt"
	"math/bits"
	"math/rand"
)

func permute(num uint32, seeds [16]uint32) uint32 {
	for i := 0; i < len(seeds); i++ {
		num ^= seeds[i]
		num = bits.RotateLeft32(num, 9)
		i++
		num += seeds[i]
		num = bits.RotateLeft32(num, 21)
		num ^= num >> 17
		num = bits.RotateLeft32(num, 13)
	}
	return num
}

func main() {
	var seeds [16]uint32
	for i := range seeds {
		seeds[i] = rand.Uint32()
	}

	// test with first 10 uint32 values
	for num := uint32(0); num < 10; num++ {
		fmt.Println(permute(num, seeds))
	}
}
