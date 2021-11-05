package ppos

import (
	"encoding/json"
	"fmt"
	"github.com/oldmanfan/platon-go-sdk/network"
	"github.com/oldmanfan/platon-go-sdk/ppos/typedefs"
	"math/big"
	"testing"
)

func TestSlashContract_CheckDoubleSign(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewSlashContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
	blockNumber := big.NewInt(583848)
	list, err := sc.CheckDoubleSign(typedefs.PREPARE_BLOCK, nodeId, blockNumber)
	if err != nil {
		t.Errorf("SlashContract.CheckDoubleSign failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	// Expect Output:
	//820bb9
	//01
	//b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050
	//8308e8a8
	//Function Data is: f84e83820bb901b842b84077fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050848308e8a8
	fmt.Println(string(result))
}

func TestSlashContract_ReportDoubleSign(t *testing.T) {
	config := network.PposMainNetParams
	sc := NewSlashContract(config, credentials)

	evidence := "{\"prepareA\":{\"epoch\":0,\"viewNumber\":0,\"blockHash\":\"0x47c0b9cc1e544e866ed3afb1a2fec5f8c0c6d97a04680f56f26b238b362482ca\",\"blockNumber\":583848,\"blockIndex\":0,\"blockData\":\"0x005e8ae4a78cd34d2c9fb08abda0e39d781e4abc58ea7f9b03c56f6a8e804027\",\"validateNode\":{\"index\":0,\"address\":\"0x0550184a50db8162c0cfe9296f06b2b1db019331\",\"nodeId\":\"77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050\",\"blsPubKey\":\"5ccd6b8c32f2713faa6c9a46e5fb61ad7b7400e53fabcbc56bdc0c16fbfffe09ad6256982c7059e7383a9187ad93a002a7cda7a75d569f591730481a8b91b5fad52ac26ac495522a069686df1061fc184c31771008c1fedfafd50ae794778811\"},\"signature\":\"0x974d787c28b7fb2ec67decdb750e4e29ace69e07ab3d1864c4fba9b7eb780868fb36966183ac6f156b99b0f1d8034d8500000000000000000000000000000000\"},\"prepareB\":{\"epoch\":0,\"viewNumber\":0,\"blockHash\":\"0x4bf9291e34fb7ae3f93eb4bb77a4b41251dce247e4a9b8e120dbf69a310f87bb\",\"blockNumber\":583848,\"blockIndex\":0,\"blockData\":\"0x4b6fd7afae0fd5ab3b2457720297165539623562eff16827981904357720b995\",\"validateNode\":{\"index\":0,\"address\":\"0x0550184a50db8162c0cfe9296f06b2b1db019331\",\"nodeId\":\"77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050\",\"blsPubKey\":\"5ccd6b8c32f2713faa6c9a46e5fb61ad7b7400e53fabcbc56bdc0c16fbfffe09ad6256982c7059e7383a9187ad93a002a7cda7a75d569f591730481a8b91b5fad52ac26ac495522a069686df1061fc184c31771008c1fedfafd50ae794778811\"},\"signature\":\"0x615967a22bb06245a76c94ce8914c95bdc9eda8fbfaa97b85d8472c7c7bb10bc6aff325798965e583130690ac7ef568c00000000000000000000000000000000\"}}"
	list, err := sc.ReportDoubleSign(typedefs.PREPARE_BLOCK, evidence)
	if err != nil {
		t.Errorf("SlashContract.ReportDoubleSign failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}
