package web3go

import (
	"fmt"
	"testing"
)

var (
	mnemonic   = "bounce half kite tuition catalog buffalo void awesome drink bunker guard vapor"
	privateKey = "7544e3822955215e1a02deeda52cccdbf3b57a015dda815ad6dd21716960d763"
	ethAccount = "0x5C267019f2022e6968CD790775f30Ac981896D4C"
	atpAccount = "atp1tsn8qx0jqghxj6xd0yrhtuc2exqcjm2vh04lhp"
	atxAccount = "atx1tsn8qx0jqghxj6xd0yrhtuc2exqcjm2vaff4yt"
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
