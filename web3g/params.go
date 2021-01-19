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
	PositionIndex big.Int
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
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Value    uint64 `json:"value"`
	Data     string `json:"data"`
	Nonce    uint64 `json:"nonce"`
}

type PlatonSendRawTransactionReq struct {
	Data string `json:"data"`
}

type PlatonCallReq struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Value    uint64 `json:"value"`
	Data     string `json:"data"`
	// FIXME(liangqin.fan): TagOrNumber 的json名称是什么?
	TagOrNumber interface{} `json:"number"`
}

type PlatonEstimateGasReq PlatonCallReq

type TransactionOnBlock struct {
	Hash string `json:"hash"`
	Nonce big.Int `json:"nonce,omitempty"`
	BlockHash string `json:"blockHash,omitempty"`
	BlockNumber big.Int `json:"blockNumber,omitempty"`
	TransactionIndex big.Int `json:"transactionIndex,omitempty"`
	From string `json:"from,omitempty"`
	To string `json:"to,omitempty"`
	Value big.Int `json:"value,omitempty"`
	Gas big.Int `json:"gas,omitempty"`
	GasPrice big.Int `json:"gasPrice,omitempty"`
	Input string `json:"input,omitempty"`
}

type PlatonBlock struct {
	Number big.Int `json:"number,omitempty"`
	Hash   string `json:"hash,omitempty"`
	ParentHash string `json:"parentHash"`
	Nonce      string `json:"nonce,omitempty"`
	Sha3Uncles string `json:"sha3Uncles"`
	LogsBloom  string `json:"logsBloom,omitempty"`
	TransactionsRoot string `json:"transactionsRoot"`
	StateRoot  string `json:"stateRoot"`
	ReceiptsRoot string `json:"receiptsRoot"`
	Miner        string `json:"miner"`
	Difficulty big.Int `json:"difficulty"`
	TotalDifficulty big.Int `json:"totalDifficulty"`
	ExtraData string `json:"extraData"`
	Size      big.Int `json:"size"`
	GasLimit  big.Int `json:"gasLimit"`
	GasUsed   big.Int `json:"gasUsed"`
	Timestamp big.Int `json:"timestamp"`
	Transactions []TransactionOnBlock `json:"transactions"`
	Uncles    []string `json:"uncles"`
}
