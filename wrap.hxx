#ifndef WRAP_H
#define WRAP_H

#include <stddef.h>

// __cplusplus tells the compiler that inside code is compiled with the c++ compiler
#ifdef __cplusplus
// extern "C" tells C++ compiler exports the symbols without a name manging.
extern "C" {
  #endif
  int variant_version();
  void cn_hash(const void* in, size_t len, void* out);
  #ifdef __cplusplus
}
#endif
#endif
