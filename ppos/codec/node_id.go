package codec

import (
	"platon-go-sdk/common/hexutil"
)

type NodeId struct {
	HexStringId string
}

func (ni NodeId) ByteEncode() BytesSlice {
	bs, err := hexutil.Decode(ni.HexStringId)
	if err != nil {
		panic(err)
	}
	return bs
}
