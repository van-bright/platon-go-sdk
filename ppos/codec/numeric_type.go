package codec

import (
	"github.com/oldmanfan/platon-go-sdk/rlp"
	"math/big"
)

type UInt16 struct {
	ValueInner uint16
}

func (u16 UInt16) GetEncodeData() BytesSlice {
	b, err := rlp.EncodeToBytes(u16.ValueInner)
	if err != nil {
		panic(err)
	}
	//fmt.Println("f: " + hexutil.Encode(b)[2:])
	return b
}

type UInt32 struct {
	ValueInner *big.Int
}

func (u32 UInt32) GetEncodeData() BytesSlice {
	b, err := rlp.EncodeToBytes(u32.ValueInner)
	if err != nil {
		panic(err)
	}
	return b
}

type UInt256 struct {
	ValueInner *big.Int
}

func (u256 UInt256) GetEncodeData() BytesSlice {
	b, err := rlp.EncodeToBytes(u256.ValueInner)
	if err != nil {
		panic(err)
	}
	return b
}

type UInt64 struct {
	ValueInner *big.Int
}

func (u64 UInt64) GetEncodeData() BytesSlice {
	b, err := rlp.EncodeToBytes(u64.ValueInner)
	if err != nil {
		panic(err)
	}
	return b
}
