package main

import (
	"fmt"
	"platon-go-sdk/common"
	"platon-go-sdk/core/types"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/typedefs"
	"platon-go-sdk/web3go"
)

func main() {
	const platonEndpoint = "http://172.16.64.132:6789"
	web3g, err := web3go.New(platonEndpoint)

	const PrivateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"
	var credentials, _ = typedefs.NewCredential(PrivateKey, network.MainNetHrp)
	tx := types.NewTransaction(0, common.MustBech32ToAddress("atp1qv5ffg7z3h42zt4e035vja65d86al7q0nr9s0g"), nil, 0, nil, nil)
	signedTx, _ := credentials.SignTx(tx, network.MainNetChainId)
	rsp, err := web3g.SendTransaction(signedTx)
	fmt.Println(rsp, err)

	//from := common.MustBech32ToAddress("atp1r93a6kug4nx637rvpkdyslpkenwq4ws0t0g884")
	//to := common.MustBech32ToAddress("atp1qv5ffg7z3h42zt4e035vja65d86al7q0nr9s0g")
	//
	//msg := platongosdk.CallMsg{
	//	From:     from,
	//	To:       &to,
	//	Gas:      0,
	//	GasPrice: nil,
	//	Value:    nil,
	//	Data:     nil,
	//}
	//
	//rsp, err := web3g.CallContract(msg, "latest")
	//fmt.Println(rsp, err)
	//tx := types.NewTransaction(
	//
	//	)
	//reeipt, err := web3g.SendTransaction(tx)
	//fmt.Println(reeipt, err)

	//p, err := web3g.GetSchnorrNIZKProve()
	//fmt.Println(p, err)

	//pv, err := web3g.AdminGetProgramVersion()
	//fmt.Println(pv, err)

	//tx, err := web3g.TransactionByBlockNumberAndIndex("latest", 0)
	//fmt.Println(tx, err)

	//block, err := web3g.BlockByNumber("latest")
	//fmt.Println(block, err)

	//block, err := web3g.BlockByHash("0x733a072fd7d074f4116585c7be7c036c86ead3f453265b8c34ba56a99b9f6bcc")
	//fmt.Println(block, err)

	//req := platongosdk.CallMsg2{
	//	From:     "atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a",
	//	To:       "atp1fgdm0vsevzc8wy2094vmw4dtpdnph25j6l9e8a",
	//	Gas:      0,
	//	GasPrice: nil,
	//	Value:    nil,
	//	Data:     nil,
	//}
	//gaslimit, err := web3g.EstimateGasLimit(req)
	//fmt.Println(gaslimit, err)

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
