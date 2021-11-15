package ppos

import (
	"context"
	"encoding/json"
	"fmt"
	platongosdk "github.com/oldmanfan/platon-go-sdk"
	common2 "github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/core/types"
	"github.com/oldmanfan/platon-go-sdk/ethclient"
	"github.com/oldmanfan/platon-go-sdk/network"
	"github.com/oldmanfan/platon-go-sdk/ppos/typedefs"
	"math/big"
)

type FunctionExecutor struct {
	httpEntry    string
	chainId      *big.Int
	contractAddr string
	credentials  *typedefs.Credentials
}

type CallResponse struct {
	Code uint64          `json:"Code"`
	Ret  json.RawMessage `json:"Ret"`
}

func (fe *FunctionExecutor) SendWithRaw(f *typedefs.Function) (json.RawMessage, error) {
	to := fe.credentials.MustBech32ToAddress(fe.contractAddr)
	data := f.ToBytes()

	gasPrice := fe.getDefaultGasPrice(f)
	gasLimit := fe.getDefaultGasLimit(f)
	chainId := fe.chainId

	r, err := fe.doSendRawTx(chainId, to, data, nil, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	//fmt.Println("[SendWithRaw] Http Response: " + string(r))
	return r, nil
}

func (fe *FunctionExecutor) SendWithResult(f *typedefs.Function, result interface{}) error {
	raw, err := fe.SendWithRaw(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &result)
	return err
}

func (fe *FunctionExecutor) getDefaultGasPrice(f *typedefs.Function) *big.Int {
	price := new(big.Int)
	switch f.Type {
	case typedefs.SUBMIT_TEXT_FUNC_TYPE:
		price.SetString("1500000000000000", 10)
	case typedefs.SUBMIT_VERSION_FUNC_TYPE:
		price.SetString("2100000000000000", 10)
	case typedefs.SUBMIT_PARAM_FUNCTION_TYPE:
		price.SetString("2000000000000000", 10)
	case typedefs.SUBMIT_CANCEL_FUNC_TYPE:
		price.SetString("3000000000000000", 10)
	default:
		price = big.NewInt(0)
	}
	return price
}

func (fe *FunctionExecutor) getDefaultGasLimit(f *typedefs.Function) uint64 {
	if typedefs.IsLocalSupportFunction(f.Type) {
		return typedefs.GetGasLimit(f)
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
		msg := platongosdk.CallMsg2{
			From:     fromAddr,
			To:       to.Bech32WithPrefix(network.MainNetHrp),
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

func (fe *FunctionExecutor) CallWithRaw(f *typedefs.Function) ([]byte, error) {
	to := fe.credentials.MustBech32ToAddress(fe.contractAddr)
	data := f.ToBytes()

	return fe.doCallRawTx(to, data)
}

func (fe *FunctionExecutor) CallWithResult(f *typedefs.Function, result interface{}) error {
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
	//fmt.Println("[doCallRawTx] HTTP RESPONSE: " + string(b))
	if err != nil {
		return nil, err
	}
	var callRsp CallResponse
	err = json.Unmarshal(b, &callRsp)
	if err != nil {
		return nil, err
	}
	if callRsp.Code != 0 {
		return nil, fmt.Errorf(string(callRsp.Ret))
	}
	return callRsp.Ret, nil
}
