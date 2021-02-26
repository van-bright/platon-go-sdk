package web3g

import "math/big"

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
	PositionIndex *big.Int
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
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      *big.Int `json:"gas"`
	GasPrice *big.Int `json:"gasPrice"`
	Value    *big.Int `json:"value"`
	Data     string `json:"data"`
	Nonce    *big.Int `json:"nonce"`
}

type PlatonSendRawTransactionReq struct {
	Data string `json:"data"`
}

type PlatonCallReq struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      *big.Int `json:"gas"`
	GasPrice *big.Int `json:"gasPrice"`
	Value    *big.Int `json:"value"`
	Data     string `json:"data"`
	// FIXME(liangqin.fan): TagOrNumber 的json名称是什么?
	TagOrNumber interface{} `json:"number"`
}

type PlatonEstimateGasReq PlatonCallReq

type PlatonTransaction struct {
	Hash string `json:"hash"`
	Nonce *big.Int `json:"nonce,omitempty"`
	BlockHash string `json:"blockHash,omitempty"`
	BlockNumber *big.Int `json:"blockNumber,omitempty"`
	TransactionIndex *big.Int `json:"transactionIndex,omitempty"`
	From string `json:"from,omitempty"`
	To string `json:"to,omitempty"`
	Value *big.Int `json:"value,omitempty"`
	Gas *big.Int `json:"gas,omitempty"`
	GasPrice *big.Int `json:"gasPrice,omitempty"`
	Input string `json:"input,omitempty"`
}

type PlatonBlock struct {
	Number *big.Int `json:"number,omitempty"`
	Hash   string `json:"hash,omitempty"`
	ParentHash string `json:"parentHash"`
	Nonce      string `json:"nonce,omitempty"`
	Sha3Uncles string `json:"sha3Uncles"`
	LogsBloom  string `json:"logsBloom,omitempty"`
	TransactionsRoot string `json:"transactionsRoot"`
	StateRoot  string                `json:"stateRoot"`
	ReceiptsRoot string              `json:"receiptsRoot"`
	Miner        string              `json:"miner"`
	Difficulty *big.Int                `json:"difficulty,omitempty"`
	TotalDifficulty *big.Int           `json:"totalDifficulty,omitempty"`
	ExtraData string                 `json:"extraData"`
	Size      *big.Int                 `json:"size"`
	GasLimit  *big.Int                 `json:"gasLimit"`
	GasUsed   *big.Int                 `json:"gasUsed"`
	Timestamp *big.Int                 `json:"timestamp"`
	Transactions []PlatonTransaction `json:"transactions"`
	Uncles    []string               `json:"uncles,omitempty"`
}

// 当Hashes不是nil时, 其它值都是nil或它们的零值
type PlatonTransactionLog  struct {
	Hashes []string `json:"hashes,omitempty"`
	Type string `json:"type,omitempty"`
	LogIndex *big.Int `json:"logIndex,omitempty"`
	TransactionIndex *big.Int `json:"transactionIndex,omitempty"`
	TransactionHash string `json:"transactionHash,omitempty"`
	BlockHash string `json:"blockHash,omitempty"`
	BlockNumber *big.Int `json:"blockNumber,omitempty"`
	Address string `json:"address,omitempty"`
	Data string `json:"data,omitempty"`
	Topics []string `json:"topics,omitempty"`
}

type PlatonTransactionReceipt struct {
	TransactionHash string `json:"transactionHash"`
	TransactionIndex *big.Int `json:"transactionIndex"`
	BlockHash string `json:"blockHash"`
	BlockNumber *big.Int `json:"blockNumber"`
	CumulativeGasUsed *big.Int `json:"cumulativeGasUsed"`
	GasUsed *big.Int `json:"gasUsed"`
	ContractAddress string `json:"contractAddress"`
	Logs []PlatonTransactionLog `json:"logs"`
}

type PlatonGetLogsReq struct {
	Topics []interface{} `json:"topics"`
}
