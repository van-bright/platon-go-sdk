package ppos

import (
	"encoding/json"
	"fmt"
	"math/big"
	common2 "platon-go-sdk/common"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/resp"
	"testing"
)

func TestRestrictingPlanContract_GetRestrictingInfo(t *testing.T) {
	config := network.PposMainNetParams
	rpc := NewRestrictingPlanContract(config, credentials)

	addr := common2.MustBech32ToAddress(MainFanAccount)
	list, err := rpc.GetRestrictingInfo(addr)
	if err != nil {
		t.Errorf("RestrictingPlanContract.GetRestrictingInfo failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// expect output
	//821004
	//94032894a3c28deaa12eb97c68c9775469f5dff80f
	//Function Data is: da838210049594032894a3c28deaa12eb97c68c9775469f5dff80f
	fmt.Println(string(result))
}

func TestRestrictingPlanContract_CreateRestrictingPlan(t *testing.T) {
	config := network.PposMainNetParams
	rpc := NewRestrictingPlanContract(config, credentials)

	addr := common2.MustBech32ToAddress(MainFanAccount)

	plan1Amount := new(big.Int)
	plan1Amount.SetString("1000000000000000000", 10)
	plan1 := resp.RestrictingPlan{
		Epoch:  big.NewInt(100),
		Amount: plan1Amount,
	}

	plan2Amount := new(big.Int)
	plan2Amount.SetString("2000000000000000000", 10)
	plan2 := resp.RestrictingPlan{
		Epoch:  big.NewInt(200),
		Amount: plan2Amount,
	}
	plans := []resp.RestrictingPlan{plan1, plan2}

	list, err := rpc.CreateRestrictingPlan(addr, plans)
	if err != nil {
		t.Errorf("RestrictingPlanContract.CreateRestrictingPlan failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	// expect output:
	//820fa0
	//94032894a3c28deaa12eb97c68c9775469f5dff80f
	//d7ca64880de0b6b3a7640000cb81c8881bc16d674ec80000
	//Function Data is: f383820fa09594032894a3c28deaa12eb97c68c9775469f5dff80f98d7ca64880de0b6b3a7640000cb81c8881bc16d674ec800000
	fmt.Println(string(result))
}
