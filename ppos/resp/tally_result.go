package resp

import "math/big"

// 投票结果
type TallyResult struct {
	/**
	 * 提案ID
	 */
	ProposalID string
	/**
	 * 赞成票
	 */
	Yeas *big.Int
	/**
	 * 反对票
	 */
	Nays *big.Int
	/**
	 * 弃权票
	 */
	Abstentions *big.Int
	/**
	 * 在整个投票期内有投票资格的验证人总数
	 */
	AccuVerifiers *big.Int
	/**
	 * 状态
	 */
	Status int
}
