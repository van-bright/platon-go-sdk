package web3g

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TResponse struct {
	ID      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type ExecFunc func(geb3 *Web3g)

func DoBridgeHttpTest(respFunc http.HandlerFunc, execFunc ExecFunc) {
	ts := httptest.NewServer(respFunc)
	defer ts.Close()

	api := ts.URL
	geb3 := New(api)

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

	var execFunc ExecFunc = func(geb3 *Web3g) {
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

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.AdminDataDir()

		if resp != expect {
			t.Errorf("web3g admin add peer failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminNodeInfo(t *testing.T) {
	var expect = NodeInfo{
		ID:     "139973edac2fb8ba6db3ea72b619efc925dda27adc306bca6b04afa04aa41453a178ce1cebfc3deb06e9ba27bc7dda2d0a9d09d92d41efc2ba0f7b550b6dfca7",
		Name:   "PlatONnetwork/platon/v0.15.0-unstable-9867ee68/linux-amd64/go1.13.4",
		BlsPub: "4844d1401953c43dd9718050e722ed02a832f6bd5b0db72e74cc14dba22091af066e0835810db07dac5e386f8c38a20866f8458490d6202d244dbc9b6b0928aa24115a10d115739542902a761b7ee0efbfc72121828f071ae680e49562f60515",
		Enode:  "\"enode://139973edac2fb8ba6db3ea72b619efc925dda27adc306bca6b04afa04aa41453a178ce1cebfc3deb06e9ba27bc7dda2d0a9d09d92d41efc2ba0f7b550b6dfca7@[::]:16789?discport=0",
		IP:     "::",
		Ports: struct {
			Discovery int `json:"discovery"`
			Listener  int `json:"listener"`
		}{0, 16789},
		ListenAddr: "[::]:16789",
		Protocols:  nil,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.AdminNodeInfo()

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
	expect := ProgramVersion{
		Sign:    "0x34b844edd4cdede3a56309148c2b4dfaf4addd787d58c78fca59c29ccfd67baa4f454175e4cd5339a62835490990476f82b17c6ef1edd166ca0d2617e38809b900",
		Version: 3840,
	}

	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.AdminGetProgramVersion()

		if resp.Sign != expect.Sign {
			t.Errorf("ProgramVersion.Sign failed.")
		}

		if resp.Version != expect.Version {
			t.Errorf("ProgramVersion.Version failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminGetSchnorrNIZKProve(t *testing.T) {
	expect := "NizkProve"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.AdminGetSchnorrNIZKProve()

		if resp != expect {
			t.Errorf("SchnorrNIZKProve failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_AdminPeers(t *testing.T) {
	expect := "[]"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.AdminPeers()

		if resp != expect {
			t.Errorf("AdminPeers failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}
