package main

import (
	"encoding/json"
	"fmt"
	"platon-go-sdk/wallet"
)

//func main() {
//	mnemonic := "bounce half kite tuition catalog buffalo void awesome drink bunker guard vapor"
//	w, _ := wallet.ImportMnemonic(mnemonic)
//	w.CreateAccount(0)
//	w.CreateAccount(1)
//
//	ta, _ := json.Marshal(w.Export())
//	fmt.Print("w: ", string(ta))
//}

/**
生成测试账号地址:
{
"mnemonics":"always brick access science decade nasty marriage attack fame topple pen add",
"accounts":
	[
		{
			"private-key":"b72faaa798c44d2359d0ccb35dd39446c9c18905fdcecede42a6570ff177ae08",
			"eth-address":"0x17d4C98a69d0141E3871E66BF79069623fdC2160",
			"mainnet-addresses":"atp1zl2vnznf6q2puwr3ue4l0yrfvglacgtqypk432",
			"testnet-addresses":"atx1zl2vnznf6q2puwr3ue4l0yrfvglacgtqw82lzq",
			"path":"m/44'/60'/0'/0/0"
		},
		{
			"private-key":"3d04b39cd4439340597fe5027c58797e824173cdab45b1f155539026f72b7029",
			"eth-address":"0x237a0C8507854E805C4a2dACF70d31F81ab944ED",
			"mainnet-addresses":"atp1ydaqepg8s48gqhz29kk0wrf3lqdtj38d8mkcz3",
			"testnet-addresses":"atx1ydaqepg8s48gqhz29kk0wrf3lqdtj38dda2j3m",
			"path":"m/44'/60'/0'/0/1"
		}
	]
}

 */

func main() {
	w, _ := wallet.Create()
	w.CreateAccount(0)
	w.CreateAccount(1)

	ta, _ := json.Marshal(w.Export())
	fmt.Print("w: ", string(ta))
}
