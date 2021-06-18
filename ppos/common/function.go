package common

import (
	"fmt"
	"platon-go-sdk/common"
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
	params := f.InputParams
	var result []codec.BytesSlice
	for _, p := range params {
		//result = append(result, p.(codec.BytesSlice))
		switch p.(type) {
		case []byte, codec.BytesSlice:
			result = append(result, codec.EncodeBytes(p.([]byte), codec.OFFSET_SHORT_STRING))
		case common.Address:
			r := codec.EncodeBytes(p.(common.Address).Bytes(), codec.OFFSET_SHORT_STRING)
			result = append(result, r)
		case []codec.BytesSlice:
			var r []byte
			for _, it := range p.([]codec.BytesSlice) {
				r = append(r, codec.EncodeBytes(it, codec.OFFSET_SHORT_STRING)...)
			}
			r = codec.EncodeBytes(r, codec.OFFSET_SHORT_LIST)
			result = append(result, r)
		default:
			fmt.Println("function parameters with not support type")
		}
	}
	return result
}

func (f *Function) ToBytes() []byte {
	ftype := f.encodeType()
	params := f.encodeParams()

	argsList := append([]codec.BytesSlice{ftype}, params...)
	fmt.Println("Function Parameters List:")
	for _, p := range argsList {
		fmt.Println(hexutil.Encode(p)[2:])
	}

	argsBytes := codec.EncodeBytesSlice(argsList)

	data := hexutil.Encode(argsBytes)
	fmt.Println("encode type: " + data)

	return argsBytes
}
