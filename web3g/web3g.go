package web3g

import (
	"fmt"
	"platon-go-sdk/rpc"
	"strconv"
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
		ipos, ok := qt.(int)
		if !ok {
			return "", fmt.Errorf("req.TagOrNumber can only be integer or 'latest', 'earliest', 'pending'");
		}
		return strconv.Itoa(ipos), nil
	}
}

func ParseHttpResponseToString(resp *rpc.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf("%s", resp.Error)
	}
	return resp.Result.(string), nil
}

// To get the client version of specific node.
func (web3g *Web3g) ClientVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(Web3ClientVersion, nil)
	if err != nil {
		return "", err
	}
	return resp.Result.(string), nil
}
// To get a hash of given hex string which starts with '0x',
// and the hash algorithm is 'keccak-256'
func (web3g *Web3g) Sha3(dates string) (string, error) {
	reqData := make([]string, 1)
	reqData[0] = dates
	resp, err := web3g.httpClient.PostAsResponse(Web3Sha3, reqData);
	if err != nil {
		return "", err
	}
	return resp.Result.(string), nil
}