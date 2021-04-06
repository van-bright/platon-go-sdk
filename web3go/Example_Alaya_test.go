package web3go

import (
	"encoding/json"
	"fmt"
)

const alayaEndpoint = "http://47.241.91.2:6789"

//const alayaEndpoint = "https://openapi.alaya.network/rpc"
var client, _ = New(alayaEndpoint)

func ExampleAlaya_ClientVersion() {
	ver, _ := client.ClientVersion()

	fmt.Println(ver)
	// Output: PlatONnetwork/alaya-47.241.91.2/v0.15.1-unstable-34c3857c/linux-amd64/go1.13.4
}

func ExampleAlaya_NetListening() {
	listening, _ := client.NetListening()

	fmt.Println(listening)
	// Output: true
}

func ExampleAlaya_NetPeerCount() {
	count, _ := client.NetPeerCount()

	fmt.Println(count)
	// Output: 0
}

func ExampleAlaya_Sha3() {
	raw := "0x68656c6c6f20776f726c64abedef768765876543"

	hash, _ := client.Sha3(raw)

	fmt.Println(hash.Hex())

	// Output: 0x71f0c36e6a89579a8c7ea07d62d120c00b326934699e918a8a41d59b0a113607
}

func ExampleAlaya_ProtocolVersion() {
	ver, err := client.ProtocolVersion()
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(ver)

	// Output: 63
}

func ExampleAlaya_Syncing() {
	syncProgress, err := client.Syncing()
	if err != nil {
		fmt.Println(err)
		// Output:
	} else if syncProgress == nil && err == nil {
		fmt.Println("", false)
		// Output: false
	} else {
		s, _ := json.Marshal(syncProgress)
		fmt.Println(string(s))
		// Output: xx
	}
}

func ExampleAlaya_GasPrice() {
	gasPrice, _ := client.GasPrice()
	fmt.Println(gasPrice)
	// Output: 1000000000
}

func ExampleAlaya_BlockNumber() {
	num, err := client.BlockNumber()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
	// Output: 1180529
}

func ExampleAlaya_BalanceOf() {
	balance, err := client.BalanceAt("atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432", "latest")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)
	// Output: 200000000000000000000
}

func ExampleAlaya_NetworkID() {
	id, err := client.NetworkID()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	// Output: 1
}
