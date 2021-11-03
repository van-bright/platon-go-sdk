package web3go

import (
	"encoding/json"
	"math/big"
	platon "platon-go-sdk"
	"platon-go-sdk/common"
	"platon-go-sdk/core/types"
	"platon-go-sdk/ethclient"
	"platon-go-sdk/params"
)

type Geb3 interface {
	// To get network id (sometimes called chain id)
	NetworkID() (*big.Int, error)
	// To check if the endpoint is listening, `true` for yes, otherwise `false`.
	NetListening() (bool, error)
	// To retrieve count of peer nodes
	NetPeerCount() (uint64, error)
	// To retrieve the client version
	ClientVersion() (string, error)
	// To calculate hash of a specific string in sha3
	Sha3(str string) (common.Hash, error)
	// To retrieve protocol version of platon
	ProtocolVersion() (uint64, error)
	// To check if the network is syncing or not,
	// if Yes, then return an object of syncing status.
	// if No, return nil
	Syncing() (*platon.SyncProgress, error)
	// To retrieve gas price currently
	GasPrice() (*big.Int, error)
	// To estimate gas limit of a CallMsg
	EstimateGasLimit(msg platon.CallMsg2) (uint64, error)
	// Returns a list of addresses owned by client.
	Accounts() ([]string, error)
	// Returns the number of most recent block.
	BlockNumber() (uint64, error)
	// Return balance of bech32Account at pos, its value can be block number, 'latest', 'pending' or 'earliest'
	BalanceAt(bech32Account string, pos interface{}) (*big.Int, error)
	// Return storage of bech32Account and key at pos, its value can be block number, 'latest', 'pending' or 'earliest'
	StorageAt(bech32Account string, key common.Hash, pos interface{}) ([]byte, error)
	// Return code of bech32Account at pos, its value can be block number, 'latest', 'pending' or 'earliest'
	CodeAt(bech32Account string, pos interface{}) ([]byte, error)
	// Return nonce of bech32Account at pos, its value can be block number, 'latest', 'pending' or 'earliest'
	NonceAt(bech32Account string, pos interface{}) (uint64, error)
	// Return count of transaction by block hash.
	TransactionCountByHash(blockHash common.Hash) (uint, error)
	// Return count of transaction by block number, or one of 'latest', 'pending', 'earliest'
	TransactionCountByNumber(option interface{}) (uint, error)
	// Return block info by block hash
	BlockByHash(hash string) (string, error)
	// Return block info by block number, or one of 'latest', 'pending', 'earliest'
	BlockByNumber(option interface{}) (string, error)
	// Returns the information about a transaction requested by transaction hash.
	TransactionByHash(hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	// Returns information about a transaction by block hash and transaction index position.
	TransactionByBlockHashAndIndex(blockHash common.Hash, index uint) (*types.Transaction, error)
	// Returns information about a transaction by block number and transaction index position.
	TransactionByBlockNumberAndIndex(option interface{}, index uint) (*types.Transaction, error)
	// Returns the receipt of a transaction by transaction hash.
	//Note That the receipt is not available for pending transactions.
	TransactionReceipt(txHash common.Hash) (*types.Receipt, error)
	// Call a contract with CallMsg at specific block number or 'latest', 'pending', 'earliest'
	// This CallMsg is never mined to block.
	CallContract(msg platon.CallMsg, option interface{}) ([]byte, error)
	// SendWithRaw a raw transaction to pool to execute.
	SendRawTransaction(tx *types.Transaction) (json.RawMessage, error)
	// SendWithRaw a signed transaction to pool to execute.
	SendTransaction(tx *types.Transaction) (json.RawMessage, error)
	// Query code version and signature.
	AdminGetProgramVersion() (*params.ProgramVersion, error)
	// Get the data directory.
	AdminDataDir() (string, error)
	// Returns the details of the peer connected to the current client.
	AdminPeers() ([]string, error)
	// Add a node. 'enode://xxxx '
	AdminAddPeer(peer string) (bool, error)
	// Returns the current client node details.
	AdminNodeInfo() (string, error)
	// Get proof of bls, returns hex string
	GetSchnorrNIZKProve() (string, error)
	// Return double sign report data.
	Evidences() (string, error)
	// to Sign a hex string, with unlocked Signer
	Sign(req *platon.SignReq) (string, error)
	// Add a filter while requirements full filled.
	NewFilter(q platon.FilterQuery) (*big.Int, error)
	// Add a filter when new block created,
	// use `GetFilterChanges` to check state logs.
	NewBlockFilter() (*big.Int, error)
	// Creates a filter in the node, to notify when new pending transactions arrive.
	// To check if the state has changed, call platon_getFilterChanges.
	NewPendingTransactionFilter() (*big.Int, error)
	// Uninstalls a filter with given id.
	// Should always be called when watch is no longer needed.
	//Additionally Filters timeout when they aren't requested with platon_getFilterChanges for a period of time.
	UninstallFilter(filterId *big.Int) bool
	// Polling method for a filter, which returns an array of logs which occurred since last poll.
	GetFilterChanges(filterId *big.Int) ([]types.Log, error)
	// Returns an array of all logs matching filter with given id.
	GetFilterLogs(filterId *big.Int) ([]types.Log, error)
	// Returns an array of all logs matching a given filter object.
	GetLogs(req platon.FilterQuery) ([]types.Log, error)
}

func New(url string) (Geb3, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	platon := &PlatonRPC{client: client}
	return platon, nil
}
