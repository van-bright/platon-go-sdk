package ppos

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	platongosdk "platon-go-sdk"
	common2 "platon-go-sdk/common"
	"platon-go-sdk/core/types"
	"platon-go-sdk/ethclient"
	"platon-go-sdk/ppos/common"
)

type FunctionExecutor struct {
	httpEntry    string
	chainId      *big.Int
	contractAddr string
	credentials  *common.Credentials
}

type CallResponse struct {
	Code uint64          `json:"Code"`
	Ret  json.RawMessage `json:"Ret"`
}

func (fe *FunctionExecutor) SendWithRaw(f *common.Function) (json.RawMessage, error) {
	to := fe.credentials.MustBech32ToAddress(fe.contractAddr)
	data := f.ToBytes()

	gasPrice := fe.getDefaultGasPrice(f)
	gasLimit := fe.getDefaultGasLimit(f)
	chainId := fe.chainId

	r, err := fe.doSendRawTx(chainId, to, data, nil, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	fmt.Println("[SendWithRaw] Http Response: " + string(r))
	return r, nil
	//return fe.doSendRawTx(chainId, to, data, big.NewInt(0), gasPrice, gasLimit)
}

func (fe *FunctionExecutor) SendWithResult(f *common.Function, result interface{}) error {
	raw, err := fe.SendWithRaw(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &result)
	return err
}

func (fe *FunctionExecutor) getDefaultGasPrice(f *common.Function) *big.Int {
	price := new(big.Int)
	switch f.Type {
	case common.SUBMIT_TEXT_FUNC_TYPE:
		price.SetString("1500000000000000", 10)
	case common.SUBMIT_VERSION_FUNC_TYPE:
		price.SetString("2100000000000000", 10)
	case common.SUBMIT_PARAM_FUNCTION_TYPE:
		price.SetString("2000000000000000", 10)
	case common.SUBMIT_CANCEL_FUNC_TYPE:
		price.SetString("3000000000000000", 10)
	default:
		price = big.NewInt(0)
	}
	return price
}

func (fe *FunctionExecutor) getDefaultGasLimit(f *common.Function) uint64 {
	if common.IsLocalSupportFunction(f.Type) {
		return common.GetGasLimit(f)
	} else {
		return 0
	}
}

func (fe *FunctionExecutor) doSendRawTx(chainId *big.Int, to common2.Address, data []byte, value *big.Int, gasPrice *big.Int, gasLimit uint64) (json.RawMessage, error) {
	client, err := ethclient.Dial(fe.httpEntry)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	if gasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
	}

	fromAddr := fe.credentials.Bech32Address()
	nonce, err := client.NonceAt(ctx, fromAddr, "pending")
	if err != nil {
		return nil, err
	}

	if gasLimit == 0 {
		msg := platongosdk.CallMsg{
			From:     fe.credentials.Address(),
			To:       &to,
			Gas:      0,
			GasPrice: gasPrice,
			Value:    value,
			Data:     data,
		}
		gasLimit, err = client.EstimateGas(ctx, msg)
		if err != nil {
			return nil, err
		}
	}

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	signedTx, err := fe.credentials.SignTx(tx, chainId)
	if err != nil {
		return nil, err
	}

	return client.SendRawTransaction(ctx, signedTx)
}

func (fe *FunctionExecutor) CallWithRaw(f *common.Function) ([]byte, error) {
	to := fe.credentials.MustBech32ToAddress(fe.contractAddr)
	data := f.ToBytes()

	return fe.doCallRawTx(to, data)
}

func (fe *FunctionExecutor) CallWithResult(f *common.Function, result interface{}) error {
	raw, err := fe.CallWithRaw(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &result)
	return err
}

func (fe *FunctionExecutor) doCallRawTx(to common2.Address, data []byte) ([]byte, error) {
	client, err := ethclient.Dial(fe.httpEntry)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	msg := platongosdk.CallMsg{
		From:     fe.credentials.Address(),
		To:       &to,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     data,
	}

	b, err := client.CallContract(ctx, msg, "latest")
	fmt.Println("[doCallRawTx] HTTP RESPONSE: " + string(b))
	if err != nil {
		return nil, err
	}
	var callRsp CallResponse
	err = json.Unmarshal(b, &callRsp)
	if err != nil {
		return nil, err
	}
	return callRsp.Ret, nil
}
