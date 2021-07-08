package resp

import (
	"math/big"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/typedefs"
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
	Verifier string
	/**
	 * 参数模块
	 */
	Module string
	/**
	 * 参数名称
	 */
	Name string
	/**
	 * 参数新值
	 */
	NewValue string
}

func CreateTextProposal(verifier string, piPid string) *Proposal {
	return &Proposal{
		ProposalId:     "",
		Proposer:       "",
		ProposalType:   typedefs.TEXT_PROPOSAL,
		PiPid:          piPid,
		SubmitBlock:    nil,
		EndVotingBlock: nil,
		NewVersion:     nil,
		ToBeCanceled:   "",
		ActiveBlock:    nil,
		Verifier:       verifier,
		Module:         "",
		Name:           "",
		NewValue:       "",
	}
}

func CreateVersionProposal(verifier string, pIDID string, newVersion *big.Int, endVotingRounds *big.Int) *Proposal {
	return &Proposal{
		ProposalId:     "",
		Proposer:       "",
		ProposalType:   typedefs.VERSION_PROPOSAL,
		PiPid:          pIDID,
		SubmitBlock:    nil,
		EndVotingBlock: endVotingRounds,
		NewVersion:     newVersion,
		ToBeCanceled:   "",
		ActiveBlock:    nil,
		Verifier:       verifier,
		Module:         "",
		Name:           "",
		NewValue:       "",
	}
}

func CreateCancelProposal(verifier string, pIDID string, endVotingRounds *big.Int, tobeCanceledProposalID string) *Proposal {
	return &Proposal{
		ProposalId:     "",
		Proposer:       "",
		ProposalType:   typedefs.CANCEL_PROPOSAL,
		PiPid:          pIDID,
		SubmitBlock:    nil,
		EndVotingBlock: endVotingRounds,
		NewVersion:     nil,
		ToBeCanceled:   tobeCanceledProposalID,
		ActiveBlock:    nil,
		Verifier:       verifier,
		Module:         "",
		Name:           "",
		NewValue:       "",
	}
}

func CreateParamProposal(verifier string, pIDID string, module string, name string, newValue string) *Proposal {
	return &Proposal{
		ProposalId:     "",
		Proposer:       "",
		ProposalType:   typedefs.PARAM_PROPOSAL,
		PiPid:          pIDID,
		SubmitBlock:    nil,
		EndVotingBlock: nil,
		NewVersion:     nil,
		ToBeCanceled:   "",
		ActiveBlock:    nil,
		Verifier:       verifier,
		Module:         module,
		Name:           name,
		NewValue:       newValue,
	}
}

func (p *Proposal) GetSubmitFunctionType() typedefs.FunctionType {
	switch p.ProposalType {
	case typedefs.TEXT_PROPOSAL:
		return typedefs.SUBMIT_TEXT_FUNC_TYPE
	case typedefs.VERSION_PROPOSAL:
		return typedefs.SUBMIT_VERSION_FUNC_TYPE
	case typedefs.CANCEL_PROPOSAL:
		return typedefs.SUBMIT_CANCEL_FUNC_TYPE
	default:
		return typedefs.SUBMIT_PARAM_FUNCTION_TYPE
	}
}

func (p *Proposal) GetSubmitInputParameters() []interface{} {
	switch p.ProposalType {
	case typedefs.TEXT_PROPOSAL:
		return []interface{}{
			codec.HexStringParam{HexStringValue: p.Verifier},
			codec.Utf8String{ValueInner: p.PiPid},
		}
	case typedefs.VERSION_PROPOSAL:
		return []interface{}{
			codec.HexStringParam{HexStringValue: p.Verifier},
			codec.Utf8String{ValueInner: p.PiPid},
			codec.UInt32{ValueInner: p.NewVersion},
			codec.UInt64{ValueInner: p.EndVotingBlock},
		}
	case typedefs.CANCEL_PROPOSAL:
		return []interface{}{
			codec.HexStringParam{HexStringValue: p.Verifier},
			codec.Utf8String{ValueInner: p.PiPid},
			codec.UInt64{ValueInner: p.EndVotingBlock},
			codec.HexStringParam{HexStringValue: p.ToBeCanceled},
		}
	default: // typedefs.PARAM_PROPOSAL:
		return []interface{}{
			codec.HexStringParam{HexStringValue: p.Verifier},
			codec.Utf8String{ValueInner: p.PiPid},
			codec.Utf8String{ValueInner: p.Module},
			codec.Utf8String{ValueInner: p.Name},
			codec.Utf8String{ValueInner: p.NewValue},
		}
	}
}
