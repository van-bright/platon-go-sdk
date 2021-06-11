package ppos

import (
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
	"platon-go-sdk/ppos/req"
	"platon-go-sdk/ppos/resp"
)

type StakingContract struct {
	executor *FunctionExecutor
}

func NewStakingContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *StakingContract {
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
	f := common.NewFunction(common.GET_STAKINGINFO_FUNC_TYPE, []interface{}{nodeId})

	var node resp.Node
	err := sc.executor.SendWithResult(f, &node)
	return node, err
}

/**
 * 查询当前结算周期的区块奖励
 *
 * @return
 */
func (sc *StakingContract) GetPackageReward() (*big.Int, error) {
	f := common.NewFunction(common.GET_PACKAGEREWARD_FUNC_TYPE, nil)

	var reward = big.NewInt(0)
	err := sc.executor.SendWithResult(f, &reward)
	return reward, err
}

/**
 * 查询当前结算周期的质押奖励
 *
 * @return
 */
func (sc *StakingContract) GetStakingReward() (*big.Int, error) {
	f := common.NewFunction(common.GET_STAKINGREWARD_FUNC_TYPE, nil)

	var reward = big.NewInt(0)
	err := sc.executor.SendWithResult(f, &reward)
	return reward, err
}

/**
 * 查询打包区块的平均时间
 *
 * @return
 */
func (sc *StakingContract) GetAvgPackTime() (*big.Int, error) {
	f := common.NewFunction(common.GET_AVGPACKTIME_FUNC_TYPE, nil)

	var reward = big.NewInt(0)
	err := sc.executor.SendWithResult(f, &reward)
	return reward, err
}

/**
 * 发起质押
 *
 * @param stakingParam
 * @return
 * @see StakingParam
 */
func (sc *StakingContract) Staking(stakingParam req.StakingParam) (common.TransactionHash, error) {
	f := common.NewFunction(common.STAKING_FUNC_TYPE, stakingParam.SubmitInputParameters())

	var receipt common.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 撤销质押
 *
 * @param nodeId 64bytes 被质押的节点Id(也叫候选人的节点Id)
 * @return
 */
func (sc *StakingContract) UnStaking(nodeId string) (common.TransactionHash, error) {
	f := common.NewFunction(common.WITHDREW_STAKING_FUNC_TYPE, []interface{}{nodeId})

	var receipt common.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 更新质押信息
 *
 * @param updateStakingParam
 * @return
 */
func (sc *StakingContract) UpdateStakingInfo(updateStakingParam req.UpdateStakingParam) (common.TransactionHash, error) {
	f := common.NewFunction(common.WITHDREW_STAKING_FUNC_TYPE, updateStakingParam.SubmitInputParameters())

	var receipt common.TransactionHash
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
func (sc *StakingContract)AddStaking(nodeId string,  stakingAmountType common.StakingAmountType, amount *big.Int) (common.TransactionHash, error) {
	params := []interface{}{nodeId, stakingAmountType.GetValue(), amount}
	f := common.NewFunction(common.ADD_STAKING_FUNC_TYPE, params)

	var receipt common.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}
