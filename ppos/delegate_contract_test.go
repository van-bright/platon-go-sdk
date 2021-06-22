package ppos

import (
	"encoding/json"
	"fmt"
	"math/big"
	"platon-go-sdk/network"
	common2 "platon-go-sdk/ppos/common"
	"testing"
)

func TestDelegateContract_Delegate(t *testing.T) {
	config := network.PposMainNetParams
	dc := NewDelegateContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"

	amount := new(big.Int)
	amount.SetString("200000000000000000000", 10)
	list, err := dc.Delegate(nodeId, common2.FREE_AMOUNT_TYPE, amount)
	if err != nil {
		t.Errorf("DelegateContract.Delegate failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}

func TestDelegateContract_UnDelegate(t *testing.T) {
	config := network.PposMainNetParams
	dc := NewDelegateContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	stakingBlockNumber := big.NewInt(129518)
	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10)
	list, err := dc.UnDelegate(nodeId, stakingBlockNumber, amount)
	if err != nil {
		t.Errorf("DelegateContract.UnDelegate failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}
