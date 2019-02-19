package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	colorGreen = "\u001b[32m"
	colorReset = "\u001b[0m"
)

var (
	alphabet   = regexp.MustCompile("^[0-9a-f]*$")
	numWorkers = runtime.NumCPU()
)

// Wallet stores private key and address containing desired substring at Index
type Wallet struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func main() {
	var one bool
	var prefix, suffix string
	flag.BoolVar(&one, "one", false, "Stop after finding first address")
	flag.StringVar(&prefix, "p", "", "Public address prefix")
	flag.StringVar(&suffix, "s", "", "Public address suffix")
	flag.Parse()
	if prefix == "" && suffix == "" {
		fmt.Printf(`
This tool generates Ethereum public and private keypair until it finds address
which contains required prefix and/or suffix.
Address part can contain only digits and letters from A to F.
For fast results suggested length of sum of preffix and suffix is 4-6 characters.
If you want more, be patient.

Usage:

`)
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !alphabet.MatchString(prefix) {
		fmt.Println("Prefix must match the alphabet:", alphabet.String())
		os.Exit(2)
	}
	if !alphabet.MatchString(suffix) {
		fmt.Println("Suffix must match the alphabet:", alphabet.String())
		os.Exit(3)
	}
	walletChan := make(chan Wallet)
	for i := 0; i < numWorkers; i++ {
		go generateWallet(prefix, suffix, walletChan)
	}
	for w := range walletChan {
		addressHex := w.Address.Hex()[2:]
		privateKeyHex := hex.EncodeToString(crypto.FromECDSA(w.PrivateKey))
		fmt.Printf(
			"Address: 0x%s%s%s%s%s%s%s PrivateKey: %s\n",
			colorGreen,
			addressHex[:len(prefix)],
			colorReset,
			addressHex[len(prefix):len(addressHex)-len(suffix)],
			colorGreen,
			addressHex[len(addressHex)-len(suffix):],
			colorReset,
			privateKeyHex)
		if one {
			break
		}
	}
}

func generateWallet(prefix, suffix string, walletChan chan Wallet) {
	for {
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			log.Fatal(err)
		}
		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		addressHex := hex.EncodeToString(address[:])
		if prefix != "" && !strings.HasPrefix(addressHex, prefix) {
			continue
		}
		if suffix != "" && !strings.HasSuffix(addressHex, suffix) {
			continue
		}
		walletChan <- Wallet{
			Address:    address,
			PrivateKey: privateKey,
		}
	}
}
