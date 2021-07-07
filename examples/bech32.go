package main

import (
	"fmt"
	"platon-go-sdk/common/bech32util"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/network"
)

func main() {
	bytes, _ := hexutil.Decode("0x1963dd5b88accDA8F86C0D9A487c36cCDC0Aba0F")
	b32Addr, _ := bech32util.ConvertAndEncode(network.MainNetHrp, bytes)
	fmt.Println("bech32 addr: ", b32Addr)

	hrp, ethAddr, _ := bech32util.DecodeAndConvert(b32Addr)
	fmt.Printf("eth hrp: %s, hex: %s\n", hrp, hexutil.Encode(ethAddr))
}
