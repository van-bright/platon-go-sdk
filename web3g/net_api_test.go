package web3g

import "testing"

func TestWeb3g_NetListening(t *testing.T) {
	expect := true
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.NetListening()

		if resp != expect {
			t.Errorf("NetListening failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NetVersion(t *testing.T) {
	expect := "net-version-0.1"
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.NetVersion()

		if resp != expect {
			t.Errorf("NetVersion failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}

func TestWeb3g_NetPeerCount(t *testing.T) {
	expect := uint32(49)
	var respFunc = GenRespFunction(expect)

	var execFunc ExecFunc = func(geb3 *Web3g) {
		resp, _ := geb3.NetPeerCount()

		if resp != expect {
			t.Errorf("NetPeerCount failed.")
		}
	}

	DoBridgeHttpTest(respFunc, execFunc)
}
