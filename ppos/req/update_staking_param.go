package req

import (
	"math/big"
	common2 "platon-go-sdk/common"
	"platon-go-sdk/ppos/codec"
)

type UpdateStakingParam struct {

	/**
	 * 64bytes 被质押的节点Id(也叫候选人的节点Id)
	 */
	NodeId string
	/**
	 * 20bytes 用于接受出块奖励和质押奖励的收益账户
	 */
	BenifitAddress string
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
	 * 奖励分成比例，采用BasePoint 1BP=0.01%
	 */
	RewardPer *big.Int
}

func (usp UpdateStakingParam) SubmitInputParameters() []interface{} {
	return []interface{}{
		common2.MustBech32ToAddress(usp.BenifitAddress),
		codec.NodeId{HexStringId: usp.NodeId},
		codec.UInt32{ValueInner: usp.RewardPer},
		codec.Utf8String{ValueInner: usp.ExternalId},
		codec.Utf8String{ValueInner: usp.NodeName},
		codec.Utf8String{ValueInner: usp.WebSite},
		codec.Utf8String{ValueInner: usp.Details},
	}
}
