package resp

import "math/big"

type Delegation struct {
	/**
	 * 委托人的账户地址
	 */
	DelegateAddress string `json:"Addr"`
	/**
	 * 验证人的节点Id
	 */
	NodeId string `json:"NodeId"`
	/**
	 * 发起质押时的区块高度
	 */
	StakingBlockNum *big.Int `json:"StakingBlockNum"`
	/**
	 * 最近一次对该候选人发起的委托时的结算周期
	 */
	DelegateEpoch *big.Int `json:"DelegateEpoch"`
	/**
	 * 发起委托账户的自由金额的锁定期委托的von
	 */
	DelegateReleased *big.Int `json:"Released"`
	/**
	 * 发起委托账户的自由金额的犹豫期委托的von
	 */
	DelegateReleasedHes *big.Int `json:"ReleaseHes"`
	/**
	 * 发起委托账户的锁仓金额的锁定期委托的von
	 */
	DelegateLocked *big.Int `json:"RestrictingPlan"`
	/**
	 * 发起委托账户的锁仓金额的犹豫期委托的von
	 */
	DelegateLockedHes *big.Int `json:"RestrictingPlanHes"`
	/**
	 * 待领取的委托收益von
	 */
	CumulativeIncome *big.Int `json:"CumulativeIncome"`
}
