package ppos

import "platon-go-sdk/web3go"

type BaseContract struct {
	geb3 *web3go.Geb3
}

func (bc *BaseContract) Send() error {
	return nil
}
