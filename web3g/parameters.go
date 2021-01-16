package web3g

type Response struct {
	Version string `json:"jsonrpc"`
	ID      uint64  `json:"id"`
	Result  string  `json:"result,omitempty"`
}
type ClientVersionResp string
type Web3Sha3Resp string