package ppos

import (
	"encoding/json"
	"github.com/oldmanfan/platon-go-sdk/network"
	"github.com/oldmanfan/platon-go-sdk/ppos/resp"
	"github.com/oldmanfan/platon-go-sdk/ppos/typedefs"
)

type NodeContract struct {
	executor *FunctionExecutor
}

func NewNodeContract(pposConfig *network.PposNetworkParameters, credentials *typedefs.Credentials) *NodeContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.StakingContract,
		credentials:  credentials,
	}
	return &NodeContract{executor}
}

func (nc *NodeContract) doExecuteFunction(f *typedefs.Function) ([]resp.Node, error) {
	raw, err := nc.executor.CallWithRaw(f)
	if err != nil {
		return nil, err
	}

	//fmt.Println("rawMessage: " + string(raw))
	var nodeList []resp.Node

	err = json.Unmarshal(raw, &nodeList)
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

/**
 * 查询当前结算周期的验证人队列
 *
 * @return
 */
func (nc *NodeContract) GetVerifierList() ([]resp.Node, error) {
	f := typedefs.NewFunction(typedefs.GET_VERIFIERLIST_FUNC_TYPE, nil)

	return nc.doExecuteFunction(f)
}

/**
 * 查询当前共识周期的验证人列表
 *
 * @return
 */
func (nc *NodeContract) GetValidatorList() ([]resp.Node, error) {
	f := typedefs.NewFunction(typedefs.GET_VALIDATORLIST_FUNC_TYPE, nil)
	return nc.doExecuteFunction(f)
}

/**
 * 查询所有实时的候选人列表
 *
 * @return
 */
func (nc *NodeContract) GetCandidateList() ([]resp.Node, error) {
	f := typedefs.NewFunction(typedefs.GET_CANDIDATELIST_FUNC_TYPE, nil)
	return nc.doExecuteFunction(f)
}
