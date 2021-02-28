package web3g

import (
	"math/big"
	"testing"
)

func TestWeb3g_PlatonAccounts(t *testing.T) {
	expect := []string{"atp123456", "atp891011"}
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonAccounts()

		if len(resp) != len(expect) {
			t.Errorf("PlatonAccounts failed.")
		}

		for i := 0; i < len(expect); i++ {
			if expect[i] != resp[i] {
				t.Errorf("PlatonAccounts item failed.")
			}
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonBlockNumber(t *testing.T) {
	expect := uint64(49)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonBlockNumber()

		if resp != expect {
			t.Errorf("PlatonBlockNumber failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonCall(t *testing.T) {
	expect := "success"
	var respFunc = GenRespFunction(expect)
	gas := big.NewInt(1000)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		var req = PlatonCallReq{
			From:        "atp1",
			To:          "atp2",
			Gas:         gas,
			GasPrice:    big.NewInt(1000000000),
			Value:       big.NewInt(9527),
			Data:        "",
			TagOrNumber: nil,
		}
		resp, _ := geb3.PlatonCall(&req)

		if resp != expect {
			t.Errorf("PlatonCall failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonEstimateGas(t *testing.T) {
	expect := uint64(49)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonEstimateGasReq{
			From:        "atp1",
			To:          "atp2",
			Gas:         big.NewInt(8888),
			GasPrice:    big.NewInt(1000000000),
			Value:       big.NewInt(9527),
			Data:        "",
			TagOrNumber: nil,
		}
		resp, _ := geb3.PlatonEstimateGas(&req)

		if resp != expect {
			t.Errorf("PlatonEstimateGas failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonEvidences(t *testing.T) {
	expect := "evidence string"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonEvidences()

		if resp != expect {
			t.Errorf("PlatonEvidences failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGasPrice(t *testing.T) {
	expect := big.NewInt(10086)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGasPrice()

		if resp.Cmp(expect) != 0 {
			t.Errorf("PlatonGasPrice failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetBalance(t *testing.T) {
	expect := big.NewInt(9527)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGetBalance("atp10098765", "latest")

		if resp.Cmp(expect) != 0 {
			t.Errorf("PlatonGetBalance failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetBlockByHash(t *testing.T) {
	expect := PlatonBlock{
		Number:           nil,
		Hash:             "",
		ParentHash:       "",
		Nonce:            "",
		Sha3Uncles:       "",
		LogsBloom:        "",
		TransactionsRoot: "",
		StateRoot:        "",
		ReceiptsRoot:     "",
		Miner:            "",
		Difficulty:       nil,
		TotalDifficulty:  nil,
		ExtraData:        "",
		Size:             nil,
		GasLimit:         nil,
		GasUsed:          nil,
		Timestamp:        nil,
		Transactions:     nil,
		Uncles:           nil,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGetBlockByHash("0x000000000001", false)

		if resp.Hash != expect.Hash {
			t.Errorf("PlatonGetBlockByHash failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetBlockByNumber(t *testing.T) {
	expect := PlatonBlock{
		Number:           nil,
		Hash:             "",
		ParentHash:       "",
		Nonce:            "",
		Sha3Uncles:       "",
		LogsBloom:        "",
		TransactionsRoot: "",
		StateRoot:        "",
		ReceiptsRoot:     "",
		Miner:            "",
		Difficulty:       nil,
		TotalDifficulty:  nil,
		ExtraData:        "",
		Size:             nil,
		GasLimit:         nil,
		GasUsed:          nil,
		Timestamp:        nil,
		Transactions:     nil,
		Uncles:           nil,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGetBlockByNumber("latest", false)

		if resp.Hash != expect.Hash {
			t.Errorf("PlatonGetBlockByHash failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetBlockTransactionCountByHash(t *testing.T) {
	expect := uint64(100)

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGetBlockTransactionCountByHash("0x000000000001")

		if resp != expect {
			t.Errorf("PlatonGetBlockTransactionCountByHash failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetBlockTransactionCountByNumber(t *testing.T) {
	expect := uint64(100)

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGetBlockTransactionCountByNumber("latest")

		if resp != expect {
			t.Errorf("PlatonGetBlockTransactionCountByNumber failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetCode(t *testing.T) {
	expect := []byte{1, 2, 3, 4}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonGetCodeReq{
			Address:     "atp000000000000",
			TagOrNumber: "latest",
		}
		resp, _ := geb3.PlatonGetCode(&req)

		if len(resp) != len(expect) {
			t.Errorf("PlatonGetCode failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetFilterChanges(t *testing.T) {
	expect := PlatonTransactionLog{
		Hashes:           nil,
		Type:             "",
		LogIndex:         nil,
		TransactionIndex: nil,
		TransactionHash:  "",
		BlockHash:        "",
		BlockNumber:      nil,
		Address:          "",
		Data:             "",
		Topics:           nil,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {

		resp, _ := geb3.PlatonGetFilterChanges(big.NewInt(10))

		if len(resp.Type) != len(expect.Type) {
			t.Errorf("PlatonGetFilterChanges failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetFilterLogs(t *testing.T) {
	expect := PlatonTransactionLog{
		Hashes:           nil,
		Type:             "",
		LogIndex:         nil,
		TransactionIndex: nil,
		TransactionHash:  "",
		BlockHash:        "",
		BlockNumber:      nil,
		Address:          "",
		Data:             "",
		Topics:           nil,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {

		resp, _ := geb3.PlatonGetFilterLogs(big.NewInt(10))

		if len(resp.Hashes) != len(expect.Hashes) {
			t.Errorf("PlatonGetFilterLogs failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetLogs(t *testing.T) {
	expect := PlatonTransactionLog{
		Hashes:           []string{"0x0001", "0x0002"},
		Type:             "",
		LogIndex:         big.NewInt(0),
		TransactionIndex: big.NewInt(0),
		TransactionHash:  "",
		BlockHash:        "",
		BlockNumber:      big.NewInt(0),
		Address:          "",
		Data:             "",
		Topics:           []string{"0xabc"},
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonGetLogsReq{Topics: []interface{}{"abc", "def"}}
		resp, _ := geb3.PlatonGetLogs(&req)

		if len(resp.Hashes) != len(expect.Hashes) {
			t.Errorf("PlatonGetLogs failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetStorageAt(t *testing.T) {
	expect := []byte{1, 2, 3, 4}
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonGetStorageAtReq{
			Address:       "atp00000000",
			PositionIndex: big.NewInt(10),
			TagOrNumber:   "latest",
		}
		resp, _ := geb3.PlatonGetStorageAt(&req)

		if len(resp) != len(expect) {
			t.Errorf("PlatonGetStorageAt failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetTransactionByBlockHashAndIndex(t *testing.T) {
	expect := PlatonTransaction{
		Hash:             "",
		Nonce:            nil,
		BlockHash:        "",
		BlockNumber:      nil,
		TransactionIndex: nil,
		From:             "",
		To:               "",
		Value:            nil,
		Gas:              nil,
		GasPrice:         nil,
		Input:            "",
	}

	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {

		resp, _ := geb3.PlatonGetTransactionByBlockHashAndIndex("0x00000000000000000000000000000000", "latest")

		if resp.Hash != expect.Hash {
			t.Errorf("PlatonGetTransactionByBlockHashAndIndex failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetTransactionByBlockNumberAndIndex(t *testing.T) {
	expect := PlatonTransaction{
		Hash:             "",
		Nonce:            nil,
		BlockHash:        "",
		BlockNumber:      nil,
		TransactionIndex: nil,
		From:             "",
		To:               "",
		Value:            nil,
		Gas:              nil,
		GasPrice:         nil,
		Input:            "",
	}

	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {

		resp, _ := geb3.PlatonGetTransactionByBlockNumberAndIndex("latest", "0")

		if resp.Hash != expect.Hash {
			t.Errorf("PlatonGetTransactionByBlockHashAndIndex failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetTransactionByHash(t *testing.T) {
	expect := PlatonTransaction{
		Hash:             "",
		Nonce:            nil,
		BlockHash:        "",
		BlockNumber:      nil,
		TransactionIndex: nil,
		From:             "",
		To:               "",
		Value:            nil,
		Gas:              nil,
		GasPrice:         nil,
		Input:            "",
	}

	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {

		resp, _ := geb3.PlatonGetTransactionByHash("0x000000000000000000000")

		if resp.Hash != expect.Hash {
			t.Errorf("PlatonGetTransactionByHash failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetTransactionCount(t *testing.T) {
	expect := uint64(9527)
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonGetTransactionCountReq{
			Address:     "atp0000000000000000000",
			TagOrNumber: "latest",
		}
		resp, _ := geb3.PlatonGetTransactionCount(&req)

		if resp != expect {
			t.Errorf("PlatonGetTransactionCount failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetTransactionReceipt(t *testing.T) {
	expect := PlatonTransactionReceipt{
		TransactionHash:   "",
		TransactionIndex:  nil,
		BlockHash:         "",
		BlockNumber:       nil,
		CumulativeGasUsed: nil,
		GasUsed:           nil,
		ContractAddress:   "",
		Logs:              nil,
	}
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonGetTransactionReceipt("0x00000000000000000000")

		if resp.TransactionHash != expect.TransactionHash {
			t.Errorf("PlatonGetTransactionReceipt failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonNewBlockFilter(t *testing.T) {
	expect := big.NewInt(100)
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonNewBlockFilter()

		if resp.Cmp(expect) != 0 {
			t.Errorf("PlatonNewBlockFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonNewFilter(t *testing.T) {
	expect := big.NewInt(100)
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonNewFilter("12345", "latest", "atp00000000000000", "abcde")

		if resp.Cmp(expect) != 0 {
			t.Errorf("PlatonNewFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonNewPendingTransactionFilter(t *testing.T) {
	expect := big.NewInt(100)
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonNewPendingTransactionFilter()

		if resp.Cmp(expect) != 0 {
			t.Errorf("PlatonNewPendingTransactionFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonProtocolVersion(t *testing.T) {
	expect := "platon protocol version"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonProtocolVersion()

		if resp != expect {
			t.Errorf("PlatonProtocolVersion failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonSendRawTransaction(t *testing.T) {
	expect := "send raw transaction success"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonSendRawTransactionReq{
			string([]byte{1, 2, 3, 4, 5}),
		}
		resp, _ := geb3.PlatonSendRawTransaction(&req)

		if resp != expect {
			t.Errorf("PlatonSendRawTransaction failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonSendTransaction(t *testing.T) {
	expect := "platon send transaction"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonSendTransactionReq{
			From:     "",
			To:       "",
			Gas:      nil,
			GasPrice: nil,
			Value:    nil,
			Data:     "",
			Nonce:    nil,
		}
		resp, _ := geb3.PlatonSendTransaction(&req)

		if resp != expect {
			t.Errorf("PlatonSendTransaction failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonSign(t *testing.T) {
	expect := "0x2ac19db245478a06032e69cdbd2b54e648b78431d0a47bd1fbab18f79f820ba407466e37adbe9e84541cab97ab7d290f4a64a5825c876d22109f3bf813254e8601"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		req := PlatonSignReq{
			Address: "atx16xk7yhxd842s5l44x2k8t89v00sfcfcewjsd7z",
			Data:    "SchoolBus",
		}
		resp, _ := geb3.PlatonSign(&req)

		if resp != expect {
			t.Errorf("PlatonSign failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonSyncing(t *testing.T) {
	var expect = SyncingInfo{
		"", "", "",
	}
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonSyncing()

		if !resp.Syncing {
			t.Errorf("PlatonSyncing failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonUninstallFilter(t *testing.T) {
	expect := true
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.PlatonUninstallFilter(big.NewInt(100))

		if resp != expect {
			t.Errorf("PlatonUninstallFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}
