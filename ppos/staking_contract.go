package ppos

import (
	"math/big"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/req"
	"platon-go-sdk/ppos/resp"
	"platon-go-sdk/ppos/typedefs"
)

type StakingContract struct {
	executor *FunctionExecutor
}

func NewStakingContract(pposConfig *network.PposNetworkParameters, credentials *typedefs.Credentials) *StakingContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.StakingContract,
		credentials:  credentials,
	}
	return &StakingContract{executor}
}

/**
 * 获取质押信息
 *
 * @param nodeId
 * @return
 */
func (sc *StakingContract) GetStakingInfo(nodeId string) (resp.Node, error) {
	f := typedefs.NewFunction(typedefs.GET_STAKINGINFO_FUNC_TYPE, []interface{}{codec.NodeId{HexStringId: nodeId}})

	var node resp.Node
	err := sc.executor.CallWithResult(f, &node)
	return node, err
}

/**
 * 查询当前结算周期的区块奖励
 *
 * @return
 */
func (sc *StakingContract) GetPackageReward() (*big.Int, error) {
	f := typedefs.NewFunction(typedefs.GET_PACKAGEREWARD_FUNC_TYPE, nil)

	var reward *hexutil.Big
	err := sc.executor.CallWithResult(f, &reward)
	return reward.ToInt(), err
}

/**
 * 查询当前结算周期的质押奖励
 *
 * @return
 */
func (sc *StakingContract) GetStakingReward() (*big.Int, error) {
	f := typedefs.NewFunction(typedefs.GET_STAKINGREWARD_FUNC_TYPE, nil)

	var reward *hexutil.Big
	err := sc.executor.CallWithResult(f, &reward)
	return reward.ToInt(), err
}

/**
 * 查询打包区块的平均时间
 *
 * @return
 */
func (sc *StakingContract) GetAvgPackTime() (*big.Int, error) {
	f := typedefs.NewFunction(typedefs.GET_AVGPACKTIME_FUNC_TYPE, nil)

	var reward = big.NewInt(0)
	err := sc.executor.CallWithResult(f, &reward)
	return reward, err
}

/**
 * 发起质押
 *
 * @param stakingParam
 * @return
 * @see StakingParam
 */
func (sc *StakingContract) Staking(stakingParam req.StakingParam) (typedefs.TransactionHash, error) {
	f := typedefs.NewFunction(typedefs.STAKING_FUNC_TYPE, stakingParam.SubmitInputParameters())

	var receipt typedefs.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 撤销质押
 *
 * @param nodeId 64bytes 被质押的节点Id(也叫候选人的节点Id)
 * @return
 */
func (sc *StakingContract) UnStaking(nodeId string) (typedefs.TransactionHash, error) {
	f := typedefs.NewFunction(typedefs.WITHDREW_STAKING_FUNC_TYPE, []interface{}{codec.NodeId{HexStringId: nodeId}})

	var receipt typedefs.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 更新质押信息
 *
 * @param updateStakingParam
 * @return
 */
func (sc *StakingContract) UpdateStakingInfo(updateStakingParam req.UpdateStakingParam) (typedefs.TransactionHash, error) {
	f := typedefs.NewFunction(typedefs.UPDATE_STAKING_INFO_FUNC_TYPE, updateStakingParam.SubmitInputParameters())

	var receipt typedefs.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 增持质押
 *
 * @param nodeId            被质押的节点Id(也叫候选人的节点Id)
 * @param stakingAmountType 表示使用账户自由金额还是账户的锁仓金额做质押，0: 自由金额； 1: 锁仓金额
 * @param amount            增持的von
 * @return
 */
func (sc *StakingContract) AddStaking(nodeId string, stakingAmountType typedefs.StakingAmountType, amount *big.Int) (typedefs.TransactionHash, error) {
	params := []interface{}{codec.NodeId{HexStringId: nodeId}, codec.UInt16{ValueInner: stakingAmountType.GetValue()}, codec.UInt256{ValueInner: amount}}
	f := typedefs.NewFunction(typedefs.ADD_STAKING_FUNC_TYPE, params)

	var receipt typedefs.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}
