package resp

import (
	"math/big"
	"platon-go-sdk/ppos/common"
)

type Proposal struct {

	/**
	 * 提案id
	 */
	ProposalId string `json:"ProposalID"`
	/**
	 * 提案节点ID
	 */
	Proposer string `json:"Proposer"`
	/**
	 * 提案类型， 0x01：文本提案； 0x02：升级提案；0x03参数提案
	 */
	ProposalType int `json:"ProposalType"`
	/**
	 * 提案PIPID
	 */
	PiPid string `json:"PIPID"`
	/**
	 * 提交提案的块高
	 */
	SubmitBlock *big.Int `json:"SubmitBlock"`
	/**
	 * 提案投票结束的块高
	 */
	EndVotingBlock *big.Int `json:"EndVotingBlock"`
	/**
	 * 升级版本
	 */
	NewVersion *big.Int `json:"NewVersion"`
	/**
	 * 提案要取消的升级提案ID
	 */
	ToBeCanceled string `json:"TobeCanceled"`
	/**
	 * （如果投票通过）生效块高（endVotingBlock + 20 + 4*250 < 生效块高 <= endVotingBlock + 20 + 10*250）
	 */
	ActiveBlock *big.Int `json:"ActiveBlock"`

	/**
	 * 提交提案的验证人
	 */
	verifier string
	/**
	 * 参数模块
	 */
	module string
	/**
	 * 参数名称
	 */
	name string
	/**
	 * 参数新值
	 */
	newValue string
}

func (p *Proposal) GetSubmitFunctionType() common.FunctionType{
	switch p.ProposalType {
	case common.TEXT_PROPOSAL:
		return common.SUBMIT_TEXT_FUNC_TYPE
	case common.VERSION_PROPOSAL:
		return common.SUBMIT_VERSION_FUNC_TYPE
	case common.CANCEL_PROPOSAL:
		return common.SUBMIT_CANCEL_FUNC_TYPE
	default:
		return common.SUBMIT_PARAM_FUNCTION_TYPE
	}
}

func (p *Proposal) GetSubmitInputParameters() []interface{} {
	switch p.ProposalType {
	case common.TEXT_PROPOSAL:
		return []interface{}{p.verifier, p.PiPid}
	case common.VERSION_PROPOSAL:
		return []interface{}{p.verifier, p.PiPid, p.NewVersion, p.EndVotingBlock}
	case common.CANCEL_PROPOSAL:
		return []interface{}{p.verifier, p.PiPid, p.EndVotingBlock, p.ToBeCanceled}
	case common.PARAM_PROPOSAL:
		return []interface{}{p.verifier, p.PiPid, p.module, p.name, p.newValue}
	default:
		return []interface{}{}
	}
}
