package web3g

import (
	"encoding/json"
	"fmt"
	"math/big"
)

func (web3g *Web3g) PlatonProtocolVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonProtocolVersion, nil)

	var version string
	err = ParseHttpResponseToResult(resp, &version, err)
	return version, err
}

func (web3g *Web3g) PlatonSyncing() (PlatonSyncingResp, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonSyncing, nil)
	if err != nil {
		return PlatonSyncingResp{}, err
	}
	if resp.Error != nil {
		return PlatonSyncingResp{}, fmt.Errorf("%s", resp.Error)
	}

	var ok = true
	if err := json.Unmarshal(resp.Result, &ok); err == nil {
		return PlatonSyncingResp{Syncing: false}, nil
	}

	var result SyncingInfo
	err = ParseHttpResponseToResult(resp, &result, err)
	return PlatonSyncingResp{
		true,
		result,
	}, err
}

func (web3g *Web3g) PlatonGasPrice() (*big.Int, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGasPrice, nil)
	var price big.Int
	err =  ParseHttpResponseToResult(resp, &price, err)
	return &price, err
}

func (web3g *Web3g) PlatonAccounts() ([]string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonAccounts, nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("%s", resp.Error)
	}

	var accounts []string
	if err := json.Unmarshal(resp.Result, &accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (web3g *Web3g) PlatonBlockNumber() (uint64, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonBlockNumber, nil)
	var number uint64
	err = ParseHttpResponseToResult(resp, &number, err)
	return number, err
}

func (web3g *Web3g) PlatonGetBalance(address string, pos interface{}) (*big.Int, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBalance, address, pos)

	var balance big.Int
	err = ParseHttpResponseToResult(resp, &balance, err)
	return &balance, err
}

func (web3g *Web3g) PlatonGetStorageAt(req PlatonGetStorageAtReq) ([]byte, error) {
	pos, err := ParseTagOrNumber(req.TagOrNumber)
	if err != nil {
		return nil, err
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetStorageAt, req.Address, req.PositionIndex, pos)
	var storage []byte
	err = ParseHttpResponseToResult(resp, &storage, err)
	return storage, err
}

func (web3g *Web3g) PlatonGetTransactionCount(req PlatonGetTransactionCountReq) (uint64, error) {
	pos, err := ParseTagOrNumber(req.TagOrNumber)
	if err != nil {
		return 0, err
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionCount, req.Address, pos)
	var count uint64
	err = ParseHttpResponseToResult(resp, &count, err)
	return count, err
}

func (web3g *Web3g) PlatonGetBlockTransactionCountByHash(blockHash string) (uint64, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBlockTransactionCountByHash, blockHash)
	var count uint64
	err = ParseHttpResponseToResult(resp, &count, err)
	return count, err
}

func (web3g *Web3g) PlatonGetBlockTransactionCountByNumber(tagOrNumber interface{}) (uint64, error) {
	pos, err := ParseTagOrNumber(tagOrNumber)
	if err != nil {
		return 0, err
	}
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBlockTransactionCountByNumber, pos)
	var count uint64
	err = ParseHttpResponseToResult(resp, &count, err)
	return count, err
}

func (web3g *Web3g) PlatonGetCode(req PlatonGetCodeReq) ([]byte, error) {
	pos, err := ParseTagOrNumber(req.TagOrNumber)
	if err != nil {
		return nil, err
	}
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetCode, req.Address, pos)
	var code []byte
	err =  ParseHttpResponseToResult(resp, &code, err)
	return code, err
}

func (web3g *Web3g) PlatonSign(req PlatonSignReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonSign, req.Address, req.Data)

	var sign string
	err =  ParseHttpResponseToResult(resp, &sign, err)
	return sign, err
}

func (web3g *Web3g) PlatonSendTransaction(req PlatonSendTransactionReq) (string, error) {
	if req.Gas == nil {
		req.Gas = (*BigInt)(big.NewInt(90000))

	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonSendTransaction, req)
	var sendResult string
	err = ParseHttpResponseToResult(resp, &sendResult, err)
	return sendResult, err
}

func (web3g *Web3g) PlatonSendRawTransaction(req PlatonSendRawTransactionReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonSendRawTransaction, req)
	var sendResult string
	err = ParseHttpResponseToResult(resp, &sendResult, err)
	return sendResult, err
}

func (web3g *Web3g) PlatonCall(req PlatonCallReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonCall, req)
	var callResult string
	err = ParseHttpResponseToResult(resp, &callResult, err)
	return callResult, err
}

func (web3g *Web3g) PlatonEstimateGas(req PlatonEstimateGasReq) (uint64, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonEstimateGas, req)

	var gas uint64
	err = ParseHttpResponseToResult(resp, &gas, err)
	return gas, err
}

func (web3g *Web3g) PlatonGetBlockByHash(blockHash string, showTxDetail bool) (PlatonBlock, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBlockByHash, blockHash, showTxDetail)
	if err != nil {
		return PlatonBlock{}, err
	}
	var result PlatonBlock
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonGetBlockByNumber(quantity interface{}, showTxDetail bool) (PlatonBlock, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBlockByNumber, quantity, showTxDetail)
	if err != nil {
		return PlatonBlock{}, err
	}
	var result PlatonBlock
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonGetTransactionByHash(txHash string) (PlatonTransaction, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionByHash, txHash)
	if err != nil {
		return PlatonTransaction{}, err
	}
	var result PlatonTransaction
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonGetTransactionByBlockHashAndIndex(txHash string, indexStr string) (PlatonTransaction, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionByBlockHashAndIndex, txHash, indexStr)
	if err != nil {
		return PlatonTransaction{}, err
	}
	var result PlatonTransaction
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonGetTransactionByBlockNumberAndIndex(number string, indexStr string) (PlatonTransaction, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionByBlockNumberAndIndex, number, indexStr)
	if err != nil {
		return PlatonTransaction{}, err
	}
	var result PlatonTransaction
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonGetTransactionReceipt(txHash string) (PlatonTransactionReceipt, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionReceipt, txHash)
	if err != nil {
		return PlatonTransactionReceipt{}, err
	}
	var result PlatonTransactionReceipt
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonNewFilter(fromBlock interface{}, toBlock interface{}, address interface{}, topics interface{}) (*BigInt, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonNewFilter, fromBlock, toBlock, address, topics)
	if err != nil {
		return &BigInt{}, err
	}
	var result BigInt
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) PlatonNewBlockFilter() (*BigInt, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonNewBlockFilter, nil)
	if err != nil {
		return &BigInt{}, err
	}
	var result BigInt
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) PlatonNewPendingTransactionFilter() (*BigInt, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonNewPendingTransactionFilter, nil)
	if err != nil {
		return &BigInt{}, err
	}
	var result BigInt
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) PlatonUninstallFilter(filterId *BigInt) (bool, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonUninstallFilter, filterId)
	if err != nil {
		return false, err
	}
	var result bool
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) PlatonGetFilterChanges(filterId *BigInt) (*PlatonTransactionLog, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetFilterChanges, filterId)
	if err != nil {
		return &PlatonTransactionLog{}, err
	}
	var result PlatonTransactionLog
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) PlatonGetFilterLogs(filterId *BigInt) (*PlatonTransactionLog, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetFilterLogs, filterId)
	if err != nil {
		return &PlatonTransactionLog{}, err
	}
	var result PlatonTransactionLog
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) PlatonGetLogs(topics *PlatonGetLogsReq) (*PlatonTransactionLog, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetLogs, topics)
	if err != nil {
		return &PlatonTransactionLog{}, err
	}
	var result PlatonTransactionLog
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) PlatonEvidences() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonEvidences, nil)

	var evidence string
	err = ParseHttpResponseToResult(resp, &evidence, err)
	return evidence, err
}
