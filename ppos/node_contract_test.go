package ppos

import (
	"encoding/json"
	"fmt"
	"platon-go-sdk/network"
	"platon-go-sdk/ppos/common"
	"testing"
)

const PrivateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"

var credentials, _ = common.NewCredential(PrivateKey, "atp")

func TestNodeContract_GetCandidateList(t *testing.T) {
	config := network.PposMainNetParams
	nc := NewNodeContract(config, credentials)

	list, err := nc.GetCandidateList()
	if err != nil {
		t.Errorf("Get Candidate list failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}

func TestNodeContract_GetValidatorList(t *testing.T) {
	config := network.PposMainNetParams
	nc := NewNodeContract(config, credentials)

	list, err := nc.GetValidatorList()
	if err != nil {
		t.Errorf("Get Validator list failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}

func TestNodeContract_GetVerifierList(t *testing.T) {
	config := network.PposMainNetParams
	nc := NewNodeContract(config, credentials)

	list, err := nc.GetVerifierList()
	if err != nil {
		t.Errorf("Get Verifier list failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}
