package ppos

import (
	"encoding/json"
	"fmt"
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
	"testing"
)

func TestProposalContract_DeclareVersion(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)
	//{"Version":4096,"Sign":"0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00"}
	version := common.ProgramVersion{
		Version: big.NewInt(4096),
		Sign:    "0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00",
	}

	verifier := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	list, err := pc.DeclareVersion(version, verifier)
	if err != nil {
		t.Errorf("Get Verifier list failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}
