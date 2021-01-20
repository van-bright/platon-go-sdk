package web3g

import (
	"encoding/json"
	"fmt"
	"math/big"
)

func (web3g *Web3g) PlatonProtocolVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonProtocolVersion, nil)
	return ParseHttpResponseToString(resp, err)
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
	if err := json.Unmarshal(resp.Result, &result); err != nil {
		return PlatonSyncingResp{}, err
	}
	return PlatonSyncingResp{true, result}, nil
}

func (web3g *Web3g) PlatonGasPrice() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGasPrice, nil)
	return ParseHttpResponseToString(resp, err)
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

func (web3g *Web3g) PlatonBlockNumber() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonBlockNumber, nil)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonGetBalance(address string, pos interface{}) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBalance, address, pos)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonGetStorageAt(req PlatonGetStorageAtReq) (string, error) {
	pos, err := ParseTagOrNumber(req.TagOrNumber)
	if err != nil {
		return "", err
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetStorageAt, req.Address, req.PositionIndex, pos)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonGetTransactionCount(req PlatonGetTransactionCountReq) (string, error) {
	pos, err := ParseTagOrNumber(req.TagOrNumber)
	if err != nil {
		return "", err
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionCount, req.Address, pos)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonGetBlockTransactionCountByHash(blockHash string) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBlockTransactionCountByHash, blockHash)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonGetBlockTransactionCountByNumber(tagOrNumber interface{}) (string, error) {
	pos, err := ParseTagOrNumber(tagOrNumber)
	if err != nil {
		return "", err
	}
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBlockTransactionCountByNumber, pos)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonGetCode(req PlatonGetCodeReq) (string, error) {
	pos, err := ParseTagOrNumber(req.TagOrNumber)
	if err != nil {
		return "", err
	}
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetCode, req.Address, pos)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonSign(req PlatonSignReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonSign, req.Address, req.Data)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonSendTransaction(req PlatonSendTransactionReq) (string, error) {
	if req.Gas == nil {
		req.Gas = (*BigInt)(big.NewInt(90000))

	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonSendTransaction, req)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonSendRawTransaction(req PlatonSendRawTransactionReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonSendRawTransaction, req)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonCall(req PlatonCallReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonCall, req)
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) PlatonEstimateGas(req PlatonEstimateGasReq) (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonEstimateGas, req)
	return ParseHttpResponseToString(resp, err)
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
