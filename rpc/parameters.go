package rpc

type Response struct {
	Version string `json:"jsonrpc"`
	ID      uint64  `json:"id"`
	Result  interface{}  `json:"result,omitempty"`
	Error   interface{}  `json:"error,omitempty"`
}