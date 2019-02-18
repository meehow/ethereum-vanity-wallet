package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	notRandom, _ = hex.DecodeString("00337e8525ef65df472ae0ed034f2247c811640e985226929c7f2e96cbef0c6361386a4cb4c4aa227855aca3c616130d17b19409a9b5524f5f0cc89f4ce2ca56")
)

//keystore.NewKeyForDirectICAP(bytes.NewReader(notRandom)).Address.Hex()

func BenchmarkNewKeyForDirectICAP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keystore.NewKeyForDirectICAP(rand.Reader).Address.Hex()
	}

}
func BenchmarkNewKeyForDirectICAPNotRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keystore.NewKeyForDirectICAP(bytes.NewReader(notRandom)).Address.Hex()
	}
}

func BenchmarkRawEcdsa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		privateKeyECDSA, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	}
}

func BenchmarkRawEcdsaNotRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		privateKeyECDSA, _ := ecdsa.GenerateKey(crypto.S256(), bytes.NewReader(notRandom))
		crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	}
}
