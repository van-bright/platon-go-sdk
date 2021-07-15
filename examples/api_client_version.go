package main

import (
	"fmt"
	platongosdk "platon-go-sdk"
	"platon-go-sdk/web3go"
)

// const alayaEndpoint = "https://openapi.alaya.network/rpc"
func main() {
	const alayaEndpoint = "http://172.16.64.132:6789"
	web3g, err := web3go.New(alayaEndpoint)

	req := platongosdk.CallMsg2{
		From:     "atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a",
		To:       "atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a",
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     nil,
	}
	gaslimit, err := web3g.EstimateGasLimit(req)
	fmt.Println(gaslimit, err)

	//code, err := web3g.CodeAt("atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a", "latest")
	//fmt.Println(code, err)

	//count, err := web3g.TransactionCountByNumber("latest")
	//fmt.Println(count, err)

	//count, err := web3g.TransactionCountByHash(common.HexToHash("0x4b3b9abff5c129788aa18916328eadc8c02465a31237df27756c7e5c6dd72ce8"))
	//fmt.Println(count, err)

	//storage, err := web3g.StorageAt(
	//	"atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a",
	//	common.HexToHash("0x4b3b9abff5c129788aa18916328eadc8c02465a31237df27756c7e5c6dd72ce8"),
	//	"latest")
	//fmt.Println(storage, err)

	//balance, err := web3g.BalanceAt("atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a", "latest")
	//fmt.Println(balance, err)

	//blockNum, err := web3g.BlockNumber()
	//fmt.Println(blockNum, err)

	//accounts, err := web3g.Accounts()
	//fmt.Println(accounts, err)

	//gasPrice, err := web3g.GasPrice()
	//fmt.Println(gasPrice, err)

	//syncing, err := web3g.Syncing()
	//fmt.Println(syncing, err)

	//protocolVer, err := web3g.ProtocolVersion()
	//fmt.Println(protocolVer, err)

	//netPeerCount, err := web3g.NetPeerCount()
	//fmt.Println(netPeerCount, err)

	//netlistening, err := web3g.NetListening()
	//fmt.Println("network version: ", netlistening, err)

	//rsp, err := web3g.ClientVersion()
	//fmt.Println("Version Info: ", rsp, err)
	//
	//info, err := web3g.AdminNodeInfo()
	//fmt.Println("admin node info: ", info, err)

	//sha3, err := web3g.Sha3("hello world")
	//fmt.Println("sha3: ", sha3.Hex(), err)

	//netver, err := web3g.NetworkID()
	//fmt.Println("network version: ", netver, err)
}
