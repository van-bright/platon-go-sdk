package codec

import (
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/rlp"
)

type NodeId struct {
	HexStringId string
}

func (ni NodeId) ByteEncode() BytesSlice {
	return HexStringParam{ni.HexStringId}.ByteEncode()
}

type HexStringParam struct {
	HexStringValue string
}

func (hsp HexStringParam) ByteEncode() BytesSlice {
	bs, err := hexutil.Decode(hsp.HexStringValue)
	if err != nil {
		panic(err)
	}

	ebs, err := rlp.EncodeToBytes(bs)
	if err != nil {
		panic(err)
	}
	return ebs
}

type Utf8String struct {
	ValueInner string
}

func (u8str Utf8String) ByteEncode() BytesSlice {
	bs, err := rlp.EncodeToBytes(u8str.ValueInner)
	if err != nil {
		panic(err)
	}

	return bs
}
