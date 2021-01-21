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

func (web3g *Web3g) AdminNodeInfo() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminNodeInfo, nil)
	if err != nil {
		return "", err
	}
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
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

func (web3g *Web3g) AdminGetProgramVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminGetProgramVersion, nil)
	if err != nil {
		return "", err
	}
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
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

func (web3g *Web3g) AdminDatadir() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(AdminDatadir, nil)
	if err != nil {
		return "", err
	}
	var result string
	e := json.Unmarshal(resp.Result, &result)
	return result, e
}
