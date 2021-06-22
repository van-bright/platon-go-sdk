package codec

import (
	"fmt"
	"math/big"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/rlp"
)

type UInt16 struct {
	ValueInner uint16
}

func (u16 UInt16) ByteEncode() BytesSlice {
	b, err := rlp.EncodeToBytes(u16.ValueInner)
	if err != nil {
		panic(err)
	}
	fmt.Println("f: " + hexutil.Encode(b)[2:])
	return b
}

type UInt256 struct {
	ValueInner *big.Int
}

func (u256 UInt256) ByteEncode() BytesSlice {
	b, err := rlp.EncodeToBytes(u256.ValueInner)
	if err != nil {
		panic(err)
	}
	return b
}

type UInt64 struct {
	ValueInner *big.Int
}

func (u64 UInt64) ByteEncode() BytesSlice {
	b, err := rlp.EncodeToBytes(u64.ValueInner)
	if err != nil {
		panic(err)
	}
	return b
}
