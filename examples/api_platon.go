package main

import (
	"fmt"
	"math/big"
	"platon-go-sdk/common"
	"platon-go-sdk/web3g"
)

func main() {
	//const alayaEndpoint = "https://openapi.alaya.network/rpc"

	const alayaEndpoint = "http://47.241.91.2:6789"

	//var masterRaw = common.HexToAddress("0x17d4C98a69d0141E3871E66BF79069623fdC2160")
	var masterAccount, _ = common.Bech32ToAddress("atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432")

	//var slaveRaw = common.HexToAddress("0x237a0C8507854E805C4a2dACF70d31F81ab944ED")
	var slaveAccount, _ = common.Bech32ToAddress("atp1ydaqepg8s48gqhz29kk0wrf3lqdtj38d8mkcz3")

	var geb3 = web3g.New(alayaEndpoint)

	//{
	//	rsp, _ := geb3.PlatonAccounts()
	//	fmt.Println("PlatonAccounts: ", rsp)
	//}
	{
		rsp, _ := geb3.PlatonGasPrice()
		fmt.Println("PlatonGasPrice: ", rsp)
	}
	//{
	//	rsp, _ := geb3.PlatonBlockNumber()
	//	fmt.Println("PlatonBlockNumber: ", rsp)
	//}
	//{
	//	if rsp, err := geb3.PlatonGetBalance("atx1gp7h8k9ynm4ct5ev73j4qlwhr4g8zqxp0rc7ym", "latest"); err != nil {
	//		fmt.Println("Error: ", err)
	//	} else {
	//		fmt.Println("PlatonBlockNumber: ", rsp)
	//	}
	//}
	//{
	//	{
	//		rsp, err := geb3.PlatonGetBlockByNumber(
	//			"0x1b4", true)
	//		if err != nil {
	//			fmt.Println("error: ", err)
	//		} else {
	//			//out, _ := json.Marshal(rsp)
	//			fmt.Println("block by number: ", rsp.Number.ToInt())
	//		}
	//	}
	//}
	{
		v, _ := big.NewInt(0).SetString("1000000000000000000", 10)
		tx := &web3g.PlatonSendTransactionReq{
			From:     masterAccount,
			To:       slaveAccount,
			Value:    (*web3g.BigInt)(v), // 1 ATP
		}
		resp, err := geb3.PlatonSendTransaction(tx)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		fmt.Printf(resp)
	}
}
