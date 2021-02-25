package web3g

func (web3g *Web3g) NetVersion() (string, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetVersion, nil)
	var version string
	err = ParseHttpResponseToResult(resp, &version, err)
	return version, err
}

func (web3g *Web3g) NetListening() (bool, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetListening, nil)

	var listening bool
	err = ParseHttpResponseToResult(resp, &listening, err)
	return listening, err
}

func (web3g *Web3g) NetPeerCount() (uint32, error) {
	resp, err := web3g.httpClient.PostAsResponse(NetPeerCount, nil)

	var count uint32
	err = ParseHttpResponseToResult(resp, &count, err)
	return count, err
}
