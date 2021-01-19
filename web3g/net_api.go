package web3g

func (web3g *Web3g) NetVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetVersion, nil);
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) NetListening() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetListening, nil);
	return ParseHttpResponseToString(resp, err)
}

func (web3g *Web3g) NetPeerCount() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetPeerCount, nil);
	return ParseHttpResponseToString(resp, err)
}
