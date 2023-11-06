package cngpugo

// #cgo CXXFLAGS: -std=c++11 -mavx2 -msse2 -maes
// #include "cn-gpu.hxx"
import "C"

import (
	"unsafe"
	"sync"
)

type Ctx struct {
	cn_ctx unsafe.Pointer
}

var ctx *Ctx
var once sync.Once

func getCtx() *Ctx{
	once.Do(func() {
		ctx = &Ctx{cn_ctx: C.new_ctx()}
	})
	return ctx
}


func Variant_version() int {
	return int(C.variant_version(getCtx().cn_ctx))
}


func Hash(input []uint8) [32]uint8 {
	out := [32]uint8{}
	var in_ptr unsafe.Pointer
	if len(input) != 0 {
		in_ptr = unsafe.Pointer(&input[0])
	}
	C.cn_hash(
		getCtx().cn_ctx,
		in_ptr,
		C.size_t(len(input)),
		unsafe.Pointer(&out[0]),
	)
	return out
}

