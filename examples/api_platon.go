package main

import (
	"fmt"
	"platon-go-sdk/web3g"
)
func main() {
	const alayaEndpoint = "https://openapi.alaya.network/rpc"
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
	{
		req := web3g.PlatonGetBalanceReq{"atx1gp7h8k9ynm4ct5ev73j4qlwhr4g8zqxp0rc7ym", "latest"}
		if rsp, err := geb3.PlatonGetBalance(req); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("PlatonBlockNumber: ", rsp)
		}
	}
}
