package main

import (
	"fmt"
	"platon-go-sdk/bech32"
)

func main() {
	//data := []byte("hello")
	b32Addr, _ := bech32.EncodeAddress(bech32.MainNet, "0x1963dd5b88accDA8F86C0D9A487c36cCDC0Aba0F")
	fmt.Println("bech32 addr: ", b32Addr)

	hrp, ethAddr, _ := bech32.DecodeAddress(b32Addr);
	fmt.Printf("eth hrp: %s, hex: %s\n", hrp, ethAddr)
}
