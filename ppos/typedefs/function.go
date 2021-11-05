package typedefs

import (
	"fmt"
	"github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/common/hexutil"
	"github.com/oldmanfan/platon-go-sdk/ppos/codec"
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
	var encodeParamItem func(p interface{}) []byte
	encodeParamItem = func(p interface{}) []byte {
		switch p.(type) {
		case []byte:
			return codec.EncodeBytes(p.([]byte), codec.OFFSET_SHORT_STRING)
		case codec.BytesSlice:
			return codec.EncodeBytes(p.(codec.BytesSlice), codec.OFFSET_SHORT_STRING)
		case common.Address:
			return codec.EncodeBytes(p.(common.Address).Bytes(), codec.OFFSET_SHORT_STRING)
		case codec.NodeId:
			return p.(codec.NodeId).GetEncodeData()
		case codec.HexStringParam:
			return p.(codec.HexStringParam).GetEncodeData()
		case codec.Utf8String:
			return p.(codec.Utf8String).GetEncodeData()
		case codec.UInt16, codec.UInt32, codec.UInt64, codec.UInt256:
			return p.(codec.ParamEncoder).GetEncodeData()
		case codec.RlpList:
			encoder := codec.RlpEncoder{}
			return encoder.Encode(p.(codec.RlpList))
		case []codec.NodeId:
			var r []byte
			for _, it := range p.([]codec.NodeId) {
				r = append(r, encodeParamItem(it)...)
			}
			r = codec.EncodeBytes(r, codec.OFFSET_SHORT_LIST)
			return r
		case []codec.BytesSlice:
			var r []byte
			for _, it := range p.([]codec.BytesSlice) {
				r = append(r, encodeParamItem(it)...)
			}
			r = codec.EncodeBytes(r, codec.OFFSET_SHORT_LIST)
			return r
		default:
			panic("function parameters with not support type")
		}
	}

	params := f.InputParams
	var result []codec.BytesSlice
	for _, p := range params {
		result = append(result, encodeParamItem(p))
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
	fmt.Println("Function Data is: " + data[2:])

	return argsBytes
}
