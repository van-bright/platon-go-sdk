package ppos

import (
	"encoding/json"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
	"platon-go-sdk/ppos/resp"
)

type NodeContract struct {
	executor *FunctionExecutor
}

func NewNodeContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *NodeContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.StakingContract,
		credentials:  credentials,
	}
	return &NodeContract{executor}
}
func (nc *NodeContract) doExecuteFunction(f *common.Function) ([]resp.Node, error) {
	raw, err := nc.executor.Send(f)
	if err != nil {
		return nil, err
	}

	var nodes []resp.Node

	err = json.Unmarshal(raw, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (nc *NodeContract) GetVerifierList() ([]resp.Node, error) {
	f := common.NewFunction(common.GET_VERIFIERLIST_FUNC_TYPE, nil)

	return nc.doExecuteFunction(f)
}

func (nc *NodeContract) GetValidatorList() ([]resp.Node, error) {
	f := common.NewFunction(common.GET_VALIDATORLIST_FUNC_TYPE, nil)
	return nc.doExecuteFunction(f)
}

func (nc *NodeContract) GetCandidateList() ([]resp.Node, error) {
	f := common.NewFunction(common.GET_CANDIDATELIST_FUNC_TYPE, nil)
	return nc.doExecuteFunction(f)
}
