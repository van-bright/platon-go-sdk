package web3g

import "encoding/json"

// FIXME(liaingqin.fan) 所以admin开头的API暂时都不能测试, 服务端返回错误

func (web3g *Web3g) AdminAddPeer(data string) (bool, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminAddPeer, data)
	if err != nil {
		return false, err
	}
	var result bool
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

// NodeInfo represents a short summary of the information known about the host.
type NodeInfo struct {
	ID     string `json:"id"`        // Unique node identifier (also the encryption key)
	Name   string `json:"name"`      // Name of the node, including client type, version, OS, custom data
	BlsPub string `json:"blsPubKey"` // BLS public key
	Enode  string `json:"enode"`     // Enode URL for adding this peer from remote peers
	IP     string `json:"ip"`        // IP address of the node
	Ports  struct {
		Discovery int `json:"discovery"` // UDP listening port for discovery protocol
		Listener  int `json:"listener"`  // TCP listening port for RLPx
	} `json:"ports"`
	ListenAddr string                 `json:"listenAddr"`
	Protocols  map[string]interface{} `json:"protocols,omitempty"`
}

func (web3g *Web3g) AdminNodeInfo() (*NodeInfo, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminNodeInfo, nil)
	if err != nil {
		return nil, err
	}
	var result NodeInfo
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) AdminPeers() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminPeers, nil)
	if err != nil {
		return "", err
	}
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

type ProgramVersion struct {
	Version uint32
	Sign    string
}
func (web3g *Web3g) AdminGetProgramVersion() (*ProgramVersion, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminGetProgramVersion, nil)
	if err != nil {
		return nil, err
	}
	var result ProgramVersion
	e := json.Unmarshal(resp.Result, &result)
	return &result, e
}

func (web3g *Web3g) AdminGetSchnorrNIZKProve() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminGetSchnorrNIZKProve, nil)
	if err != nil {
		return "", err
	}
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}

func (web3g *Web3g) AdminDataDir() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminDataDir, nil)
	if err != nil {
		return "", err
	}
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}
