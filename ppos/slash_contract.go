package ppos

import (
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/codec"
	"platon-go-sdk/ppos/common"
)

type SlashContract struct {
	executor *FunctionExecutor
}

func NewSlashContract(pposConfig *network.PposNetworkParameters, credentials *common.Credentials) *SlashContract {
	executor := &FunctionExecutor{
		httpEntry:    pposConfig.Url,
		chainId:      pposConfig.ChainId,
		contractAddr: pposConfig.SlashContract,
		credentials:  credentials,
	}
	return &SlashContract{executor}
}

/**
 * 举报双签
 *
 * @param data 证据的json值
 * @return
 */
func (sc *SlashContract) ReportDoubleSign(duplicateSignType common.DuplicateSignType, data string) (common.TransactionHash, error) {
	params := []interface{}{codec.UInt32{ValueInner: duplicateSignType.GetValue()}, codec.Utf8String{ValueInner: data}}
	f := common.NewFunction(common.REPORT_DOUBLESIGN_FUNC_TYPE, params)

	var receipt common.TransactionHash
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

/**
 * 查询节点是否已被举报过多签
 *
 * @param doubleSignType 代表双签类型，1：prepare，2：viewChange
 * @param nodeId         举报的节点Id
 * @param blockNumber    多签的块高
 * @return
 */
func (sc *SlashContract) CheckDoubleSign(doubleSignType common.DuplicateSignType, nodeId string, blockNumber *big.Int) (string, error) {
	params := []interface{}{codec.UInt32{ValueInner: doubleSignType.GetValue()}, codec.NodeId{HexStringId: nodeId}, codec.UInt64{ValueInner: blockNumber}}
	f := common.NewFunction(common.CHECK_DOUBLESIGN_FUNC_TYPE, params)

	var result string
	err := sc.executor.CallWithResult(f, &result)
	return result, err
}
