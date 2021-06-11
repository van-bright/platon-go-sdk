package req

import "math/big"

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
	return []interface{}{usp.BenifitAddress,
		usp.NodeId,
		usp.RewardPer,
		usp.ExternalId,
		usp.NodeName,
		usp.WebSite,
		usp.Details,
	}
}
