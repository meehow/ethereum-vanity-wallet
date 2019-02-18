# Ethereum Vanity Wallet Generator

This tool generates Ethereum public and private keypair until it finds an address
which contains required substring.

## Installation

```
go get github.com/meehow/ethereum-vanity-wallet
```

## Usage


➜  ~ ethereum-vanity-wallet **beef**

Address: 0x6BE5fC81**BEeF**1114C40f70D24656cE2A2CB2aD6d PrivateKey: 508927a75f04808a545106c0188ada5eee65ea26c479082fe017c587dfa5a868

Address: 0x2E6f8dc5aF609949D11a4fbfb648e333A5**bEEf**33 PrivateKey: 3d73e940b74c113ff9f8152c90b50366b1e0437f0d50e848b5e88a20621298cc

Address: 0xd5cFBe8aFB24872299810273dA3FAD35dA31**BEEF** PrivateKey: e8569d67cae4744b0abab829ad1cd83016496eb98e4b15cb26122a958b9b4ba4

Address: 0xE2E8**beEF**BA5bE50C7215C3540674972C6734814A PrivateKey: 5f9ab72c2728c3c7c6a75c6f8be5d76e8e5fa9b298902a59d7665986ac6758d7

^C

## Benchmarks

My laptop with i7-6500U CPU can generate around 6500 keys per second per core.

```
pkg: github.com/meehow/ethereum-vanity-wallet
BenchmarkNewKeyForDirectICAP-4            	      30	  35699828 ns/op
BenchmarkNewKeyForDirectICAPNotRandom-4   	   10000	    157199 ns/op
BenchmarkRawEcdsa-4                       	   10000	    152757 ns/op
BenchmarkRawEcdsaNotRandom-4              	   10000	    152196 ns/op
```
