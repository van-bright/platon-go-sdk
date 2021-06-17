package web3go

import (
	"encoding/json"
	"fmt"
	"math/big"
	"platon-go-sdk/common"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/core/types"
	"platon-go-sdk/crypto"
	"platon-go-sdk/network"
	"testing"
	"time"
)

var (
	mnemonic   = "prevent scissors box assist enroll bean cup mushroom tragic steel best move"
	seed       = "0x9dfc7e3f52c4438d04db5488e13672faa37920ec62bacdc333a83974cb07bfdd893bfd46940dedfeb7ef30a142c4d07d552dd6589b40d3a58b941b7e9d6dae7e"
	privateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"
	atpAccount = "atp1v0jmfxmmq4mhv97rt5x8pwsfmd67594g5jrl72"
	atxAccount = "atx1v0jmfxmmq4mhv97rt5x8pwsfmd67594g75l4dq"
	path       = "m/44'/60'/0'/0/0"
)

func TestNewWalletByMnemonics(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)

	account := w.Accounts()[0]
	r, _ := w.hd.Address(account)
	fmt.Println("raw: ", r.Hex())
	mpk, _ := w.MainNetAddress(account)
	if mpk != atpAccount {
		t.Error("main net failed.")
	}
}

func TestNewWallet(t *testing.T) {
	w, _ := NewWallet()

	defaultAccount := w.Accounts()

	fmt.Println("new wallet: ", w.ExportHdWallet())
	if len(defaultAccount) != 1 {
		t.Error("new wallet failed.")
	}
}

func TestNewWalletBySeed(t *testing.T) {
	seedBytes, _ := hexutil.Decode(seed)
	w, _ := NewWalletBySeed(seedBytes)

	account := w.Accounts()[0]
	atp, _ := account.ToMainNetAddress()

	if atp != atpAccount {
		t.Error("wallet generate default account failed.")
	}
}

func TestAlayaWallet_Accounts(t *testing.T) {
	w, _ := NewWallet()
	w.NewAccount(1)
	w.ImportPrivateKey(crypto.HexMustToECDSA(privateKey), "./tmp", "123")

	accounts := w.Accounts()

	if len(accounts) != 3 {
		t.Error("list accounts of wallet failed.")
	}
}

func TestAlayaWallet_BalanceOf(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	w.SetNetworkCfg(&network.DefaultTestNetConfig)

	addr := common.MustBech32ToAddress("atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432")

	balance, _ := w.BalanceOf(addr)
	if balance.Cmp(big.NewInt(0)) == 0 {
		t.Error("read balance of atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432 failed.")
	}
}

func TestAlayaWallet_Export(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	info := w.ExportHdWallet()

	if len(info) == 0 {
		t.Error("export hd wallet failed.")
	}
}

func TestAlayaWallet_ExportMnemonic(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	info, _ := w.ExportMnemonic()
	if info != mnemonic {
		t.Error("export mnemonic failed.")
	}
}

func TestAlayaWallet_ExportPrivateKey(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)

	a, err := w.AccountByBech32("atp1v0jmfxmmq4mhv97rt5x8pwsfmd67594g5jrl72")
	if err != nil {
		t.Error("find account failed: ", err)
	}

	pk, err := w.ExportPrivateKey(a, "")
	if err != nil {
		t.Error("export private key failed: ", err)
	}

	pkStr := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	if pkStr != privateKey {
		t.Error("export private key failed")
	}
}

func TestAlayaWallet_ImportPrivateKey(t *testing.T) {
	w, _ := NewWallet()
	a, _ := w.ImportPrivateKey(crypto.HexMustToECDSA("ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"), "./", "")

	pk, _ := w.ExportPrivateKey(a, "")
	pkStr := hexutil.Encode(crypto.FromECDSA(pk))[2:]
	if pkStr != privateKey {
		t.Error("import private key failed.")
	}
}

func TestAlayaWallet_MainNetAddress(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	a := w.Accounts()[0]

	mainAddress, _ := a.ToMainNetAddress()
	if mainAddress != atpAccount {
		t.Error("export main net address failed.")
	}
}

func TestAlayaWallet_NewAccount(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	a, _ := w.NewAccount(1)

	aInfo := w.ToString(a)

	if len(aInfo) == 0 {
		t.Error("new account failed.")
	}
}

func TestAlayaWallet_TestNetAddress(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	a := w.Accounts()[0]

	testAddress, _ := a.ToTestNetAddress()
	if testAddress != atxAccount {
		t.Error("export main net address failed.")
	}
}

func TestAlayaWallet_Transfer(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	w.SetNetworkCfg(&network.DefaultTestNetConfig)

	fromPrivateKey := crypto.HexMustToECDSA("b72faaa798c44d2359d0ccb35dd39446c9c18905fdcecede42a6570ff177ae08")

	w.ImportPrivateKey(fromPrivateKey, "./", "")

	from := common.MustBech32ToAddress("atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432")
	to := common.MustBech32ToAddress("atp1v0jmfxmmq4mhv97rt5x8pwsfmd67594g5jrl72")

	balanceBefore, _ := w.BalanceOf(from)

	w.UnlockBech32("atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432", "")
	_, err := w.Transfer(from, to, big.NewInt(1000000))
	if err != nil {
		t.Error("transfer error: ", err)
	}

	time.Sleep(4 * time.Second)

	balanceAfter, _ := w.BalanceOf(from)

	diff := big.NewInt(0).Sub(balanceBefore, balanceAfter)
	// transfer value + transfer fee
	if diff.Cmp(big.NewInt(21000001000000)) != 0 {
		t.Error("transfer balance failed.")
	}
}

func TestAlayaWallet_SignTx(t *testing.T) {
	w, _ := NewWalletByMnemonics(mnemonic)
	w.NewAccount(1)
	fromAccount := w.Accounts()[0]
	toAccount := w.Accounts()[1]

	nonce := uint64(1)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(5000000000)
	tx := types.NewTransaction(nonce, toAccount.Address, big.NewInt(100000), gasLimit, gasPrice, nil)

	signedTx, _ := w.SignTx(tx, fromAccount)

	s, _ := json.Marshal(signedTx)

	fmt.Println("signed Tx: ", string(s))
}
