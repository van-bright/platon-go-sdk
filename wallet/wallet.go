package wallet

import "crypto"

type AlayaAccount struct {
	crypto.PublicKey
}

func ImportPrivateKey(hexPriv string) (bool, error) {
	return true, nil
}

func CreateAlayaAccount() (AlayaAccount, error) {
	return AlayaAccount{}, nil
}
