package common

import (
	"fmt"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/ppos/codec"
)

type Function struct {
	Type        FunctionType
	InputParams []interface{}
}

func NewFunction(typ FunctionType, params []interface{}) *Function {
	return &Function{typ, params}
}

func (f *Function) encodeType() codec.BytesSlice {
	typ := f.Type.GetType()
	encoded := codec.MinimalByteArray(typ)
	//fmt.Println("ftype0: " + hexutil.Encode(encoded))
	b := codec.EncodeBytes(encoded, codec.OFFSET_SHORT_STRING)
	//fmt.Println("ftype1: " + hexutil.Encode(b))
	return b
}

func (f *Function) encodeParams() []codec.BytesSlice {
	return []codec.BytesSlice{}
}

func (f *Function) ToBytes() []byte {
	ftype := f.encodeType()
	params := f.encodeParams()

	argsList := append([]codec.BytesSlice{ftype}, params...)
	argsBytes := codec.EncodeBytesSlice(argsList)

	data := hexutil.Encode(argsBytes)
	fmt.Println("encode type: " + data)

	return argsBytes
}
