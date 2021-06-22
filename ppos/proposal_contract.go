package ppos

import (
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/common"
	"platon-go-sdk/ppos/resp"
)

type ProposalContract struct {
	executor *FunctionExecutor
}

func NewProposalContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *ProposalContract {
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
	f := common.NewFunction(common.GET_PROPOSAL_FUNC_TYPE, []interface{}{proposalId})

	var proposal resp.Proposal
	err := pc.executor.SendWithResult(f, &proposal)
	return proposal, err
}

/**
 * 查询提案结果
 *
 * @param proposalId
 * @return
 */
func (pc *ProposalContract) GetTallyResult(proposalId string) (resp.TallyResult, error) {
	f := common.NewFunction(common.GET_TALLY_RESULT_FUNC_TYPE, []interface{}{proposalId})

	var tallyResult resp.TallyResult
	err := pc.executor.SendWithResult(f, &tallyResult)
	return tallyResult, err
}

/**
 * 获取提案列表
 *
 * @return
 */
func (pc *ProposalContract) GetProposalList() ([]resp.Proposal, error) {
	f := common.NewFunction(common.GET_PROPOSAL_LIST_FUNC_TYPE, nil)

	var proposals []resp.Proposal
	err := pc.executor.SendWithResult(f, &proposals)
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
func (pc *ProposalContract) Vote(programVersion common.ProgramVersion, voteOption common.VoteOption, proposalID string, verifier string) (common.TransactionHash, error) {
	params := []interface{}{verifier, proposalID, programVersion, voteOption.GetValue(), programVersion.Version, programVersion.Sign}
	f := common.NewFunction(common.VOTE_FUNC_TYPE, params)

	var receipt common.TransactionHash
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
func (pc *ProposalContract) DeclareVersion(programVersion common.ProgramVersion, verifier string) (common.TransactionHash, error) {
	params := []interface{}{codec.HexStringParam{HexStringValue: verifier}, codec.UInt32{ValueInner: programVersion.Version}, codec.HexStringParam{HexStringValue: programVersion.Sign}}
	f := common.NewFunction(common.DECLARE_VERSION_FUNC_TYPE, params)

	var receipt common.TransactionHash
	err := pc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 提交提案
 *
 * @param proposal 包括文本提案和版本提案
 * @return
 */
func (pc *ProposalContract) SubmitProposal(proposal *resp.Proposal) (common.TransactionHash, error) {
	f := common.NewFunction(proposal.GetSubmitFunctionType(), proposal.GetSubmitInputParameters())

	var receipt common.TransactionHash
	err := pc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 查询已生效的版本
 *
 * @return
 */
func (pc *ProposalContract) GetActiveVersion() (*big.Int, error) {
	f := common.NewFunction(common.GET_ACTIVE_VERSION, nil)

	var ver = big.NewInt(0)
	err := pc.executor.SendWithResult(f, &ver)
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
	params := []interface{}{module, name}
	f := common.NewFunction(common.GET_GOVERN_PARAM_VALUE, params)

	var value string
	err := pc.executor.SendWithResult(f, &value)
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
	params := []interface{}{proposalId, blockHash}
	f := common.NewFunction(common.GET_ACCUVERIFIERS_COUNT, params)

	var value = big.NewInt(0)
	err := pc.executor.SendWithResult(f, &value)
	return value, err
}

/**
 * 查询可治理参数列表
 *
 */
func (pc *ProposalContract) GetParamList(module string) ([]resp.GovernParam, error) {
	params := []interface{}{module}
	f := common.NewFunction(common.GET_PARAM_LIST, params)

	var value []resp.GovernParam
	err := pc.executor.SendWithResult(f, &value)
	return value, err
}
