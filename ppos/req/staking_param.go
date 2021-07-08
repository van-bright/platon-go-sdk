package req

import (
	"math/big"
	common2 "platon-go-sdk/common"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/typedefs"
)

type StakingParam struct {

	/**
	 * 64bytes 被质押的节点Id(也叫候选人的节点Id)
	 */
	NodeId string
	/**
	 * 质押的von
	 */
	Amount *big.Int
	/**
	 * 表示使用账户自由金额还是账户的锁仓金额做质押，0: 自由金额； 1: 锁仓金额
	 */
	StakingAmountType typedefs.StakingAmountType
	/**
	 * 20bytes 用于接受出块奖励和质押奖励的收益账户
	 */
	BenefitAddress string
	/**
	 * 外部Id(有长度限制，给第三方拉取节点描述的Id)
	 */
	ExternalId string
	/**
	 * 被质押节点的名称(有长度限制，表示该节点的名称)
	 */
	NodeName string
	/**
	 * 节点的第三方主页(有长度限制，表示该节点的主页)
	 */
	WebSite string
	/**
	 * 节点的描述(有长度限制，表示该节点的描述)
	 */
	Details string
	/**
	 * 程序的真实版本，治理rpc获取
	 */
	ProcessVersion typedefs.ProgramVersion
	/**
	 * bls的公钥
	 */
	BlsPubKey string

	/**
	 * bls的证明
	 */
	BlsProof string

	/**
	 * 奖励分成比例，采用BasePoint 1BP=0.01%
	 */
	RewardPer *big.Int
}

func (sp StakingParam) SubmitInputParameters() []interface{} {
	return []interface{}{
		codec.UInt16{ValueInner: sp.StakingAmountType.GetValue()},
		common2.MustBech32ToAddress(sp.BenefitAddress),
		codec.NodeId{HexStringId: sp.NodeId},
		codec.Utf8String{ValueInner: sp.ExternalId},
		codec.Utf8String{ValueInner: sp.NodeName},
		codec.Utf8String{ValueInner: sp.WebSite},
		codec.Utf8String{ValueInner: sp.Details},
		codec.UInt256{ValueInner: sp.Amount},
		codec.UInt32{ValueInner: sp.RewardPer},
		codec.UInt32{ValueInner: sp.ProcessVersion.Version},
		codec.HexStringParam{HexStringValue: sp.ProcessVersion.Sign},
		codec.HexStringParam{HexStringValue: sp.BlsPubKey},
		codec.HexStringParam{HexStringValue: sp.BlsProof},
	}
}
