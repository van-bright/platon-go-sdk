package main

import (
	"fmt"
	"github.com/oldmanfan/platon-go-sdk/web3go"
	"os"
	"time"
)

//http://35.247.155.162:6789 以及 ws://35.247.155.162:6790
func main() {
	const platonEndpoint = "ws://35.247.155.162:6790"
	web3g, err := web3go.New(platonEndpoint)

	if err != nil {
		fmt.Println("new web3 instance failed: ", err)
		os.Exit(1)
	}

	filterId, err := web3g.NewBlockFilter()
	if err != nil {
		fmt.Println("new filter failed: ", err)
		os.Exit(1)
	}

	time.Sleep(time.Duration(20) * time.Second)

	logs, err := web3g.GetFilterChanges(filterId)
	if err != nil {
		fmt.Println("get filter changes failed: ", err)
		os.Exit(1)
	}

	fmt.Println("block hash: ", string(logs))

	suc := web3g.UninstallFilter(filterId)
	fmt.Println("uninstall filter result: ", suc)
}
