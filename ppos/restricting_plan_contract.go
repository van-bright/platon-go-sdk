package ppos

import (
	common2 "platon-go-sdk/common"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
	"platon-go-sdk/ppos/resp"
)

type RestrictingPlanContract struct {
	executor *FunctionExecutor
}

func NewRestrictingPlanContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *RestrictingPlanContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.RestrictingPlanContract,
		credentials:  credentials,
	}
	return &RestrictingPlanContract{executor}
}

/**
 * 创建锁仓计划
 *
 * @param account             锁仓释放到账账户
 * @param restrictingPlanList 其中，Epoch：表示结算周期的倍数。与每个结算周期出块数的乘积表示在目标区块高度上释放锁定的资金。
 *                            如果 account 是激励池地址，那么 period 值是 120（即，30*4） 的倍数。
 *                            另外，period * 每周期的区块数至少要大于最高不可逆区块高度。Amount：表示目标区块上待释放的金额。
 * @return
 */
func (rc *RestrictingPlanContract) CreateRestrictingPlan(account common2.Address, restrictingPlanList []resp.RestrictingPlan) (common.TransactionHash, error) {
	params := []interface{}{account, restrictingPlanList}
	f := common.NewFunction(common.CREATE_RESTRICTINGPLAN_FUNC_TYPE, params)

	var receipt common.TransactionHash
	err := rc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 获取锁仓信息
 *
 * @param account 锁仓释放到账账户
 * @return
 */
func (rc *RestrictingPlanContract) GetRestrictingInfo(account common2.Address) (resp.RestrictingItem, error) {
	f := common.NewFunction(common.GET_RESTRICTINGINFO_FUNC_TYPE, []interface{}{account})
	var item resp.RestrictingItem
	err := rc.executor.SendWithResult(f, &item)
	return item, err
}
