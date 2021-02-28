package main

import (
	"fmt"
	"platon-go-sdk/web3g"
)

func main() {
	const alayaEndpoint = "https://openapi.alaya.network/rpc"
	var web3g = web3g.New(alayaEndpoint)

	dates := "0x68656c6c6f20776f726c64abedef768765876543"
	rsp, _ := web3g.Sha3(dates)
	fmt.Println("hash: ", rsp)
}
