package web3g

import (
	"encoding/json"
	"fmt"
	"math/big"
	"platon-go-sdk/rpc"
)

type Web3g struct {
	httpClient *rpc.HttpClient
}

func New(url string) *Web3g {
	client := rpc.NewHttpClient(url)
	return &Web3g{ client}
}

func ParseTagOrNumber(qt interface{}) (string, error) {
	spos,ok := qt.(string)

	if ok && (spos == "latest" || spos == "earliest" || spos == "pending") {
		return spos, nil
	} else {
		ipos, ok := qt.(big.Int)
		if !ok {
			return "", fmt.Errorf("req.TagOrNumber can only be integer or 'latest', 'earliest', 'pending'");
		}

		if ipos.BitLen() == 0 {
			return "0x0", nil
		}
		return fmt.Sprintf("%#x", ipos), nil
	}
}

func ParseHttpResponseToString(resp *rpc.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf("%s", resp.Error)
	}
	//return resp.Result.(string), nil
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

// To get the client version of specific node.
func (web3g *Web3g) ClientVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(Web3ClientVersion, nil)
	return ParseHttpResponseToString(resp, err)
}
// To get a hash of given hex string which starts with '0x',
// and the hash algorithm is 'keccak-256'
func (web3g *Web3g) Sha3(dates string) (string, error) {
	reqData := make([]string, 1)
	reqData[0] = dates
	resp, err := web3g.httpClient.PostAsResponse(Web3Sha3, reqData);
	return ParseHttpResponseToString(resp, err)
}
