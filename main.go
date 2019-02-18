package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	alphabet   = regexp.MustCompile("^[0-9a-f]+$")
	numWorkers = runtime.NumCPU()
)

// Wallet stores private key and address containing desired substring at Index
type Wallet struct {
	AddressHex    string
	PrivateKeyHex string
	Index         int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(`
This tool generates Ethereum public and private keypair until it finds address
which contains required substring.

Usage:
	%s {part}
		Address part can contain only digits and letters from A to F.
		For fast results suggested length of address part is 4-6 characters.
`, os.Args[0])
		os.Exit(1)
	}
	part := strings.ToLower(os.Args[1])
	if !alphabet.MatchString(part) {
		fmt.Println("{part} must match the alphabet:", alphabet.String())
		os.Exit(2)
	}
	walletChan := make(chan Wallet)
	for i := 0; i < numWorkers; i++ {
		go generateWallet(part, walletChan)
	}
	for wallet := range walletChan {
		fmt.Printf(
			"Address: %s\u001b[32m%s\u001b[0m%s PrivateKey: %s\n",
			wallet.AddressHex[:wallet.Index],
			wallet.AddressHex[wallet.Index:len(part)+wallet.Index],
			wallet.AddressHex[len(part)+wallet.Index:],
			wallet.PrivateKeyHex)
	}
}

func generateWallet(part string, walletChan chan Wallet) {
	for {
		privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			log.Fatal(err)
		}
		address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
		idx := strings.Index(hex.EncodeToString(address[:]), part)
		if idx == -1 {
			continue
		}
		walletChan <- Wallet{
			AddressHex:    address.Hex(),
			PrivateKeyHex: hex.EncodeToString(crypto.FromECDSA(privateKeyECDSA)),
			Index:         idx + 2, // Hex() will add `0x` preffix
		}
	}
}
