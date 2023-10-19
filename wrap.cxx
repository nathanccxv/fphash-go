#include "wrap.hxx"
#include "cn_slow_hash.hpp"

cn_v4_hash_t ctx;
void cn_hash(const void* in, size_t len, void* out) {
	ctx.hash(in, len, out);
}

int variant_version() {
	return ctx.variant_version();
}
