package ppos

import (
	"encoding/json"
	"fmt"
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
	"platon-go-sdk/ppos/req"
	"testing"
)

func TestStakingContract_GetAvgPackTime(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	list, err := sc.GetAvgPackTime()
	if err != nil {
		t.Errorf("StakingContract.GetAvgPackTime failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8204b2
	//Function Data is: c4838204b2
	fmt.Println(string(result))
}

func TestStakingContract_GetPackageReward(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	list, err := sc.GetPackageReward()
	if err != nil {
		t.Errorf("StakingContract.GetPackageReward failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8204b0
	//Function Data is: c4838204b0
	fmt.Println(string(result))
}

func TestStakingContract_GetStakingInfo(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	list, err := sc.GetStakingInfo(SNZPoolId)
	if err != nil {
		t.Errorf("StakingContract.GetStakingInfo failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//820451
	//b840c0dc97ee57ba202faf012ecb72bf30aebcd2cf7c161d7012017e0320e0db15925c107998bd833d61ec4c2689172d7e34a0371f4511773641e00814c2632b0e66
	//Function Data is: f84883820451b842b840c0dc97ee57ba202faf012ecb72bf30aebcd2cf7c161d7012017e0320e0db15925c107998bd833d61ec4c2689172d7e34a0371f4511773641e00814c2632b0e66
	fmt.Println(string(result))
}

func TestStakingContract_GetStakingReward(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	list, err := sc.GetStakingReward()
	if err != nil {
		t.Errorf("StakingContract.GetStakingReward failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8204b1
	//Function Data is: c4838204b1
	fmt.Println(string(result))
}

func TestStakingContract_Staking(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	sp := req.StakingParam{
		NodeId:            "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050",
		Amount:            big.NewInt(10),
		StakingAmountType: common.FREE_AMOUNT_TYPE,
		BenefitAddress:    MainFanAccount,
		ExternalId:        "",
		NodeName:          "chendai-node3",
		WebSite:           "www.baidu.com",
		Details:           "chendai-node3-details",
		ProcessVersion: common.ProgramVersion{
			Version: big.NewInt(4096),
			Sign:    "0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00",
		},
		BlsPubKey: "0x5ccd6b8c32f2713faa6c9a46e5fb61ad7b7400e53fabcbc56bdc0c16fbfffe09ad6256982c7059e7383a9187ad93a002a7cda7a75d569f591730481a8b91b5fad52ac26ac495522a069686df1061fc184c31771008c1fedfafd50ae794778811",
		BlsProof:  "0xa8fadadfc215f4f73fcdd539f5c2c8228a948f9d9f1f840329965a4abaec284be94d76f02839a0dd73d5a446dd5cd415c10b6ce621f0b1226924b8f3633f055a94b86446bb57dba0c9e85fb2e4b065773be39d8435352b6542a43c953afa0981",
		RewardPer: big.NewInt(1000),
	}

	list, err := sc.Staking(sp)
	if err != nil {
		t.Errorf("StakingContract.Staking failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8203e8
	//80
	//94032894a3c28deaa12eb97c68c9775469f5dff80f
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//80
	//8d6368656e6461692d6e6f646533
	//8d7777772e62616964752e636f6d
	//956368656e6461692d6e6f6465332d64657461696c73
	//0a
	//8203e8
	//821000
	//b8410dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00
	//b8605ccd6b8c32f2713faa6c9a46e5fb61ad7b7400e53fabcbc56bdc0c16fbfffe09ad6256982c7059e7383a9187ad93a002a7cda7a75d569f591730481a8b91b5fad52ac26ac495522a069686df1061fc184c31771008c1fedfafd50ae794778811
	//b860a8fadadfc215f4f73fcdd539f5c2c8228a948f9d9f1f840329965a4abaec284be94d76f02839a0dd73d5a446dd5cd415c10b6ce621f0b1226924b8f3633f055a94b86446bb57dba0c9e85fb2e4b065773be39d8435352b6542a43c953afa0981
	//Function Data is: f901ad838203e881809594032894a3c28deaa12eb97c68c9775469f5dff80fb842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c8205549905081808e8d6368656e6461692d6e6f6465338e8d7777772e62616964752e636f6d96956368656e6461692d6e6f6465332d64657461696c730a838203e883821000b843b8410dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00b862b8605ccd6b8c32f2713faa6c9a46e5fb61ad7b7400e53fabcbc56bdc0c16fbfffe09ad6256982c7059e7383a9187ad93a002a7cda7a75d569f591730481a8b91b5fad52ac26ac495522a069686df1061fc184c31771008c1fedfafd50ae794778811b862b860a8fadadfc215f4f73fcdd539f5c2c8228a948f9d9f1f840329965a4abaec284be94d76f02839a0dd73d5a446dd5cd415c10b6ce621f0b1226924b8f3633f055a94b86446bb57dba0c9e85fb2e4b065773be39d8435352b6542a43c953afa0981

	fmt.Println(string(result))
}

func TestStakingContract_AddStaking(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"

	list, err := sc.AddStaking(nodeId, common.FREE_AMOUNT_TYPE, big.NewInt(10))
	if err != nil {
		t.Errorf("StakingContract.AddStaking failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8203ea
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//80
	//0a
	//Function Data is: f84b838203eab842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c8205549905081800a
	fmt.Println(string(result))
}

func TestStakingContract_UnStaking(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"

	list, err := sc.UnStaking(nodeId)
	if err != nil {
		t.Errorf("StakingContract.UnStaking failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8203eb
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//Function Data is: f848838203ebb842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	fmt.Println(string(result))
}

func TestStakingContract_UpdateStakingInfo(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewStakingContract(config, credentials)

	req := req.UpdateStakingParam{
		NodeId:         "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050",
		BenifitAddress: MainFanAccount,
		ExternalId:     "",
		NodeName:       "chendai-node3",
		WebSite:        "www.baidu.com",
		Details:        "chendai-node3-details",
		RewardPer:      big.NewInt(1000),
	}
	list, err := sc.UpdateStakingInfo(req)
	if err != nil {
		t.Errorf("StakingContract.UpdateStakingInfo failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//8203e9
	//94032894a3c28deaa12eb97c68c9775469f5dff80f
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//8203e8
	//80
	//8d6368656e6461692d6e6f646533
	//8d7777772e62616964752e636f6d
	//956368656e6461692d6e6f6465332d64657461696c73
	//Function Data is: f899838203e99594032894a3c28deaa12eb97c68c9775469f5dff80fb842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050838203e881808e8d6368656e6461692d6e6f6465338e8d7777772e62616964752e636f6d96956368656e6461692d6e6f6465332d64657461696c73

	fmt.Println(string(result))
}
