package web3go

import (
	"context"
	"encoding/json"
	"fmt"
	platon "github.com/oldmanfan/platon-go-sdk"
	"github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/common/hexutil"
	"github.com/oldmanfan/platon-go-sdk/core/types"
	"github.com/oldmanfan/platon-go-sdk/ethclient"
	"github.com/oldmanfan/platon-go-sdk/params"
	"math/big"
)

type PlatonRPC struct {
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
func (platon *PlatonRPC) NetworkID() (*big.Int, error) {
	return platon.client.NetworkID(ctx)
}

// to check if network is working by net_listening
func (platon *PlatonRPC) NetListening() (bool, error) {
	return platon.client.NetListening(ctx)
}

// to get peer count by net_peerCount
func (platon *PlatonRPC) NetPeerCount() (uint64, error) {
	return platon.client.NetPeerCount(ctx)
}

// to get client version by web3_clientVersion
func (platon *PlatonRPC) ClientVersion() (string, error) {
	return platon.client.ClientVersion(ctx)
}

// to get hash of input string by web3_sha3
func (platon *PlatonRPC) Sha3(str string) (common.Hash, error) {
	hexWith0x := hexutil.Encode([]byte(str))
	return platon.client.Sha3(ctx, hexWith0x)
}

// to get platon protocol version by platon_protocolVersion
func (platon *PlatonRPC) ProtocolVersion() (uint64, error) {
	return platon.client.ProtocolVersion(ctx)
}

// to get syncing status of network by platon_syncing
// if returns (nil, nil) means is not syncing.
// if returns (*pointer, nil) means is syncing.
// if returns (any, error) means error happens.
func (platon *PlatonRPC) Syncing() (*platon.SyncProgress, error) {
	return platon.client.SyncProgress(ctx)
}

// to get the suggest gas price at now by platon_gasPrice
func (platon *PlatonRPC) GasPrice() (*big.Int, error) {
	return platon.client.SuggestGasPrice(ctx)
}

// to estimate gas count by platon_estimateGas
func (platon *PlatonRPC) EstimateGasLimit(msg platon.CallMsg2) (uint64, error) {
	return platon.client.EstimateGas(ctx, msg)
}

// to get accounts by platon_accounts
func (platon *PlatonRPC) Accounts() ([]string, error) {
	return platon.client.Accounts(ctx)
}

// to get the latest block number by platon_blockNumber
func (platon *PlatonRPC) BlockNumber() (uint64, error) {
	return platon.client.BlockNumber(ctx)
}

// to get the balance of account by platon_getBalance
// @param account an account on platon starts with 'atp' or 'atx'
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) BalanceAt(account string, option interface{}) (*big.Int, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}

	balance, err := platon.client.BalanceAt(ctx, account, blockNumber)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// to get the storage of account at pos with key by platon_getStorageAt
// @param account an account on platon starts with 'atp' or 'atx'
// @param key     which value of this key will be returned
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) StorageAt(account string, key common.Hash, option interface{}) ([]byte, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}

	storage, err := platon.client.StorageAt(ctx, common.MustBech32ToAddress(account), key, blockNumber)
	if err != nil {
		return nil, err
	}
	return storage, nil
}

// to get the code of account at pos with key by platon_getCode
// @param account an account on platon starts with 'atp' or 'atx'
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) CodeAt(account string, option interface{}) ([]byte, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}

	code, err := platon.client.CodeAt(ctx, account, blockNumber)
	if err != nil {
		return nil, err
	}
	return code, nil
}

// to get nonce of account at pos with key by platon_getTransactionCount
// @param account an account on platon starts with 'atp' or 'atx'
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) NonceAt(account string, option interface{}) (uint64, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return 0, err
	}

	nonce, err := platon.client.NonceAt(ctx, account, blockNumber)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// to get transaction count by block hash by platon_getBlockTransactionCountByHash
func (platon *PlatonRPC) TransactionCountByHash(blockHash common.Hash) (uint, error) {
	count, err := platon.client.TransactionCount(ctx, blockHash)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// to get transaction count by block number  by platon_getBlockTransactionCountByNumber
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) TransactionCountByNumber(option interface{}) (uint, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return 0, err
	}

	count, err := platon.client.TransactionCountByNumber(ctx, blockNumber)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// to get block info by block hash  by platon_getBlockByHash
// @param hash  block hash
func (platon *PlatonRPC) BlockByHash(hash string) (string, error) {
	return platon.client.BlockByHash(ctx, hash)
}

// to get block info by block number  by platon_getBlockByNumber
// @param option     the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) BlockByNumber(option interface{}) (string, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return "", err
	}
	return platon.client.BlockByNumber(ctx, blockNumber)
}

// to get transaction info by block hash by platon_getTransactionByHash
// @param hash  the block hash.
func (platon *PlatonRPC) TransactionByHash(hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return platon.client.TransactionByHash(ctx, hash)
}

// to get transaction info by block hash and index by platon_getTransactionByBlockHashAndIndex
// @param hash  the block hash.
// @param index the specific index
func (platon *PlatonRPC) TransactionByBlockHashAndIndex(blockHash common.Hash, index uint) (*types.Transaction, error) {
	return platon.client.TransactionByBlockHashAndIndex(ctx, blockHash, index)
}

// to get transaction info by block number and index by platon_getTransactionByBlockNumberAndIndex
// @param option  the block number as bignumber, or 'latest', 'pending', 'earliest'.
// @param index the specific index
func (platon *PlatonRPC) TransactionByBlockNumberAndIndex(option interface{}, index uint) (*types.Transaction, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}
	return platon.client.TransactionByBlockNumberAndIndex(ctx, blockNumber, index)
}

// to get transaction receipt by platon_getTransactionReceipt
// @param txHash  the transaction hash.
func (platon *PlatonRPC) TransactionReceipt(txHash common.Hash) (*types.Receipt, error) {
	return platon.client.TransactionReceipt(ctx, txHash)
}

// to get transaction receipt by platon_getLogs
// @param q  the filter query
func (platon *PlatonRPC) FilterLogs(q platon.FilterQuery) ([]types.Log, error) {
	return platon.client.GetLogs(ctx, q)
}

// to call a contract by msg at option by platon_call,
// notice this msg is never mined to block.
// @param msg  the call message
// @param option  the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) CallContract(msg platon.CallMsg, option interface{}) ([]byte, error) {
	blockNumber, err := checkBlockNumber(option)
	if err != nil {
		return nil, err
	}
	return platon.client.CallContract(ctx, msg, blockNumber)
}

// to send a transaction by platon_sendRawTransaction
// notice this msg is never mined to block.
// @param msg  the call message
// @param option  the block number as bignumber, or 'latest', 'pending', 'earliest'.
func (platon *PlatonRPC) SendRawTransaction(tx *types.Transaction) (json.RawMessage, error) {
	return platon.client.SendRawTransaction(ctx, tx)
}

func (platon *PlatonRPC) SendTransaction(tx *types.Transaction) (json.RawMessage, error) {
	return platon.client.SendTransaction(ctx, tx)
}

// to get program version by admin_getProgramVersion
func (platon *PlatonRPC) AdminGetProgramVersion() (*params.ProgramVersion, error) {
	return platon.client.GetProgramVersion(ctx)
}

// to get proof of bls by platon_getSchnorrNIZKProve
func (platon *PlatonRPC) GetSchnorrNIZKProve() (string, error) {
	return platon.client.GetSchnorrNIZKProve(ctx)
}

// to get data dir of admin by admin_dataDir
func (platon *PlatonRPC) AdminDataDir() (string, error) {
	return platon.client.AdminDataDir(ctx)
}

// to get peers info by admin_peers
func (platon *PlatonRPC) AdminPeers() ([]string, error) {
	return platon.client.AdminPeers(ctx)
}

// to get node info by admin_nodeInfo
func (platon *PlatonRPC) AdminNodeInfo() (string, error) {
	return platon.client.AdminNodeInfo(ctx)
}

// to add a peer by admin_addPeer
func (platon *PlatonRPC) AdminAddPeer(peer string) (bool, error) {
	return platon.client.AdminAddPeer(ctx, peer)
}

// to sign a hex string with unlocked wallet signer
func (platon *PlatonRPC) Sign(req *platon.SignReq) (string, error) {
	return platon.client.Sign(ctx, req)
}

// Return double sign report data.
func (platon *PlatonRPC) Evidences() (string, error) {
	return platon.client.Evidences(ctx)
}

func (platon *PlatonRPC) NewFilter(q platon.FilterQuery) (*big.Int, error) {
	return platon.client.NewFilter(ctx, q)
}

// Add a filter when new block created,
// use `GetFilterChanges` to check state logs.
func (platon *PlatonRPC) NewBlockFilter() (*big.Int, error) {
	return platon.client.NewBlockFilter(ctx)
}

// Creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, call platon_getFilterChanges.
func (platon *PlatonRPC) NewPendingTransactionFilter() (*big.Int, error) {
	return platon.client.NewPendingTransactionFilter(ctx)
}

// Uninstalls a filter with given id.
// Should always be called when watch is no longer needed.
//Additionally Filters timeout when they aren't requested with platon_getFilterChanges for a period of time.
func (platon *PlatonRPC) UninstallFilter(filterId *big.Int) bool {
	return platon.client.UninstallFilter(ctx, filterId)
}

// Polling method for a filter, which returns an array of logs which occurred since last poll.
func (platon *PlatonRPC) GetFilterChanges(filterId *big.Int) ([]types.Log, error) {
	return platon.client.GetFilterChanges(ctx, filterId)
}

// Returns an array of all logs matching filter with given id.
func (platon *PlatonRPC) GetFilterLogs(filterId *big.Int) ([]types.Log, error) {
	return platon.client.GetFilterLogs(ctx, filterId)
}

// Returns an array of all logs matching a given filter object.
func (platon *PlatonRPC) GetLogs(q platon.FilterQuery) ([]types.Log, error) {
	return platon.client.GetLogs(ctx, q)
}
