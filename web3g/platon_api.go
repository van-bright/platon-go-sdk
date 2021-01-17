package web3g

import (
	"fmt"
	"strconv"
)

func (web3g *Web3g) PlatonProtocolVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonProtocolVersion, nil)
	return ParseHttpResponseToString(resp, err)
}

type SyncingInfo struct {
	StartingBlock string
	CurrentBlock  string
	HighestBlock  string
}

type PlatonSyncingResp struct {
	Syncing bool
	SyncingInfo
}

func (web3g *Web3g) PlatonSyncing() (PlatonSyncingResp, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonSyncing, nil)
	if err != nil {
		return PlatonSyncingResp{}, err
	}
	if resp.Error != nil {
		return PlatonSyncingResp{}, fmt.Errorf("%s", resp.Error)
	}

	if _, ok := resp.Result.(bool); ok {
		return PlatonSyncingResp{Syncing: false}, nil
	}

	info := resp.Result.(SyncingInfo)
	return PlatonSyncingResp{true, info}, nil
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

	accounts := resp.Result.([]interface{})
	result := make([]string, len(accounts))
	for i,v := range accounts {
		result[i] = v.(string)
	}
	return result, nil
}

func (web3g *Web3g) PlatonBlockNumber() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(PlatonBlockNumber, nil)
	return ParseHttpResponseToString(resp, err)
}

type PlatonGetBalanceReq struct {
	Address     string
	TagOrNumber interface{}
}

func (web3g *Web3g) PlatonGetBalance(req PlatonGetBalanceReq) (string, error) {
	reqp := make([]string, 2)
	reqp[0] = req.Address
	if pos, err := ParseTagOrNumber(req.TagOrNumber); err != nil {
		return "", err
	} else {
		reqp[1] = pos
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetBalance, reqp)
	return ParseHttpResponseToString(resp, err)
}

type PlatonGetStorageAtReq struct {
	Address       string
	PositionIndex int
	TagOrNumber   interface{}
}
func (web3g *Web3g) PlatonGetStorageAt(req PlatonGetStorageAtReq) (string, error) {
	reqp := make([]string, 3)
	reqp[0] = req.Address
	reqp[1] = strconv.Itoa(req.PositionIndex)
	if pos, err := ParseTagOrNumber(req.TagOrNumber); err != nil {
		return "", err
	} else {
		reqp[3] = pos
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetStorageAt, reqp)
	return ParseHttpResponseToString(resp, err)
}

type PlatonGetTransactionCountReq struct {
	Address       string
	TagOrNumber   interface{}
}
func (web3g *Web3g) PlatonGetTransactionCount(req PlatonGetTransactionCountReq) (string, error) {
	reqp := make([]string, 2)
	reqp[0] = req.Address
	if pos, err := ParseTagOrNumber(req.TagOrNumber); err != nil {
		return "", err
	} else {
		reqp[1] = pos
	}

	resp, err := web3g.httpClient.PostAsResponse(PlatonGetTransactionCount, reqp)
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

type PlatonGetCodeReq struct {
	Address string
	TagOrNumber interface{}
}
func (web3g *Web3g) PlatonGetCode(req PlatonGetCodeReq) (string, error) {
	reqp := make([]string, 2)
	reqp[0] = req.Address
	if pos, err := ParseTagOrNumber(req.TagOrNumber); err != nil {
		return "", err
	} else {
		reqp[1] = pos
	}
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetCode, reqp)
	return ParseHttpResponseToString(resp, err)
}

type PlatonSignReq struct {
	Address string
	Data    string
}
func (web3g *Web3g) PlatonSign(req PlatonSignReq) (string, error) {
	reqp := make([]string, 2)
	reqp[0] = req.Address
	reqp[1] = req.Data
	resp, err := web3g.httpClient.PostAsResponse(PlatonGetCode, reqp)
	return ParseHttpResponseToString(resp, err)
}