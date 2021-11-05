package codec

import (
	"github.com/oldmanfan/platon-go-sdk/common/hexutil"
	"github.com/oldmanfan/platon-go-sdk/rlp"
)

type NodeId struct {
	HexStringId string
}

func (ni NodeId) GetEncodeData() BytesSlice {
	return HexStringParam{ni.HexStringId}.GetEncodeData()
}

type HexStringParam struct {
	HexStringValue string
}

func (hsp HexStringParam) GetEncodeData() BytesSlice {
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

func (u8str Utf8String) GetEncodeData() BytesSlice {
	bs, err := rlp.EncodeToBytes(u8str.ValueInner)
	if err != nil {
		panic(err)
	}

	return bs
}
