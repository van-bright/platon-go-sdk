package web3g

func (web3g *Web3g) NetVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetVersion, nil);
	if err != nil {
		return "", err
	}
	return resp.Result.(string), nil
}

func (web3g *Web3g) NetListening() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetListening, nil);
	if err != nil {
		return "", err
	}
	return resp.Result.(string), nil
}

func (web3g *Web3g) NetPeerCount() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetPeerCount, nil);
	if err != nil {
		return "", err
	}
	return resp.Result.(string), nil
}
