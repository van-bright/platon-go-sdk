package ppos

import (
	"encoding/json"
	"fmt"
	"github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/network"
	"testing"
)

const (
	MainFanAccount = "lat1qzqykpmfwvuj9lf2mx88r6vx2q0lmeals83aw8"
	SNZPoolId      = "0xc0dc97ee57ba202faf012ecb72bf30aebcd2cf7c161d7012017e0320e0db15925c107998bd833d61ec4c2689172d7e34a0371f4511773641e00814c2632b0e66"
	XiaoYiId       = "0x0423022e05633d8d2a80cb66d65d7b8bb267ef3c5deeb6cc124726219ddcd05194f6e958207443c3d39fb81b1072025b24525c7f2a14a2a2076b59a9af7bb70e"
)

func TestRewardContract_GetDelegateReward(t *testing.T) {
	config := network.PposMainNetParams
	rc := NewRewardContract(config, credentials)

	addr := common.MustBech32ToAddress(MainFanAccount)
	nodeId := []string{SNZPoolId, XiaoYiId}

	list, err := rc.GetDelegateReward(addr, nodeId)
	if err != nil {
		t.Errorf("Get Delegate Reward info failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}

func TestRewardContract_WithdrawDelegateReward(t *testing.T) {
	config := network.PposMainNetParams
	rc := NewRewardContract(config, credentials)

	receipt, err := rc.WithdrawDelegateReward()
	if err != nil {
		t.Errorf("Get Delegate Reward info failed: %s", err)
	}

	fmt.Println(receipt)
}
