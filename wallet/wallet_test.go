package wallet

import "testing"

var (
	mnemonic   = "bounce half kite tuition catalog buffalo void awesome drink bunker guard vapor"
	privateKey = "7544e3822955215e1a02deeda52cccdbf3b57a015dda815ad6dd21716960d763"
	ethAccount = "0x5C267019f2022e6968CD790775f30Ac981896D4C"
	atpAccount = "atp1tsn8qx0jqghxj6xd0yrhtuc2exqcjm2vh04lhp"
	atxAccount = "atx1tsn8qx0jqghxj6xd0yrhtuc2exqcjm2vaff4yt"
	path       = "m/44'/60'/0'/0/0"
)

func TestCreateAlayaWalletWithMnemonic(t *testing.T) {

}

func TestCreateAlayaWalletByMnemonic(t *testing.T) {
	alayawallet, err := CreateAlayaWalletByMnemonic(mnemonic)

	if err != nil {
		t.Fatal("create alaya wallet failed")
	}

	account, err := alayawallet.CreateAccount(0)
	if err != nil {
		t.Fatal("create account failed: ", err)
	}

	if account.Address.Hex() != ethAccount {
		t.Error("create eth account failed.")
	}

	if pk, _ := alayawallet.PrivateKeyHex(account); pk != privateKey {
		t.Error("convert to private key failed.")
	}

	if atp, _ := alayawallet.MainNetAccount(account); atp != atpAccount {
		t.Error("convert to main net address failed")
	}

	if atx, _ := alayawallet.TestNetAccount(account); atx != atxAccount {
		t.Error("convert to test net address failed")
	}

	if pi, _ := alayawallet.Path(account); pi != path {
		t.Error("retrieve account path failed.")
	}
}
