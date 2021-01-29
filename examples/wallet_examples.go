package main

import (
	"encoding/hex"
	"fmt"
	"platon-go-sdk/common"
	"platon-go-sdk/crypto"
)

func main() {
	//entropy, _ := bip39.NewEntropy(128)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//mnemonic := "target eight invite uncover obtain donate damage shield shield witness brown enable"

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	//seed, err := bip39.EntropyFromMnemonic(mnemonic)
	//fmt.Println(len(seed))
	//if err != nil {
	//	fmt.Println("mne error: ", err)
	//	return
	//}
	//masterKey, _ := bip32.NewMasterKey(seed)
	//publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	//fmt.Println("Mnemonic: ", mnemonic)
	//fmt.Println("Master private key: ", masterKey)
	//fmt.Println("Master public key: ", publicKey)
	//addr, err := crypto.ToECDSA(seed)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//pubkey := addr.PublicKey
	addr, _ := crypto.GenerateKey()
	//fmt.Println("mnemonic: ", mnemonic)
	fmt.Println("private key: ", hex.EncodeToString(addr.D.Bytes()))
	//fmt.Println("public key: ", crypto.PubkeyToAddress(addr.PublicKey).Hex())
	fmt.Println("public key: ", crypto.PubkeyToAddress(addr.PublicKey).Bech32WithPrefix(common.MainNetAddressPrefix))
}
