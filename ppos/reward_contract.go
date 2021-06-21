package ppos

import (
	common2 "platon-go-sdk/common"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/common"
	"platon-go-sdk/ppos/resp"
)

type RewardContract struct {
	executor *FunctionExecutor
}

func NewRewardContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *RewardContract {
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
func (rc *RewardContract) WithdrawDelegateReward() (common.TransactionHash, error) {
	f := common.NewFunction(common.WITHDRAW_DELEGATE_REWARD_FUNC_TYPE, nil)

	var receipt common.TransactionHash
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
	f := common.NewFunction(common.GET_DELEGATE_REWARD_FUNC_TYPE, params)

	var rewards []resp.Reward
	err := rc.executor.CallWithResult(f, &rewards)
	return rewards, err
}
