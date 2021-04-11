package main

import (
	"fmt"
	"platon-go-sdk/web3go"
)

func main() {
	// const alayaEndpoint = "https://openapi.alaya.network/rpc"
	const alayaEndpoint = "http://172.16.64.132:6789"
	web3g, err := web3go.New(alayaEndpoint)

	rsp, err := web3g.ClientVersion()
	fmt.Println("Version Info: ", rsp, err)

	info, err := web3g.AdminNodeInfo()
	fmt.Println("admin node info: ", info, err)
}
