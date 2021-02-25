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
		//return ipos.Text(16), nil
		return fmt.Sprintf("0x%x", &ipos), nil
	}
}

func ParseHttpResponseToResult(resp *rpc.Response, result interface{}, err error) error {
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf("%s", resp.Error)
	}

	e := json.Unmarshal(resp.Result, &result)
	return e
}

// To get the client version of specific node.
func (web3g *Web3g) ClientVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(Web3ClientVersion, nil)

	var version string
	err = ParseHttpResponseToResult(resp, &version, err)
	return version, err
}
// To get a hash of given hex string which starts with '0x',
// and the hash algorithm is 'keccak-256'
func (web3g *Web3g) Sha3(dates string) (string, error) {
	reqData := make([]string, 1)
	reqData[0] = dates
	resp, err := web3g.httpClient.PostAsResponse(Web3Sha3, reqData);
	var sha3 string
	err = ParseHttpResponseToResult(resp, &sha3, err)
	return sha3, err
}
