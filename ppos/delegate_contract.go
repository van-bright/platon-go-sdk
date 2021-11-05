package ppos

import (
	"github.com/oldmanfan/platon-go-sdk/network"
	"github.com/oldmanfan/platon-go-sdk/ppos/codec"
	"github.com/oldmanfan/platon-go-sdk/ppos/typedefs"
	"math/big"
)

type DelegateContract struct {
	executor *FunctionExecutor
}

func NewDelegateContract(pposConfig *network.PposNetworkParameters, credentials *typedefs.Credentials) *DelegateContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.StakingContract,
		credentials:  credentials,
	}
	return &DelegateContract{executor}
}

/**
 * 发起委托
 *
 * @param nodeId            被质押的节点的NodeId
 * @param stakingAmountType 表示使用账户自由金额还是账户的锁仓金额做委托，0: 自由金额； 1: 锁仓金额
 * @param amount            委托的金额(按照最小单位算，1LAT = 10**18 von)
 * @return
 */
func (dc *DelegateContract) Delegate(nodeId string, stakingAmountType typedefs.StakingAmountType, amount *big.Int) (typedefs.TransactionHash, error) {
	params := []interface{}{codec.UInt16{ValueInner: stakingAmountType.GetValue()}, codec.NodeId{HexStringId: nodeId}, codec.UInt256{ValueInner: amount}}
	function := typedefs.NewFunction(typedefs.DELEGATE_FUNC_TYPE, params)

	var receipt typedefs.TransactionHash
	err := dc.executor.SendWithResult(function, &receipt)
	return receipt, err
}

/**
 * 减持/撤销委托(全部减持就是撤销)
 *
 * @param nodeId          被质押的节点的NodeId
 * @param stakingBlockNum 代表着某个node的某次质押的唯一标示
 * @param amount          减持委托的金额(按照最小单位算，1LAT = 10**18 von)
 * @return
 */
func (dc *DelegateContract) UnDelegate(nodeId string, stakingBlockNum *big.Int, amount *big.Int) (typedefs.TransactionHash, error) {
	params := []interface{}{codec.UInt64{ValueInner: stakingBlockNum}, codec.NodeId{HexStringId: nodeId}, codec.UInt256{ValueInner: amount}}
	function := typedefs.NewFunction(typedefs.WITHDREW_DELEGATE_FUNC_TYPE, params)

	var receipt typedefs.TransactionHash
	err := dc.executor.SendWithResult(function, &receipt)
	return receipt, err
}
