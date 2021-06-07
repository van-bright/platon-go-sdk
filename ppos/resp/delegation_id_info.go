package resp

import "math/big"

type DelegationIdInfo struct {

	/**
	 * 验证人节点的地址
	 */
	Address string `json:"Addr"`
	/**
	 * 验证人的节点Id
	 */
	NodeId string `json:"NodeId"`
	/**
	 * 发起质押时的区块高度
	 */
	StakingBlockNum *big.Int `json:"StakingBlockNum"`
}
