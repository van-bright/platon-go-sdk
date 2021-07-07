package main

import (
	"fmt"
	"github.com/AlayaNetwork/Alaya-Go/accounts/abi/bind"
	"log"
	"platon-go-sdk/contracts"
	"platon-go-sdk/examples/store"
	"time"

	"github.com/AlayaNetwork/Alaya-Go/common"
)

const AlayaEndpoint = "http://172.16.64.132:6789"
const privateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"

func toDeployContract(opts *bind.TransactOpts, client bind.ContractBackend) string {
	input := "1.0"
	address, tx, _, err := store.DeployStore(opts, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	// wait until tx confirmed
	time.Sleep(time.Duration(10) * time.Second)
	return address.Hex()
}

func toCallContractMethod(hexContractAddr string, opts *bind.TransactOpts, client bind.ContractBackend) {
	addr := common.HexToAddress(hexContractAddr)
	instance, err := store.NewStore(addr, client)
	if err != nil {
		log.Fatal("new instance failed: ", err)
	}
	// to query version
	ver, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("version of contract is ", ver)

	// to set new items
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("hello"))
	copy(value[:], []byte("kitty"))

	tx, err := instance.SetItem(opts, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	time.Sleep(time.Duration(10) * time.Second)

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // kitty
}

func main() {
	contract := contracts.Contract{
		Url:        AlayaEndpoint,
		PrivateKey: privateKey,
	}

	opts, client, err := contract.Init()
	if err != nil {
		log.Fatal(err)
	}

	addr := toDeployContract(opts, client)
	toCallContractMethod(addr, opts, client)
}
