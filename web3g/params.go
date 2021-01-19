package web3g

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
	PositionIndex int
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

type PlatonBlockInfo struct {

}
