package main

import (
	"encoding/json"
	"fmt"
	"platon-go-sdk/wallet"
)

func main() {
	mnemonic := "bounce half kite tuition catalog buffalo void awesome drink bunker guard vapor"
	w, _ := wallet.CreateAlayaWalletByMnemonic(mnemonic)
	w.CreateAccount(0)
	w.CreateAccount(1)

	ta, _ := json.Marshal(w.Export())
	fmt.Print("w: ", string(ta))
}
