package ppos

import (
	"encoding/json"
	"math/big"
	"platon-go-sdk/network"
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

func (pc *ProposalContract) doExecuteFunction(f *common.Function, result interface{}) error {
	raw, err := pc.executor.Send(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &result)
	return err
}

func (pc *ProposalContract) GetProposal(proposalId string) (resp.Proposal, error) {
	f := common.NewFunction(common.GET_PROPOSAL_FUNC_TYPE, []interface{}{proposalId})

	var proposal resp.Proposal
	err := pc.doExecuteFunction(f, &proposal)
	return proposal, err
}

func (pc *ProposalContract) GetTallyResult(proposalId string) (resp.TallyResult, error) {
	f := common.NewFunction(common.GET_TALLY_RESULT_FUNC_TYPE, []interface{}{proposalId})

	var tallyResult resp.TallyResult
	err := pc.doExecuteFunction(f, &tallyResult)
	return tallyResult, err
}

func (pc *ProposalContract) GetProposalList() ([]resp.Proposal, error) {
	f := common.NewFunction(common.GET_PROPOSAL_LIST_FUNC_TYPE, nil)

	var proposals []resp.Proposal
	err := pc.doExecuteFunction(f, &proposals)
	return proposals, err
}

func (pc *ProposalContract) Vote(programVersion common.ProgramVersion, voteOption common.VoteOption, proposalID string, verifier string) (common.TransactionReceipt, error) {
	params := []interface{}{verifier, proposalID, programVersion, voteOption.GetValue(), programVersion.Version, programVersion.Sign}
	f := common.NewFunction(common.VOTE_FUNC_TYPE, params)

	var receipt common.TransactionReceipt
	err := pc.doExecuteFunction(f, &receipt)
	return receipt, err
}

func (pc *ProposalContract) DeclareVersion(programVersion common.ProgramVersion, verifier string) (common.TransactionReceipt, error) {
	params := []interface{}{verifier, programVersion.Version, programVersion.Sign}
	f := common.NewFunction(common.DECLARE_VERSION_FUNC_TYPE, params)

	var receipt common.TransactionReceipt
	err := pc.doExecuteFunction(f, &receipt)
	return receipt, err
}

func (pc *ProposalContract) SubmitProposal(proposal *resp.Proposal) (common.TransactionReceipt, error) {
	f := common.NewFunction(proposal.GetSubmitFunctionType(), proposal.GetSubmitInputParameters())

	var receipt common.TransactionReceipt
	err := pc.doExecuteFunction(f, &receipt)
	return receipt, err
}

func (pc *ProposalContract) GetActiveVersion() (*big.Int, error) {
	f := common.NewFunction(common.GET_ACTIVE_VERSION, nil)

	var ver = big.NewInt(0)
	err := pc.doExecuteFunction(f, &ver)
	return ver, err
}

func (pc *ProposalContract) GetGovernParamValue(module string, name string) (string, error) {
	params := []interface{}{module, name}
	f := common.NewFunction(common.GET_GOVERN_PARAM_VALUE, params)

	var value string
	err := pc.doExecuteFunction(f, &value)
	return value, err
}

func (pc *ProposalContract) GetAccuVerifiersCount(proposalId string, blockHash string) (*big.Int, error) {
	params := []interface{}{proposalId, blockHash}
	f := common.NewFunction(common.GET_ACCUVERIFIERS_COUNT, params)

	var value = big.NewInt(0)
	err := pc.doExecuteFunction(f, &value)
	return value, err
}

func (pc *ProposalContract) GetParamList(module string) ([]resp.GovernParam, error) {
	params := []interface{}{module}
	f := common.NewFunction(common.GET_PARAM_LIST, params)

	var value []resp.GovernParam
	err := pc.doExecuteFunction(f, &value)
	return value, err
}
