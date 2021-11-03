package main

import (
	"fmt"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/web3go"
)

func main() {
	const platonEndpoint = "http://127.0.0.1:16789"
	web3g, _ := web3go.New(platonEndpoint)

	dates := "0x68656c6c6f20776f726c64abedef768765876543"
	rsp, _ := web3g.Sha3(dates)
	fmt.Println("hash: ", hexutil.Encode(rsp.Bytes()))
}
