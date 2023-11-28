# Cryptonight-GPU-light Hash

Cryptonight-GPU-light is a variant of [Cryptonight-GPU](https://ryo-currency.com/cn-gpu/) POW (Proof of work) algorithm, with parameters below:

```
MEMORY = 32 * 1024
ITER = 0x300
```

This project provides Go bindings for the original C++ implementation of the Cryptonight-GPU algorithm can be found at [ryo-currency/ryo-currency](https://github.com/ryo-currency/ryo-currency/tree/master/src/crypto/pow_hash).


## Usage

To use these bindings in your Go project, import the package:

```go
import cngpu "github.com/nexis-dev/cn-gpu-go"

intput := []uint8{1, 2, 3, 4, 5}
result := cngpu.Hash(input) // [32]uint8
```


## Other resources

[CyberChain](https://github.com/cyberchain/ccx): Blockchain using Cryptonight-GPU-light hash algorithm.

[CCXminer](https://github.com/cyberchain/ccxminer): Cryptonight-GPU-light miner.
