package codec

import (
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/rlp"
)

type NodeId struct {
	HexStringId string
}

func (ni NodeId) ByteEncode() BytesSlice {
	bs, err := hexutil.Decode(ni.HexStringId)
	if err != nil {
		panic(err)
	}

	ebs, err := rlp.EncodeToBytes(bs)
	if err != nil {
		panic(err)
	}
	return ebs
}
