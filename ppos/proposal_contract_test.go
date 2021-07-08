package ppos

import (
	"encoding/json"
	"fmt"
	"math/big"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/resp"
	"platon-go-sdk/ppos/typedefs"
	"testing"
)

func TestProposalContract_DeclareVersion(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)
	//{"Version":4096,"Sign":"0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00"}
	version := typedefs.ProgramVersion{
		Version: big.NewInt(4096),
		Sign:    "0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00",
	}

	verifier := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	list, err := pc.DeclareVersion(version, verifier)
	if err != nil {
		t.Errorf("ProposalContract.DeclareVersion failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//8207d4
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//821000
	//b8410dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00
	//Function Data is: f891838207d4b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c8205549905083821000b843b8410dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00

	fmt.Println(string(result))
}

func TestProposalContract_GetAccuVerifiersCount(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	proposalId := "0x1234"
	blockHash := "0x5678"
	list, err := pc.GetAccuVerifiersCount(proposalId, blockHash)
	if err != nil {
		t.Errorf("ProposalContract.GetAccuVerifiersCount failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//820839
	//821234
	//825678
	//Function Data is: cc838208398382123483825678
	fmt.Println(string(result))
}

func TestProposalContract_GetActiveVersion(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	list, err := pc.GetActiveVersion()
	if err != nil {
		t.Errorf("ProposalContract.GetActiveVersion failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//820837
	//Function Data is: c483820837
	fmt.Println(string(result))
}

func TestProposalContract_GetGovernParamValue(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	module := "staking"
	name := "operatingThreshold"
	list, err := pc.GetGovernParamValue(module, name)
	if err != nil {
		t.Errorf("ProposalContract.GetGovernParamValue failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//820838
	//877374616b696e67
	//926f7065726174696e675468726573686f6c64
	//Function Data is: e18382083888877374616b696e6793926f7065726174696e675468726573686f6c64

	fmt.Println(string(result))
}

func TestProposalContract_GetParamList(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	module := "staking"
	list, err := pc.GetParamList(module)
	if err != nil {
		t.Errorf("ProposalContract.GetParamList failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//82083a
	//877374616b696e67
	//Function Data is: cd8382083a88877374616b696e67

	fmt.Println(string(result))
}

func TestProposalContract_GetProposal(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	proposalId := "0x261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed"
	list, err := pc.GetProposal(proposalId)
	if err != nil {
		t.Errorf("ProposalContract.GetProposal failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//820834
	//a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed
	//Function Data is: e683820834a1a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed

	fmt.Println(string(result))
}

func TestProposalContract_GetProposalList(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	list, err := pc.GetProposalList()
	if err != nil {
		t.Errorf("ProposalContract.GetProposalList failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//820836
	//Function Data is: c483820836

	fmt.Println(string(result))
}

func TestProposalContract_GetTallyResult(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	proposalId := "0x261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed"
	list, err := pc.GetTallyResult(proposalId)
	if err != nil {
		t.Errorf("ProposalContract.GetTallyResult failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//820835
	//a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed
	//Function Data is: e683820835a1a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed

	fmt.Println(string(result))
}

func TestProposalContract_Vote(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	proposalId := "0x261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed"
	pv := typedefs.ProgramVersion{
		Version: big.NewInt(4096),
		Sign:    "0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00",
	}
	voption := typedefs.YEAS
	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"

	list, err := pc.Vote(pv, voption, proposalId, nodeId)
	if err != nil {
		t.Errorf("ProposalContract.Vote failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//8207d3
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed
	//01
	//821000
	//b8410dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00
	//Function Data is: f8b4838207d3b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050a1a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed0183821000b843b8410dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00

	fmt.Println(string(result))
}

func TestProposalContract_SubmitTextProposal(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	proposal := resp.CreateTextProposal(nodeId, "1000")

	list, err := pc.SubmitProposal(proposal)
	if err != nil {
		t.Errorf("ProposalContract.SubmitTextProposal failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//8207d0
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//8431303030
	//Function Data is: f84e838207d0b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050858431303030

	fmt.Println(string(result))
}

func TestProposalContract_SubmitVersionProposal(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	proposal := resp.CreateVersionProposal(nodeId, "1000", big.NewInt(100), big.NewInt(200))

	list, err := pc.SubmitProposal(proposal)
	if err != nil {
		t.Errorf("ProposalContract.SubmitVersionProposal failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//8207d1
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//8431303030
	//64
	//81c8
	//Function Data is: f852838207d1b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050858431303030648281c8

	fmt.Println(string(result))
}

func TestProposalContract_SubmitCancelProposal(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	toBeCancelProposal := "0x261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed"
	proposal := resp.CreateCancelProposal(nodeId, "1000", big.NewInt(100), toBeCancelProposal)

	list, err := pc.SubmitProposal(proposal)
	if err != nil {
		t.Errorf("ProposalContract.SubmitCancelProposal failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//8207d5
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//8431303030
	//64
	//a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed
	//Function Data is: f871838207d5b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c8205549905085843130303064a1a0261cf6c0f518aeddffb2aa5536685af6f13f8ba763c77b42f12ce025ef7170ed

	fmt.Println(string(result))
}

func TestProposalContract_SubmitParamProposal(t *testing.T) {
	config := network.PposMainNetParams
	pc := NewProposalContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	proposal := resp.CreateParamProposal(nodeId, "1000", "module", "name", "newValue")

	list, err := pc.SubmitProposal(proposal)
	if err != nil {
		t.Errorf("ProposalContract.SubmitParamProposal failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect
	//8207d2
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//8431303030
	//866d6f64756c65
	//846e616d65
	//886e657756616c7565
	//Function Data is: f866838207d2b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c8205549905085843130303087866d6f64756c6585846e616d6589886e657756616c7565

	fmt.Println(string(result))
}
