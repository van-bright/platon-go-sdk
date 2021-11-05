package web3go

import (
	"bytes"
	"encoding/json"
	platon_go_sdk "github.com/oldmanfan/platon-go-sdk"
	"github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/common/hexutil"
	"github.com/oldmanfan/platon-go-sdk/core/types"
	"github.com/oldmanfan/platon-go-sdk/params"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TResponse struct {
	ID      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type ExecFunc func(geb3 Geb3)

func DoBridgeHttpTest(respFunc http.HandlerFunc, execFunc ExecFunc) {
	ts := httptest.NewServer(respFunc)
	defer ts.Close()

	api := ts.URL
	geb3, _ := New(api)

	execFunc(geb3)
}

func GenRespFunction(expect interface{}) func(w http.ResponseWriter, r *http.Request) {
	var respFunc = func(w http.ResponseWriter, r *http.Request) {
		var rsp = TResponse{74, "2.0", expect}
		rspBytes, _ := json.Marshal(rsp)
		w.WriteHeader(http.StatusOK)
		w.Write(rspBytes)
	}

	return respFunc
}

func TestWeb3g_AdminAddPeer(t *testing.T) {
	var respFunc = GenRespFunction(true)

	var execFunc ExecFunc = func(geb3 Geb3) {
		params := "\"enode: //acb2281452fb9fc25d40113fb6afe82b498361de0ee4ce69f55c180bb2afce2c5a00f97bfbe0270fadba174264cdf6da76ba334a6380c0005a84e8a6449c2502@127.0.0.1: 14789\""
		resp, _ := geb3.AdminAddPeer(params)

		if !resp {
			t.Errorf("web3g admin add peer failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminDataDir(t *testing.T) {
	expect := "/home/platon/network/data"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.AdminDataDir()

		if resp != expect {
			t.Errorf("web3g admin add peer failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminNodeInfo(t *testing.T) {
	var expect = platon_go_sdk.NodeInfo{
		ID:     "139973edac2fb8ba6db3ea72b619efc925dda27adc306bca6b04afa04aa41453a178ce1cebfc3deb06e9ba27bc7dda2d0a9d09d92d41efc2ba0f7b550b6dfca7",
		Name:   "PlatONnetwork/platon/v0.15.0-unstable-9867ee68/linux-amd64/go1.13.4",
		BlsPub: "4844d1401953c43dd9718050e722ed02a832f6bd5b0db72e74cc14dba22091af066e0835810db07dac5e386f8c38a20866f8458490d6202d244dbc9b6b0928aa24115a10d115739542902a761b7ee0efbfc72121828f071ae680e49562f60515",
		ENode:  "\"enode://139973edac2fb8ba6db3ea72b619efc925dda27adc306bca6b04afa04aa41453a178ce1cebfc3deb06e9ba27bc7dda2d0a9d09d92d41efc2ba0f7b550b6dfca7@[::]:16789?discport=0",
		IP:     "::",
		Ports: struct {
			Discovery int `json:"discovery"`
			Listener  int `json:"listener"`
		}{0, 16789},
		ListenAddr: "[::]:16789",
		Protocols:  nil,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		rawResp, _ := geb3.AdminNodeInfo()

		resp := platon_go_sdk.NodeInfo{}
		json.Unmarshal([]byte(rawResp), &resp)
		if resp.ID != expect.ID {
			t.Errorf("nodeInfo.ID failed.")
		}

		if resp.Name != expect.Name {
			t.Errorf("nodeInfo.Name failed.")
		}

		if resp.BlsPub != expect.BlsPub {
			t.Errorf("nodeInfo.BlsPub failed.")
		}

		if resp.Ports.Listener != expect.Ports.Listener {
			t.Errorf("nodeInfo.Ports.Listener failed.")
		}

		if resp.ListenAddr != expect.ListenAddr {
			t.Errorf("nodeInfo.ListenAddr failed.")
		}

		if resp.IP != expect.IP {
			t.Errorf("nodeInfo.IP failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminGetProgramVersion(t *testing.T) {
	expect := params.ProgramVersion{
		Sign:    "0x34b844edd4cdede3a56309148c2b4dfaf4addd787d58c78fca59c29ccfd67baa4f454175e4cd5339a62835490990476f82b17c6ef1edd166ca0d2617e38809b900",
		Version: 3840,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.AdminGetProgramVersion()

		if resp.Sign != expect.Sign {
			t.Errorf("AdminGetProgramVersion.Sign failed.")
		}

		if resp.Version != expect.Version {
			t.Errorf("AdminGetProgramVersion.Version failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_GetSchnorrNIZKProve(t *testing.T) {
	expect := "NizkProve"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.GetSchnorrNIZKProve()

		if resp != expect {
			t.Errorf("SchnorrNIZKProve failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminPeers(t *testing.T) {
	expect := []string{""}
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.AdminPeers()

		if len(resp) != len(expect) {
			t.Errorf("AdminPeers failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NetListening(t *testing.T) {
	expect := true
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.NetListening()

		if resp != expect {
			t.Errorf("NetListening failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NetVersion(t *testing.T) {
	expect := "201030" //big.NewInt(201030)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.NetworkID()
		ret := new(big.Int)
		ret.SetString(expect, 10)
		if resp.Cmp(ret) != 0 {
			t.Errorf("NetVersion failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NetPeerCount(t *testing.T) {
	expect := hexutil.Uint64(49)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.NetPeerCount()

		if resp != uint64(expect) {
			t.Errorf("NetPeerCount failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_Accounts(t *testing.T) {
	expect := []string{"atp123456", "atp891011"}
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.Accounts()

		if len(resp) != len(expect) {
			t.Errorf("Accounts failed.")
		}

		for i := 0; i < len(expect); i++ {
			if expect[i] != resp[i] {
				t.Errorf("Accounts item failed.")
			}
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_BlockNumber(t *testing.T) {
	expect := hexutil.Uint64(49)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.BlockNumber()

		if resp != (uint64)(expect) {
			t.Errorf("BlockNumber failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_CallContract(t *testing.T) {
	expect := []byte("success")
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		gas := uint64(1000)
		to := common.MustBech32ToAddress("lat1506p5uyejv2eq9xa35vg6g4puam6xhwf073xcp")
		var req = platon_go_sdk.CallMsg{
			From:     common.MustBech32ToAddress("lat1t3jsgu5km95aeqfqxx396k46e2ejxcg442ltq0"),
			To:       &to,
			Gas:      gas,
			GasPrice: big.NewInt(1000000000),
			Value:    big.NewInt(9527),
			Data:     []byte(""),
		}
		resp, _ := geb3.CallContract(req, "latest")

		if bytes.Compare(expect, resp) == 0 {
			t.Errorf("CallContract failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_EstimateGas(t *testing.T) {
	expect := hexutil.Uint64(49)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		gas := uint64(1000)
		var req = platon_go_sdk.CallMsg2{
			From:     "atp1",
			To:       "atp2",
			Gas:      gas,
			GasPrice: big.NewInt(1000000000),
			Value:    big.NewInt(9527),
			Data:     []byte(""),
		}
		resp, _ := geb3.EstimateGasLimit(req)

		if resp != uint64(expect) {
			t.Errorf("EstimateGasLimit failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonEvidences(t *testing.T) {
	expect := "evidence string"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.Evidences()

		if resp != expect {
			t.Errorf("Evidences failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGasPrice(t *testing.T) {
	expect := (*hexutil.Big)(big.NewInt(100))
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.GasPrice()

		if resp.Cmp(expect.ToInt()) != 0 {
			t.Errorf("GasPrice failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonGetBalance(t *testing.T) {
	expect := (*hexutil.Big)(big.NewInt(9527))
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.BalanceAt("atp10098765", "latest")

		if resp.Cmp(expect.ToInt()) != 0 {
			t.Errorf("BalanceAt failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

//func TestWeb3g_PlatonGetBlockByHash(t *testing.T) {
//	expect := "123"
//
//	var respFunc = GenRespFunction(expect)
//
//	var execFunc ExecFunc = func(geb3 Geb3) {
//		resp, _ := geb3.BlockByHash("0x000000000001")
//
//		if resp != resp {
//			t.Errorf("BlockByHash failed.")
//		}
//	}
//
//	DoBridgeHttpTest(respFunc, execFunc)
//}
//
//func TestWeb3g_PlatonGetBlockByNumber(t *testing.T) {
//	expect := ""
//
//	var respFunc = GenRespFunction(expect)
//
//	var execFunc ExecFunc = func(geb3 Geb3) {
//		resp, _ := geb3.BlockByNumber("latest")
//
//		if resp != expect {
//			t.Errorf("BlockByNumber failed.")
//		}
//	}
//
//	DoBridgeHttpTest(respFunc, execFunc)
//}

func TestWeb3g_TransactionCountByHash(t *testing.T) {
	expect := hexutil.Uint(100)

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.TransactionCountByHash(common.HexToHash("0x000000000001"))

		if resp != uint(expect) {
			t.Errorf("TransactionCountByHash failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_TransactionCountByNumber(t *testing.T) {
	expect := hexutil.Uint(100)

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.TransactionCountByNumber("latest")

		if resp != uint(expect) {
			t.Errorf("TransactionCountByNumber failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_CodeAt(t *testing.T) {
	expect := hexutil.Bytes{1, 2, 3, 4}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.CodeAt("lat1t3jsgu5km95aeqfqxx396k46e2ejxcg442ltq0", "latest")

		if len(resp) != len(expect) {
			t.Errorf("CodeAt failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_GetFilterChanges(t *testing.T) {
	expect := []types.Log{types.Log{}}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {

		resp, _ := geb3.GetFilterChanges(big.NewInt(10))

		if len(resp) != len(expect) {
			t.Errorf("GetFilterChanges failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_GetFilterLogs(t *testing.T) {
	expect := []types.Log{types.Log{}}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {

		resp, _ := geb3.GetFilterLogs(big.NewInt(10))

		if len(resp) != len(expect) {
			t.Errorf("GetFilterLogs failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_GetLogs(t *testing.T) {
	expect := []types.Log{types.Log{}}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 Geb3) {
		req := platon_go_sdk.FilterQuery{
			BlockHash: nil,
			FromBlock: nil,
			ToBlock:   nil,
			Addresses: nil,
			Topics:    nil,
		}
		resp, _ := geb3.GetLogs(req)

		if len(resp) != len(expect) {
			t.Errorf("GetLogs failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_StorageAt(t *testing.T) {
	expect := hexutil.Bytes{1, 2, 3, 4}
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.StorageAt("lat1t3jsgu5km95aeqfqxx396k46e2ejxcg442ltq0", common.HexToHash("0x00000000000"), "latest")

		if len(resp) != len(expect) {
			t.Errorf("StorageAt failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_TransactionByBlockHashAndIndex(t *testing.T) {
	expect := types.Transaction{}

	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {

		resp, _ := geb3.TransactionByBlockHashAndIndex(common.HexToHash("0x00000000000000000000000000000000"), 0)

		if resp.Hash() != expect.Hash() {
			t.Errorf("TransactionByBlockHashAndIndex failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_TransactionByBlockNumberAndIndex(t *testing.T) {
	expect := types.Transaction{}

	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {

		resp, _ := geb3.TransactionByBlockNumberAndIndex("latest", 0)

		if resp.Hash() != expect.Hash() {
			t.Errorf("TransactionByBlockNumberAndIndex failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

//func TestWeb3g_TransactionByHash(t *testing.T) {
//	expect := types.Transaction{}
//
//	var respFunc = GenRespFunction(expect)
//	var execFunc ExecFunc = func(geb3 Geb3) {
//
//		resp, _, _ := geb3.TransactionByHash(common.HexToHash("0x000000000000000000000"))
//
//		if resp.Hash() != expect.Hash() {
//			t.Errorf("TransactionByHash failed.")
//		}
//	}
//
//	DoBridgeHttpTest(respFunc, execFunc)
//}

func TestWeb3g_TransactionReceipt(t *testing.T) {
	expect := types.Receipt{}
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.TransactionReceipt(common.HexToHash("0x00000000000000000000"))

		if resp.TxHash != expect.TxHash {
			t.Errorf("TransactionHash failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NewBlockFilter(t *testing.T) {
	expect := (*hexutil.Big)(big.NewInt(100))
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.NewBlockFilter()

		if resp.Cmp(expect.ToInt()) != 0 {
			t.Errorf("NewBlockFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NewFilter(t *testing.T) {
	expect := (*hexutil.Big)(big.NewInt(100))
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		q := platon_go_sdk.FilterQuery{
			BlockHash: nil,
			FromBlock: nil,
			ToBlock:   nil,
			Addresses: nil,
			Topics:    nil,
		}
		resp, _ := geb3.NewFilter(q)

		if resp.Cmp(expect.ToInt()) != 0 {
			t.Errorf("NewFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NewPendingTransactionFilter(t *testing.T) {
	expect := (*hexutil.Big)(big.NewInt(100))
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.NewPendingTransactionFilter()

		if resp.Cmp(expect.ToInt()) != 0 {
			t.Errorf("NewPendingTransactionFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_ProtocolVersion(t *testing.T) {
	expect := hexutil.Uint64(1)
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.ProtocolVersion()

		if resp != uint64(expect) {
			t.Errorf("ProtocolVersion failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_SendRawTransaction(t *testing.T) {
	expect := "send raw transaction success"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		req := types.Transaction{}
		_, err := geb3.SendRawTransaction(&req)

		if err != nil {
			t.Errorf("SendRawTransaction failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_SendTransaction(t *testing.T) {
	expect := "platon send transaction"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		req := types.Transaction{}
		_, err := geb3.SendTransaction(&req)

		if err != nil {
			t.Errorf("PlatonSendTransaction failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_PlatonSign(t *testing.T) {
	expect := "0x2ac19db245478a06032e69cdbd2b54e648b78431d0a47bd1fbab18f79f820ba407466e37adbe9e84541cab97ab7d290f4a64a5825c876d22109f3bf813254e8601"
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		req := platon_go_sdk.SignReq{
			Signer: "atx16xk7yhxd842s5l44x2k8t89v00sfcfcewjsd7z",
			Data:   "SchoolBus",
		}
		resp, _ := geb3.Sign(&req)

		if resp != expect {
			t.Errorf("PlatonSign failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_Syncing(t *testing.T) {
	var expect = platon_go_sdk.SyncProgress{StartingBlock: 0, CurrentBlock: 10, HighestBlock: 10, PulledStates: 20, KnownStates: 30}
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp, _ := geb3.Syncing()

		if resp.HighestBlock != 10 {
			t.Errorf("Syncing failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_UninstallFilter(t *testing.T) {
	expect := true
	var respFunc = GenRespFunction(expect)
	var execFunc ExecFunc = func(geb3 Geb3) {
		resp := geb3.UninstallFilter(big.NewInt(100))

		if resp != expect {
			t.Errorf("UninstallFilter failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}
