package main

import (
	"fmt"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/crypto"
	"platon-go-sdk/web3go"
)

/*
"mnemonics":"always brick access science decade nasty marriage attack fame topple pen add",
[
		{
			"private-key":"b72faaa798c44d2359d0ccb35dd39446c9c18905fdcecede42a6570ff177ae08",
			"eth-address":"0x17d4C98a69d0141E3871E66BF79069623fdC2160",
			"mainnet-addresses":"atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432",
			"testnet-addresses":"atx1zl2vnznf6q2puwr3ue4l0yrfvglacgtqw82lzq",
			"path":"m/44'/60'/0'/0/0"
		},
		{
			"private-key":"3d04b39cd4439340597fe5027c58797e824173cdab45b1f155539026f72b7029",
			"eth-address":"0x237a0C8507854E805C4a2dACF70d31F81ab944ED",
			"mainnet-addresses":"atp1ydaqepg8s48gqhz29kk0wrf3lqdtj38d8mkcz3",
			"testnet-addresses":"atx1ydaqepg8s48gqhz29kk0wrf3lqdtj38dda2j3m",
			"path":"m/44'/60'/0'/0/1"
		}
]
*/
const mnemonic = "always brick access science decade nasty marriage attack fame topple pen add"

func write() {
	wallet, err := web3go.NewWalletByMnemonics(mnemonic)
	if err != nil {
		fmt.Println("new wallet by mnemonic failed: ", err)
		return
	}

	accounts := wallet.Accounts()
	mainKey, _ := accounts[0].ToMainNetAddress()
	fmt.Println("account 0 main net address: ", mainKey)

	err = wallet.ExportToKeyStore(accounts[0], "/Users/liangqin/temp/keystore", "123456")
	if err != nil {
		fmt.Println("save account to key store failed: ", err)
	}
}

func load() {
	file := "/Users/liangqin/temp/keystore/UTC--2021-04-06T06-56-16.093914000Z--17d4c98a69d0141e3871e66bf79069623fdc2160"
	w1, _ := web3go.NewWalletByMnemonics(mnemonic)
	a, err := w1.ImportFromKeyStore(file, "123456", "123456")
	if err != nil {
		fmt.Println("import keystore error: ", err)
		return
	}
	fmt.Println("imported account ", a.Address.Hex())
	key, err := w1.ExportPrivateKey(a, "123456")
	if err != nil {
		fmt.Println("export private key error: ", err)
		return
	}
	fmt.Println("imported private key of account: ", hexutil.Encode(crypto.FromECDSA(key)))
}

func main() {
	//write()
	load()
}
