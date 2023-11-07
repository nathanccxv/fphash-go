package cngpugo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHash(t *testing.T) {

	fmt.Println("variant_version:", Variant_version())

	input := []uint8{12, 89, 67, 90, 3, 156}
	fmt.Println(Hash(input))

	var tt = time.Now()

	const ITER = 100
	for i := 0; i < ITER; i++ {
		in_len := rand.Intn(50)
		input := make([]uint8, in_len, in_len)
		for i := 0; i < in_len; i++ {
			input[i] = uint8(rand.Intn(256))
		}

		Hash(input)
		// fmt.Println(out)
	}

	fmt.Println("time", time.Since(tt)/ITER)

}
