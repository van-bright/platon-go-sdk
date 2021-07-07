package main

import (
	"fmt"
	"log"
	"math/big"
	"platon-go-sdk/common"
	"platon-go-sdk/core/types"
	"platon-go-sdk/crypto"
	"platon-go-sdk/web3go"
)

/*
"mnemonics":"always brick access science decade nasty marriage attack fame t    opple pen add",
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
//const AlayaEndpoint = "https://openapi.alaya.network/rpc"
//const AlayaEndpoint = "http://47.241.91.2:6789"

func main() {
	geb3, err := web3go.New(AlayaEndpoint)
	if err != nil {
		log.Fatal("NewCredential error:", err)
	}

	privateKey, err := crypto.HexToECDSA("b72faaa798c44d2359d0ccb35dd39446c9c18905fdcecede42a6570ff177ae08")
	if err != nil {
		log.Fatal(err)
	}

	sp, _ := geb3.GasPrice()
	fmt.Println("sp: ", sp.String())

	nwid, _ := geb3.NetworkID()
	fmt.Println("nwid: ", nwid.String())

	nonce, _ := geb3.NonceAt("atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432", "pending")
	fmt.Println("nonce: ", nonce)

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice := big.NewInt(1000000000)

	toAddress := common.MustBech32ToAddress("atp1ydaqepg8s48gqhz29kk0wrf3lqdtj38d8mkcz3")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID := big.NewInt(201030)

	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	from, err := types.Sender(signer, signedTx)
	log.Println("from: ", from.Bech32WithPrefix("atp"))

	_, err = geb3.SendRawTransaction(signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
