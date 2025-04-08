package fphash

// #cgo CFLAGS: -I${SRCDIR}/include_boost_107400
// #cgo CXXFLAGS: -I${SRCDIR}/include_boost_107400
// #cgo CXXFLAGS: -std=c++11 -mavx2 -msse2 -maes
// #cgo !darwin LDFLAGS: -static-libstdc++ -static-libgcc
// #cgo darwin LDFLAGS: -static-libstdc++
// #include "fphash.hxx"
import "C"

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// For testing purposes
var (
	ctxCreationCount  int32
	ctxFinalizerCount int32
)

// GetCtxStats returns the number of contexts created and finalized
func GetCtxStats() (created, finalized int32) {
	return atomic.LoadInt32(&ctxCreationCount), atomic.LoadInt32(&ctxFinalizerCount)
}

type Ctx struct {
	cn_ctx unsafe.Pointer
}

func (c *Ctx) finalizer() {
	if c == nil || c.cn_ctx == nil {
		return
	}
	atomic.AddInt32(&ctxFinalizerCount, 1)
	C.del_ctx(c.cn_ctx)
}

func newCtx() any {
	atomic.AddInt32(&ctxCreationCount, 1)
	c := &Ctx{cn_ctx: C.new_ctx()}
	runtime.SetFinalizer(c, (*Ctx).finalizer)
	return c
}

// Ctx pool
var ctxPool = sync.Pool{
	New: newCtx,
}

// Singleton ctx
var ctx *Ctx
var once sync.Once

func getCtx() *Ctx {
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
	ctx := ctxPool.Get().(*Ctx)
	C.cn_hash(
		ctx.cn_ctx,
		in_ptr,
		C.size_t(len(input)),
		unsafe.Pointer(&out[0]),
	)
	ctxPool.Put(ctx)
	return out
}
