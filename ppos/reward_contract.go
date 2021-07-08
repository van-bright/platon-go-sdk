package ppos

import (
	common2 "platon-go-sdk/common"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/resp"
	"platon-go-sdk/ppos/typedefs"
)

type RewardContract struct {
	executor *FunctionExecutor
}

func NewRewardContract(pposConfig *network.PposNetworkParameters, credentials *typedefs.Credentials) *RewardContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.RewardContract,
		credentials:  credentials,
	}
	return &RewardContract{executor}
}

/**
 * 提取账户当前所有的可提取的委托奖励
 *
 * @return
 */
func (rc *RewardContract) WithdrawDelegateReward() (typedefs.TransactionHash, error) {
	f := typedefs.NewFunction(typedefs.WITHDRAW_DELEGATE_REWARD_FUNC_TYPE, nil)

	var receipt typedefs.TransactionHash
	err := rc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 查询当前账户地址所委托的节点的NodeID和质押Id
 *
 * @param address 查询的地址
 * @param nodeList 节点id列表
 * @return
 */
func (rc *RewardContract) GetDelegateReward(address common2.Address, nodeIdList []string) ([]resp.Reward, error) {
	var nodeIds []codec.NodeId
	for _, n := range nodeIdList {
		nodeIds = append(nodeIds, codec.NodeId{HexStringId: n})
	}

	params := []interface{}{address, nodeIds}
	f := typedefs.NewFunction(typedefs.GET_DELEGATE_REWARD_FUNC_TYPE, params)

	var rewards []resp.Reward
	err := rc.executor.CallWithResult(f, &rewards)
	return rewards, err
}
