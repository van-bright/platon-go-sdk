package main

import (
	"fmt"
	"github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/web3go"
)

func main() {
	const platonEndpoint = "http://127.0.0.1:6789"
	web3g, err := web3go.New(platonEndpoint)

	p, err := web3g.GetSchnorrNIZKProve()
	fmt.Println("GetSchnorrNIZKProve: ", p, err)

	pv, err := web3g.AdminGetProgramVersion()
	fmt.Println("AdminGetProgramVersion: ", pv, err)

	tx, err := web3g.TransactionByBlockNumberAndIndex("latest", 0)
	fmt.Println("TransactionByBlockNumberAndIndex: ", tx, err)

	block, err := web3g.BlockByNumber("latest")
	fmt.Println("BlockByNumber: ", block, err)

	block, err = web3g.BlockByHash("0x733a072fd7d074f4116585c7be7c036c86ead3f453265b8c34ba56a99b9f6bcc")
	fmt.Println("BlockByHash: ", block, err)

	count, err := web3g.TransactionCountByNumber("latest")
	fmt.Println("TransactionCountByNumber: ", count, err)

	count, err = web3g.TransactionCountByHash(common.HexToHash("0x4b3b9abff5c129788aa18916328eadc8c02465a31237df27756c7e5c6dd72ce8"))
	fmt.Println("TransactionCountByHash: ", count, err)

	balance, err := web3g.BalanceAt("lat1t3jsgu5km95aeqfqxx396k46e2ejxcg442ltq0", "latest")
	fmt.Println("BalanceAt: ", balance, err)

	blockNum, err := web3g.BlockNumber()
	fmt.Println("BlockNumber: ", blockNum, err)

	accounts, err := web3g.Accounts()
	fmt.Println("Accounts: ", accounts, err)

	gasPrice, err := web3g.GasPrice()
	fmt.Println("GasPrice: ", gasPrice, err)

	syncing, err := web3g.Syncing()
	fmt.Println("Syncing: ", syncing != nil, err)

	protocolVer, err := web3g.ProtocolVersion()
	fmt.Println("ProtocolVersion: ", protocolVer, err)

	netPeerCount, err := web3g.NetPeerCount()
	fmt.Println("NetPeerCount: ", netPeerCount, err)

	netlistening, err := web3g.NetListening()
	fmt.Println("NetListening: ", netlistening, err)

	rsp, err := web3g.ClientVersion()
	fmt.Println("ClientVersion: ", rsp, err)

	info, err := web3g.AdminNodeInfo()
	fmt.Println("AdminNodeInfo: ", info, err)

	sha3, err := web3g.Sha3("hello world")
	fmt.Println("Sha3: ", sha3.Hex(), err)

	netver, err := web3g.NetworkID()
	fmt.Println("NetworkID: ", netver, err)
}
