package main

import (
	"fmt"
	"platon-go-sdk/web3g"
)
func main() {
	const alayaEndpoint = "https://openapi.alaya.network/rpc"
	var web3g = web3g.New(alayaEndpoint)

	rsp, err := web3g.ClientVersion()
	fmt.Println("Version Info: ", rsp, err)
}
