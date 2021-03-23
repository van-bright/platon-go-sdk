package web3g

import "platon-go-sdk/common"

type SyncingInfo struct {
	StartingBlock string
	CurrentBlock  string
	HighestBlock  string
}

type PlatonSyncingResp struct {
	Syncing bool
	SyncingInfo
}

type PlatonGetBalanceReq struct {
	Address     string
	TagOrNumber interface{}
}

type PlatonGetStorageAtReq struct {
	Address       string
	PositionIndex *BigInt
	TagOrNumber   interface{}
}

type PlatonGetTransactionCountReq struct {
	Address     string
	TagOrNumber interface{}
}

type PlatonGetCodeReq struct {
	Address     string
	TagOrNumber interface{}
}

type PlatonSignReq struct {
	Address string
	Data    string
}

type PlatonSendTransactionReq struct {
	From     common.Address   `json:"from"`
	To       common.Address   `json:"to"`
	Gas      *BigInt `json:"gas"`
	GasPrice *BigInt `json:"gasPrice"`
	Value    *BigInt `json:"value"`
	Data     string   `json:"data"`
	Nonce    *BigInt `json:"nonce"`
}

type PlatonSendRawTransactionReq struct {
	Data string `json:"data"`
}

type PlatonCallReq struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Gas      *BigInt `json:"gas"`
	GasPrice *BigInt `json:"gasPrice"`
	Value    *BigInt `json:"value"`
	Data     string   `json:"data"`
	// FIXME(liangqin.fan): TagOrNumber 的json名称是什么?
	TagOrNumber interface{} `json:"number"`
}

type PlatonEstimateGasReq PlatonCallReq

type PlatonTransaction struct {
	Hash             string   `json:"hash"`
	Nonce            *BigInt `json:"nonce,omitempty"`
	BlockHash        string   `json:"blockHash,omitempty"`
	BlockNumber      *BigInt `json:"blockNumber,omitempty"`
	TransactionIndex *BigInt `json:"transactionIndex,omitempty"`
	From             string   `json:"from,omitempty"`
	To               string   `json:"to,omitempty"`
	Value            *BigInt `json:"value,omitempty"`
	Gas              *BigInt `json:"gas,omitempty"`
	GasPrice         *BigInt `json:"gasPrice,omitempty"`
	Input            string   `json:"input,omitempty"`
}

type PlatonBlock struct {
	Number           *BigInt            `json:"number,omitempty"`
	Hash             string              `json:"hash,omitempty"`
	ParentHash       string              `json:"parentHash"`
	Nonce            string              `json:"nonce,omitempty"`
	Sha3Uncles       string              `json:"sha3Uncles"`
	LogsBloom        string              `json:"logsBloom,omitempty"`
	TransactionsRoot string              `json:"transactionsRoot"`
	StateRoot        string              `json:"stateRoot"`
	ReceiptsRoot     string              `json:"receiptsRoot"`
	Miner            string              `json:"miner"`
	Difficulty       *BigInt            `json:"difficulty,omitempty"`
	TotalDifficulty  *BigInt            `json:"totalDifficulty,omitempty"`
	ExtraData        string              `json:"extraData"`
	Size             *BigInt            `json:"size"`
	GasLimit         *BigInt            `json:"gasLimit"`
	GasUsed          *BigInt            `json:"gasUsed"`
	Timestamp        *BigInt            `json:"timestamp"`
	Transactions     []PlatonTransaction `json:"transactions"`
	Uncles           []string            `json:"uncles,omitempty"`
}

// 当Hashes不是nil时, 其它值都是nil或它们的零值
type PlatonTransactionLog struct {
	Hashes           []string `json:"hashes,omitempty"`
	Type             string   `json:"type,omitempty"`
	LogIndex         *BigInt `json:"logIndex,omitempty"`
	TransactionIndex *BigInt `json:"transactionIndex,omitempty"`
	TransactionHash  string   `json:"transactionHash,omitempty"`
	BlockHash        string   `json:"blockHash,omitempty"`
	BlockNumber      *BigInt `json:"blockNumber,omitempty"`
	Address          string   `json:"address,omitempty"`
	Data             string   `json:"data,omitempty"`
	Topics           []string `json:"topics,omitempty"`
}

type PlatonTransactionReceipt struct {
	TransactionHash   string                 `json:"transactionHash"`
	TransactionIndex  *BigInt               `json:"transactionIndex"`
	BlockHash         string                 `json:"blockHash"`
	BlockNumber       *BigInt               `json:"blockNumber"`
	CumulativeGasUsed *BigInt               `json:"cumulativeGasUsed"`
	GasUsed           *BigInt               `json:"gasUsed"`
	ContractAddress   string                 `json:"contractAddress"`
	Logs              []PlatonTransactionLog `json:"logs"`
}

type PlatonGetLogsReq struct {
	Topics []interface{} `json:"topics"`
}
