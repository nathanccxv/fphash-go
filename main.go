package main

// #cgo CXXFLAGS: -std=c++11 -mavx2 -msse2 -maes
// #include "wrap.hxx"
import "C"

import "unsafe"
import "fmt"
import "time"

func main() {

	fmt.Println("variant_version:", C.variant_version())

	// var in = [10]uint8{2, 2, 3, 4, 5, 9, 7, 8, 9, 10};
	// var in = [3]uint8{34, 89, 11}
	var in = [0]uint8{}
	var out = [32]uint8{}

	var t = time.Now()

	const ITER = 100

	for i := 0; i < ITER; i++ {
		C.cn_hash(
			unsafe.Pointer(&in),
			0,
			unsafe.Pointer(&out[0]),
		)
	}

	fmt.Println("time", time.Since(t)/ITER)

	fmt.Println(in)
	fmt.Println(out)
}
