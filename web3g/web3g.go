package web3g

import (
	"encoding/json"
	"platon-go-sdk/rpc"
)

type Web3g struct {
	httpClient *rpc.HttpClient
}

func New(url string) *Web3g {
	client := rpc.NewHttpClient(url)
	return &Web3g{ client}
}
// To get the client version of specific node.
func (web3g *Web3g) ClientVersion() (ClientVersionResp, error) {
	resp, _ := web3g.httpClient.Post(Web3_clientVersion, nil)

	var body = &Response{};
	if err := json.Unmarshal(resp, body); err != nil {
		return "", err
	} else {
		return ClientVersionResp(body.Result), nil
	}
}
// To get a hash of given hex string which starts with '0x',
// and the hash algorithm is 'keccak-256'
func (web3g *Web3g) Sha3(dates string) (Web3Sha3Resp, error) {
	reqData := make([]string, 1)
	reqData[0] = dates
	resp, _ := web3g.httpClient.Post(Web3_sha3, reqData);
	var body = &Response{}
	if err := json.Unmarshal(resp, body); err != nil {
		return "", err
	} else {
		return Web3Sha3Resp(body.Result), nil
	}

}