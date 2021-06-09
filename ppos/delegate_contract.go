package ppos

import (
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
)

type DelegateContract struct {
	executor *FunctionExecutor
}

func NewDelegateContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *DelegateContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.StakingContract,
		credentials:  credentials,
	}
	return &DelegateContract{executor}
}

func (dc *DelegateContract) Delegate(nodeId string, stakingAmountType common.StakingAmountType, amount *big.Int) error {
	params := []interface{}{nodeId, stakingAmountType, amount}
	function := common.NewFunction(common.DELEGATE_FUNC_TYPE, params)

	_, err := dc.executor.Send(function)
	return err
}

func (dc *DelegateContract) UnDelegate(nodeId string, stakingBlockNum *big.Int, amount *big.Int) error {
	params := []interface{}{nodeId, stakingBlockNum, amount}
	function := common.NewFunction(common.WITHDREW_DELEGATE_FUNC_TYPE, params)

	_, err := dc.executor.Send(function)
	return err
}
