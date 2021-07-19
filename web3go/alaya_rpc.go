package web3go

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	platon "platon-go-sdk"
	"platon-go-sdk/common"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/core/types"
	"platon-go-sdk/ethclient"
	"platon-go-sdk/params"
)

type AlayaRPC struct {
	client *ethclient.Client
}

var ctx = context.Background()

func checkBlockNumber(pos interface{}) (interface{}, error) {
	var e = fmt.Errorf("option should be 'latest', 'pending', 'earliest' or a big number")
	if pos == nil {
		return "latest", nil
	}

	if opt, ok := pos.(string); ok {
		if opt == "latest" || opt == "pending" || opt == "earliest" {
			return opt, nil
		} else {
			return nil, e
		}
	} else {
		if blockNumber, ok := pos.(*big.Int); !ok {
			return nil, e
		} else {
			return blockNumber, nil
		}
	}
}

// to get net version by net_version
func (alaya *AlayaRPC) NetworkID() (*big.Int, error) {
	return alaya.client.NetworkID(ctx)
}

// to check if network is working by net_listening
func (alaya *AlayaRPC) NetListening() (bool, error) {
	return alaya.client.NetListening(ctx)
}

// to get peer count by net_peerCount
func (alaya *AlayaRPC) NetPeerCount() (uint64, error) {
	return alaya.client.NetPeerCount(ctx)
}

// to get client version by web3_clientVersion
func (alaya *AlayaRPC) ClientVersion() (string, error) {
	return alaya.client.ClientVersion(ctx)
}

// to get hash of input string by web3_sha3
func (alaya *AlayaRPC) Sha3(str string) (common.Hash, error) {
	hexWith0x := hexutil.Encode([]byte(str))
	return alaya.client.Sha3(ctx, hexWith0x)
}

// to get platon protocol version by platon_protocolVersion
func (alaya *AlayaRPC) ProtocolVersion() (uint64, error) {
	return alaya.client.ProtocolVersion(ctx)
}

// to get syncing status of network by platon_syncing
// if returns (nil, nil) means is not syncing.
// if returns (*pointer, nil) means is syncing.
// if returns (any, error) means error happens.
func (alaya *AlayaRPC) Syncing() (*platon.SyncProgress, error) {
	return alaya.client.SyncProgress(ctx)
}

// to get the suggest gas price at now by platon_gasPrice
func (alaya *AlayaRPC) GasPrice() (*big.Int, error) {
	return alaya.client.SuggestGasPrice(ctx)
}

// to estimate gas count by platon_estimateGas
func (alaya *AlayaRPC) EstimateGasLimit(msg platon.CallMsg2) (uint64, error) {
	return alaya.client.EstimateGas(ctx, msg)
}

// to get accounts by platon_accounts
func (alaya *AlayaRPC) Accounts() ([]string, error) {
	return alaya.client.Accounts(ctx)
}

// to get the latest block number by platon_blockNumber
func (alaya *AlayaRPC) BlockNumber() (uint64, error) {
	return alaya.client.BlockNumber(ctx)
}

// to get the balance of account by platon_getBalance
// @param account an account on platon starts with 'atp' or 'atx'
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) BalanceAt(account string, option interface{}) (*big.Int, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}

	balance, err := alaya.client.BalanceAt(ctx, account, blockNumber)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// to get the storage of account at pos with key by platon_getStorageAt
// @param account an account on platon starts with 'atp' or 'atx'
// @param key     which value of this key will be returned
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) StorageAt(account string, key common.Hash, option interface{}) ([]byte, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}

	storage, err := alaya.client.StorageAt(ctx, common.MustBech32ToAddress(account), key, blockNumber)
	if err != nil {
		return nil, err
	}
	return storage, nil
}

// to get the code of account at pos with key by platon_getCode
// @param account an account on platon starts with 'atp' or 'atx'
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) CodeAt(account string, option interface{}) ([]byte, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}

	code, err := alaya.client.CodeAt(ctx, account, blockNumber)
	if err != nil {
		return nil, err
	}
	return code, nil
}

// to get nonce of account at pos with key by platon_getTransactionCount
// @param account an account on platon starts with 'atp' or 'atx'
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) NonceAt(account string, option interface{}) (uint64, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return 0, err
	}

	nonce, err := alaya.client.NonceAt(ctx, account, blockNumber)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// to get transaction count by block hash by platon_getBlockTransactionCountByHash
func (alaya *AlayaRPC) TransactionCountByHash(blockHash common.Hash) (uint, error) {
	count, err := alaya.client.TransactionCount(ctx, blockHash)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// to get transaction count by block number  by platon_getBlockTransactionCountByNumber
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) TransactionCountByNumber(option interface{}) (uint, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return 0, err
	}

	count, err := alaya.client.TransactionCountByNumber(ctx, blockNumber)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// to get block info by block hash  by platon_getBlockByHash
// @param hash  block hash
func (alaya *AlayaRPC) BlockByHash(hash string) (string, error) {
	return alaya.client.BlockByHash(ctx, hash)
}

// to get block info by block number  by platon_getBlockByNumber
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) BlockByNumber(option interface{}) (string, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return "", err
	}
	return alaya.client.BlockByNumber(ctx, blockNumber)
}

// to get transaction info by block hash by platon_getTransactionByHash
// @param hash  the block hash.
func (alaya *AlayaRPC) TransactionByHash(hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return alaya.client.TransactionByHash(ctx, hash)
}

// to get transaction info by block hash and index by platon_getTransactionByBlockHashAndIndex
// @param hash  the block hash.
// @param index the specific index
func (alaya *AlayaRPC) TransactionByBlockHashAndIndex(blockHash common.Hash, index uint) (*types.Transaction, error) {
	return alaya.client.TransactionByBlockHashAndIndex(ctx, blockHash, index)
}

// to get transaction info by block number and index by platon_getTransactionByBlockNumberAndIndex
// @param option  the block number as bignumber, or 'latest', 'pending', 'earliest'.
// @param index the specific index
func (alaya *AlayaRPC) TransactionByBlockNumberAndIndex(option interface{}, index uint) (*types.Transaction, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}
	return alaya.client.TransactionByBlockNumberAndIndex(ctx, blockNumber, index)
}

// to get transaction receipt by platon_getTransactionReceipt
// @param txHash  the transaction hash.
func (alaya *AlayaRPC) TransactionReceipt(txHash common.Hash) (*types.Receipt, error) {
	return alaya.client.TransactionReceipt(ctx, txHash)
}

// to get transaction receipt by platon_getLogs
// @param q  the filter query
func (alaya *AlayaRPC) FilterLogs(q platon.FilterQuery) ([]types.Log, error) {
	return alaya.client.GetLogs(ctx, q)
}

// to call a contract by msg at option by platon_call,
// notice this msg is never mined to block.
// @param msg  the call message
// @param option  the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) CallContract(msg platon.CallMsg, option interface{}) ([]byte, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}
	return alaya.client.CallContract(ctx, msg, blockNumber)
}

// to send a transaction by platon_sendRawTransaction
// notice this msg is never mined to block.
// @param msg  the call message
// @param option  the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (alaya *AlayaRPC) SendRawTransaction(tx *types.Transaction) (json.RawMessage, error) {
	return alaya.client.SendRawTransaction(ctx, tx)
}

func (alaya *AlayaRPC) SendTransaction(tx *types.Transaction) (json.RawMessage, error) {
	return alaya.client.SendTransaction(ctx, tx)
}

// to get program version by admin_getProgramVersion
func (alaya *AlayaRPC) AdminGetProgramVersion() (*params.ProgramVersion, error) {
	return alaya.client.GetProgramVersion(ctx)
}

// to get proof of bls by platon_getSchnorrNIZKProve
func (alaya *AlayaRPC) GetSchnorrNIZKProve() (string, error) {
	return alaya.client.GetSchnorrNIZKProve(ctx)
}

// to get data dir of admin by admin_dataDir
func (alaya *AlayaRPC) AdminDataDir() (string, error) {
	return alaya.client.AdminDataDir(ctx)
}

// to get peers info by admin_peers
func (alaya *AlayaRPC) AdminPeers() ([]string, error) {
	return alaya.client.AdminPeers(ctx)
}

// to get node info by admin_nodeInfo
func (alaya *AlayaRPC) AdminNodeInfo() (string, error) {
	return alaya.client.AdminNodeInfo(ctx)
}

// to add a peer by admin_addPeer
func (alaya *AlayaRPC) AdminAddPeer(peer string) (bool, error) {
	return alaya.client.AdminAddPeer(ctx, peer)
}

// to sign a hex string with unlocked wallet signer
func (alaya *AlayaRPC) Sign(req *platon.SignReq) (string, error) {
	return alaya.client.Sign(ctx, req)
}

// Return double sign report data.
func (alaya *AlayaRPC) Evidences() (string, error) {
	return alaya.client.Evidences(ctx)
}

func (alaya *AlayaRPC) NewFilter(q platon.FilterQuery) (*big.Int, error) {
	return alaya.client.NewFilter(ctx, q)
}

// Add a filter when new block created,
// use `GetFilterChanges` to check state logs.
func (alaya *AlayaRPC) NewBlockFilter() (*big.Int, error) {
	return alaya.client.NewBlockFilter(ctx)
}

// Creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, call platon_getFilterChanges.
func (alaya *AlayaRPC) NewPendingTransactionFilter() (*big.Int, error) {
	return alaya.client.NewPendingTransactionFilter(ctx)
}

// Uninstalls a filter with given id.
// Should always be called when watch is no longer needed.
//Additionally Filters timeout when they aren't requested with platon_getFilterChanges for a period of time.
func (alaya *AlayaRPC) UninstallFilter(filterId *big.Int) bool {
	return alaya.client.UninstallFilter(ctx, filterId)
}

// Polling method for a filter, which returns an array of logs which occurred since last poll.
func (alaya *AlayaRPC) GetFilterChanges(filterId *big.Int) ([]types.Log, error) {
	return alaya.client.GetFilterChanges(ctx, filterId)
}

// Returns an array of all logs matching filter with given id.
func (alaya *AlayaRPC) GetFilterLogs(filterId *big.Int) ([]types.Log, error) {
	return alaya.client.GetFilterLogs(ctx, filterId)
}

// Returns an array of all logs matching a given filter object.
func (alaya *AlayaRPC) GetLogs(q platon.FilterQuery) ([]types.Log, error) {
	return alaya.client.GetLogs(ctx, q)
}
