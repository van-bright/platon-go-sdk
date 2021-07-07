package main

import (
	"fmt"
	"math/big"
	"platon-go-sdk/common"
	"platon-go-sdk/network"
	"platon-go-sdk/web3go"
)

func main() {
	const mnemonic = "always brick access science decade nasty marriage attack fame topple pen add"
	w, err := web3go.NewWalletByMnemonics(mnemonic)
	if err != nil {
		fmt.Println("import wall error: ", err)
		return
	}

	w.SetNetworkCfg(&network.DefaultTestNetConfig)

	accounts := w.Accounts()
	for _, account := range accounts {
		b, _ := w.BalanceOf(account.Address)
		addr, _ := account.ToMainNetAddress()
		fmt.Printf("balance of %s is %s\n", addr, b.String())
	}

	digest, err := w.Transfer(accounts[0].Address, common.MustBech32ToAddress("atp1ydaqepg8s48gqhz29kk0wrf3lqdtj38d8mkcz3"), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("transfer failed: ", err)
		return
	}

	fmt.Println("tx send: ", digest)
}
