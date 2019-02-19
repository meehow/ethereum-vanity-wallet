# Ethereum Vanity Wallet Generator

This tool generates Ethereum public and private keypair until it finds address
which contains required prefix and/or suffix.
Address part can contain only digits and letters from A to F.

## Installation

```
go get github.com/meehow/ethereum-vanity-wallet
```

## Usage

➜  ~ ethereum-vanity-wallet -p ab -s cd

Address: 0x**Ab**719Dc27043be01e9ef02026407163455Ae7f**Cd** PrivateKey: 0a2ab8789247f97b1a54919c3ff45dda4ec26ecc4f9a1b3df82f9c7090dafe83

Address: 0x**ab**56b21826316f56f1Ee7e2e69C489A51EB8A5**cd** PrivateKey: f9fd0529afeeb74e3918447cf69f4fadc3684306e613f92f38380548134e2760

Address: 0x**ab**f6AC87BBad30215359739A6EE4159F4c4a23**cd** PrivateKey: c5bc2fac7fb08cc883adf37c0e0edeb142978859668546ac859bf7b1e8fe8767

Address: 0x**aB**29C8117002497fe99926895DF3dCD2a8259C**cD** PrivateKey: 9713a384ab4514ebef0bf1c064ddd14b7aebbff22aee1ee0791776bf8ab37298

Address: 0x**aB**a8e517e35A695DA093D0272CE05A53b6BE05**cD** PrivateKey: 63d352dac2ce38679cb0a258c8820858046fd983a528110446466cb3a66c7752

Address: 0x**aB**3Bba66FCEF2Ba55ac0b7631fbee352362f3B**cd** PrivateKey: f9205f2f91f5a612f41f0348889f8b558ac8ebd2b5234627267c295fb9587aac

^C

[![asciicast](https://asciinema.org/a/228369.svg)](https://asciinema.org/a/228369)

## Benchmarks

My laptop with i7-6500U CPU can generate around 6500 keys per second per core.

```
pkg: github.com/meehow/ethereum-vanity-wallet
BenchmarkNewKeyForDirectICAP-4            	      30	  35699828 ns/op
BenchmarkNewKeyForDirectICAPNotRandom-4   	   10000	    157199 ns/op
BenchmarkRawEcdsa-4                       	   10000	    152757 ns/op
BenchmarkRawEcdsaNotRandom-4              	   10000	    152196 ns/op
```
