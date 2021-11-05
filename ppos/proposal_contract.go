package ppos

import (
	"github.com/oldmanfan/platon-go-sdk/common/hexutil"
	"github.com/oldmanfan/platon-go-sdk/network"
	"github.com/oldmanfan/platon-go-sdk/ppos/codec"
	"github.com/oldmanfan/platon-go-sdk/ppos/resp"
	"github.com/oldmanfan/platon-go-sdk/ppos/typedefs"
	"math/big"
)

type ProposalContract struct {
	executor *FunctionExecutor
}

func NewProposalContract(pposConfig *network.PposNetworkParameters, credentials *typedefs.Credentials) *ProposalContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.ProposalContract,
		credentials:  credentials,
	}
	return &ProposalContract{executor}
}

/**
 * 查询提案
 *
 * @param proposalId
 * @return
 */
func (pc *ProposalContract) GetProposal(proposalId string) (resp.Proposal, error) {
	f := typedefs.NewFunction(typedefs.GET_PROPOSAL_FUNC_TYPE, []interface{}{codec.HexStringParam{HexStringValue: proposalId}})

	var proposal resp.Proposal
	err := pc.executor.CallWithResult(f, &proposal)
	return proposal, err
}

/**
 * 查询提案结果
 *
 * @param proposalId
 * @return
 */
func (pc *ProposalContract) GetTallyResult(proposalId string) (resp.TallyResult, error) {
	f := typedefs.NewFunction(typedefs.GET_TALLY_RESULT_FUNC_TYPE, []interface{}{codec.HexStringParam{HexStringValue: proposalId}})

	var tallyResult resp.TallyResult
	err := pc.executor.CallWithResult(f, &tallyResult)
	return tallyResult, err
}

/**
 * 获取提案列表
 *
 * @return
 */
func (pc *ProposalContract) GetProposalList() ([]resp.Proposal, error) {
	f := typedefs.NewFunction(typedefs.GET_PROPOSAL_LIST_FUNC_TYPE, nil)

	var proposals []resp.Proposal
	err := pc.executor.CallWithResult(f, &proposals)
	return proposals, err
}

/**
 * 给提案投票
 *
 * @param programVersion
 * @param voteOption     投票选项
 * @param proposalID     提案ID
 * @param verifier       投票验证人
 * @return
 */
func (pc *ProposalContract) Vote(programVersion typedefs.ProgramVersion, voteOption typedefs.VoteOption, proposalID string, nodeId string) (typedefs.TransactionHash, error) {
	params := []interface{}{
		codec.HexStringParam{HexStringValue: nodeId},
		codec.HexStringParam{HexStringValue: proposalID},
		codec.UInt16{ValueInner: voteOption.GetValue()},
		codec.UInt32{ValueInner: programVersion.Version},
		codec.HexStringParam{HexStringValue: programVersion.Sign},
	}
	f := typedefs.NewFunction(typedefs.VOTE_FUNC_TYPE, params)

	var receipt typedefs.TransactionHash
	err := pc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 版本声明
 *
 * @param programVersion
 * @param verifier       声明的节点，只能是验证人/候选人
 * @param gasProvider
 * @return
 */
func (pc *ProposalContract) DeclareVersion(programVersion typedefs.ProgramVersion, verifier string) (typedefs.TransactionHash, error) {
	params := []interface{}{codec.HexStringParam{HexStringValue: verifier}, codec.UInt32{ValueInner: programVersion.Version}, codec.HexStringParam{HexStringValue: programVersion.Sign}}
	f := typedefs.NewFunction(typedefs.DECLARE_VERSION_FUNC_TYPE, params)

	var receipt typedefs.TransactionHash
	err := pc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 提交提案
 *
 * @param proposal 包括文本提案和版本提案
 * @return
 */
func (pc *ProposalContract) SubmitProposal(proposal *resp.Proposal) (typedefs.TransactionHash, error) {
	f := typedefs.NewFunction(proposal.GetSubmitFunctionType(), proposal.GetSubmitInputParameters())

	var receipt typedefs.TransactionHash
	err := pc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 查询已生效的版本
 *
 * @return
 */
func (pc *ProposalContract) GetActiveVersion() (uint64, error) {
	f := typedefs.NewFunction(typedefs.GET_ACTIVE_VERSION, nil)

	var ver uint64
	err := pc.executor.CallWithResult(f, &ver)
	return ver, err
}

/**
 * 查询当前块高的治理参数值
 *
 * @param module 参数模块
 * @param name   参数名称
 * @return
 */
func (pc *ProposalContract) GetGovernParamValue(module string, name string) (string, error) {
	params := []interface{}{codec.Utf8String{ValueInner: module}, codec.Utf8String{ValueInner: name}}
	f := typedefs.NewFunction(typedefs.GET_GOVERN_PARAM_VALUE, params)

	var value string
	err := pc.executor.CallWithResult(f, &value)
	return value, err
}

/**
 * 查询提案的累积可投票人数
 *
 * @param proposalId 提案ID
 * @param blockHash  块hash
 * @return
 */
func (pc *ProposalContract) GetAccuVerifiersCount(proposalId string, blockHash string) (*big.Int, error) {
	params := []interface{}{codec.HexStringParam{HexStringValue: proposalId}, codec.HexStringParam{HexStringValue: blockHash}}
	f := typedefs.NewFunction(typedefs.GET_ACCUVERIFIERS_COUNT, params)

	var value *hexutil.Big
	err := pc.executor.CallWithResult(f, &value)
	return value.ToInt(), err
}

/**
 * 查询可治理参数列表
 *
 */
func (pc *ProposalContract) GetParamList(module string) ([]resp.GovernParam, error) {
	params := []interface{}{codec.Utf8String{ValueInner: module}}
	f := typedefs.NewFunction(typedefs.GET_PARAM_LIST, params)

	var value []resp.GovernParam
	err := pc.executor.CallWithResult(f, &value)
	return value, err
}
