#include "wrap.hxx"
#include "cn_slow_hash.hpp"

cn_v3_hash_t ctx;
void cn_hash(const void* in, size_t len, void* out) {
	ctx.hash(in, len, out);
}

