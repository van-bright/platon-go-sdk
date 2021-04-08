package main

import (
	"fmt"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/web3go"
)

func main() {
	//const alayaEndpoint = "https://openapi.alaya.network/rpc"
	const alayaEndpoint = "http://47.241.91.2:6789"
	web3g, _ := web3go.New(alayaEndpoint)

	dates := "0x68656c6c6f20776f726c64abedef768765876543"
	rsp, _ := web3g.Sha3(dates)
	fmt.Println("hash: ", hexutil.Encode(rsp.Bytes()))
}
