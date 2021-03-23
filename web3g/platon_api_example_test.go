package web3g

import (
	"fmt"
	"math/big"
)

const alayaTestNet = "http://47.241.91.2:6789"

const masterAccount = "atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432"
const slaveAccount = ""
func ExampleWeb3g_PlatonGetBalance() {
	w3g := New(alayaTestNet)

	resp, _ := w3g.PlatonGetBalance(masterAccount, "latest")
	fmt.Println(resp) //
	// Output: 200000000000000000000
}

func ExampleWeb3g_PlatonSendTransaction() {
	w3g := New(alayaTestNet)
	tx := &PlatonSendTransactionReq{
		From:     masterAccount,
		To:       slaveAccount,
		Value:    (*BigInt)(big.NewInt(1000000000000000000)), // 1 ATP
	}
	resp, err := w3g.PlatonSendTransaction(tx)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf(resp)
	// Output: xxxxx
}
