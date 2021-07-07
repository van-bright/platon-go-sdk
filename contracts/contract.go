package contracts

import (
	"github.com/AlayaNetwork/Alaya-Go/accounts/abi/bind"
	"github.com/AlayaNetwork/Alaya-Go/crypto"
	"github.com/AlayaNetwork/Alaya-Go/ethclient"
)

type Contract struct {
	Url        string
	PrivateKey string
}

func (c *Contract) Init() (*bind.TransactOpts, bind.ContractBackend, error) {
	client, err := ethclient.Dial(c.Url)
	if err != nil {
		return nil, nil, err
	}

	if len(c.PrivateKey) == 0 {
		return nil, client, nil
	} else {
		privateKey, err := crypto.HexToECDSA(c.PrivateKey)
		if err != nil {
			return nil, nil, err
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = nil
		auth.Value = nil
		auth.GasLimit = 0
		auth.GasPrice = nil

		return auth, client, nil
	}
}
