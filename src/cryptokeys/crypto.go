/**
 * Copyright (c) Peter Robinson 2024
 *
 * Program to take a BIP39 Mnemonic (a seed phrase) and a BIP44 path 
 * and print out the resulting key information.
 *
 */
package main

import (
    "fmt"
	"os"
	"strings"
    "github.com/LinX-OpenNetwork/coinutil/bip39"
    "github.com/LinX-OpenNetwork/coinutil/bip44"
	"github.com/LinX-OpenNetwork/coinutil/wallet"
)

func main() {
	fmt.Println("Crypto Keys")
	
	words := os.Getenv("WORDS")
	if words == "" {
		usage()
	}

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		usage()
	}
	rawPath := argsWithoutProg[0]

	wordsAsArray := strings.Split(words, " ")
	if len(wordsAsArray) != 12 && len(wordsAsArray) != 24 {
		fmt.Printf("WARNING: Mnemonics are typically 12 or 24 words long. Yours is %v words long.\n", len(wordsAsArray))
	}

	mnemonic := &bip39.Mnemonic{
		Words: wordsAsArray,
		Language: bip39.English,
	}
	path := bip44.DerivePath(rawPath)


    fmt.Println("Input:")
    fmt.Printf(" Seed phrase: %v\n", mnemonic.Sentence())
    fmt.Printf(" BIP44 Path: %v\n", path)

	hdWallet, _ := wallet.NewHDWallet(mnemonic)
	bip32Key, _ := hdWallet.KeyForDerivePath(path)
	key, _ := wallet.NewEthereumWalletFromKey(bip32Key)

    fmt.Println("Output:")
	fmt.Printf(" BIP39 Seed: %v\n", mnemonic.GenerateSeed(""))
	fmt.Printf(" Address:     %v\n", key.Address())
	//fmt.Printf(" Public Key:  %v\n", key.PublicKey())
	fmt.Printf(" Private Key: 0x%x\n", key.PrivateKey().D)
}


func usage() {
	fmt.Println("Usage:")
	fmt.Println(" Specify BIP39 mnemonic (i.e the seed phrase) is the environment variable WORDS")
	fmt.Println(" Specify the BIP44 path as the first parameter. For instance: \"m/44'/60'/0'/0/0\"")
	os.Exit(1)
}